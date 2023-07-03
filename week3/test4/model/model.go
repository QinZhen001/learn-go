package model

type Product struct {
	ID   int    `uri:"id" json:"id" binding:"required"`
	Name string ` uri:"name" json:"name" binding:"required"`
}
