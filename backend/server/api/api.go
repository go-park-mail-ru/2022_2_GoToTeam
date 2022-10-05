package api

import (
	"2022_2_GoTo_team/server/api/models"
	"2022_2_GoTo_team/server/storage"
	"encoding/json"
	"log"
	"net/http"
)

type Api struct {
	usersStorage *storage.UsersStorage
	feedStorage  *storage.FeedStorage
	sessions     []models.Session
}

func GetApi() *Api {
	authApi := &Api{
		usersStorage: storage.GetUsersStorage(),
		feedStorage:  storage.GetFeedStorage(),
	}
	authApi.usersStorage.PrintUsers()
	authApi.feedStorage.PrintArticles()

	return authApi
}

func (api *Api) SignupUserHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

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

func (api *Api) CreateSessionHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	parsedInput := new(models.Session)
	err := json.NewDecoder(r.Body).Decode(parsedInput)
	if err != nil {
		log.Println("Error while decode JSON:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	log.Println(parsedInput)

	api.sessions = append(api.sessions, *parsedInput)

	api.printSessions()

	w.WriteHeader(http.StatusOK)
}

func (api *Api) FeedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	articles, err := api.feedStorage.GetArticles()
	if err != nil {
		log.Println("Error while decode JSON:", err)
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

	w.WriteHeader(http.StatusOK)
}

func (api *Api) printSessions() {
	log.Println("Sessions in storage:")
	for _, v := range api.sessions {
		log.Printf("%#v ", v)
	}
}
