package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int32 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Category struct {
	BaseModel
	Name             string `gorm:"type:varchar(32);not null"`
	ParentCategoryID int32
	ParentCategory   *Category
	Level            int32
	SubCategory      []*Category `gorm:"foreignKey:ParentCategoryID;references:ID"`
}

type ProductCategoryBrand struct {
	BaseModel

	BrandID    int32 `gorm:"type:int;not null"`
	Brand      Brand
	CategoryID int32 `gorm:"type:int;not null"`
	Category   Category
}

type Brand struct {
	BaseModel
	Name string `gorm:"type:varchar(32);not null"`
	Logo string `gorm:"type:varchar(128);not null"`
}

type Advertise struct {
	BaseModel
	Index int32  `gorm:"type:int;not null;default:1"`
	Image string `gorm:"type:varchar(256);not null"`
	Url   string `gorm:"type:varchar(256);not null"`
	Sort  int32  `gorm:"type:int;not null;default:1"`
}

type Product struct {
	BaseModel
	CateGoryID int `gorm:"type:int;not null"`
	CateGory   Category

	BrandID int `gorm:"type:int;not null"`
	Brand   Brand

	Selling    bool   `gorm:"default:false"`
	IsShipFree bool   `gorm:"default:false"`
	IsPop      bool   `gorm:"default:false"`
	IsNew      bool   `gorm:"default:false"`
	KeyWord    string `gorm:"type:varchar(128);not null"`

	Name       string  `gorm:"type:varchar(128);not null"`
	SN         string  `gorm:"type:varchar(128);not null"`
	FavNum     int32   `gorm:"type:int;default:0"`
	SoldNum    int32   `gorm:"type:int;default:0"`
	Price      float32 `gorm:"not null"`
	RealPrice  float32 `gorm:"not null"`
	ShortDesc  string  `gorm:"type:varchar(256);not null"`
	Images     MyList  `gorm:"type:varchar(1024);not null"`
	DesImages  MyList  `gorm:"type:varchar(1024);not null"`
	CoverImage string  `gorm:"type:varchar(256);not null"`
}

// -------------------------------------------

type MyList []string

func (myList MyList) Value() (driver.Value, error) {
	return json.Marshal(myList)
}

func (myList MyList) Scan(data interface{}) error {
	return json.Unmarshal(data.([]byte), &myList)
}

func (p *Product) AfterCreate(tx *gorm.Tx) error {
	// esProduct := ESProduct{
	// 	ID:         0,
	// 	BrandID:    0,
	// 	CategoryID: 0,
	// 	Selling:    false,
	// 	ShipFree:   false,
	// 	IsPop:      false,
	// 	IsNew:      false,
	// 	Name:       "",
	// 	FavNum:     0,
	// 	SoldNum:    0,
	// 	Price:      0,
	// 	RealPrice:  0,
	// 	ShortDesc:  "",
	// }
	// _, err := internal.ESClient.Index().Index(GetIndex()).BodyJson(esProduct).Id(strconv.Itoa(int(p.ID))).Do(context.Background())
	// if err != nil {
	// 	panic(err)
	// }
	return nil
}
