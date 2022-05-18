package models

import "gorm.io/gorm"

// User schema of the user table
type User struct {
	gorm.Model
	ID       int64   `json:"id" gorm:"primary_key;not null"`
	Name     string  `json:"name" gorm:"not null"`
	Location float64 `json:"location" gorm:"not null"`
	Gender   string  `json:"gender"`
	Email    string  `json:"email"`

	Likes []User `gorm:"many2many:user_likes;"`
}

// Like schema of the like table

// type UserLikes struct {
// 	ID      int64 `json:"id" gorm:"primary_key;not null"`
// 	Liked   int64 `json:"who_likes"`
// 	LikedBy int64 `json:"who_is_liked"`
// }
