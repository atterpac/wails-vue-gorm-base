package controllers

import (
	"changeme/api/models"
	"testing"

	"github.com/google/uuid"
)

var id uuid.UUID

func TestCreateUser(t *testing.T) {
	app := initApp()
    // Setup test data
    user := models.Register{
        Username:  "testuser",
        Password:  "testpassword",
        Password2: "testpassword",
        Email:     "test@example.com",
    }
    // Call the CreateUser function
    err := app.CreateUser(user)
    // Check if there was an error
    if err != nil {
        t.Errorf("CreateUser failed: %v", err)
    }
}

func TestGetUserByEmail(t *testing.T) {
	app := initApp()
    // Setup test data
    userEmail := "test@example.com"
    // Call the GetUserByEmail function
    user, err := app.GetUserByEmail(userEmail)
    // Check if there was an error
    if err != nil {
        t.Errorf("GetUserByEmail failed: %v", err)
    }
	id = user.Id 
}

func TestUpdateUser(t *testing.T) {
	app := initApp()
    // Setup test data
    user := models.User{
		Id: id,
		Username:  "testuser",
		Password:  "testpasswordUpdated",
		Email:     "test@example.com",
	}
    // Call the UpdateUser function
    err := app.UpdateUser(user)
    // Check if there was an error
    if err != nil {
        t.Errorf("UpdateUser failed: %v", err)
    }
}

func TestGetUserById(t *testing.T) {
	app := initApp()
    // Call the GetUserById function
    _, err := app.GetUserById(id.String())
    // Check if there was an error
    if err != nil {
        t.Errorf("GetUserById failed: %v", err)
    }
}

func TestGetUsers(t *testing.T) {
	app := initApp()
    // Call the GetUsers function
    _, err := app.GetUsers()
    // Check if there was an error
    if err != nil {
        t.Errorf("GetUsers failed: %v", err)
    }
}

func TestDeleteUserById(t *testing.T) {
	app := initApp()
    // Call the DeleteUserById function
    err := app.DeleteUserById(id.String())
    // Check if there was an error
    if err != nil {
        t.Errorf("DeleteUserById failed: %v", err)
    }
}
