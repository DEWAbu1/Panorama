package model

type UpdateYear struct {
	Oldyear int `json:"oldyear"`
	Newyear int `json:"newyear"`
}

type UpdatePicture struct {
	Oldpicture string `json:"oldpicture"`
	Newpicture string `json:"newpicture"`
	Year       int    `json:"year"`
	Place      string `json:"place"`
}

type InsertPicture struct {
	Year    int    `json:"year"`
	Place   string `json:"place"`
	Picture string `json:"picture"`
}

type UpdatePlace struct {
	Oldplace string `json:"oldplace"`
	Newplace string `json:"newplace"`
}

type InsertTag struct {
	Picture string  `json:"picture"`
	Tag     string  `json:"tag"`
	Date    string  `json:"date"`
	Xval    float64 `json:"xval"`
	YVal    float64 `json:"yval"`
	Theta   float64 `json:"theta"`
	Phi     float64 `json:"phi"`
}

type UpdateTag struct {
	Id      int     `json:"id"`
	Picture string  `json:"picture"`
	Tag     string  `json:"tag"`
	Date    string  `json:"date"`
	Xval    float64 `json:"xval"`
	YVal    float64 `json:"yval"`
	Theta   float64 `json:"theta"`
	Phi     float64 `json:"phi"`
}

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
	Total float64 `json:"total"`
}
