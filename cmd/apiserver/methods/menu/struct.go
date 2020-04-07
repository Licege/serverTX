package menu

import "test2/cmd/apiserver/methods/files"

type Dish struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Weight      int        `json:"weight"`
	Price       int        `json:"price"`
	CategoryId  int        `json:"category_id"`
	File        files.File `json:"file"`
}
