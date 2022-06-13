package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/zainabmohammed9949/golang-sql-store/database"
	"github.com/zainabmohammed9949/golang-sql-store/models"
)

type Address struct {
	ID uint `json:"useraddress_id"`

	City  string `json:"city"`
	Regin string `json:"regin"`
	User  User   `json:"user"`
}

func CreateResponseAddress(A models.UserAddress, user User) Address {
	return Address{ID: A.ID, City: A.City, Regin: A.Regin, User: user}
}

func findAddress(id int, A *models.UserAddress) error {

	database.DB.Find(&A, "id =?", id)
	if A.ID == 0 {
		return errors.New("Address does not exist")
	}
	return nil
}
func AddAddress(c *fiber.Ctx) error {
	var address models.UserAddress
	if err := c.BodyParser(&address); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var user models.User
	if err := findUser(address.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())

	}
	database.DB.Create(&address)
	ResponseAddress := CreateResponseAddress(address, CreateResponseUser(user))

	return c.Status(200).JSON(ResponseAddress)

}

func GetAddresses(c *fiber.Ctx) error {

	addresses := []models.UserAddress{}
	var user models.User

	database.DB.Find(&addresses)
	database.DB.Find(&user)
	ResponseAddresses := []Address{}

	for _, add := range addresses {

		ResponseAddress := CreateResponseAddress(add, CreateResponseUser(user))
		ResponseAddresses = append(ResponseAddresses, ResponseAddress)

	}
	return c.Status(200).JSON(ResponseAddresses)

}

func DeleteAddress(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.JSON("please insure the id is int value")
	}
	var address models.UserAddress
	var user models.User

	database.DB.Find(address, "id =?", id)
	if err := findUser(address.UserRefer, &user); err != nil {
		return c.JSON("the user is not found , so the address will be deleted from the database")
	}

	database.DB.Delete(address)
	return c.JSON("address deleted")
}
func GetAddressByID(c *fiber.Ctx) error {
	var address models.UserAddress

	var user models.User
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Please insure that id is integer")

	}

	if err := findAddress(id, &address); err != nil {
		return c.Status(400).JSON(err.Error())

	}

	if err := findUser(address.UserRefer, &user); err != nil {
		database.DB.Delete(address)
		return c.JSON("the user is not found , so the Responsed  address will be deleted from the database")

	}
	resaddress := CreateResponseAddress(address, CreateResponseUser(user))
	return c.Status(200).JSON(resaddress)

}
func UpdateAddress(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var address models.UserAddress
	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please insure that id is integer")

	}
	if err := findUser(address.UserRefer, &user); err != nil {
		database.DB.Delete(address)
		return c.JSON("the user is not found , so the Responsed  address will be deleted from the database")

	}

	if err := findAddress(id, &address); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type UpdateAddress struct {
		City  string `json:"city"`
		Regin string `json:"regin"`
	}
	var updateData UpdateAddress
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	address.City = updateData.City
	address.Regin = updateData.Regin
	database.DB.Save(&address)

	res := CreateResponseAddress(address, CreateResponseUser(user))
	return c.Status(200).JSON(res)
}
