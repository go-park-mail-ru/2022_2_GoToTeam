package api

import (
	"2022_2_GoTo_team/internal/serverRestAPI/api/models"
	"2022_2_GoTo_team/internal/serverRestAPI/storage"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"net/mail"
	"os"
	"strconv"
	"unicode"
)

const ARTICLE_NUMBER_IN_FEED = 10

type Api struct {
	usersStorage    *storage.UsersStorage
	sessionsStorage *storage.SessionsStorage
	feedStorage     *storage.FeedStorage
	logger          *logrus.Logger
}

func GetApi() *Api {
	authApi := &Api{
		usersStorage:    storage.GetUsersStorage(),
		feedStorage:     storage.GetFeedStorage(),
		sessionsStorage: storage.GetSessionsStorage(),
		logger:          logrus.New(),
	}
	authApi.usersStorage.PrintUsers()
	authApi.feedStorage.PrintArticles()
	authApi.sessionsStorage.PrintSessions()

	return authApi
}

func (api *Api) ConfigureLogger(logLVL, logPath string) error {
	level, err := logrus.ParseLevel(logLVL)
	if err != nil {
		return err
	}
	api.logger.SetLevel(level)
	if len(logPath) != 0 {
		logfile, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		api.logger.SetOutput(logfile)
	}
	return nil
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

func (api *Api) isAuthorized(c echo.Context) bool {
	authorized := false
	if session, err := c.Cookie(api.sessionsStorage.GetSessionHeaderName()); err == nil && session != nil {
		authorized = api.sessionsStorage.SessionExists(session.Value)
	}

	return authorized
}

func (api *Api) SignupUserHandler(c echo.Context) error {
	defer c.Request().Body.Close()

	parsedInput := new(models.User)
	if err := c.Bind(parsedInput); err != nil {
		c.Logger().Printf("Error: %s", err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	log.Println("Parsed input user data:", parsedInput)

	// TODO VALIDATOR

	if api.usersStorage.UserIsExistByLogin(parsedInput.NewUserData.Login) || api.usersStorage.UserIsExistByEmail(parsedInput.NewUserData.Email) {
		c.Logger().Printf("Error: %s", "user with this login or email exist")
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
		c.Logger().Printf("Error: %s", err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	cookie := api.sessionsStorage.CreateSessionForUser(parsedInput.NewUserData.Email)
	c.SetCookie(cookie)
	api.sessionsStorage.PrintSessions()

	return c.NoContent(http.StatusOK)
}

func (api *Api) CreateSessionHandler(c echo.Context) error {
	defer c.Request().Body.Close()

	parsedInput := new(models.SessionCreate)
	if err := c.Bind(parsedInput); err != nil {
		c.Logger().Printf("Error: %s", err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	log.Println("parsedInput = ", parsedInput)

	email := parsedInput.UserData.Email
	password := parsedInput.UserData.Password
	log.Println("URL", c.Request().URL)
	log.Println("email", email)
	log.Println("password ", password)

	user, err := api.usersStorage.GetUserByEmail(email)
	if err != nil {
		c.Logger().Printf("Error: %s", err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	if user.Password != password {
		c.Logger().Printf("Error: %s", "invalid password.")
		return c.NoContent(http.StatusBadRequest)
	}

	cookie := api.sessionsStorage.CreateSessionForUser(user.Email)
	c.SetCookie(cookie)
	api.sessionsStorage.PrintSessions()

	return c.NoContent(http.StatusOK)
}

func (api *Api) RemoveSessionHandler(c echo.Context) error {
	if !api.isAuthorized(c) {
		c.Logger().Printf("Error: %s", "unauthorized")
		return c.NoContent(http.StatusUnauthorized)
	}
	cookie, err := c.Cookie(api.sessionsStorage.GetSessionHeaderName())
	if err != nil {
		c.Logger().Printf("Error: %s", err.Error())
		return c.NoContent(http.StatusUnauthorized)
	}

	api.sessionsStorage.RemoveSession(cookie)
	api.sessionsStorage.PrintSessions()
	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}

func (api *Api) SessionInfoHandler(c echo.Context) error {
	if !api.isAuthorized(c) {
		c.Logger().Printf("Error: %s", "unauthorized")
		return c.NoContent(http.StatusUnauthorized)
	}
	cookie, err := c.Cookie(api.sessionsStorage.GetSessionHeaderName())
	if err != nil {
		c.Logger().Printf("Error: %s", err.Error())
		return c.NoContent(http.StatusUnauthorized)
	}

	email := api.sessionsStorage.GetEmailByCookie(cookie)
	user, err := api.usersStorage.GetUserByEmail(email)

	sessionInfo := models.SessionInfo{}
	sessionInfo.Username = user.Username

	log.Println("Formed sessionInfo = ", sessionInfo)

	return c.JSON(http.StatusOK, sessionInfo)
}

func (api *Api) FeedHandler(c echo.Context) error {
	startFromArticleOfNumberStr := c.QueryParam("startFromArticleOfNumber")
	if startFromArticleOfNumberStr == "" {
		startFromArticleOfNumberStr = "0"
	}

	startFromArticleOfNumber, err := strconv.Atoi(startFromArticleOfNumberStr)
	if err != nil {
		c.Logger().Printf("Error: %s", err.Error())
		return c.NoContent(http.StatusBadRequest)
	}
	if startFromArticleOfNumber < 0 {
		c.Logger().Printf("Error: startFromArticleOfNumber = %d < 0", startFromArticleOfNumber)
		return c.NoContent(http.StatusBadRequest)
	}

	articles, err := api.feedStorage.GetArticles()
	if err != nil {
		c.Logger().Printf("Error: %s", err.Error())
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

	feed := models.Feed{}
	for _, v := range articles {
		article := models.Article{
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

	log.Println("Formed feed = ", feed)

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
