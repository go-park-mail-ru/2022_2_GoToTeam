package repository

import (
	"2022_2_GoTo_team/internal/authSessionService/domain/customErrors/userComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/authSessionService/domain/models"
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

func TestGetAllUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewUserPostgreSQLRepository(db, loggerMock)

	// good query
	rows := sqlmock.NewRows([]string{"user_id", "email", "login", "password", "username", "sex", "date_of_birth", "avatar_img_path", "registration_date", "subscribers_count", "subscriptions_count"})
	items := []*models.User{
		&models.User{
			UserId:             1,
			Email:              "asd@asd.asd",
			Login:              "asd",
			Password:           "asdP",
			Username:           "asdU",
			Sex:                "F",
			DateOfBirth:        "dob",
			AvatarImgPath:      "/asd.jpeg",
			RegistrationDate:   "dob",
			SubscribersCount:   1,
			SubscriptionsCount: 2,
		},
	}
	for _, user := range items {
		rows = rows.AddRow(user.UserId, user.Email, user.Login, user.Password, user.Username, user.Sex, user.DateOfBirth, user.AvatarImgPath, user.RegistrationDate, user.SubscribersCount, user.SubscriptionsCount)
	}

	mock.
		ExpectQuery("SELECT U.user_id, U.email, U.login, U.password,").
		WithArgs().
		WillReturnRows(rows)

	res, err := repo.GetAllUsers(context.Background())
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(res, items) {
		t.Errorf("results not match, want %v, have %#v", items, res)
		return
	}

	// query error
	mock.
		ExpectQuery("SELECT U.user_id, U.email, U.login, U.password,").
		WithArgs().
		WillReturnError(fmt.Errorf("db_error"))

	res, err = repo.GetAllUsers(context.Background())
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err != repositoryToUsecaseErrors.UserRepositoryError {
		t.Errorf("expected error repositoryToUsecaseErrors.UserRepositoryErro")
		return
	}
	if res != nil {
		t.Errorf("results not match, want %v, have %v", nil, res)
		return
	}

	// row scan error
	rows = sqlmock.NewRows([]string{"user_id"}).AddRow(1)
	mock.
		ExpectQuery("SELECT U.user_id, U.email, U.login, U.password,").
		WithArgs().
		WillReturnRows(rows)

	res, err = repo.GetAllUsers(context.Background())
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err != repositoryToUsecaseErrors.UserRepositoryError {
		t.Errorf("expected error repositoryToUsecaseErrors.UserRepositoryErro")
		return
	}
	if res != nil {
		t.Errorf("results not match, want %v, have %v", nil, res)
		return
	}
}

func TestCheckUserEmailAndPassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewUserPostgreSQLRepository(db, loggerMock)

	email := "asd@asd.asd"
	password := "pass"

	// good query true
	rows := sqlmock.NewRows([]string{"email"}).AddRow(email)

	mock.
		ExpectQuery("SELECT U.email").
		WithArgs(email, password).
		WillReturnRows(rows)

	res, err := repo.CheckUserEmailAndPassword(context.Background(), email, password)
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
		ExpectQuery("SELECT U.email").
		WithArgs(email, password).
		WillReturnError(fmt.Errorf("db_error"))

	res, err = repo.CheckUserEmailAndPassword(context.Background(), email, password)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if !reflect.DeepEqual(res, false) {
		t.Errorf("results not match, want %v, have %v", false, res)
		return
	}

	// ErrNoRows
	mock.
		ExpectQuery("SELECT U.email").
		WithArgs(email, password).
		WillReturnError(sql.ErrNoRows)

	res, err = repo.CheckUserEmailAndPassword(context.Background(), email, password)
	if err != nil {
		t.Errorf("expected error, got nil")
		return
	}
	if !reflect.DeepEqual(res, false) {
		t.Errorf("results not match, want %v, have %v", false, res)
		return
	}
}

func TestGetUserInfoForSessionComponentByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewUserPostgreSQLRepository(db, loggerMock)

	// good query
	email := "asd@asd.asd"
	rows := sqlmock.NewRows([]string{"username", "login", "avatar_img_path"})
	items := []*models.User{
		&models.User{
			Login:         "asd",
			Username:      "asdU",
			AvatarImgPath: "/asd.jpeg",
		},
	}
	for _, user := range items {
		rows = rows.AddRow(user.Username, user.Login, user.AvatarImgPath)
	}

	mock.
		ExpectQuery("SELECT").
		WithArgs().
		WillReturnRows(rows)

	res, err := repo.GetUserInfoForSessionComponentByEmail(context.Background(), email)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(res, items[0]) {
		t.Errorf("results not match, want %v, have %#v", items[0], res)
		return
	}

	// query error
	mock.
		ExpectQuery("SELECT").
		WithArgs().
		WillReturnError(fmt.Errorf("db_error"))

	res, err = repo.GetUserInfoForSessionComponentByEmail(context.Background(), email)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err != repositoryToUsecaseErrors.UserRepositoryError {
		t.Errorf("expected error repositoryToUsecaseErrors.UserRepositoryErro")
		return
	}
	if res != nil {
		t.Errorf("results not match, want %v, have %v", nil, res)
		return
	}

	// row scan error
	mock.
		ExpectQuery("SELECT").
		WithArgs().
		WillReturnError(sql.ErrNoRows)

	res, err = repo.GetUserInfoForSessionComponentByEmail(context.Background(), email)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err != repositoryToUsecaseErrors.UserRepositoryEmailDoesntExistError {
		t.Errorf("expected error repositoryToUsecaseErrors.UserRepositoryEmailDoesntExistError")
		return
	}
	if res != nil {
		t.Errorf("results not match, want %v, have %v", nil, res)
		return
	}
}
