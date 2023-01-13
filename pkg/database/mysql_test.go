package database

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConnectDB_Success(t *testing.T) {
	db, err := ConnectSQLDB()

	require.Nil(t, err)
	require.NotNil(t, db)
}
