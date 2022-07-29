package main

import "github.com/google/uuid"

var customers = map[string]Customer{
	"test-9651-4a8c-9c9e-03a5849d94dd": {
		Name:      "Eric",
		Role:      "elite",
		Email:     "eric@mail.com",
		Phone:     31463388,
		Contacted: true,
	},
	uuid.New().String(): {
		Name:      "Anna",
		Role:      "premium",
		Email:     "anna@mail.com",
		Phone:     33673637,
		Contacted: false,
	},
	uuid.New().String(): {
		Name:      "Gabrielle",
		Role:      "standard",
		Email:     "gabrielle@mail.com",
		Phone:     337983783,
		Contacted: true,
	},
}
