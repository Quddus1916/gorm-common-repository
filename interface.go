package gorm_common_repository

// CommonRepositoryInterface is an interface that defines common database operations for any type Model using the GORM library.
type CommonRepositoryInterface[Model any] interface {
	// CreateRecord inserts a new record into the database.
	// It takes a dataObject of type Model and returns the newly created record
	// or an error if any is encountered during the operation.
	CreateRecord(dataObject Model) (Model, error)

	// CreateBulkRecords inserts multiple records into the database.
	// It takes a slice of dataObjects of type Model and returns the newly created records
	// or an error if any is encountered during the operation.
	CreateBulkRecords(dataObjects []Model) ([]Model, error)

	// GetRecordByID retrieves a record from the database by its unique identifier.
	// It takes a queryID of any type and returns the found record
	// or an error if any is encountered during the operation.
	// If no record is found, it should return an error indicating the record was not found.
	GetRecordByID(queryID interface{}) (Model, error)

	// GetRecordByAttributes retrieves a record from the database based on query parameters.
	// It takes queryParams as a map where keys represent attribute names
	// and values are the target attribute values. It returns the found record
	// or an error if any is encountered during the operation.
	// If no record is found, it should return an error indicating the record was not found.
	GetRecordByAttributes(queryParams map[string]interface{}) (Model, error)

	// GetRecordsForMultipleIDs retrieves a list of records from the database based on multiple IDs.
	// It takes queryIDs as a slice of any type and returns the found records
	// or an error if any is encountered during the operation.
	GetRecordsForMultipleIDs(queryIDs []interface{}) ([]Model, error)

	// GetRecordsByMultipleAttributeValues retrieves a list of records from the database
	// based on multiple attribute values.
	// It takes queryValues as a map where keys represent attribute names
	// and values are slices of attribute values. It returns the found records
	// or an error if any is encountered during the operation.
	GetRecordsByMultipleAttributeValues(queryValues map[string][]interface{}) ([]Model, error)

	// GetRecordsByQueryParams retrieves a paginated list of records from the database
	// based on query parameters and sorting criteria.
	// It takes queryParams of any type, pagination information (PageReq), and sorting information (SortReq).
	// It returns the found records or an error if any is encountered during the operation.
	GetRecordsByQueryParams(queryParams *QueryParams) ([]Model, error)

	// GetRecordCount retrieves the count of records from the database based on query parameters.
	// It takes queryParams of type *QueryParams and returns the count of matching records
	// or an error if any is encountered during the operation.
	GetRecordCount(queryParams *QueryParams) (int64, error)

	// UpdateRecordByID updates a record in the database by its unique identifier.
	// It takes queryID of any type and a map of data to update.
	// It returns an error if no records match the criteria or if any error is encountered during the operation.
	// If no record is found to update, it should return an error indicating the record was not found.
	UpdateRecordByID(queryID interface{}, data map[string]interface{}) error

	// UpdateRecordsByAttributes updates records in the database based on query parameters.
	// It takes queryParams as a map where keys represent attribute names and values are the target attribute values to query records.
	// It also takes a map of data to update the matching records.
	// It returns an error if no records match the criteria or if any error is encountered during the operation.
	// If no records are found to update, it should return an error indicating no records were found.
	UpdateRecordsByAttributes(queryParams map[string]interface{}, data map[string]interface{}) error

	// DeleteRecordByID deletes a record from the database by its unique identifier.
	// It takes queryID of any type and returns an error if no records match the criteria
	// or if any error is encountered during the operation.
	// If no record is found to delete, it should return an error indicating the record was not found.
	DeleteRecordByID(queryID interface{}) error

	// DeleteRecordsByAttributes deletes records from the database based on query parameters.
	// It takes queryParams as a map where keys represent attribute names and values are the target attribute values to query records.
	// It returns an error if no records match the criteria or if any error is encountered during the operation.
	// If no records are found to delete, it should return an error indicating no records were found.
	DeleteRecordsByAttributes(queryParams map[string]interface{}) error
}
