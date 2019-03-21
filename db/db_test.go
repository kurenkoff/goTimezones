package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	bd := New()
	defer bd.db.Close()
	err := bd.db.Ping()
	assert.Equal(t, err,nil, "Error is not nil")
}

func TestDatabase_Initialize(t *testing.T) {
	bd := New()
	bd.Initialize([]string{})
	defer bd.db.Close()
	//possible := []error{nil, sql.ErrNoRows}
	_, err := bd.db.Exec("select * from users")
	assert.NoError(t, err)
	_, err = bd.db.Exec("select * from timezones")
	assert.NoError(t, err)
	_, err = bd.db.Exec("select * from user_data")
	assert.NoError(t, err)
}
