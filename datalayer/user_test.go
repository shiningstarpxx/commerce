package datalayer

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestUserDatalayer_UpdateUserEmailByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.Nilf(t, err, "an error '%s' was not expected when opening a stub database connection", err)
	defer db.Close()

	gdb, err := gorm.Open(mysql.New(mysql.Config{SkipInitializeWithVersion: true, Conn: db}), &gorm.Config{})
	assert.Nilf(t, err, "an error '%s' was not expected when opening a stub gorm connection", err)

	userID := uint(1)
	updateNewEmail := "456"
	timestamp := sqlmock.AnyArg()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `users`").
		WithArgs(updateNewEmail, timestamp, userID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	userInterface := &UserDatalayer{gdb}
	err = userInterface.UpdateUserEmailByID(userID, updateNewEmail)

	assert.Nilf(t, err, "an error '%s' was not expected when updating user email", err)

	assert.Nilf(t, mock.ExpectationsWereMet(), "there were unfulfilled expectations: %s", mock.ExpectationsWereMet())
}
