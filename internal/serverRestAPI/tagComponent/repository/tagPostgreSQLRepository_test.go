package repository

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/tagComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
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

func TestGetAllTags(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewTagPostgreSQLRepository(db, loggerMock)

	// good query
	rows := sqlmock.NewRows([]string{"tag_name"})
	items := []*models.Tag{
		&models.Tag{
			TagName: "java",
		},
	}
	for _, item := range items {
		rows = rows.AddRow(item.TagName)
	}

	mock.
		ExpectQuery("SELECT tag_name FROM tags").
		WithArgs().
		WillReturnRows(rows)

	res, err := repo.GetAllTags(context.Background())
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
		ExpectQuery("SELECT tag_name FROM tags").
		WithArgs().
		WillReturnError(fmt.Errorf("db_error"))

	res, err = repo.GetAllTags(context.Background())
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err != repositoryToUsecaseErrors.TagRepositoryError {
		t.Errorf("expected error repositoryToUsecaseErrors.UserRepositoryErro")
		return
	}
	if res != nil {
		t.Errorf("results not match, want %v, have %v", nil, res)
		return
	}

	// row scan error
	rows = sqlmock.NewRows([]string{}).AddRow()
	mock.
		ExpectQuery("SELECT tag_name FROM tags").
		WithArgs().
		WillReturnRows(rows)

	res, err = repo.GetAllTags(context.Background())
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err != repositoryToUsecaseErrors.TagRepositoryError {
		t.Errorf("expected error repositoryToUsecaseErrors.UserRepositoryErro")
		return
	}
	if res != nil {
		t.Errorf("results not match, want %v, have %v", nil, res)
		return
	}

}
