package model

import (
	"gorm.io/gorm"
	"time"
)

type Address struct {
	ID         int
	CustomerId int
	Address    string
	City       string
	Country    string
	CreateDate time.Time
	UpdateDate time.Time
}

type AddressDto struct {
	ID      int    `json:"id"`
	Address string `json:"address"`
	City    string `json:"city"`
	Country string `json:"country"`
}

func (a Address) Save(db *gorm.DB) (error, int) {
	tx := db.Debug().Save(&a)
	return tx.Error, a.ID
}

func SaveAllAddresses(db *gorm.DB, addresses []Address) (error, []int) {
	tx := db.Debug().Save(&addresses)
	ids := make([]int, len(addresses))
	for i, c := range addresses {
		ids[i] = c.ID
	}
	return tx.Error, ids
}
func FindAddressById(db *gorm.DB, id int) (error, *Address) {
	a := new(Address)
	tx := db.Debug().Find(a, id)
	if tx.RowsAffected == 0 {
		a = nil
	}
	return tx.Error, a
}
func FindAllAddresses(db *gorm.DB, customerId int) (error, *[]Address) {
	addresses := new([]Address)
	tx := db.Debug().Where("customer_id=?", customerId).Find(&addresses)
	return tx.Error, addresses
}

func DeleteAddressById(db *gorm.DB, id int) error {
	err, address := FindAddressById(db, id)
	if err == nil {
		tx := db.Debug().Delete(address)
		return tx.Error
	}
	return err
}
func DeleteAddressesByCustomerId(db *gorm.DB, customerId int) error {
	tx := db.Debug().Where("customer_id=?", customerId).Delete(Address{})
	return tx.Error
}
