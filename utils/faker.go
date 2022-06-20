package utils

import (
	"github.com/bxcodec/faker/v3"
	"github.com/nadirbasalamah/go-simple-inventory/models"
)

func CreateItemFaker() (models.Item, error) {
	var item models.Item = models.Item{}
	err := faker.FakeData(&item)
	if err != nil {
		return models.Item{}, err
	}

	return item, nil
}

func CreateUserFaker() (models.User, error) {
	var user models.User = models.User{}
	err := faker.FakeData(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
