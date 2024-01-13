package controllers

import (
	"changeme/api/models"
	"log/slog"
)

func (a *App) CreateUser(user models.Register) error {
	slog.Warn("Creating User", "email", user.Email, "username", user.Username, "password", user.Password)
	err := user.Verify()
	if err != nil {
		slog.Error("Verifying User", "err", err.Error())
		return err
	}
	newUser := models.User {
		Email: user.Email,
		Username: user.Username,
		Password: user.Password,
	}
	err = newUser.Create(a.Db)
	if err != nil {
		slog.Error("Creating User", "err", err.Error())
		return err
	}
	slog.Info("User Created", "id", newUser.Id)
	return nil
}

func (a *App) UpdateUser(user models.User) error {
	err := user.Update(a.Db)
	if err != nil {
		slog.Error("Updating User", "err", err.Error())
		return err
	}
	slog.Info("User Updated", "id", user.Id)
	return nil
}

func (a *App) DeleteUserById(id string) error {
	user, err := models.GetUserById(a.Db, id)
	if err != nil {
		slog.Error("Getting User", "err", err.Error())
		return err
	}
	err = user.Delete(a.Db)
	if err != nil {
		slog.Error("Deleting User", "err", err.Error())
		return err
	}
	slog.Info("User Deleted", "id", user.Id)
	return nil
}

func (a *App) GetUserById(id string) (models.User, error) {
	user, err := models.GetUserById(a.Db, id)
	if err != nil {
		slog.Error("Getting User", "err", err.Error())
		return models.User{}, err
	}
	slog.Info("User Retrieved", "id", user.Id)
	return user, nil
}

func (a *App) GetUserByEmail(email string) (models.User, error) {
	user, err := models.GetUserByEmail(a.Db, email)
	if err != nil {
		slog.Error("Getting User", "err", err.Error())
		return models.User{}, err
	}
	slog.Info("User Retrieved", "id", user.Id)
	return user, nil
}

func (a *App) GetUsers() ([]models.User, error) {
	users, err := models.GetAllUsers(a.Db)
	if err != nil {
		slog.Error("Getting Users", "err", err.Error())
		return nil, err
	}
	slog.Info("Users Retrieved", "count", len(users))
	return users, nil
}
