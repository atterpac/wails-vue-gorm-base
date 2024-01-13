package models

import (
	"changeme/api/auth"
	"errors"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id         uuid.UUID      `json:"id" gorm:"primaryKey;not_null"`
	Name       string         `json:"first_name"`
	Username   string         `json:"username"`
	Email      string         `json:"email"`
	Jwt        string         `json:"jwt"`
	Password   string         `json:"password"`
	ModifiedAt time.Time      `json:"modified_at"`
	CreatedAt  time.Time      `json:"created_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

type Register struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
	Email     string `json:"email"`
}

type Login struct {
	Username string `json:"username"` // Could be email
	Password string `json:"password"`
}

func (u *User) Verify() (bool, error) {
	if u.Username == "" {
		return false, errors.New("Username is required")
	}
	if u.Email == "" {
		return false, errors.New("Email is required")
	}
	if u.Password == "" {
		return false, errors.New("Password is required")
	}
	return true, nil
}

func (u *Login) Verify() error {
	if u.Username == "" {
		return errors.New("Username is required")
	}
	if u.Password == "" {
		return errors.New("Password is required")
	}
	return nil
}

func (r *Register) Verify() error {
	if r.Username == "" {
		return errors.New("Username is required")
	}
	if r.Password == "" {
		return errors.New("Password is required")
	}
	if r.Password2 == "" {
		return errors.New("Confirm Password")
	}
	if r.Password != r.Password2 {
		return errors.New("Passwords do not match")
	}
	return nil
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	u.Id = uuid.New()
	u.Password, err = auth.Hash(u.Password)
	if err != nil {
		return err
	}
	u.CreatedAt = time.Now()
	slog.Debug("BeforeCreate", "user", u)
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.ModifiedAt = time.Now()
	return
}

func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	u.DeletedAt = gorm.DeletedAt{
		Time:  time.Now(),
		Valid: true,
	}
	return
}

func (u *User) Create(db *gorm.DB) (err error) {
	err = db.Save(u).Error
	return
}

func (u *User) Update(db *gorm.DB) (err error) {
	err = db.Save(u).Error
	return
}

func (u *User) Delete(db *gorm.DB) (err error) {
	err = db.Delete(u).Error
	return
}

func (u *User) Get(db *gorm.DB) (err error) {
	err = db.First(u).Error
	return
}

func GetAllUsers(db *gorm.DB) ([]User, error) {
	var u []User
	err := db.Find(&u).Error
	return u, err
}

func GetUserByEmail(db *gorm.DB, email string) (User, error) {
	var u User
	err := db.First(&u).Where("email = ?", email).Error
	return u, err
}

func GetUserByUsername(db *gorm.DB, user string) (User, error) {
	var u User
	err := db.First(&u).Where("username = ?", user).Error
	return u, err
}

func GetUserById(db *gorm.DB, id string) (user User, err error) {
	err = db.Where("id = ?", id).First(&user).Error
	return
}
