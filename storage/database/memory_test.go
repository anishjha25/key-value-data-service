package database

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMemoryDB_Set(t *testing.T) {
	db := newMemoryDB()
	err := db.Set("mykey", "myvalue")
	require.NoError(t, err)
}

func TestMemoryDB_Get(t *testing.T) {
	db := newMemoryDB()
	err := db.Set("mykey", "myvalue")
	require.NoError(t, err)

	value, err := db.Get("mykey")
	require.NoError(t, err)
	require.Equal(t, "myvalue", value)

}

func TestMemoryDB_Delete(t *testing.T) {
	db := newMemoryDB()
	err := db.Set("mykey", "myvalue")
	require.NoError(t, err)

	value, err := db.Get("mykey")
	require.NoError(t, err)
	require.Equal(t, "myvalue", value)

	err = db.Delete("mykey")
	require.NoError(t, err)

	value, err = db.Get("mykey")
	require.Error(t, err)
	require.Equal(t, "not found", err.Error())

}
