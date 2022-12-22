package repository

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/articleComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"database/sql"
	"reflect"

	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sirupsen/logrus"
	"testing"
)

var loggerMock = &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
	"requestId": "qwerty",
	"userEmail": "asd@asd.asd",
})}

func TestGetArticleById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewArticlePostgreSQLRepository(db, loggerMock)

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
		fmt.Printf("%#v", item)
		rows = rows.AddRow(item.ArticleId, item.Title, item.Description, item.Rating, item.CommentsCount, item.Content, item.CoverImgPath, item.CoAuthor.Username, item.CoAuthor.Login, item.Publisher.Username, item.Publisher.Login, item.CategoryName, item.Liked)
	}

	mock.
		ExpectQuery("SELECT A.article_id,").
		WithArgs().
		WillReturnRows(rows)

	_, err = repo.GetArticleById(context.Background(), 1, "asd@asd.asd")
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	// query error

	mock.
		ExpectQuery("SELECT A.article_id,").
		WithArgs().
		WillReturnError(fmt.Errorf("db_error"))

	res, err := repo.GetArticleById(context.Background(), 1, "asd@asd.asd")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err != repositoryToUsecaseErrors.ArticleRepositoryError {
		t.Errorf("expected error repositoryToUsecaseErrors.ArticleRepositoryErro")
		return
	}
	if res != nil {
		t.Errorf("results not match, want %v, have %v", nil, res)
		return
	}

	// empty row
	mock.
		ExpectQuery("SELECT A.article_id,").
		WithArgs().
		WillReturnError(sql.ErrNoRows)

	res, err = repo.GetArticleById(context.Background(), 1, "asd@asd.asd")
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

func TestAddArticle(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewArticlePostgreSQLRepository(db, loggerMock)

	// good query
	mock.
		ExpectExec("INSERT INTO articles (title, description, content, cover_img_path, co_author_id, publisher_id, category_id)  VALUES ").
		WithArgs(true, 3, "asd@asd.asd").
		WillReturnResult(sqlmock.NewResult(1, 1))

	res, err := repo.AddArticle(context.Background(), &models.Article{})
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if res != 0 {
		t.Errorf("bad id: want %v, have %v", 1, res)
		return
	}
}

func TestUpdateArticle(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewArticlePostgreSQLRepository(db, loggerMock)

	// good query
	mock.
		ExpectQuery("UPDATE articles SET title = ").
		WithArgs()

	err = repo.UpdateArticle(context.Background(), &models.Article{})
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestDeleteArticleById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewArticlePostgreSQLRepository(db, loggerMock)

	// good query
	mock.
		ExpectExec("DELETE FROM articles WHERE article_id = ").
		WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.DeleteArticleById(context.Background(), 3)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	// query error
	mock.
		ExpectExec("DELETE FROM articles WHERE article_id = ").
		WithArgs().WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.DeleteArticleById(context.Background(), 3)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestGetAuthorEmailForArticle(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewArticlePostgreSQLRepository(db, loggerMock)

	// good query
	rows := sqlmock.NewRows([]string{"email"})

	items := []string{
		"asd@asd.asd",
	}

	for _, item := range items {
		fmt.Printf("%#v", item)
		rows = rows.AddRow(item)
	}

	mock.
		ExpectQuery("SELECT U.email FROM users U").
		WithArgs().
		WillReturnRows(rows)

	res, err := repo.GetAuthorEmailForArticle(context.Background(), 1)
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
		ExpectQuery("SELECT U.email FROM users U").
		WithArgs().
		WillReturnError(fmt.Errorf("db_error"))

	res, err = repo.GetAuthorEmailForArticle(context.Background(), 1)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err != repositoryToUsecaseErrors.ArticleRepositoryError {
		t.Errorf("expected error repositoryToUsecaseErrors.ArticleRepositoryErro")
		return
	}
	if res != "" {
		t.Errorf("results not match, want %v, have %v", "", res)
		return
	}

	// row scan error

	rows = sqlmock.NewRows([]string{}).AddRow()
	mock.
		ExpectQuery("SELECT U.email FROM users U").
		WithArgs().
		WillReturnRows(rows)

	res, err = repo.GetAuthorEmailForArticle(context.Background(), 1)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err != repositoryToUsecaseErrors.ArticleRepositoryError {
		t.Errorf("expected error repositoryToUsecaseErrors.ArticleRepositoryErro")
		return
	}
	if res != "" {
		t.Errorf("results not match, want %v, have %v", "", res)
		return
	}

}

func TestAddLike(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewArticlePostgreSQLRepository(db, loggerMock)

	// good query
	mock.
		ExpectExec("INSERT INTO articles_likes (is_like, article_id, user_id)  VALUES ").
		WithArgs(true, 3, "asd@asd.asd").
		WillReturnResult(sqlmock.NewResult(1, 1))

	res, err := repo.AddLike(context.Background(), true, 3, "asd@asd.asd")
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if res != 0 {
		t.Errorf("bad id: want %v, have %v", 1, res)
		return
	}
}

func TestRemoveLike(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewArticlePostgreSQLRepository(db, loggerMock)

	// good query
	mock.
		ExpectExec("DELETE FROM articles_likes WHERE article_id = ").
		WithArgs()

	_, err = repo.RemoveLike(context.Background(), 3, "asd@asd.asd")
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestGetArticleRating(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewArticlePostgreSQLRepository(db, loggerMock)

	// good query
	rows := sqlmock.NewRows([]string{"rating"})
	items := []int{
		15,
	}
	for _, item := range items {
		rows = rows.AddRow(item)
	}

	mock.
		ExpectQuery("SELECT rating FROM articles").
		WithArgs().
		WillReturnRows(rows)

	res, err := repo.GetArticleRating(context.Background(), 1)
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

	// multiple rows error
	rows = sqlmock.NewRows([]string{"rating"})
	items = []int{
		15,
		20,
	}
	for _, item := range items {
		rows = rows.AddRow(item)
	}

	mock.
		ExpectQuery("SELECT rating FROM articles").
		WithArgs().
		WillReturnRows(rows)

	res, err = repo.GetArticleRating(context.Background(), 1)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err != repositoryToUsecaseErrors.ArticleRepositoryError {
		t.Errorf("expected error repositoryToUsecaseErrors.ArticleRepositoryErro")
		return
	}
	if res != 0 {
		t.Errorf("results not match, want %v, have %v", 0, res)
		return
	}

	// query error
	mock.
		ExpectQuery("SELECT rating FROM articles").
		WithArgs().
		WillReturnError(fmt.Errorf("db_error"))

	res, err = repo.GetArticleRating(context.Background(), 1)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err != repositoryToUsecaseErrors.ArticleRepositoryError {
		t.Errorf("expected error repositoryToUsecaseErrors.ArticleRepositoryErro")
		return
	}
	if res != 0 {
		t.Errorf("results not match, want %v, have %v", 0, res)
		return
	}

	// row scan error
	rows = sqlmock.NewRows([]string{}).AddRow()
	mock.
		ExpectQuery("SELECT rating FROM articles").
		WithArgs().
		WillReturnRows(rows)

	res, err = repo.GetArticleRating(context.Background(), 1)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err != repositoryToUsecaseErrors.ArticleRepositoryError {
		t.Errorf("expected error repositoryToUsecaseErrors.ArticleRepositoryErro")
		return
	}
	if res != 0 {
		t.Errorf("results not match, want %v, have %v", 0, res)
		return
	}
}
