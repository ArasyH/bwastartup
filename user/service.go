package user

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
	GetUserByID(ID int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	user.Created_at = time.Now()
	user.Updated_at = time.Now()
	//var PasswordHash []byte
	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password_hash = string(PasswordHash)
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.Id == 0 {
		return user, errors.New("User tidak ditemukan")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password_hash), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}
	// check logic
	if user.Id == 0 {
		return true, nil
	}
	return false, nil
}

func (s *service) SaveAvatar(ID int, fileLocation string) (User, error) {
	//1. dapatkan user berdasarkan id
	//2. update atribut avatar file name
	//3. simpan perubahan avatar file name
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	user.Avatar_file_name = fileLocation
	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil

}

func (s *service) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.Id == 0 {
		return user, errors.New("no user found on with that id")
	}

	return user, nil
}

//MAPPING STRUCT INPUT KE STRUCT USER
//SIMPAN STRUCT USER MELALUI repository
