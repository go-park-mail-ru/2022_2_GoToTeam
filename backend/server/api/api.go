package api

import (
	"2022_2_GoTo_team/server/api/models"
	"2022_2_GoTo_team/server/storage"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Api struct {
	usersStorage *storage.UsersStorage
	feedStorage  *storage.FeedStorage
	//sessions     []models.Session
	sessions_ map[string]int
}

func GetApi() *Api {
	authApi := &Api{
		usersStorage: storage.GetUsersStorage(),
		feedStorage:  storage.GetFeedStorage(),
		sessions_:    map[string]int{},
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

func (api *Api) RootHandler(w http.ResponseWriter, r *http.Request) {
	authorized := false
	session, err := r.Cookie("session_id")
	if err == nil && session != nil {
		_, authorized = api.sessions_[session.Value]
	}

	if authorized {
		w.Write([]byte("autrorized"))
	} else {
		w.Write([]byte("not autrorized"))
	}
}

func (api *Api) LoginHandler(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email")
	password := r.FormValue("password")
	log.Println("email", email)
	log.Println("password ", password)

	//user, ok := api.users[r.FormValue("login")]
	user, err := api.usersStorage.GetUserByEmail(email)
	if err != nil {
		log.Println("Error ", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if user.Password != password {
		log.Println("Error ", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	log.Println("____________________________________________________________________________________")

	SID := RandStringRunes(32)

	api.sessions_[SID] = user.UserId

	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   SID,
		Expires: time.Now().Add(10 * time.Hour),
	}
	http.SetCookie(w, cookie)

	api.printSessions()

	w.Write([]byte(SID))

}

func (api *Api) LogoutHandler(w http.ResponseWriter, r *http.Request) {

	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		http.Error(w, `no sess`, 401)
		return
	}

	if _, ok := api.sessions_[session.Value]; !ok {
		http.Error(w, `no sess`, 401)
		return
	}

	delete(api.sessions_, session.Value)

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
}

func (api *Api) SignupUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Request-Headers", r.Header.Get("Access-Control-Request-Headers"))
		w.Header().Set("Access-Control-Allow-Headers", r.Header.Get("Access-Control-Request-Headers"))
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:3004")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Request-Method", "POST")
		w.WriteHeader(200)
	} else {
		defer r.Body.Close()
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:3004")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		parsedInput := new(models.User)
		err := json.NewDecoder(r.Body).Decode(parsedInput)
		if err != nil {
			log.Println("Error while decode JSON:", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		log.Println(parsedInput)
		err = api.usersStorage.AddUser(
			parsedInput.NewUserData.Username,
			parsedInput.NewUserData.Email,
			parsedInput.NewUserData.Login,
			parsedInput.NewUserData.Password,
		)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		api.usersStorage.PrintUsers()
		w.WriteHeader(http.StatusOK)
	}
}

func (api *Api) CreateSessionHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Request-Headers", r.Header.Get("Access-Control-Request-Headers"))
		w.Header().Set("Access-Control-Allow-Headers", r.Header.Get("Access-Control-Request-Headers"))
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:3004")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Request-Method", "POST")
		w.WriteHeader(http.StatusOK)
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:3004")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		parsedInput := new(models.Session)
		err := json.NewDecoder(r.Body).Decode(parsedInput)
		if err != nil {
			log.Println("Error while decode JSON:", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		log.Println(parsedInput)

		//api.sessions = append(api.sessions, *parsedInput)

		api.printSessions()

		w.WriteHeader(http.StatusOK)
	}

}

func (api *Api) FeedHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Request-Headers", r.Header.Get("Access-Control-Request-Headers"))
		w.Header().Set("Access-Control-Allow-Headers", r.Header.Get("Access-Control-Request-Headers"))
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:3004")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Request-Method", "POST")
		w.WriteHeader(200)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:3004")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		articles, err := api.feedStorage.GetArticles()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
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

		log.Println(feed)

		err = json.NewEncoder(w).Encode(&feed)
		if err != nil {
			log.Println("Error while encode JSON:", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
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
