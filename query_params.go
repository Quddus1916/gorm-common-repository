package gorm_common_repository

import (
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type GormQueryBuilder func(dbQuery *gorm.DB) *gorm.DB

// Page represents pagination settings.
type Page struct {
	Number int `json:"number" query:"number"` // Number represents the requested page number.
	Limit  int `json:"limit" query:"limit"`   // Limit represents the number of items per page.
}

// DefaultPage returns a new Page with default values.
func DefaultPage() Page {
	return Page{Number: 1, Limit: 10}
}

// SortDirectionType represents the direction of sorting.
type SortDirectionType string

const (
	Ascending  SortDirectionType = "ASC"  // Ascending represents ascending order.
	Descending SortDirectionType = "DESC" // Descending represents descending order.
)

// Sort represents sorting settings.
type Sort struct {
	By        string            `json:"by" query:"by"`               // By represents the field to sort by.
	Direction SortDirectionType `json:"direction" query:"direction"` // Direction represents the sorting direction.
}

// DefaultSort returns a new Sort with default values.
func DefaultSort() Sort {
	return Sort{By: "created_at", Direction: Descending}
}

// FilterAction represents available filtering actions.
type FilterAction string

const (
	Equals           FilterAction = "equals"
	Like             FilterAction = "like"
	In               FilterAction = "in"
	GreaterThan      FilterAction = "greater-than"
	GreaterThanEqual FilterAction = "greater-than-equal"
	LessThan         FilterAction = "less-than"
	LessThanEqual    FilterAction = "less-than-equal"
)

// FilterParam represents an individual filter parameter.
type FilterParam struct {
	Attribute string       `json:"attribute"` // Attribute is the name of the attribute to filter on.
	Action    FilterAction `json:"action"`    // Action is the filtering action to perform.
	Value     string       `json:"value"`     // Value is the value to filter by.
}

// QueryParams represents additional query parameters for filtering data.
type QueryParams struct {
	Page         Page          `json:"page"`          // Page represents the pagination settings.
	Sort         Sort          `json:"sort"`          // Sort represents the sorting settings.
	FilterParams []FilterParam `json:"filter_params"` // FilterParams represents additional filtering criteria.
}

// SortByDirection returns a scope for sorting based on Sort settings.
func (queryParams *QueryParams) SortByDirection() GormQueryBuilder {
	return func(dbQuery *gorm.DB) *gorm.DB {
		return dbQuery.Order(fmt.Sprintf("%s %s", queryParams.Sort.By, queryParams.Sort.Direction))
	}
}

// Paginate returns a scope for pagination based on Page settings.
func (queryParams *QueryParams) Paginate() GormQueryBuilder {
	return func(dbQuery *gorm.DB) *gorm.DB {
		offset := (queryParams.Page.Number - 1) * queryParams.Page.Limit
		return dbQuery.Offset(offset).Limit(queryParams.Page.Limit)
	}
}

// FilterByParams returns a scope for querying based on filter parameters.
func (queryParams *QueryParams) FilterByParams(tableName string) GormQueryBuilder {
	return func(dbQuery *gorm.DB) *gorm.DB {
		for _, filterParam := range queryParams.FilterParams {
			attribute := filterParam.Attribute
			if !IsEmpty(tableName) {
				attribute = tableName + "." + filterParam.Attribute
			}
			action := filterParam.Action
			value := filterParam.Value

			switch action {
			case Equals:
				whereQuery := fmt.Sprintf("%s = ?", attribute)
				dbQuery = dbQuery.Where(whereQuery, value)
			case Like:
				whereQuery := fmt.Sprintf("%s LIKE ?", attribute)
				dbQuery = dbQuery.Where(whereQuery, "%"+value+"%")
			case In:
				whereQuery := fmt.Sprintf("%s IN (?)", attribute)
				queryArray := strings.Split(value, ",")
				dbQuery = dbQuery.Where(whereQuery, queryArray)
			case GreaterThan:
				whereQuery := fmt.Sprintf("%s > ?", attribute)
				dbQuery = dbQuery.Where(whereQuery, value)
			case GreaterThanEqual:
				whereQuery := fmt.Sprintf("%s >= ?", attribute)
				dbQuery = dbQuery.Where(whereQuery, value)
			case LessThan:
				whereQuery := fmt.Sprintf("%s < ?", attribute)
				dbQuery = dbQuery.Where(whereQuery, value)
			case LessThanEqual:
				whereQuery := fmt.Sprintf("%s <= ?", attribute)
				dbQuery = dbQuery.Where(whereQuery, value)
			}
		}
		return dbQuery
	}
}

// QueryKeyType represents query parameter keys.
type QueryKeyType string

const (
	PageLimit     QueryKeyType = "limit"
	PageNumber    QueryKeyType = "page"
	SortBy        QueryKeyType = "sort_by"
	SortDirection QueryKeyType = "sort_direction"
)

// SetQueryParams parses query parameters and constructs a QueryParams object.
func SetQueryParams(queryParams map[string][]string) *QueryParams {
	page := DefaultPage()
	sort := DefaultSort()
	var filterParams []FilterParam
	for key, values := range queryParams {
		queryValue := values[len(values)-1]
		switch QueryKeyType(key) {
		case PageLimit:
			page.Limit, _ = strconv.Atoi(queryValue)
		case PageNumber:
			page.Number, _ = strconv.Atoi(queryValue)
		case SortBy:
			sort.By = queryValue
		case SortDirection:
			sort.Direction = SortDirectionType(queryValue)
		}
		// Check if the query parameter key contains a dot.
		if strings.Contains(key, ".") {
			// Split the query parameter key by dot.
			searchKeys := strings.Split(key, ".")
			// Create a FilterParam object.
			filterParam := FilterParam{Attribute: searchKeys[0], Action: FilterAction(searchKeys[1]), Value: queryValue}
			// Add the FilterParam object to filterParams.
			filterParams = append(filterParams, filterParam)
		}
	}
	return &QueryParams{
		Page:         page,
		Sort:         sort,
		FilterParams: filterParams,
	}
}
