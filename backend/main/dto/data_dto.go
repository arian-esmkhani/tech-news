package dto

//For Getting Data that Cards need
type DataDto struct {
	Id		 uint64  `json:"id" db:"id"`
	Title    string  `json:"title"  db:"title"`
	ImgUrl 	 *string `json:"imgUrl"  db:"imag_url"`
}