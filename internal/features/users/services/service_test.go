package services_test

import (
	"newsapps/internal/features/users"
	"newsapps/internal/features/users/services"
	"newsapps/mocks"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func ImportMocks(t *testing.T) (*mocks.UQuery, *mocks.PassUtilInterface, *mocks.TokenUtilInterface, users.UServices) {
	qry := mocks.NewUQuery(t)
	pu := mocks.NewPassUtilInterface(t)
	tu := mocks.NewTokenUtilInterface(t)
	srv := services.NewUsersSrv(qry, pu, tu)

	return qry, pu, tu, srv
}

func TestLogin(t *testing.T) {
	qry, pu, tu, srv := ImportMocks(t)

	input := users.Users{
		Username: "admin",
		Password: "hashed_admin",
	}

	loginData := users.Users{
		ID:       1,
		Username: "admin",
		Password: "hashed_admin",
		Email:    "admin@gmail.com",
		Phone:    "089123456879",
	}

	t.Run("Query Error", func(t *testing.T) {
		qry.On("Login", input.Username).Return(users.Users{}, gorm.ErrInvalidData).Once()
		_, err := srv.Login(input.Username, input.Password)

		qry.AssertExpectations(t)
		assert.Error(t, err)
	})

	t.Run("Password Error", func(t *testing.T) {
		qry.On("Login", input.Username).Return(loginData, nil).Once()
		pu.On("ComparePassword", []byte(loginData.Password), []byte("hashed_admin")).Return(bcrypt.ErrMismatchedHashAndPassword).Once()
		_, err := srv.Login(input.Username, input.Password)

		qry.AssertExpectations(t)
		pu.AssertExpectations(t)
		assert.Error(t, err)
	})

	t.Run("Token Error", func(t *testing.T) {
		qry.On("Login", input.Username).Return(loginData, nil).Once()
		pu.On("ComparePassword", []byte(loginData.Password), []byte("hashed_admin")).Return(nil).Once()
		tu.On("GenerateToken", loginData).Return("", jwt.ErrTokenInvalidClaims).Once()
		_, err := srv.Login(input.Username, input.Password)

		qry.AssertExpectations(t)
		pu.AssertExpectations(t)
		assert.Error(t, err)
	})

	t.Run("Login Success", func(t *testing.T) {
		token := "hasil_token"

		qry.On("Login", input.Username).Return(loginData, nil).Once()
		pu.On("ComparePassword", []byte(loginData.Password), []byte("hashed_admin")).Return(nil).Once()
		tu.On("GenerateToken", loginData).Return(token, nil).Once()

		_, err := srv.Login(input.Username, input.Password)

		qry.AssertExpectations(t)
		assert.Nil(t, err)
	})
}

func TestCreateUser(t *testing.T) {
	qry, pu, _, srv := ImportMocks(t)

	input := users.Users{
		ID:       1,
		Username: "admin",
		Password: "admin",
		Email:    "admin@gmail.com",
		Phone:    "089123456879",
	}

	t.Run("Password Error", func(t *testing.T) {
		pu.On("GeneratePassword", input.Password).Return(nil, bcrypt.ErrPasswordTooLong).Once()
		err := srv.CreateUser(input)

		pu.AssertExpectations(t)
		assert.Error(t, err)
	})

	t.Run("Query Error", func(t *testing.T) {
		inputQry := users.Users{
			ID:       1,
			Username: "admin",
			Password: "hashed_admin",
			Email:    "admin@gmail.com",
			Phone:    "089123456879",
		}
		pu.On("GeneratePassword", input.Password).Return([]byte("hashed_admin"), nil).Once()
		qry.On("CreateUser", inputQry).Return(gorm.ErrInvalidData).Once()
		err := srv.CreateUser(input)

		pu.AssertExpectations(t)
		qry.AssertExpectations(t)
		assert.Error(t, err)
	})

	t.Run("Register Success", func(t *testing.T) {
		inputQry := users.Users{
			ID:       1,
			Username: "admin",
			Password: "hashed_admin",
			Email:    "admin@gmail.com",
			Phone:    "089123456879",
		}
		pu.On("GeneratePassword", input.Password).Return([]byte("hashed_admin"), nil).Once()
		qry.On("CreateUser", inputQry).Return(nil).Once()
		err := srv.CreateUser(input)

		pu.AssertExpectations(t)
		qry.AssertExpectations(t)
		assert.Nil(t, err)
	})
}

func TestReadUser(t *testing.T) {
	qry, _, _, srv := ImportMocks(t)

	var userID uint = 1

	t.Run("Query Error", func(t *testing.T) {
		qry.On("ReadUser", userID).Return(users.Users{}, gorm.ErrInvalidData).Once()
		_, err := srv.ReadUser(userID)

		qry.AssertExpectations(t)
		assert.Error(t, err)
	})

	t.Run("Read User Success", func(t *testing.T) {
		qry.On("ReadUser", userID).Return(users.Users{}, nil).Once()
		_, err := srv.ReadUser(userID)

		qry.AssertExpectations(t)
		assert.Nil(t, err)
	})
}

func TestUpdateUser(t *testing.T) {
	qry, pu, _, srv := ImportMocks(t)

	var userID uint = 1
	input := users.Users{
		Username: "admin",
		Password: "admin",
		Email:    "admin@gmail.com",
		Phone:    "089123456879",
	}

	t.Run("Password Error", func(t *testing.T) {
		pu.On("GeneratePassword", input.Password).Return(nil, bcrypt.ErrPasswordTooLong).Once()

		_, err := pu.GeneratePassword(input.Password)

		pu.AssertExpectations(t)
		assert.Error(t, err)
	})

	t.Run("Query Error", func(t *testing.T) {
		pu.On("GeneratePassword", input.Password).Return([]byte("admin"), nil).Once()
		qry.On("UpdateUser", userID, input).Return(gorm.ErrInvalidData).Once()

		err := srv.UpdateUser(userID, input)

		pu.AssertExpectations(t)
		qry.AssertExpectations(t)
		assert.Error(t, err)
	})

	t.Run("Update User Success", func(t *testing.T) {
		pu.On("GeneratePassword", input.Password).Return([]byte("admin"), nil).Once()
		qry.On("UpdateUser", userID, input).Return(nil).Once()

		err := srv.UpdateUser(userID, input)

		pu.AssertExpectations(t)
		qry.AssertExpectations(t)
		assert.Nil(t, err)
	})
}

func TestDeleteUser(t *testing.T) {
	qry, _, _, srv := ImportMocks(t)

	var userID uint = 1

	t.Run("Query Error", func(t *testing.T) {
		qry.On("DeleteUser", userID).Return(gorm.ErrInvalidData).Once()

		err := srv.DeleteUser(userID)

		qry.AssertExpectations(t)
		assert.Error(t, err)
	})

	t.Run("Delete User Success", func(t *testing.T) {
		qry.On("DeleteUser", userID).Return(nil).Once()

		err := srv.DeleteUser(userID)

		qry.AssertExpectations(t)
		assert.Nil(t, err)
	})
}
