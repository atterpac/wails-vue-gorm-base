package controllers

import (
	"changeme/api/auth"
	"changeme/api/models"
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

func (a *App) LoginUser(login models.Login) (string, error) {
	err := login.Verify()
	if err != nil {
		return "", err
	}
	u, err := models.GetUserByUsername(a.Db, login.Username)
	if err != nil {
		return "", err
	}
	err = auth.VerifyPassword(u.Password, login.Password)
	if err != nil {
		return "", err
	}
	token, err := auth.CreateToken(u.Id)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *App) RefreshToken(staleToken string) (string, error) {
	token, err := auth.RefreshToken(staleToken)
	if err != nil {
		return "", err
	}
	return token, nil
}
