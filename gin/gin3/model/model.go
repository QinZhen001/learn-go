package model

type Product struct {
	ID   int32  `uri:"id" binding:"required,gt=0"`
	Name string `uri:"name" binding:"required"`
}
