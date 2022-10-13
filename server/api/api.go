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
	"time"
)

const serverAddress = "95.163.213.142:3004"

type Api struct {
	serverAddress string
	usersStorage  *storage.UsersStorage
	feedStorage   *storage.FeedStorage
	//sessions     []models.Session
	sessions_ map[string]int
}

func GetApi() *Api {
	authApi := &Api{
		serverAddress: serverAddress,
		usersStorage:  storage.GetUsersStorage(),
		feedStorage:   storage.GetFeedStorage(),
		sessions_:     map[string]int{},
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

func (api *Api) IsAuthorized(w http.ResponseWriter, r *http.Request) bool {
	authorized := false
	session, err := r.Cookie("session_id")
	if err == nil && session != nil {
		_, authorized = api.sessions_[session.Value]
	}

	return authorized
}

//func (api *Api) RootHandler(w http.ResponseWriter, r *http.Request) {
//	log.Println("Called RootHandler.")
//	log.Println(api.serverAddress)
//
//	authorized := false
//	session, err := r.Cookie("session_id")
//	if err == nil && session != nil {
//		_, authorized = api.sessions_[session.Value]
//	}
//
//	if authorized {
//		w.Write([]byte("Authrorized"))
//		//w.WriteHeader(http.StatusOK)
//	} else {
//		//w.Write([]byte("not autrorized"))
//		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
//	}
//}

func (api *Api) RootHandler(c echo.Context) error {
	return nil
}

//func (api *Api) LoginHandler(w http.ResponseWriter, r *http.Request) {
//	log.Println("Called LoginHandler.")
//
//	email := r.FormValue("email")
//	password := r.FormValue("password")
//	log.Println("LoginHandler")
//	log.Println("URL", r.URL)
//	log.Println("email", email)
//	log.Println("password ", password)
//
//	//user, ok := api.users[r.FormValue("login")]
//	user, err := api.usersStorage.GetUserByEmail(email)
//	if err != nil {
//		log.Println("Error ", err)
//		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
//		return
//	}
//
//	if user.Password != password {
//		log.Println("Error ", err)
//		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
//		return
//	}
//
//	log.Println("____________________________________________________________________________________")
//
//	SID := RandStringRunes(32)
//
//	api.sessions_[SID] = user.UserId
//
//	cookie := &http.Cookie{
//		Name:    "session_id",
//		Path:    "/",
//		Value:   SID,
//		Expires: time.Now().Add(10 * time.Hour),
//	}
//	http.SetCookie(w, cookie)
//
//	api.printSessions()
//
//	w.Write([]byte(SID))
//
//}

func (api *Api) LoginHandler(c echo.Context) error {
	return nil
}

//func (api *Api) LogoutHandler(w http.ResponseWriter, r *http.Request) {
//	log.Println("Called LogoutHandler.")
//
//	session, err := r.Cookie("session_id")
//	if err == http.ErrNoCookie {
//		http.Error(w, `no sess`, 401)
//		return
//	}
//
//	if _, ok := api.sessions_[session.Value]; !ok {
//		http.Error(w, `no sess`, 401)
//		return
//	}
//
//	delete(api.sessions_, session.Value)
//
//	session.Expires = time.Now().AddDate(0, 0, -1)
//	http.SetCookie(w, session)
//}

func (api *Api) LogoutHandler(c echo.Context) error {
	return nil
}

//func (api *Api) SignupUserHandler(w http.ResponseWriter, r *http.Request) {
//	log.Println("Called SignupUserHandler.")
//
//	if r.Method == "OPTIONS" {
//		w.Header().Set("Access-Control-Request-Headers", r.Header.Get("Access-Control-Request-Headers"))
//		w.Header().Set("Access-Control-Allow-Headers", r.Header.Get("Access-Control-Request-Headers"))
//		w.Header().Set("Access-Control-Allow-Origin", api.serverAddress)
//		w.Header().Set("Access-Control-Allow-Credentials", "true")
//		w.Header().Set("Access-Control-Request-Method", "POST")
//		w.WriteHeader(200)
//	} else {
//		defer r.Body.Close()
//		w.Header().Set("Access-Control-Allow-Origin", api.serverAddress)
//		w.Header().Set("Access-Control-Allow-Credentials", "true")
//
//		parsedInput := new(models.User)
//		err := json.NewDecoder(r.Body).Decode(parsedInput)
//		if err != nil {
//			log.Println("Error while decode JSON:", err)
//			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
//			return
//		}
//
//		log.Println("Parsed input:", parsedInput)
//		err = api.usersStorage.AddUser(
//			parsedInput.NewUserData.Username,
//			parsedInput.NewUserData.Email,
//			parsedInput.NewUserData.Login,
//			parsedInput.NewUserData.Password,
//		)
//
//		if err != nil {
//			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
//			return
//		}
//
//		api.usersStorage.PrintUsers()
//		w.WriteHeader(http.StatusOK)
//	}
//}

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

	cc, _ := c.Cookie("session")

	//если правильно понял про мапу sessions,но это надо будет переписать
	if _, ok := api.sessions_[cc.Value]; ok {
		return c.JSON(models.ErrUserAuthorised.Status, models.ErrUserAuthorised.Message)
	}

	//проверки на валидность

	id := api.usersStorage.AddUser(*newUser)
	cookie := makeCookie()
	c.SetCookie(cookie)

	//добавил сессию
	api.sessions_[cookie.Value] = id

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

//func (api *Api) CreateSessionHandler(w http.ResponseWriter, r *http.Request) {
//	log.Println("Called CreateSessionHandler.")
//
//	defer r.Body.Close()
//	if r.Method == "OPTIONS" {
//		w.Header().Set("Access-Control-Request-Headers", r.Header.Get("Access-Control-Request-Headers"))
//		w.Header().Set("Access-Control-Allow-Headers", r.Header.Get("Access-Control-Request-Headers"))
//		w.Header().Set("Access-Control-Allow-Origin", api.serverAddress)
//		w.Header().Set("Access-Control-Allow-Credentials", "true")
//		w.Header().Set("Access-Control-Request-Method", "POST")
//		w.WriteHeader(http.StatusOK)
//	} else {
//		w.Header().Set("Access-Control-Allow-Origin", api.serverAddress)
//		w.Header().Set("Access-Control-Allow-Credentials", "true")
//		parsedInput := new(models.Session)
//		err := json.NewDecoder(r.Body).Decode(parsedInput)
//		if err != nil {
//			log.Println("Error while decode JSON:", err)
//			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
//			return
//		}
//
//		log.Println(parsedInput)
//
//		//api.sessions = append(api.sessions, *parsedInput)
//
//		// ======================================
//
//		email := parsedInput.UserData.Email
//		password := parsedInput.UserData.Password
//		log.Println("URL", r.URL)
//		log.Println("email", email)
//		log.Println("password ", password)
//
//		//user, ok := api.users[r.FormValue("login")]
//		user, err := api.usersStorage.GetUserByEmail(email)
//		if err != nil {
//			log.Println("Error ", err)
//			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
//			return
//		}
//
//		if user.Password != password {
//			log.Println("Error ", err)
//			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
//			return
//		}
//
//		log.Println("____________________________________________________________________________________")
//
//		SID := RandStringRunes(32)
//
//		api.sessions_[SID] = user.UserId
//
//		cookie := &http.Cookie{
//			Name:    "session_id",
//			Path:    "/",
//			Value:   SID,
//			Expires: time.Now().Add(10 * time.Hour),
//		}
//		http.SetCookie(w, cookie)
//
//		api.printSessions()
//
//		w.Write([]byte(SID))
//
//		//w.WriteHeader(http.StatusOK)
//	}
//
//}

func (api *Api) CreateSessionHandler(c echo.Context) error {
	return nil
}

//func (api *Api) FeedHandler(w http.ResponseWriter, r *http.Request) {
//	log.Println("Called FeedHandler.")
//
//	if r.Method == "OPTIONS" {
//		w.Header().Set("Access-Control-Request-Headers", r.Header.Get("Access-Control-Request-Headers"))
//		w.Header().Set("Access-Control-Allow-Headers", r.Header.Get("Access-Control-Request-Headers"))
//		w.Header().Set("Access-Control-Allow-Origin", api.serverAddress)
//		w.Header().Set("Access-Control-Allow-Credentials", "true")
//		w.Header().Set("Access-Control-Request-Method", "POST")
//		w.WriteHeader(200)
//	} else {
//		//if api.IsAuthorized(w, r) {
//		if true {
//			w.Header().Set("Content-Type", "application/json")
//			w.Header().Set("Access-Control-Allow-Origin", api.serverAddress)
//			w.Header().Set("Access-Control-Allow-Credentials", "true")
//			articles, err := api.feedStorage.GetArticles()
//			if err != nil {
//				log.Println("Error api.feedStorage.GetArticles", err)
//				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
//				return
//			}
//
//			feed := models.Feed{}
//			for _, v := range articles {
//				article := models.Article{
//					Id:          v.Id,
//					Title:       v.Title,
//					Description: v.Description,
//					Tags:        v.Tags,
//					Category:    v.Category,
//					Rating:      v.Rating,
//					Authors:     v.Authors,
//					Content:     v.Content,
//				}
//				feed.Articles = append(feed.Articles, article)
//			}
//
//			log.Println(feed)
//
//			err = json.NewEncoder(w).Encode(&feed)
//			if err != nil {
//				log.Println("Error while encode JSON:", err)
//				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
//				return
//			}
//		} else {
//			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
//		}
//	}
//}

func (api *Api) UserHandler(c echo.Context) error {
	return nil
}

func (api *Api) FeedHandler(c echo.Context) error {
	return nil
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