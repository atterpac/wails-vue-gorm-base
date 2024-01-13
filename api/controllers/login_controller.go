package controllers

import (
	"changeme/api/auth"
	"changeme/api/models"
	"log/slog"
)


func (a *App) RegisterUser(reg models.Register) error {
	err := reg.Verify()
	if err != nil {
		return err
	}
	newUser := models.User {
		Username: reg.Username,
		Password: reg.Password,
		Email: reg.Email,
	}
	err = newUser.Create(a.Db)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) LoginUser(login models.Login) error {
	slog.Debug("Attempting to LoginUser", "username", login.Username)
	err := login.Verify()
	if err != nil {
		return err
	}
	slog.Debug("User Verified", "username", login.Username)
	u, err := models.GetUserByUsername(a.Db, login.Username)
	if err != nil {
		slog.Error("Getting User", "err", err.Error(), "username", login.Username)
		return err
	}
	err = auth.VerifyPassword(u.Password, login.Password)
	if err != nil {
		slog.Error("Verifying Password", "err", err.Error(), "username", login.Username)
		return err
	}
	token, err := auth.CreateToken(u.Id)
	if err != nil {
		slog.Error("Creating Token", "err", err.Error(), "username", login.Username)
		return err
	}
	u.Jwt = token
	err = u.Update(a.Db)
	if err != nil {
		slog.Error("Updating User", "err", err.Error())
		return err
	}
	slog.Debug("User Logged In", "username", login.Username)
	return nil
}

func (a *App) RefreshToken(staleToken string) (string, error) {
	token, err := auth.RefreshToken(staleToken)
	if err != nil {
		return "", err
	}
	return token, nil
}
