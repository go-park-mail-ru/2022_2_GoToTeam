package api

import (
	models3 "2022_2_GoTo_team/internal/serverRestAPI/feedComponent/delivery/models"
	"2022_2_GoTo_team/internal/serverRestAPI/repository"
	repository2 "2022_2_GoTo_team/internal/serverRestAPI/sessionComponent/repository"
	models4 "2022_2_GoTo_team/internal/serverRestAPI/userComponent/delivery/models"
	repository3 "2022_2_GoTo_team/internal/serverRestAPI/userComponent/repository"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"net/mail"
	"strconv"
	"unicode"
)

const ARTICLE_NUMBER_IN_FEED = 10

type Api struct {
	usersStorage    *repository3.UsersStorage
	sessionsStorage *repository2.SessionsStorage
	feedStorage     *repository.FeedStorage
	logger          *logrus.Logger
}

func GetApi() *Api {
	authApi := &Api{
		usersStorage:    repository3.GetUsersStorage(),
		feedStorage:     repository.GetFeedStorage(),
		sessionsStorage: repository2.NewSessionsRepository(),
		logger:          logrus.New(),
	}
	authApi.usersStorage.PrintUsers()
	authApi.feedStorage.PrintArticles()
	authApi.sessionsStorage.PrintSessions()

	return authApi
}

func (api *Api) LogInfo(info string) {
	api.logger.Info(info)
}

func emailIsValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func loginIsValid(login string) bool {
	if len(login) < 8 {
		return false
	}
	for _, sep := range login {
		if !unicode.IsLetter(sep) || sep != '_' {
			return false
		}
	}
	return true
}

// at least 8 symbols
// at least 1 upper symbol
// at least 1 special symbol
func passwordIsValid(password string) (eightOrMore, upper, special bool) {
	letters := 0
	for _, c := range password {
		switch {
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
			//return false, false, false
		}
	}
	eightOrMore = letters >= 8
	return
}

func usernameIsValid(uname string) bool {
	if len(uname) == 0 {
		return false
	}
	en := unicode.Is(unicode.Latin, rune(uname[0]))

	for _, sep := range uname {
		if (en && !unicode.Is(unicode.Latin, sep)) || (!en && unicode.Is(unicode.Latin, sep)) {
			return false
		}
	}
	return true
}

func (api *Api) SignupUserHandler(c echo.Context) error {
	defer c.Request().Body.Close()

	parsedInput := new(models4.User)
	if err := c.Bind(parsedInput); err != nil {
		c.Logger().Printf("Error: %s", err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	log.Println("Parsed input user data:", parsedInput)

	// TODO VALIDATOR
	if !loginIsValid(parsedInput.NewUserData.Login) {
		api.logger.Error("Incorrect login")
		return c.NoContent(http.StatusBadRequest)
	}

	if eightOrMore, upper, special := passwordIsValid(parsedInput.NewUserData.Password); !(eightOrMore && upper && special) {
		api.logger.Error("Password is incorrect")
		return c.NoContent(http.StatusConflict)
	}

	if !emailIsValid(parsedInput.NewUserData.Email) {
		api.logger.Error("Email is incorrect")
		return c.NoContent(http.StatusConflict)
	}

	if !usernameIsValid(parsedInput.NewUserData.Username) {
		api.logger.Error("Username is incorrect")
		return c.NoContent(http.StatusConflict)
	}

	if api.usersStorage.UserIsExistByLogin(parsedInput.NewUserData.Login) || api.usersStorage.UserIsExistByEmail(parsedInput.NewUserData.Email) {
		//c.LogrusLogger().Printf("Error: %s", "user with this login or email exist")
		api.logger.Error("User with this login or email exist")
		return c.NoContent(http.StatusConflict)
	}

	if err := api.usersStorage.AddUser(
		api.usersStorage.CreateUserInstanceFromData(
			parsedInput.NewUserData.Username,
			parsedInput.NewUserData.Email,
			parsedInput.NewUserData.Login,
			parsedInput.NewUserData.Password,
		),
	); err != nil {
		//c.LogrusLogger().Printf("Error: %s", err.Error())
		api.logger.Error(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	cookie := api.sessionsStorage.CreateSessionForUser(parsedInput.NewUserData.Email)
	c.SetCookie(cookie)
	api.sessionsStorage.PrintSessions()

	api.logger.Info("User register successful!")
	return c.NoContent(http.StatusOK)
}

func (api *Api) FeedHandler(c echo.Context) error {
	startFromArticleOfNumberStr := c.QueryParam("startFromArticleOfNumber")
	if startFromArticleOfNumberStr == "" {
		startFromArticleOfNumberStr = "0"
	}

	startFromArticleOfNumber, err := strconv.Atoi(startFromArticleOfNumberStr)
	if err != nil {
		//c.LogrusLogger().Printf("Error: %s", err.Error())
		api.logger.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}
	if startFromArticleOfNumber < 0 {
		//c.LogrusLogger().Printf("Error: startFromArticleOfNumber = %d < 0", startFromArticleOfNumber)
		api.logger.Error("startFromArticleOfNumber = ", startFromArticleOfNumber, " < 0")
		return c.NoContent(http.StatusBadRequest)
	}

	articles, err := api.feedStorage.GetArticles()
	if err != nil {
		//c.LogrusLogger().Printf("Error: %s", err.Error())
		api.logger.Error(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	if startFromArticleOfNumber+ARTICLE_NUMBER_IN_FEED <= len(articles) {
		articles = articles[startFromArticleOfNumber : startFromArticleOfNumber+ARTICLE_NUMBER_IN_FEED]
	} else if startFromArticleOfNumber < len(articles) {
		articles = articles[startFromArticleOfNumber:]
	} else {
		var startTmp = len(articles) - ARTICLE_NUMBER_IN_FEED
		if startTmp < 0 {
			startTmp = 0
		}
		articles = articles[startTmp:]
	}

	feed := models3.Feed{}
	for _, v := range articles {
		article := models3.Article{
			Id:          v.Id,
			Title:       v.Title,
			Description: v.Description,
			Tags:        v.Tags,
			Category:    v.Category,
			Rating:      v.Rating,
			Authors:     v.Authors,
			Content:     v.Content,
		}
		feed.Articles = append(feed.Articles, article)
	}

	//log.Println("Formed feed = ", feed)

	api.logger.Info("Formed feed = ", feed)
	return c.JSON(http.StatusOK, feed)
}

func (api *Api) CreateArticleHandler(c echo.Context) error {
	return nil
}

func (api *Api) UpdateArticleHandler(c echo.Context) error {
	return nil
}

func (api *Api) UserInfoHandler(c echo.Context) error {
	return nil
}

func (api *Api) UserFeedHandler(c echo.Context) error {
	return nil
}

func (api *Api) CategoryInfoHandler(c echo.Context) error {
	return nil
}

func (api *Api) CategoryFeedHandler(c echo.Context) error {
	return nil
}
