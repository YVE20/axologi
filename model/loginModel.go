package model

import "time"

type LoginModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterModel struct {
	CustomerID    string     `json:"customerID"`
	FirstName     string     `json:"firstName"`
	LastName      string     `json:"lastName"`
	Email         string     `json:"email"`
	Password      string     `json:"password"`
	DeviceID      string     `json:"deviceID"`
	IpAddress     string     `json:"ipAddress"`
	BirthDate     *time.Time `json:"birthDate"`
	Photo         string     `json:"photo"`
	StatusAccount string     `json:"statusAccount"`
	BuyerID       string     `json:"buyerID"`
	SellerID      string     `json:"sellerID"`
	CreatedBy     string     `json:"createdBy"`
	CreatedDate   time.Time  `json:"createdDate"`
	UpdatedBy     string     `json:"updatedBy"`
	UpdatedDate   time.Time  `json:"updatedDate"`
}
