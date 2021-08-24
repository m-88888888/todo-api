package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Product is struct
type Product struct {
	ID              int `gorm:"primary_key"`
	Code            string
	Price           uint
	ProductServices []ProductService `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// ProductService is struct
type ProductService struct {
	ID        int `gorm:"primary_key"`
	ProductID int
	Name      string
}

func initDB() (db *gorm.DB) {
	DNS := "root:@tcp(localhost:3306)/product?charset=utf8&parseTime=True&loc=Local"
	fmt.Println(DNS)

	db, err := gorm.Open(mysql.Open(DNS), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{}, &ProductService{})

	fmt.Println("db connection is successful")
	return
}

func main() {
	db := initDB()
	insertInitData(db)
	//deleteProduct(db)
}

func insertInitData(db *gorm.DB) {
	product := Product{
		Code:  "D002",
		Price: 20000,
		ProductServices: []ProductService{
			{
				Name: "出張見積もりサービス",
			},
			{
				Name: "保守運用サービス",
			},
		},
	}
	db.Create(&product)
	fmt.Println(product)
}

func deleteProduct(db *gorm.DB) {
	var product Product
	db.Debug().Find(&product)
	var productService ProductService
	db.Debug().Find(&productService)
	fmt.Println(product)
	fmt.Println(productService)
	//　外部キー制約が削除されるだけ。
	//db.Debug().Model(&product).Association("ProductService").Delete(&product.ProductServices)
	//db.Debug().Model(&product).Association("ProductService")

	db.Debug().Select("ProductServices").Delete(&product)
	fmt.Println(product)
}
