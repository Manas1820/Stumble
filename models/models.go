package models

// User schema of the user table
type User struct {
	ID       int64   `json:"id" gorm:"primary_key;not null"`
	Name     string  `json:"name" gorm:"not null"`
	Location float64 `json:"location" gorm:"not null"`
	Gender   string  `json:"gender"`
	Email    string  `json:"email"`
}

// Like schema of the like table

type Like struct {
	ID int64 `json:"id"`
	
}