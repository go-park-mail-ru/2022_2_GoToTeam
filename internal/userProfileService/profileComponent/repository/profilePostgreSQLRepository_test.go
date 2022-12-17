package repository

import (
	"2022_2_GoTo_team/internal/userProfileService/domain/models"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sirupsen/logrus"
	"reflect"
	"testing"
)

var loggerMock = &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
	"requestId": "qwerty",
	"userEmail": "asd@asd.asd",
})}

func TestGetProfileByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewProfilePostgreSQLRepository(db, loggerMock)

	email := "asd@asd.asd"

	// good query
	rows := sqlmock.NewRows([]string{"email", "login", "username", "avatar_img_path"})
	items := []*models.Profile{
		{Email: email, Login: "asd", Username: "asd_name", AvatarImgPath: "/asd.jpeg"},
	}
	for _, item := range items {
		rows = rows.AddRow(item.Email, item.Login, item.Username, item.AvatarImgPath)
	}

	mock.
		ExpectQuery("SELECT email, login").
		WithArgs(email).
		WillReturnRows(rows)

	res, err := repo.GetProfileByEmail(context.Background(), email)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(res, items[0]) {
		t.Errorf("results not match, want %v, have %v", items[0], res)
		return
	}

	// query error
	mock.
		ExpectQuery("SELECT email, login").
		WithArgs(email).
		WillReturnError(fmt.Errorf("db_error"))

	res, err = repo.GetProfileByEmail(context.Background(), email)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if res != nil {
		t.Errorf("results not match, want %v, have %v", nil, res)
		return
	}

	// ErrNoRows
	mock.
		ExpectQuery("SELECT email, login").
		WithArgs(email).
		WillReturnError(sql.ErrNoRows)

	res, err = repo.GetProfileByEmail(context.Background(), email)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if res != nil {
		t.Errorf("results not match, want %v, have %v", nil, res)
		return
	}
}

func TestUserExistsByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &profilePostgreSQLRepository{
		database: db,
		logger:   loggerMock,
	}

	email := "asd@asd.asd"

	// good query true
	rows := sqlmock.NewRows([]string{"email"})
	items := []string{
		"asd@asd.asd",
	}
	for _, item := range items {
		rows = rows.AddRow(item)
	}

	mock.
		ExpectQuery("SELECT U.email FROM users U WHERE U.email =").
		WithArgs(email).
		WillReturnRows(rows)

	res, err := repo.userExistsByEmail(context.Background(), email)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(res, true) {
		t.Errorf("results not match, want %v, have %v", true, res)
		return
	}

	// query error
	mock.
		ExpectQuery("SELECT U.email FROM users U WHERE U.email =").
		WithArgs(email).
		WillReturnError(fmt.Errorf("db_error"))

	res, err = repo.userExistsByEmail(context.Background(), email)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if !reflect.DeepEqual(res, true) {
		t.Errorf("results not match, want %v, have %v", true, res)
		return
	}

	// ErrNoRows
	mock.
		ExpectQuery("SELECT U.email FROM users U WHERE U.email =").
		WithArgs(email).
		WillReturnError(sql.ErrNoRows)

	res, err = repo.userExistsByEmail(context.Background(), email)
	if err != nil {
		t.Errorf("expected error, got nil")
		return
	}
	if !reflect.DeepEqual(res, false) {
		t.Errorf("results not match, want %v, have %v", false, res)
		return
	}
}

func TestUserExistsByLoginWithIgnoringRowsWithEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &profilePostgreSQLRepository{
		database: db,
		logger:   loggerMock,
	}

	login := "asd"
	emailToIgnore := "asd@asd.asd"

	// good query true
	rows := sqlmock.NewRows([]string{"email"})
	items := []string{
		login,
	}
	for _, item := range items {
		rows = rows.AddRow(item)
	}

	mock.
		ExpectQuery("SELECT U.login FROM users U WHERE U.login =").
		WithArgs(login, emailToIgnore).
		WillReturnRows(rows)

	res, err := repo.userExistsByLoginWithIgnoringRowsWithEmail(context.Background(), login, emailToIgnore)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(res, true) {
		t.Errorf("results not match, want %v, have %v", true, res)
		return
	}

	// query error
	mock.
		ExpectQuery("SELECT U.login FROM users U WHERE U.login =").
		WithArgs(login, emailToIgnore).
		WillReturnError(fmt.Errorf("db_error"))

	res, err = repo.userExistsByLoginWithIgnoringRowsWithEmail(context.Background(), login, emailToIgnore)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if !reflect.DeepEqual(res, true) {
		t.Errorf("results not match, want %v, have %v", true, res)
		return
	}

	// ErrNoRows
	mock.
		ExpectQuery("SELECT U.login FROM users U WHERE U.login =").
		WithArgs(login, emailToIgnore).
		WillReturnError(sql.ErrNoRows)

	res, err = repo.userExistsByLoginWithIgnoringRowsWithEmail(context.Background(), login, emailToIgnore)
	if err != nil {
		t.Errorf("expected error, got nil")
		return
	}
	if !reflect.DeepEqual(res, false) {
		t.Errorf("results not match, want %v, have %v", false, res)
		return
	}

}
