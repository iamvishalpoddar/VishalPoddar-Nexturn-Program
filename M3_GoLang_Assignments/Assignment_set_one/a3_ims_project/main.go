package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

var inventory []Product

func AddProduct(id int, name string, price interface{}, stock int) error {
	for _, product := range inventory {
		if product.ID == id {
			return errors.New("product ID must be unique")
		}
	}

	priceFloat, ok := price.(float64)
	if !ok {
		return errors.New("price must be a valid float64 value")
	}

	if stock < 0 {
		return errors.New("stock cannot be negative")
	}

	inventory = append(inventory, Product{
		ID:    id,
		Name:  name,
		Price: priceFloat,
		Stock: stock,
	})
	return nil
}

func UpdateStock(id int, newStock int) error {
	if newStock < 0 {
		return errors.New("stock cannot be negative")
	}
	for i, product := range inventory {
		if product.ID == id {
			inventory[i].Stock = newStock
			return nil
		}
	}
	return errors.New("product not found")
}

func SearchProduct(query string) (*Product, error) {
	for _, product := range inventory {
		if strings.EqualFold(fmt.Sprint(product.ID), query) || strings.EqualFold(product.Name, query) {
			return &product, nil
		}
	}
	return nil, errors.New("product not found")
}

func DisplayInventory() {
	if len(inventory) == 0 {
		fmt.Println("Inventory is empty.")
		return
	}
	fmt.Println("ID\tName\t\tPrice\tStock")
	fmt.Println("-----------------------------------")
	for _, product := range inventory {
		fmt.Printf("%d\t%-10s\t%.2f\t%d\n", product.ID, product.Name, product.Price, product.Stock)
	}
}

func SortInventory(by string) error {
	switch strings.ToLower(by) {
	case "price":
		sort.Slice(inventory, func(i, j int) bool {
			return inventory[i].Price < inventory[j].Price
		})
	case "stock":
		sort.Slice(inventory, func(i, j int) bool {
			return inventory[i].Stock < inventory[j].Stock
		})
	default:
		return errors.New("invalid sort criteria, use 'price' or 'stock'")
	}
	return nil
}

func main() {
	if err := AddProduct(1, "Laptop", 50000.0, 10); err != nil {
		fmt.Println("Error:", err)
	}
	if err := AddProduct(2, "Phone", 20000.0, 25); err != nil {
		fmt.Println("Error:", err)
	}
	if err := AddProduct(3, "Headphones", 1500.0, 50); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\nCurrent Inventory:")
	DisplayInventory()

	if err := UpdateStock(2, 30); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("\nStock updated.")
	}

	DisplayInventory()

	if product, err := SearchProduct("Phone"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("\nProduct found: %+v\n", *product)
	}

	if err := SortInventory("price"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("\nInventory sorted by price:")
		DisplayInventory()
	}

	if err := SortInventory("stock"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("\nInventory sorted by stock:")
		DisplayInventory()
	}
}
