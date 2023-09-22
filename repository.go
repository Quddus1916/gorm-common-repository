package gorm_common_repository

import (
	"gorm.io/gorm"
)

// CommonRepository is a generic repository for common database operations using GORM.
// It is parameterized with a generic type Model, representing the model structure used in the repository.
type CommonRepository[Model any] struct {
	genericModelStruct Model    // The generic model structure for this repository.
	tableName          string   // The name of the database table associated with the model.
	dbClient           *gorm.DB // The GORM database client used for database operations.
}

// NewCommonRepository creates a new instance of CommonRepository for a specific model type.
// It takes the tableName as the name of the associated database table and dbClient as the GORM database client.
// The returned repository instance can be used to perform common database operations for the specified model type.
func NewCommonRepository[Model any](tableName string, dbClient *gorm.DB) CommonRepositoryInterface[Model] {
	return CommonRepository[Model]{
		tableName: tableName,
		dbClient:  dbClient,
	}
}

func (repo CommonRepository[Model]) CreateRecord(dataObject Model) (Model, error) {
	err := repo.dbClient.
		Table(repo.tableName).
		Create(&dataObject).Error
	if err != nil {
		// Check if the error is a duplicate entry error, detected by SQL unique constraint.
		if ok, _ := ParseDuplicateEntry(err); ok {
			return repo.genericModelStruct, gorm.ErrDuplicatedKey
		}
		return repo.genericModelStruct, err
	}
	return dataObject, nil
}

func (repo CommonRepository[Model]) CreateBulkRecords(dataObjects []Model) ([]Model, error) {
	err := repo.dbClient.
		Table(repo.tableName).
		Create(&dataObjects).Error
	if err != nil {
		// Check if the error is a duplicate entry error, detected by SQL unique constraint.
		if ok, _ := ParseDuplicateEntry(err); ok {
			return nil, gorm.ErrDuplicatedKey
		}
		return nil, err
	}
	return dataObjects, nil
}

func (repo CommonRepository[Model]) GetRecordByID(queryID interface{}) (Model, error) {
	return repo.GetRecordByAttributes(map[string]interface{}{
		"id": queryID,
	})
}

func (repo CommonRepository[Model]) GetRecordByAttributes(queryParams map[string]interface{}) (Model, error) {
	var dataObject Model
	err := repo.dbClient.
		Table(repo.tableName).
		Where(queryParams).
		First(&dataObject).Error
	if err != nil {
		return repo.genericModelStruct, err
	}
	return dataObject, nil
}

func (repo CommonRepository[Model]) GetRecordsForMultipleIDs(queryIDs []interface{}) ([]Model, error) {
	return repo.GetRecordsByMultipleAttributeValues(map[string][]interface{}{
		"id": queryIDs,
	})
}

func (repo CommonRepository[Model]) GetRecordsByMultipleAttributeValues(queryValues map[string][]interface{}) ([]Model, error) {
	var dataObjects []Model
	err := repo.dbClient.
		Table(repo.tableName).
		Where(queryValues).
		Find(&dataObjects).Error
	if err != nil {
		return nil, err
	}
	return dataObjects, nil
}

func (repo CommonRepository[Model]) GetRecordsByQueryParams(queryParams *QueryParams) ([]Model, error) {
	var dataObjects []Model

	// Preload all related data
	dbQuery := repo.dbClient.Table(repo.tableName)

	if queryParams != nil {
		// Apply Filtering
		dbQuery.Scopes(queryParams.FilterByParams(repo.tableName))

		// Apply pagination
		dbQuery.Scopes(queryParams.Paginate())

		// Apply sorting
		dbQuery.Scopes(queryParams.SortByDirection())
	}

	err := dbQuery.Find(&dataObjects).Error
	if err != nil {
		return nil, err
	}
	return dataObjects, nil
}

func (repo CommonRepository[Model]) GetRecordCount(queryParams *QueryParams) (int64, error) {
	var totalCount int64

	// Preload all related data
	dbQuery := repo.dbClient.Table(repo.tableName)

	if queryParams != nil {
		// Apply Filtering
		dbQuery.Scopes(queryParams.FilterByParams(repo.tableName))
	}

	err := dbQuery.Count(&totalCount).Error
	if err != nil {
		return 0, err
	}
	return totalCount, nil
}

func (repo CommonRepository[Model]) UpdateRecordByID(queryID interface{}, data map[string]interface{}) error {
	return repo.UpdateRecordsByAttributes(map[string]interface{}{"id": queryID}, data)
}

func (repo CommonRepository[Model]) UpdateRecordsByAttributes(queryParams map[string]interface{}, data map[string]interface{}) error {
	res := repo.dbClient.
		Table(repo.tableName).
		Where(queryParams).
		UpdateColumns(data)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (repo CommonRepository[Model]) DeleteRecordByID(queryID interface{}) error {
	return repo.DeleteRecordsByAttributes(map[string]interface{}{"id": queryID})
}

func (repo CommonRepository[Model]) DeleteRecordsByAttributes(queryParams map[string]interface{}) error {
	res := repo.dbClient.
		Table(repo.tableName).
		Where(queryParams).
		Delete(&repo.genericModelStruct)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
