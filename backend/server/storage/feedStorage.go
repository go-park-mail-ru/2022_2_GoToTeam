package storage

import (
	"2022_2_GoTo_team/server/storage/models"
	"log"
	"sync"
)

type FeedStorage struct {
	articles []*models.Article
	mu       sync.RWMutex
	//nextID uint
}

func GetFeedStorage() *FeedStorage {
	return &FeedStorage{
		articles: articlesData,
		mu:       sync.RWMutex{},
	}
}

func (o *FeedStorage) PrintArticles() {
	log.Println("Articles in storage:")
	for _, v := range o.articles {
		log.Printf("%#v ", v)
	}
}

func (o *FeedStorage) GetArticles() ([]*models.Article, error) {
	log.Println("Storage GetArticles called.")

	o.mu.RLock()
	defer o.mu.RUnlock()

	return o.articles, nil
}
