package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Dish struct {
	ID     uint `gorm:"primaryKey"`
	Name   string
	Price  uint
	Number uint
}

func createColumn(db *gorm.DB, name string, price uint, number uint) {
	db.Create(&Dish{Name: name, Price: price, Number: number})
}

func main() {

	dsn := "root:qweqwe@tcp(127.0.0.1:3306)/foods?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("db not open!")
	}

	// Migrate the schema
	db.AutoMigrate(&Dish{})

	// Create
	createColumn(db, "sushi", 400, 500)
	createColumn(db, "pizza", 700, 300)
	createColumn(db, "doner", 300, 180)
	createColumn(db, "wok", 400, 350)
	createColumn(db, "burger", 126, 250)
	createColumn(db, "avocado", 30, 120)

	// Read
	var dish Dish
	db.First(&dish, 1) // find product with id 1

	fmt.Println(dish.Name, dish.Price, dish.Number)

	// Update
	db.Model(&dish).Update("Price", 500)

	fmt.Println(dish.Name, dish.Price, dish.Number)

	// Delete
	var dishToDelete Dish
	db.First(&dishToDelete, "Name = ?", "avocado")
	fmt.Println(dishToDelete.Name, dishToDelete.Price, dishToDelete.Number)
	db.Delete(&dishToDelete)

}
