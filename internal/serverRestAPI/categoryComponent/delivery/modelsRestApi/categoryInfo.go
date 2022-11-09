package modelsRestApi

type CategoryInfo struct {
	CategoryName     string `json:"category_name"`
	Description      string `json:"description"`
	SubscribersCount int    `json:"subscribers_count"`
}
