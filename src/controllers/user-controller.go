package controllers

import (
	"backend/api/src/models"
	"backend/api/src/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	users, err := services.GetAllUsers()

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
		"message": "success fetched all users",
		"data": users,
	})
}

func GetDetailUser(c *gin.Context) {
	userId := c.Param("id")
	id, err := strconv.Atoi(userId)

	fmt.Println("err : ", err)

	user, err := services.GetDetailUser(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"message": "user not found",
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"status": "success",
		"message": "user found succesfully",
		"data": user,
	})

}

func CreateUser(c *gin.Context) {
	var userInput models.User

	fmt.Println("user : ", c.Request.Body)
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"message": "user created failed",
			"error": err.Error()})
		return
	}

	user := models.User{Name: userInput.Name, BirthDate: userInput.BirthDate,Address: userInput.Address, Phone: userInput.Phone}

	users, err := services.CreateUser(&user)
	if err != nil {
		// Jika terjadi error saat menyimpan, kirim response error
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "User creation failed",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created succesfully",
		"data": users,
	})
}

func UpdateUser(c *gin.Context) {

	var userInput models.User

	userId := c.Param("id")
	id, err := strconv.Atoi(userId)

	fmt.Println(err)

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid input data",
			"error":   err.Error(),
		})
		return
	}

	user, errUpdate := services.UpdateUser(id, userInput)

	if errUpdate != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "User update failed",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user updated succesfully",
		"data": user,
	})
}

func DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	id, err := strconv.Atoi(userId)

	fmt.Println(err)

	user, err := services.DeleteUser(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "success",
			"statusCode": http.StatusInternalServerError,
			"message": "user deleted failed",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status" : "success",
		"statusCode": http.StatusOK,
		"message": "user deleted successfully",
		"data": user,
	})
}