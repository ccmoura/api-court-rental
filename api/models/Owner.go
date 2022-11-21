package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Owner struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Email       string       `json:"email"`
	Phone       string       `json:"phone"`
	AreaCode    string       `json:"area_code"`
	Password    string       `json:"password"`
	Cpf         string       `json:"cpf"`
	IsConfirmed bool         `json:"is_confirmed"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
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

func (owner *Owner) FindOwnerById(db *gorm.DB, id string) (*Owner, error) {
	db = db.Debug().Model(&Owner{}).Where("id = ?", id).First(&owner)

	if db.Error != nil {
		return nil, db.Error
	}
	return owner, nil
}

func (owner *Owner) FindDuplicatedEmail(db *gorm.DB) error {
	var result Owner
	db = db.Debug().Model(&Owner{}).Where("email = ?", owner.Email).First(&result)

	if db.Error != nil {
		return nil
	}

	return errors.New("Email already exists")
}

func (owner *Owner) FindDuplicatedCpf(db *gorm.DB) error {
	var result Owner
	db = db.Debug().Model(&Owner{}).Where("cpf = ?", owner.Cpf).First(&result)

	if db.Error != nil {
		return nil
	}

	return errors.New("Cpf already exists")
}

func (owner *Owner) FindDuplicatedPhone(db *gorm.DB, except string) error {
	var result Owner
	if len(except) == 0 {
		db = db.Debug().Model(&Owner{}).Where("phone = ?", owner.Phone).First(&result)
	} else {
		db = db.Debug().Model(&Owner{}).Where("phone = ? AND id != ?", owner.Phone, except).First(&result)
	}

	if db.Error != nil {
		return nil
	}

	return errors.New("Phone already exists")
}

func (owner *Owner) SaveOwner(db *gorm.DB) (*Owner, error) {
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

func (owner *Owner) UpdateOwner(db *gorm.DB, uid string) (*Owner, error) {
	err := owner.BeforeSave()
	if err != nil {
		return &Owner{}, err
	}

	db = db.Debug().Model(&Owner{}).Where("id = ?", uid).Take(&Owner{}).UpdateColumns(
		map[string]interface{}{
			"password":   owner.Password,
			"name":       owner.Name,
			"phone":      owner.Phone,
			"area_code":  owner.AreaCode,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Owner{}, db.Error
	}

	err = db.Debug().Model(&Owner{}).Where("id = ?", uid).Take(&owner).Error
	if err != nil {
		return &Owner{}, err
	}

	return owner, nil
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
