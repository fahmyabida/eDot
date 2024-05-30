package repository_test

import (
	"testing"

	"github.com/fahmyabida/eDot/internal/app/repository"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestApplySortByQuery(t *testing.T) {
	db := gorm.DB{}

	sortBy := "status; DROP table products; --"
	_, err := repository.ApplySortByQuery(&db, sortBy)
	require.Error(t, err)
}

func TestDetermineOrdering(t *testing.T) {
	s := "-created_at"
	query, err := repository.DetermineOrdering(s)
	require.Equal(t, "created_at ASC", query)
	require.NoError(t, err)
	s = "created_at"
	query, err = repository.DetermineOrdering(s)
	require.Equal(t, "created_at DESC", query)
	require.NoError(t, err)
	s = "     "
	_, err = repository.DetermineOrdering(s)
	require.Error(t, err)
	s = ""
	_, err = repository.DetermineOrdering(s)
	require.Error(t, err)
	s = "created_at; DROP table products; --"
	_, err = repository.DetermineOrdering(s)
	require.Error(t, err)
}
