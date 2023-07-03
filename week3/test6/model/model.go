package model

type Accountlogin struct {
	AccountName string `json:"accountName" binding:"required,min=8,max=32"`
	Password    string `json:"password" binding:"required,min=4,max=32"`
}
