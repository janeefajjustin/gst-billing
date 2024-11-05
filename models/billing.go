package models

import (
	"github.com/janeefajjustin/gst-billing/db"
)


type Billing struct {
	ID       int64
	Name     string  `binding:"required"`
	Code     int64   `binding:"required"`
	Quantity int64   `binding:"required"`
	Amount   float64 `binding:"required"`
}

func (p *Product) SaveBilling(quantity int64, total float64) error {
	query := `
	INSERT INTO billings(product_name,product_code,quantity,amount) 
	VALUES (?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Name, p.Code, quantity, total)
	if err != nil {
		return err
	}
	return err
}



func GetAllBilling() ([]Billing, error) {
	query := "SELECT * FROM billings"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var billings []Billing

	for rows.Next() {
		var billing Billing
		err := rows.Scan(&billing.ID, &billing.Name, &billing.Code, &billing.Quantity, &billing.Amount)
		if err != nil {
			return nil, err
		}

		billings = append(billings, billing)
	}

	return billings, nil
}
