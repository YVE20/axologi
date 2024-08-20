package service

import (
	"axologi/model"
	"axologi/model/entity"
	"axologi/model/repo"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type LoginService interface {
	LoginService(LoginModel model.LoginModel) ([]model.LoginModel, error)
	RegisterService(RegisterModel model.RegisterModel) ([]model.RegisterModel, error)
}

type AccountStructModel struct {
	dataLogin    []model.LoginModel
	dataRegister []model.RegisterModel
}

// Construct
func NewLoginService() LoginService {
	return &AccountStructModel{}
}

func (s *AccountStructModel) LoginService(LoginModel model.LoginModel) ([]model.LoginModel, error) {

	//check data DB
	var loginEntity entity.Customers
	err := repo.DB.Model(&entity.Customers{}).Where("email = ?", LoginModel.Email).First(&loginEntity).Error
	if err != nil {
		return nil, err
	}

	// Compare password input (plaintext) dengan password hash yang disimpan di database
	err = bcrypt.CompareHashAndPassword([]byte(loginEntity.Password), []byte(LoginModel.Password))

	if err != nil {
		return nil, err
	}

	//Success
	loginModel := model.LoginModel{
		Email:    loginEntity.Email,
		Password: loginEntity.Password,
	}
	return []model.LoginModel{loginModel}, nil
}

func (s *AccountStructModel) RegisterService(RegisterModel model.RegisterModel) ([]model.RegisterModel, error) {

	//Check firstName, lastName and email
	var count int64
	repo.DB.Model(&entity.Customers{}).Where("first_name= ?", RegisterModel.FirstName).Count(&count)
	isFirstNameRegistered := count > 0

	count = 0 // Reset count
	repo.DB.Model(&entity.Customers{}).Where("last_name = ?", RegisterModel.LastName).Count(&count)
	isLastNameRegistered := count > 0

	count = 0 // Reset count
	repo.DB.Model(&entity.Customers{}).Where("email = ?", RegisterModel.Email).Count(&count)
	isEmailRegistered := count > 0

	var errorMessage string
	if isFirstNameRegistered {
		errorMessage += "First name is already registered. "
	}
	if isLastNameRegistered {
		errorMessage += "Last name is already registered. "
	}
	if isEmailRegistered {
		errorMessage += "Email is already registered. "
	}

	if errorMessage != "" {
		return nil, errors.New(errorMessage)
	}

	var RegisterEntity entity.Customers
	//Hasing password
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(RegisterModel.Password), bcrypt.DefaultCost)
	RegisterModel.Password = string(hashPassword)

	//Populate DataRegister
	RegisterEntity = PopulateDataRegister(RegisterModel)

	// Insert ke database
	if err := repo.DB.Create(&RegisterEntity).Error; err != nil {
		return nil, err // Mengembalikan error jika insert ke database gagal
	}

	s.dataRegister = append(s.dataRegister, RegisterModel)

	return s.dataRegister, nil

}

func PopulateDataRegister(RegisterModel model.RegisterModel) (RegisterEntity entity.Customers) {
	var registerEntity entity.Customers
	registerEntity.FirstName = RegisterModel.FirstName
	registerEntity.LastName = RegisterModel.LastName
	registerEntity.Email = RegisterModel.Email
	registerEntity.Password = RegisterModel.Password
	registerEntity.BirthDate = RegisterModel.BirthDate
	registerEntity.DeviceID = RegisterModel.DeviceID
	registerEntity.IpAddress = RegisterModel.IpAddress
	registerEntity.Photo = RegisterModel.Photo
	registerEntity.BuyerID = RegisterModel.BuyerID
	registerEntity.SellerID = RegisterModel.SellerID
	registerEntity.CreatedBy = RegisterModel.Email
	registerEntity.UpdatedBy = RegisterModel.Email

	return registerEntity
}
