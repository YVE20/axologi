package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customers struct {
	CustomerID    string     `gorm:"type:varchar(36);primary_key" json:"customerID"`
	FirstName     string     `gorm:"type:varchar(30);default:null" json:"firstName"`
	LastName      string     `gorm:"type:varchar(30);default:null" json:"lastName"`
	Email         string     `gorm:"type:varchar(30);default:null" json:"email"`
	Password      string     `gorm:"type:varchar(100);default:null" json:"password"`
	DeviceID      string     `gorm:"type:varchar(36);default:null" json:"deviceID"`
	IpAddress     string     `gorm:"type:varchar(36);default:null" json:"ipAddress"`
	BirthDate     *time.Time `gorm:"type:datetime" json:"birthDate"`
	Photo         string     `gorm:"type:varchar(100);default:null" json:"photo"`
	StatusAccount string     `gorm:"type:varchar(30);default:Not Active" json:"statusAccount"`
	BuyerID       string     `gorm:"type:varchar(36);default:null" json:"buyerID"`
	SellerID      string     `gorm:"type:varchar(36);default:null" json:"sellerID"`
	CreatedBy     string     `gorm:"type:varchar(30)" json:"createdBy"`
	CreatedDate   *time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"createdDate"`
	UpdatedBy     string     `gorm:"type:varchar(30)" json:"updatedBy"`
	UpdatedDate   *time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"updatedDate"`
}

// BeforeCreate hook will be triggered before the record is created in the database
func (c *Customers) BeforeCreate(tx *gorm.DB) (err error) {
	c.CustomerID = uuid.New().String() // Generate a new UUID for CustomerID
	return
}
