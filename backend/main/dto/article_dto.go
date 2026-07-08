package dto

//Article structure
type ArticleDto struct {
	Head        string      `json:"head"`
	Body		string		`json:"body"`
	ImgUrl		string		`json:"imgUrl"`
	Sec			int64		`json:"sec"`
}

type ArticleResponse struct {
	TopicDto    TopicDto	 `json:"topicDto"`
	ArticleDto	[]ArticleDto `json:"articleDto"`
}

type ArticleRow struct {
	Title          string `db:"title"`
	Description    string `db:"description"`
	TopicImage     *string `db:"topic_imag_url"`
	FirstCategory  string `db:"first_category"`
	SecondCategory string `db:"second_category"`
	Head           string `db:"head"`
	Body           string `db:"body"`
	ArticleImage   *string `db:"article_imag_url"`
	Sec            int64  `db:"sec"`
}
