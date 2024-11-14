package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

// You need to install docker
// You need to install docker compose

// COMMANDS TO RUN DOCKER
// Start a container

// docker-compose up -d
// docker-compose exec mysql bash
// mysql -uroot -p goexpert
// Type the database password

// Creating a table on mysql - create table products (id varchar(255), name varchar(80), price decimal(10,2), primary key(id));

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	product := NewProduct("Notebook", 1889.90)

	// Inserting product
	err = InsertProduct(db, product)
	if err != nil {
		panic(err)
	}

	// Updating product
	product.Price = 100.0

	err = UpdateProduct(db, product)
	if err != nil {
		panic(err)
	}

	// Selecting product
	p, err := SelectProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Product %v, has %.2f price\n", p.Name, p.Price)

	// Selecting all products
	products, err := SelectAllProduct(db)
	if err != nil {
		panic(err)
	}
	for _, product := range products {
		fmt.Printf("Product %v, has %.2f price\n", product.Name, product.Price)
	}
}

func UpdateProduct(db *sql.DB, product *Product) error {
	statement, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()
	_, err = statement.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func InsertProduct(db *sql.DB, product *Product) error {
	statement, err := db.Prepare("insert into products(id, name, price) values(?, ?, ?)")
	if err != nil {
		return err
	}

	defer statement.Close()
	_, err = statement.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func SelectProduct(db *sql.DB, id string) (*Product, error) {
	statement, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	var product Product
	err = statement.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func SelectAllProduct(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product

	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)

	}
	return products, nil
}
