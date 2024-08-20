package service

import "axologi/model"

//Interface supaya bisa diakses yang lain
type ProvinceService interface {
	GetAllProvinceService() ([]model.ProvinceModel, error)
}

//Return dari service ini, bisa isi apapun
type ProvinceStructModel struct {
	DataProvinces []model.ProvinceModel
}

//Construct untuk execute service ini di main.go
func NewProvinceService() ProvinceService {
	return &ProvinceStructModel{}
}

func (s ProvinceStructModel) GetAllProvinceService() ([]model.ProvinceModel, error) {
	return s.DataProvinces, nil
}
