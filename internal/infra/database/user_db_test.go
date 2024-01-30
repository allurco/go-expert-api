package database

import (
	"testing"

	"github.com/allurco/go-expert-api/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	db, err := CreateConnection(&entity.User{})
	if err != nil {
		t.Error(err)
	}

	user, _ := entity.NewUser("jon", "j@j.com", "123456")
	userDB := NewUser(db)

	err = userDB.CreateUser(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)

}

func TestGetUserByEmail(t *testing.T) {
	db, err := CreateConnection(&entity.User{})
	if err != nil {
		t.Error(err)
	}
	user, _ := entity.NewUser("jon", "bolas@gmail.com", "123456")
	userDB := NewUser(db)

	err = userDB.CreateUser(user)
	assert.Nil(t, err)

	userFound, err := userDB.GetUserByEmail("bolas@gmail.com")
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
}
