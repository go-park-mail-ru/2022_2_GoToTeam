package api

import (
	"2022_2_GoTo_team/server/storage"
	"2022_2_GoTo_team/server/storage/models"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const serverAddress = "95.163.213.142:3004"
const articleNumber = 3

type Api struct {
	serverAddress string
	usersStorage  *storage.UsersStorage
	feedStorage   *storage.FeedStorage
	//sessions     []models.Session
	sessions_ map[string]string
}

func GetApi() *Api {
	authApi := &Api{
		serverAddress: serverAddress,
		usersStorage:  storage.GetUsersStorage(),
		feedStorage:   storage.GetFeedStorage(),
		sessions_:     map[string]string{},
	}
	authApi.usersStorage.PrintUsers()
	authApi.feedStorage.PrintArticles()

	return authApi
}

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

func (api *Api) IsAuthorized(c echo.Context) bool {
	authorized := false
	session, err := c.Cookie("session_id")
	if err == nil && session != nil {
		_, authorized = api.sessions_[session.Value]
	}

	return authorized
}

func (api *Api) RootHandler(c echo.Context) error {
	return nil
}

func (api *Api) UserHandler(c echo.Context) error {
	if !api.IsAuthorized(c) {
		return c.JSON(models.ErrUserNotAuthorised.Status, models.ErrUserNotAuthorised.Message)
	}
	cookie, _ := c.Cookie("session_id")
	userLogin, ok := api.sessions_[cookie.Value]
	if !ok {
		return c.JSON(models.ErrUserNotExist.Status, models.ErrUserNotExist.Message)
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
		return c.JSON(models.ErrUnpackingJSON.Status, models.ErrUnpackingJSON.Message+"1")
	}
	//fmt.Println(string(formData))
	err = json.Unmarshal(formData, &userForm)
	if err != nil {
		return c.JSON(models.ErrUnpackingJSON.Status, models.ErrUnpackingJSON.Message+"2")
	}
	// можно добавить проверки на валидность логина и пароля

	userFromBD, err := api.usersStorage.GetUserByLogin(userForm.Login)
	if err != nil {
		return c.JSON(models.ErrUserNotExist.Status, models.ErrUserNotExist.Message)
	}
	if userFromBD.Password != userForm.Password {
		return c.JSON(models.ErrWrongPassword.Status, models.ErrWrongPassword.Message)
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
		return c.JSON(models.ErrAlreadyLogout.Status, models.ErrAlreadyLogout.Message)
	}
	cookie, _ := c.Cookie("session_id")
	delete(api.sessions_, cookie.Value)
	cookie.Expires = time.Now().Local().Add(-1 * time.Hour)
	c.SetCookie(cookie)

	return c.JSON(models.LogoutResponse.Status, models.LogoutResponse.Message)
}

func (api *Api) SignupUserHandler(c echo.Context) error {
	newUser := new(models.User)
	requestData, err := ioutil.ReadAll(c.Request().Body)
	defer c.Request().Body.Close()
	if err != nil {
		return c.JSON(models.ErrUnpackingJSON.Status, models.ErrUnpackingJSON.Message)
	}

	err = json.Unmarshal(requestData, &newUser)
	if err != nil {
		return c.JSON(models.ErrUnpackingJSON.Status, models.ErrUnpackingJSON.Message)
	}

	// проверка есть ли такой пользователь
	_, err = api.usersStorage.GetUserByLogin(newUser.Login)
	if err == nil {
		return c.JSON(models.ErrUserExist.Status, models.ErrUserExist.Message)
	}

	//если правильно понял про мапу sessions,но это надо будет переписать
	if api.IsAuthorized(c) {
		return c.JSON(models.ErrUserAuthorised.Status, models.ErrUserAuthorised.Message)
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

func (api *Api) CreateSessionHandler(c echo.Context) error {
	return nil
}

func (api *Api) FeedHandler(c echo.Context) error {
	strId := c.QueryParam("idLastLoaded")
	if strId == "" {
		strId = "0"
	}

	startId, err := strconv.Atoi(strId)
	if err != nil {
		return c.JSON(models.ErrNoNextFeedId.Status, models.ErrNoNextFeedId.Message)
	}

	var ArticlesData []*models.Article
	AllArticles := storage.GetFeedStorage()
	testData, _ := AllArticles.GetArticles()
	if startId >= 0 && startId+articleNumber < len(testData) {
		ArticlesData = testData[startId : startId+articleNumber]
	} else {
		startId = 0
		if len(testData) > articleNumber {
			startId = len(testData) - articleNumber
		}
		ArticlesData = testData[startId : len(testData)-1]
	}

	response := models.ArticleResponse{
		Status: http.StatusOK,
		Data:   ArticlesData,
	}
	return c.JSON(response.Status, response.Data)
}

func makeCookie() *http.Cookie {
	SID := RandStringRunes(32)
	cookie := new(http.Cookie)
	cookie.Name = "session_id"
	cookie.Value = SID
	cookie.HttpOnly = true
	cookie.Expires = time.Now().Add(10 * time.Hour)
	return cookie
}

func (api *Api) printSessions() {
	log.Println("Sessions in storage:")
	/*
		for _, v := range api.sessions {
			log.Printf("%#v ", v)
		}

	*/
	for k, v := range api.sessions_ {
		log.Printf("cook: %#v for user id: %#v", k, v)
	}
}
