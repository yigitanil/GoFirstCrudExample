package model

import (
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	ID          int
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
	CreateDate  time.Time
	UpdateDate  time.Time
}

type CustomerDto struct {
	ID          int          `json:"id"`
	FirstName   string       `json:"firstName"`
	LastName    string       `json:"lastName"`
	PhoneNumber string       `json:"phoneNumber"`
	Email       string       `json:"email"`
	Addresses   []AddressDto `json:"addresses"`
}

func (c Customer) Save(db *gorm.DB) (error, int) {
	tx := db.Debug().Save(&c)
	return tx.Error, c.ID
}

func SaveAllCustomers(db *gorm.DB, customers []Customer) (error, []int) {
	tx := db.Debug().Save(&customers)
	ids := make([]int, len(customers))
	for i, c := range customers {
		ids[i] = c.ID
	}
	return tx.Error, ids
}
func FindCustomerById(db *gorm.DB, id int) (error, *Customer) {
	c := new(Customer)
	tx := db.Debug().Find(c, id)
	if tx.RowsAffected == 0 {
		c = nil
	}
	return tx.Error, c
}
func FindCustomerByEmail(db *gorm.DB, email string) (error, *Customer) {
	c := new(Customer)
	tx := db.Debug().Where("email=?", email).Find(c)
	if tx.RowsAffected == 0 {
		c = nil
	}
	return tx.Error, c
}
func FindAllCustomers(db *gorm.DB, limit int, offset int) (error, *[]Customer) {
	customers := new([]Customer)
	tx := db.Debug().Limit(limit).Offset(offset).Find(&customers)
	return tx.Error, customers
}

func DeleteCustomerById(db *gorm.DB, id int) error {
	err, customer := FindCustomerById(db, id)
	if err == nil {
		tx := db.Debug().Delete(customer)
		return tx.Error
	}
	return err
}
