package controllers

import (
	"changeme/api/auth"
	"changeme/api/models"
)


func (a *App) RegisterUser(user models.User) error {
	ok, err := user.Verify()
	if !ok {
		return err
	}
	err = user.Create(a.Db)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) LoginUser(user models.User) (string, error) {
	err := user.VerifyLogin()
	if err != nil {
		return "", err
	}
	u, err := models.GetUserByEmail(a.Db, user.Email)
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
