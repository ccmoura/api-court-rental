package models

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Owner struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	AreaCode string `json:"area_code"`
	Password string `json:"password"`
	Cpf string `json:"cpf"`
	IsConfirmed bool `json:"is_confirmed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (owner *Owner) Prepare() {
	owner.Name = owner.Name
	owner.Email = owner.Email
	owner.Phone = owner.Phone
	owner.AreaCode = owner.AreaCode
	owner.Cpf = owner.Cpf
	owner.IsConfirmed = false
}

func (owner *Owner) BeforeSave() error {
	hashedPassword, err := Hash(owner.Password)
	if err != nil {
		return err
	}
	owner.Password = string(hashedPassword)
	return nil
}

func (owner *Owner) SaveOwner(db *gorm.DB) (*Owner, error) {
	owner.BeforeSave()

	var err error
	err = db.Debug().Create(&owner).Error
	if err != nil {
		return &Owner{}, err
	}
	return owner, nil
}

func (owner *Owner) DeleteOwner(db *gorm.DB, id string) (int64, error) {
	db = db.Debug().Model(&Owner{}).Where("id = ?", id).Take(&Owner{}).Delete(&Owner{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
