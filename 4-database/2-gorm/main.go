package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryId int
	Category   Category
	// This will create created_at, updated_at and deleted_at columns - Soft delete
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// Create
	// db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 100.0,
	// })

	// products := []Product{
	// 	{Name: "Notebook", Price: 2000.00},
	// 	{Name: "Mouse", Price: 150.00},
	// 	{Name: "Keyboard", Price: 223.00},
	// }

	// db.Create(&products)

	// Select ONE
	var product1 Product
	var product2 Product
	db.First(&product1, 2)
	fmt.Println(product1)

	db.First(&product2, "name = ?", "Mouse")
	fmt.Println(product2)

	// Select ALL
	var products []Product
	db.Limit(2).Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}

	// Select with CONDITIONS
	var products2 []Product
	db.Where("price < ?", 200).Find(&products2)
	for _, product := range products2 {
		fmt.Println(product)
	}

	var products3 []Product
	db.Where("name LIKE ?", "%book%").Find(&products3)
	for _, product := range products3 {
		fmt.Println(product)
	}

	// Update
	var product3 Product
	db.First(&product3, 1)
	product3.Name = "New Mouse"
	db.Save(&product3)

}
