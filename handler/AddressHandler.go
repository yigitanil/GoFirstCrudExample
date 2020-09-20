package handler

import (
	errors "GoFirstCrudExample/error"
	"GoFirstCrudExample/model"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

func SaveAddress(db *gorm.DB, c echo.Context) error {
	addressDto := new(model.AddressDto)
	if err := c.Bind(addressDto); err != nil {
		return err
	}

	idStr := c.Param("id")
	customerId, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	err, customer := model.FindCustomerById(db, customerId)
	if err != nil {
		return err
	}

	if customer == nil {
		return errors.BusinessError{Reason: "No customer found for id: " + idStr}
	}

	address := model.Address{
		CustomerId: customerId,
		Address:    addressDto.Address,
		City:       addressDto.City,
		Country:    addressDto.Country,
		CreateDate: time.Now(),
		UpdateDate: time.Now(),
	}

	err, _ = address.Save(db)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func UpdateAddress(db *gorm.DB, c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	err, address := model.FindAddressById(db, id)
	if err != nil {
		return err
	}

	if address == nil {
		e := model.Error{
			Message: "No address found for id: " + idStr,
		}
		return c.JSON(http.StatusBadRequest, e)
	}

	addressBody := new(model.Address)
	if err := c.Bind(addressBody); err != nil {
		return err
	}

	address.Address = addressBody.Address
	address.Country = addressBody.Country
	address.City = addressBody.City
	address.UpdateDate = time.Now()

	err, _ = address.Save(db)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func FindAddressById(db *gorm.DB, c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	err, address := model.FindAddressById(db, id)
	if err != nil {
		return err
	}
	if address != nil {
		addressDto := model.AddressDto{
			ID:      address.ID,
			Address: address.Address,
			City:    address.City,
			Country: address.Country,
		}
		return c.JSON(http.StatusOK, addressDto)
	}
	return c.NoContent(http.StatusNoContent)
}

func DeleteAddressById(db *gorm.DB, c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	err = model.DeleteAddressById(db, id)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func FindAddresses(db *gorm.DB, c echo.Context) error {
	idStr := c.Param("id")
	customerId, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	err, addresses := model.FindAllAddresses(db, customerId)
	if err != nil {
		return err
	}
	addressDtos := make([]model.AddressDto, len(*addresses))
	for i, address := range *addresses {
		addressDtos[i] = model.AddressDto{
			ID:      address.ID,
			Address: address.Address,
			City:    address.City,
			Country: address.Country,
		}
	}

	return c.JSON(http.StatusOK, addressDtos)
}
