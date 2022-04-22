package schema

// Customer schema
type Customer struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	State       bool   `json:"state"`
}

// CustomerQueryParam schema
type CustomerQueryParam struct {
	Country  string `json:"country,omitempty" form:"country,omitempty"`
	State    *bool  `json:"state,omitempty" form:"state,omitempty"`
	Page     int    `json:"-" form:"page,omitempty" validator:"numeric"`
	PageSize int    `json:"-" form:"page_size,omitempty" validator:"numeric"`
}
