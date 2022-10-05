package models

type Article struct {
	Id          int
	Title       string
	Description string
	Tags        []string
	Category    string
	Rating      int
	Authors     []string
	Content     string
}
