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

func SaveCustomer(db *gorm.DB, c echo.Context) error {
	customerDto := new(model.CustomerDto)
	if err := c.Bind(customerDto); err != nil {
		return err
	}

	err, savedCustomer := model.FindCustomerByEmail(db, customerDto.Email)
	if err != nil {
		return err
	}
	if savedCustomer != nil {
		businessError := errors.BusinessError{
			Reason: "Customer with given email already exists",
		}
		return businessError
	}

	customer := model.Customer{
		FirstName:   customerDto.FirstName,
		LastName:    customerDto.LastName,
		PhoneNumber: customerDto.PhoneNumber,
		Email:       customerDto.Email,
		CreateDate:  time.Now(),
		UpdateDate:  time.Now(),
	}
	err, customerId := customer.Save(db)
	if err != nil {
		return err
	}
	if customerDto.Addresses != nil && len(customerDto.Addresses) > 0 {
		addresses := make([]model.Address, len(customerDto.Addresses))
		for i, address := range customerDto.Addresses {
			addresses[i] = model.Address{
				CustomerId: customerId,
				Address:    address.Address,
				City:       address.City,
				Country:    address.Country,
				CreateDate: time.Now(),
				UpdateDate: time.Now(),
			}
		}
		err, _ := model.SaveAllAddresses(db, addresses)
		if err != nil {
			err := model.DeleteCustomerById(db, customerId)
			return err
		}
	}

	return c.NoContent(http.StatusCreated)
}

func UpdateCustomer(db *gorm.DB, c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	err, customer := model.FindCustomerById(db, id)
	if err != nil {
		return err
	}

	if customer == nil {
		e := model.Error{
			Message: "No customer found for id: " + idStr,
		}
		return c.JSON(http.StatusBadRequest, e)
	}

	customerBody := new(model.Customer)
	if err := c.Bind(customerBody); err != nil {
		return err
	}

	customer.FirstName = customerBody.FirstName
	customer.LastName = customerBody.LastName
	customer.PhoneNumber = customerBody.PhoneNumber
	customer.Email = customerBody.Email
	customer.UpdateDate = time.Now()

	err, _ = customer.Save(db)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func FindCustomerById(db *gorm.DB, c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	err, customer := model.FindCustomerById(db, id)
	if err != nil {
		return err
	}
	if customer == nil {
		return c.NoContent(http.StatusNoContent)
	}

	err, addresses := model.FindAllAddresses(db, customer.ID)

	customerDto := model.CustomerDto{
		ID:          customer.ID,
		FirstName:   customer.FirstName,
		LastName:    customer.LastName,
		PhoneNumber: customer.PhoneNumber,
		Email:       customer.Email,
		Addresses:   make([]model.AddressDto, 0),
	}

	if addresses != nil && len(*addresses) > 0 {
		addressDtos := make([]model.AddressDto, len(*addresses))
		for i, address := range *addresses {
			addressDtos[i] = model.AddressDto{
				ID:      address.ID,
				Address: address.Address,
				City:    address.City,
				Country: address.Country,
			}
		}
		customerDto.Addresses = addressDtos
	}
	return c.JSON(http.StatusOK, customerDto)

}

func DeleteCustomerById(db *gorm.DB, c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	err = model.DeleteAddressesByCustomerId(db, id)
	err = model.DeleteCustomerById(db, id)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func FindCustomers(db *gorm.DB, c echo.Context) error {
	offsetStr := c.QueryParam("offset")
	limitStr := c.QueryParam("limit")
	if len(offsetStr) == 0 {
		offsetStr = "0"
	}
	if len(limitStr) == 0 {
		limitStr = "10"
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return err
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return err
	}
	err, customers := model.FindAllCustomers(db, limit, offset)
	if err != nil {
		return err
	}
	customerDtos := make([]model.CustomerDto, len(*customers))
	for i, customer := range *customers {
		customerDtos[i] = model.CustomerDto{
			ID:          customer.ID,
			FirstName:   customer.FirstName,
			LastName:    customer.LastName,
			PhoneNumber: customer.PhoneNumber,
			Email:       customer.Email,
		}
	}

	return c.JSON(http.StatusOK, customerDtos)
}
