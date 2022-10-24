package api

import (
	"2022_2_GoTo_team/server/api/models"
	"2022_2_GoTo_team/server/storage"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

const ARTICLE_NUMBER_IN_FEED = 10

type Api struct {
	usersStorage    *storage.UsersStorage
	sessionsStorage *storage.SessionsStorage
	feedStorage     *storage.FeedStorage
}

func GetApi() *Api {
	authApi := &Api{
		usersStorage:    storage.GetUsersStorage(),
		feedStorage:     storage.GetFeedStorage(),
		sessionsStorage: storage.GetSessionsStorage(),
	}
	authApi.usersStorage.PrintUsers()
	authApi.feedStorage.PrintArticles()
	authApi.sessionsStorage.PrintSessions()

	return authApi
}

func (api *Api) isAuthorized(c echo.Context) bool {
	authorized := false
	if session, err := c.Cookie(api.sessionsStorage.GetSessionHeaderName()); err == nil && session != nil {
		authorized = api.sessionsStorage.SessionExists(session.Value)
	}

	return authorized
}

/*
func (api *Api) RootHandler(c echo.Context) error {
	return nil
}

*/

/*
func (api *Api) UserHandler(c echo.Context) error {
	if !api.IsAuthorized(c) {
		return c.JSON(ErrUserNotAuthorised.Status, ErrUserNotAuthorised.Message)
	}
	cookie, _ := c.Cookie("session_id")
	userLogin, ok := api.sessions_[cookie.Value]
	if !ok {
		return c.JSON(ErrUserNotExist.Status, ErrUserNotExist.Message)
	}
	user, _ := api.usersStorage.GetUserByLogin(userLogin)
	data := models.SignupData{
		UserName:   user.Username,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		MiddleName: user.MiddleName,
		Email:      user.Email,
		Login:      user.Login,
	}
	response := models.SignupResponse{
		Data:    data,
		Message: "Hello! Its your profile",
	}
	return c.JSON(http.StatusOK, response)
}

func (api *Api) LoginHandler(c echo.Context) error {
	cookie, _ := c.Cookie("session_id")
	if api.IsAuthorized(c) {
		userLogin, _ := api.sessions_[cookie.Value]
		user, _ := api.usersStorage.GetUserByLogin(userLogin)
		data := models.SignupData{
			UserName:   user.Username,
			FirstName:  user.FirstName,
			LastName:   user.LastName,
			MiddleName: user.MiddleName,
			Email:      user.Email,
			Login:      user.Login,
		}
		response := models.SignupResponse{
			Data:    data,
			Message: "Hello",
		}
		return c.JSON(http.StatusOK, response)
	}
	userForm := new(models.LoginForm)

	formData, err := ioutil.ReadAll(c.Request().Body)
	defer c.Request().Body.Close()
	if err != nil {
		return c.JSON(ErrUnpackingJSON.Status, ErrUnpackingJSON.Message+"1")
	}
	//fmt.Println(string(formData))
	err = json.Unmarshal(formData, &userForm)
	if err != nil {
		return c.JSON(ErrUnpackingJSON.Status, ErrUnpackingJSON.Message+"2")
	}
	// можно добавить проверки на валидность логина и пароля

	userFromBD, err := api.usersStorage.GetUserByLogin(userForm.Login)
	if err != nil {
		return c.JSON(ErrUserNotExist.Status, ErrUserNotExist.Message)
	}
	if userFromBD.Password != userForm.Password {
		return c.JSON(ErrWrongPassword.Status, ErrWrongPassword.Message)
	}
	cookie = makeCookie()
	c.SetCookie(cookie)
	responseData := models.SignupData{
		UserName:   userFromBD.Username,
		FirstName:  userFromBD.FirstName,
		LastName:   userFromBD.LastName,
		MiddleName: userFromBD.MiddleName,
		Email:      userFromBD.Email,
		Login:      userFromBD.Login,
	}
	response := models.SignupResponse{
		Data:    responseData,
		Message: "Hello",
	}
	return c.JSON(http.StatusOK, response)
}

func (api *Api) LogoutHandler(c echo.Context) error {
	if !api.IsAuthorized(c) {
		return c.JSON(ErrAlreadyLogout.Status, ErrAlreadyLogout.Message)
	}
	cookie, _ := c.Cookie("session_id")
	delete(api.sessions_, cookie.Value)
	cookie.Expires = time.Now().Local().Add(-1 * time.Hour)
	c.SetCookie(cookie)

	return c.JSON(LogoutResponse.Status, LogoutResponse.Message)
}

*/

/*
func (api *Api) SignupUserHandler(c echo.Context) error {
	newUser := new(models.User)
	requestData, err := ioutil.ReadAll(c.Request().Body)
	defer c.Request().Body.Close()
	if err != nil {
		return c.JSON(ErrUnpackingJSON.Status, ErrUnpackingJSON.Message)
	}

	err = json.Unmarshal(requestData, &newUser)
	if err != nil {
		return c.JSON(ErrUnpackingJSON.Status, ErrUnpackingJSON.Message)
	}

	// проверка есть ли такой пользователь
	_, err = api.usersStorage.GetUserByLogin(newUser.Login)
	if err == nil {
		return c.JSON(ErrUserExist.Status, ErrUserExist.Message)
	}

	//если правильно понял про мапу sessions,но это надо будет переписать
	if api.IsAuthorized(c) {
		return c.JSON(ErrUserAuthorised.Status, ErrUserAuthorised.Message)
	}

	//проверки на валидность

	_ = api.usersStorage.AddUser(*newUser)
	cookie := makeCookie()
	c.SetCookie(cookie)

	//добавил сессию
	api.sessions_[cookie.Value] = newUser.Login

	res := models.SignupData{
		UserName:   newUser.Username,
		FirstName:  newUser.FirstName,
		LastName:   newUser.LastName,
		MiddleName: newUser.MiddleName,
		Email:      newUser.Email,
		Login:      newUser.Login,
	}
	response := models.SignupResponse{
		Data:    res,
		Message: "You have successfully registered",
	}
	return c.JSON(http.StatusOK, response)
}

*/

func (api *Api) SignupUserHandler(c echo.Context) error {
	defer c.Request().Body.Close()

	parsedInput := new(models.User)
	if err := c.Bind(parsedInput); err != nil {
		c.Logger().Printf("Error: %s", err.Error())
		return c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	log.Println("Parsed input user data:", parsedInput)

	// TODO VALIDATOR

	if api.usersStorage.UserIsExistByLogin(parsedInput.NewUserData.Login) || api.usersStorage.UserIsExistByEmail(parsedInput.NewUserData.Email) {
		c.Logger().Printf("Error: %s", "user with this login or email exist")
		return c.JSON(http.StatusConflict, http.StatusText(http.StatusConflict))
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
		return c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	cookie := api.sessionsStorage.CreateSessionForUser(parsedInput.NewUserData.Email)
	c.SetCookie(cookie)
	api.sessionsStorage.PrintSessions()

	return c.JSON(http.StatusOK, http.StatusText(http.StatusOK))
}

func (api *Api) CreateSessionHandler(c echo.Context) error {
	defer c.Request().Body.Close()

	parsedInput := new(models.SessionCreate)
	if err := c.Bind(parsedInput); err != nil {
		c.Logger().Printf("Error: %s", err.Error())
		return c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
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
		return c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	if user.Password != password {
		c.Logger().Printf("Error: %s", "invalid password.")
		return c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	cookie := api.sessionsStorage.CreateSessionForUser(user.Email)
	c.SetCookie(cookie)
	api.sessionsStorage.PrintSessions()

	return c.JSON(http.StatusOK, http.StatusText(http.StatusOK))
}

func (api *Api) RemoveSessionHandler(c echo.Context) error {
	if !api.isAuthorized(c) {
		c.Logger().Printf("Error: %s", "unauthorized")
		return c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	}
	cookie, err := c.Cookie(api.sessionsStorage.GetSessionHeaderName())
	if err != nil {
		c.Logger().Printf("Error: %s", err.Error())
		return c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	}

	api.sessionsStorage.RemoveSession(cookie)
	api.sessionsStorage.PrintSessions()
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, http.StatusText(http.StatusOK))
}

func (api *Api) SessionInfoHandler(c echo.Context) error {
	if !api.isAuthorized(c) {
		c.Logger().Printf("Error: %s", "unauthorized")
		return c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	}
	cookie, err := c.Cookie(api.sessionsStorage.GetSessionHeaderName())
	if err != nil {
		c.Logger().Printf("Error: %s", err.Error())
		return c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
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
		return c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	if startFromArticleOfNumber < 0 {
		c.Logger().Printf("Error: startFromArticleOfNumber = %d < 0", startFromArticleOfNumber)
		return c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	articles, err := api.feedStorage.GetArticles()
	if err != nil {
		c.Logger().Printf("Error: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
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
