package main

type Customer struct {
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     uint   `json:"phone"`
	Contacted bool   `json:"contacted"`
}
