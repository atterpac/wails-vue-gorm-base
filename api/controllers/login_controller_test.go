package controllers

import (
	"changeme/api/models"
	"testing"
)

var token string

func TestRegisterUser(t *testing.T) {
	app := initApp()
    // Setup test data
    user := models.Register{
		Username:  "testuser",
		Password:  "testpassword",
		Password2: "testpassword",
    }
    // Call the RegisterUser function
    err := app.RegisterUser(user)
    // Check if there was an error
    if err != nil {
        t.Errorf("RegisterUser failed: %v", err)
    }
}

func TestLoginUser(t *testing.T) {
	var err error
	app := initApp()
    // Setup test data
    login := models.Login{
        Username: "testuser",
        Password: "testpassword",
    }
    // Call the LoginUser function
    token, err = app.LoginUser(login)
    // Check if there was an error
    if err != nil {
        t.Errorf("LoginUser failed: %v", err)
    }
}

func TestRefreshToken(t *testing.T) {
	app := initApp()
    // Call the RefreshToken function
    _, err := app.RefreshToken(token)

    // Check if there was an error
    if err != nil {
        t.Errorf("RefreshToken failed: %v", err)
    }
}



