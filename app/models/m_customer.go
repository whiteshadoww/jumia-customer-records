package models

// Customer model customer
type Customer struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	State       *bool  `json:"state"`
}

func (c Customer) TableName() string {
	return "customer"
}
