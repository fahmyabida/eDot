package repository

import (
	"fmt"
	"strings"

	pkgErrors "github.com/fahmyabida/eDot/pkg/errors"
	"gorm.io/gorm"
)

// ApplySortQuery will applying query order based on sortByQueryParams
func ApplySortByQuery(db *gorm.DB, sortByQueryParams string) (*gorm.DB, error) {
	sortBys := strings.Split(sortByQueryParams, ",")
	for _, sortBy := range sortBys {
		queryOrder, err := DetermineOrdering(sortBy)
		if err != nil {
			err = pkgErrors.InvalidColumnError(err.Error())
			return db, err
		}
		db = db.Order(queryOrder)
	}
	return db, nil
}

// DetermineOrdering will given query order
func DetermineOrdering(querySort string) (string, error) {
	if strings.ContainsAny(querySort, ";") ||
		strings.TrimSpace(querySort) == "" {
		return "", pkgErrors.InvalidColumnError(fmt.Sprintf("%v '%v'", pkgErrors.ErrColumnInvalid, querySort))
	}
	if isAscendingOrder(querySort) {
		return strings.TrimPrefix(querySort, "-") + " ASC", nil
	}
	return querySort + " DESC", nil
}

// isAscendingOrder will check has (-) for ASCending order
func isAscendingOrder(s string) bool {
	return strings.HasPrefix(s, "-")
}
