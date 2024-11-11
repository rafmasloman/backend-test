package services

import (
	"backend/api/src/config"
	"backend/api/src/models"
	"fmt"
)

func GetAllCustomers() ([]models.Customer, error) {
	var customers []models.Customer
	
	fmt.Println("all customers : ", customers)
	result := config.DB.Find(&customers)



	fmt.Println("customers : ", customers)

	if result.Error != nil {
		return nil, result.Error
	}

	return customers, nil
}

func GetDetailCustomer(customerId int) (*models.Customer, error) {
	var customer models.Customer

	fmt.Println("customer detail 1 : ", customer)
	findcustomer := config.DB.Where("id = ?", customerId).First(&customer)

	
	if findcustomer.Error != nil {
		return nil, findcustomer.Error
	}

	fmt.Println("customer detail 2 : ", customer)

	return &customer, nil
}

func CreateCustomer(customer *models.Customer) (*models.Customer, error) {
	result := config.DB.Create(customer)

	if result.Error != nil {
		return nil, result.Error
	}

	return customer, nil
}

func UpdateCustomer(customerId int, payload models.Customer) (*models.Customer, error) {

	var customer models.Customer

	findCustomer := config.DB.Where("id = ?",customerId).First(&customer)

	if findCustomer.Error != nil {
		return nil, findCustomer.Error
	}

	body := models.Customer{Name: payload.Name, BirthDate: payload.BirthDate,Address: payload.Address,Phone: payload.Phone}


	updateCustomer := config.DB.Model(&customer).Updates(body)

	if updateCustomer.Error != nil {
		return nil, updateCustomer.Error
	}

	fmt.Println("customer update : ", customer)

	return &customer, nil
}

func DeleteCustomer(customerId int) (*models.Customer, error) {

	var customer models.Customer

	findCustomer := config.DB.Where("id = ?",customerId).First(&customer)

	if findCustomer.Error != nil {
		return nil, findCustomer.Error
	}
	
	result := config.DB.Delete(&customer, customerId)

	if result.Error != nil {
		fmt.Println("error : ", result.Error)
		return nil, result.Error
	}

	return &customer, nil
}