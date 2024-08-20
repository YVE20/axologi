package model

//It's simmiliar to Model to wrap the data
type ProvinceModel struct {
	ProvinceID string `json:"provinceID"`
	Province   string `json:"province"`
	CreatedBy  string `json:"createdBy"`
	UpdatedBy  string `json:"updatedBy"`
}
