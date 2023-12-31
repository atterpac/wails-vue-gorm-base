package controllers

import "changeme/api/models"

func (a *App) CreateUser(user models.User) error {
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

func (a *App) UpdateUser(user models.User) error {
	err := user.Update(a.Db)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) DeleteUserById(id string) error {
	user, err := models.GetUserById(a.Db, id)
	if err != nil {
		return err
	}
	err = user.Delete(a.Db)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) GetUserById(id string) (models.User, error) {
	user, err := models.GetUserById(a.Db, id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (a *App) GetUserByEmail(email string) (models.User, error) {
	user, err := models.GetUserByEmail(a.Db, email)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (a *App) GetUsers() ([]models.User, error) {
	users, err := models.GetAllUsers(a.Db)
	if err != nil {
		return nil, err
	}
	return users, nil
}
