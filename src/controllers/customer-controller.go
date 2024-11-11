package controllers

import (
	"backend/api/src/models"
	"backend/api/src/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCustomers(c *gin.Context) {
	customers, err := services.GetAllCustomers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"message": "server error",
			"error": err.Error(),
		})

		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "success fetched all customers",
		"data": customers,
	})
}

func GetDetailCustomer(c *gin.Context) {
	customerId := c.Param("id")
	id, err := strconv.Atoi(customerId)

	fmt.Println("err : ", err)

	customer, err := services.GetDetailCustomer(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"message": "customer not found",
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "customer found succesfully",
		"data": customer,
	})

}

func CreateCustomer(c *gin.Context) {
	var customerInput models.Customer
	

	fmt.Println("customer : ", c.Request.Body)
	if err := c.ShouldBindJSON(&customerInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"message": "customer created failed",
			"error": err.Error()})
		return
	}

	customer := models.Customer{Name: customerInput.Name, BirthDate: customerInput.BirthDate,Address: customerInput.Address, Phone: customerInput.Phone}

	customers, err := services.CreateCustomer(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Customer creation failed",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Customer created succesfully",
		"data": customers,
	})
}

func UpdateCustomer(c *gin.Context) {

	var customerInput models.Customer

	customerId := c.Param("id")
	id, err := strconv.Atoi(customerId)

	fmt.Println(err)

	if err := c.ShouldBindJSON(&customerInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid input data",
			"error":   err.Error(),
		})
		return
	}

	customer, errUpdate := services.UpdateCustomer(id, customerInput)

	if errUpdate != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Customer update failed",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Customer updated succesfully",
		"data": customer,
	})
}

func DeleteCustomer(c *gin.Context) {
	customerId := c.Param("id")
	id, err := strconv.Atoi(customerId)

	fmt.Println(err)

	customer, err := services.DeleteCustomer(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "success",
			"statusCode": http.StatusInternalServerError,
			"message": "customer deleted failed",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status" : "success",
		"statusCode": http.StatusOK,
		"message": "customer deleted successfully",
		"data": customer,
	})
}