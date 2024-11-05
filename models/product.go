package models

import "github.com/janeefajjustin/gst-billing/db"

type Product struct {
	Code  int64   `binding:"required"`
	Name  string  `binding:"required"`
	Price float64 `binding:"required"`
	Gst   float64 `binding:"required"`
}

func (p *Product) Save() error {
	query := `
	INSERT INTO products(product_code,product_name,product_price,product_gst) 
	VALUES (?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Code, p.Name, p.Price, p.Gst)
	if err != nil {
		return err
	}
	return err

}

func FetchProductByCode(productcode int64) (*Product, error) {
	query := "SELECT * FROM products WHERE product_code = ?"
	row := db.DB.QueryRow(query, productcode)
	var product Product
	err := row.Scan(&product.Code, &product.Name, &product.Price, &product.Gst)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func FetchProductByName(productname string) (*Product, error) {
	query := "SELECT * FROM products WHERE product_name = ?"
	row := db.DB.QueryRow(query, productname)
	var product Product
	err := row.Scan(&product.Code, &product.Name, &product.Price, &product.Gst)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
