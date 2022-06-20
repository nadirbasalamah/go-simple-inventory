package utils

import (
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/nadirbasalamah/go-simple-inventory/models"
)

func CreateItemFaker() (models.Item, error) {
	type ItemData struct {
		ID       string `faker:"uuid_hyphenated"`
		Name     string `faker:"name"`
		Price    int    `faker:"oneof: 15, 27, 61"`
		Quantity int    `faker:"oneof: 15, 27, 61"`
	}

	itemData := ItemData{}
	err := faker.FakeData(&itemData)
	if err != nil {
		return models.Item{}, err
	}

	var newItem models.Item = models.Item{
		ID:        itemData.ID,
		Name:      itemData.Name,
		Price:     itemData.Price,
		Quantity:  itemData.Quantity,
		CreatedAt: time.Now(),
	}

	return newItem, nil
}

func CreateUserFaker() (models.User, error) {
	type UserData struct {
		ID       string `faker:"uuid_hyphenated"`
		Email    string `faker:"email"`
		Password string `faker:"password"`
	}

	userData := UserData{}
	err := faker.FakeData(&userData)
	if err != nil {
		return models.User{}, err
	}

	var newUser models.User = models.User{
		ID:       userData.ID,
		Email:    userData.Email,
		Password: userData.Password,
	}

	return newUser, nil
}
