package main

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	UserID   string `gorm:"column:user_id;primaryKey"`
	Password string
}

type Level struct {
	UserID string `gorm:"column:user_id;primaryKey"`
	Level  int
}

type Quiz struct {
	QuizID          int `gorm:"column:quiz_id;primaryKey"`
	Chapter         int
	CorrectChoiceID int    `gorm:"column:correct_choice_id"`
	QuizTitle       string `gorm:"column:quiz_title"`
	ProblemNumber   int    `gorm:"column:problem_number"`
}

type Choice struct {
	ChoiceID   int    `gorm:"column:choice_id;primaryKey"`
	QuizID     int    `gorm:"column:quiz_id"`
	ChoiceText string `gorm:"column:choice_text"`
}

func ConnectDB() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&User{})
	return db, nil
}

func CreateUser(db *gorm.DB, user *User) error {
	result := db.Create(user)
	return result.Error
}

func GetUser(db *gorm.DB, userId string) (User, error) {
	var user User
	result := db.First(&user, "user_id = ?", userId)
	return user, result.Error
}

func UpdateUser(db *gorm.DB, user *User) error {
	result := db.Save(user)
	return result.Error
}

func DeleteUser(db *gorm.DB, userId string) error {
	var user User
	result := db.Delete(&user, "user_id = ?", userId)
	return result.Error
}
