package models

type User struct {
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"size:64;not null;unique"`
}

type Blog struct {
	ID   int    `gorm:"primaryKey"`
	Body string `gorm:"size:64;not null"`
}

type UserLikeBlog struct {
	ID     int `gorm:"primaryKey"`
	UserID int `gorm:"not null"`
	BlogID int `gorm:"not null"`
	User   User
	Blog   Blog
}
