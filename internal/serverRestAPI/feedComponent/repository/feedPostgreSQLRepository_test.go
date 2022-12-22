package repository

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/feedComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"database/sql"
	"github.com/stretchr/testify/assert"

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

func TestGetArticlesString(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &feedPostgreSQLRepository{
		database: db,
		logger:   loggerMock,
	}

	res := repo.getArticlesString([]*models.Article{&models.Article{}})

	assert.NotEqual(t, "", res)
}

func TestGetTagsForArticle(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &feedPostgreSQLRepository{
		database: db,
		logger:   loggerMock,
	}

	// good query
	rows := sqlmock.NewRows([]string{"tag_name"})
	items := []string{
		"java",
	}
	for _, item := range items {
		rows = rows.AddRow(item)
	}

	mock.
		ExpectQuery("SELECT T.tag_name").
		WithArgs().
		WillReturnRows(rows)

	res, err := repo.GetTagsForArticle(context.Background(), 1)
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
		ExpectQuery("SELECT T.tag_name").
		WithArgs().
		WillReturnError(fmt.Errorf("db_error"))

	res, err = repo.GetTagsForArticle(context.Background(), 1)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err != repositoryToUsecaseErrors.FeedRepositoryError {
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
		ExpectQuery("SELECT T.tag_name").
		WithArgs().
		WillReturnRows(rows)

	res, err = repo.GetTagsForArticle(context.Background(), 1)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err != repositoryToUsecaseErrors.FeedRepositoryError {
		t.Errorf("expected error repositoryToUsecaseErrors.UserRepositoryErro")
		return
	}
	if res != nil {
		t.Errorf("results not match, want %v, have %v", nil, res)
		return
	}
}

func TestGetFeed(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewFeedPostgreSQLRepository(db, loggerMock)

	// good query
	rows := sqlmock.NewRows([]string{"article_id", "title", "description", "rating", "comments_count", "content", "cover_img_path", "UCusername", "UClogin", "UPusername", "UPlogin", "category_name", "is_like"})

	items := []*models.Article{
		&models.Article{
			ArticleId:     1,
			Title:         "abracodabra",
			Description:   "descr_codabr",
			Rating:        15,
			CommentsCount: 1,
			Content:       "super content",
			CoverImgPath:  "",

			CoAuthor: models.CoAuthor{
				Email:    "asd@asd.asd",
				Username: "asd",
			},
			Publisher: models.Publisher{
				Email:    "asd@asd.asd",
				Username: "asd",
			},
			CategoryName: "java",
			Liked:        0,
		},
	}

	for _, item := range items {
		rows = rows.AddRow(item.ArticleId, item.Title, item.Description, item.Rating, item.CommentsCount, item.Content, item.CoverImgPath, item.CoAuthor.Username, item.CoAuthor.Login, item.Publisher.Username, item.Publisher.Login, item.CategoryName, item.Liked)
	}

	mock.
		ExpectQuery("SELECT A.article_id, A.title,").
		WithArgs().
		WillReturnRows(rows)

	res, err := repo.GetFeed(context.Background(), "asd@asd.asd")
	if err == nil {
		t.Errorf("expected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	// query error

	mock.
		ExpectQuery("SELECT A.article_id, A.title,").
		WithArgs().
		WillReturnError(fmt.Errorf("db_error"))

	res, err = repo.GetFeed(context.Background(), "asd@asd.asd")
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

func TestGetFeedForUserByLogin(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewFeedPostgreSQLRepository(db, loggerMock)

	// good query
	rows := sqlmock.NewRows([]string{"article_id", "title", "description", "rating", "comments_count", "content", "cover_img_path", "UCusername", "UClogin", "UPusername", "UPlogin", "category_name", "is_like"})

	items := []*models.Article{
		&models.Article{
			ArticleId:     1,
			Title:         "abracodabra",
			Description:   "descr_codabr",
			Rating:        15,
			CommentsCount: 1,
			Content:       "super content",
			CoverImgPath:  "",

			CoAuthor: models.CoAuthor{
				Email:    "asd@asd.asd",
				Username: "asd",
			},
			Publisher: models.Publisher{
				Email:    "asd@asd.asd",
				Username: "asd",
			},
			CategoryName: "java",
			Liked:        0,
		},
	}

	for _, item := range items {
		rows = rows.AddRow(item.ArticleId, item.Title, item.Description, item.Rating, item.CommentsCount, item.Content, item.CoverImgPath, item.CoAuthor.Username, item.CoAuthor.Login, item.Publisher.Username, item.Publisher.Login, item.CategoryName, item.Liked)
	}

	mock.
		ExpectQuery("SELECT A.article_id, A.title,").
		WithArgs().
		WillReturnRows(rows)

	res, err := repo.GetFeedForUserByLogin(context.Background(), "asd", "asd@asd.asd")
	if err == nil {
		t.Errorf("expected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	// query error

	mock.
		ExpectQuery("SELECT A.article_id, A.title,").
		WithArgs().
		WillReturnError(fmt.Errorf("db_error"))

	res, err = repo.GetFeedForUserByLogin(context.Background(), "asd", "asd@asd.asd")
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

func TestGetFeedGetFeedForCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewFeedPostgreSQLRepository(db, loggerMock)

	// good query
	rows := sqlmock.NewRows([]string{"article_id", "title", "description", "rating", "comments_count", "content", "cover_img_path", "UCusername", "UClogin", "UPusername", "UPlogin", "category_name", "is_like"})

	items := []*models.Article{
		&models.Article{
			ArticleId:     1,
			Title:         "abracodabra",
			Description:   "descr_codabr",
			Rating:        15,
			CommentsCount: 1,
			Content:       "super content",
			CoverImgPath:  "",

			CoAuthor: models.CoAuthor{
				Email:    "asd@asd.asd",
				Username: "asd",
			},
			Publisher: models.Publisher{
				Email:    "asd@asd.asd",
				Username: "asd",
			},
			CategoryName: "java",
			Liked:        0,
		},
	}

	for _, item := range items {
		rows = rows.AddRow(item.ArticleId, item.Title, item.Description, item.Rating, item.CommentsCount, item.Content, item.CoverImgPath, item.CoAuthor.Username, item.CoAuthor.Login, item.Publisher.Username, item.Publisher.Login, item.CategoryName, item.Liked)
	}

	mock.
		ExpectQuery("SELECT A.article_id, A.title,").
		WithArgs().
		WillReturnRows(rows)

	res, err := repo.GetFeedForCategory(context.Background(), "asd", "asd@asd.asd")
	if err == nil {
		t.Errorf("expected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	// query error

	mock.
		ExpectQuery("SELECT A.article_id, A.title,").
		WithArgs().
		WillReturnError(fmt.Errorf("db_error"))

	res, err = repo.GetFeedForCategory(context.Background(), "asd", "asd@asd.asd")
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

func TestUserExistsByLogin(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewFeedPostgreSQLRepository(db, loggerMock)

	// good query
	rows := sqlmock.NewRows([]string{"login"})
	items := []string{
		"asd",
	}
	for _, item := range items {
		rows = rows.AddRow(item)
	}

	mock.
		ExpectQuery("SELECT U.login").
		WithArgs().
		WillReturnRows(rows)

	res, err := repo.UserExistsByLogin(context.Background(), "asd")
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(res, true) {
		t.Errorf("results not match, want %v, have %#v", true, res)
		return
	}

	// query error

	mock.
		ExpectQuery("SELECT U.login").
		WithArgs().
		WillReturnError(fmt.Errorf("db_error"))

	res, err = repo.UserExistsByLogin(context.Background(), "asd")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err != repositoryToUsecaseErrors.FeedRepositoryError {
		t.Errorf("expected error repositoryToUsecaseErrors.UserRepositoryErro")
		return
	}
	if res != true {
		t.Errorf("results not match, want %v, have %v", true, res)
		return
	}

	// empty row error

	mock.
		ExpectQuery("SELECT U.login").
		WithArgs().
		WillReturnError(sql.ErrNoRows)

	res, err = repo.UserExistsByLogin(context.Background(), "asd")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err != nil {
		t.Errorf("expected error, got nil")
		return
	}

	if res != false {
		t.Errorf("results not match, want %v, have %v", false, res)
		return
	}
}

func TestCategoryExists(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewFeedPostgreSQLRepository(db, loggerMock)

	// good query
	rows := sqlmock.NewRows([]string{"category_name"})
	items := []string{
		"asd",
	}
	for _, item := range items {
		rows = rows.AddRow(item)
	}

	mock.
		ExpectQuery("SELECT C.category_name").
		WithArgs().
		WillReturnRows(rows)

	res, err := repo.CategoryExists(context.Background(), "asd")
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(res, true) {
		t.Errorf("results not match, want %v, have %#v", true, res)
		return
	}

	// query error

	mock.
		ExpectQuery("SELECT C.category_name").
		WithArgs().
		WillReturnError(fmt.Errorf("db_error"))

	res, err = repo.CategoryExists(context.Background(), "asd")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err != repositoryToUsecaseErrors.FeedRepositoryError {
		t.Errorf("expected error repositoryToUsecaseErrors.UserRepositoryErro")
		return
	}
	if res != true {
		t.Errorf("results not match, want %v, have %v", true, res)
		return
	}

	// empty row error

	mock.
		ExpectQuery("SELECT C.category_name").
		WithArgs().
		WillReturnError(sql.ErrNoRows)

	res, err = repo.CategoryExists(context.Background(), "asd")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err != nil {
		t.Errorf("expected error, got nil")
		return
	}
	if res != false {
		t.Errorf("results not match, want %v, have %v", false, res)
		return
	}
}

func TestGetNewArticlesFromIdForSubscriber(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewFeedPostgreSQLRepository(db, loggerMock)

	// good query
	rows := sqlmock.NewRows([]string{"article_id"})
	items := []int{
		1,
	}
	for _, item := range items {
		rows = rows.AddRow(item)
	}

	mock.
		ExpectQuery("SELECT A.article_id").
		WithArgs().
		WillReturnRows(rows)

	res, err := repo.GetNewArticlesFromIdForSubscriber(context.Background(), 1, "asd@asd.asd")
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
		ExpectQuery("SELECT A.article_id").
		WithArgs().
		WillReturnError(fmt.Errorf("db_error"))

	res, err = repo.GetNewArticlesFromIdForSubscriber(context.Background(), 1, "asd@asd.asd")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err != repositoryToUsecaseErrors.FeedRepositoryError {
		t.Errorf("expected error repositoryToUsecaseErrors.UserRepositoryErro")
		return
	}
	if res != nil {
		t.Errorf("results not match, want %v, have %v", true, res)
		return
	}
}
