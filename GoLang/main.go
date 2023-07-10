package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	id := uuid.New().String()
	return &Product{
		ID:    id,
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "test:test@tcp(localhost:3306)/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	product := NewProduct("Apple", 100)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}
	product.Name = "Orange"
	product.Price = 200
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}
	product, err = getProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(product.ID, product.Name, product.Price)
}

// private

func insertProduct(db *sql.DB, product *Product) error {
	smt, err := db.Prepare("INSERT INTO products(id, name, price) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer smt.Close()
	_, err = smt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	smt, err := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer smt.Close()
	_, err = smt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func getProduct(db *sql.DB, id string) (*Product, error) {
	smt, err := db.Prepare("SELECT id, name, price FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer smt.Close()
	var product Product
	err = smt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
