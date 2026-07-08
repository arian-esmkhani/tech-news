package dto

type TopicDto struct {
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	ImgUrl        string   `json:"imgUrl"`
	FirstCategory string   `json:"firstCategory"`
	SecondCategory string  `json:"secondCategory"`
}

type TopicResponse struct {
    DataDto   []DataDto `json:"dataDto"`
    HasNext  bool	  	`json:"hasNext"`
}