package controllers

import (
	"errors"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vladimiroff/jwt-go/v3"
	"github.com/zainabmohammed9949/golang-sql-store/database"
	"github.com/zainabmohammed9949/golang-sql-store/models"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint   `json:"user_id" gorm:" primaryKey"`
	User_Name string `json:"user_name" `
	Email     string `gorm:"unique" json:"email"`
	Phone     string `json:"phone" gorm:"unique" `
	Password  []byte `json:"password"`
}

const SecretKey = "Secret"

func CreateResponseUser(u models.User) User {

	return User{ID: u.ID, User_Name: u.User_Name, Phone: u.Phone, Email: u.Email}
}

func Signup(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	var user models.User

	database.DB.Where("email =?", data["email"]).First(&user)

	if user.ID != 0 {
		c.Status(fiber.StatusBadRequest)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "user alredy exist, try another email"})

	}
	user = models.User{
		User_Name:  data["user_name"],
		Email:      data["email"],
		Password:   password,
		Phone:      data["phone"],
		Created_at: time.Now(),
	}

	database.DB.Create(&user)

	return c.Status(200).JSON(user)

}
func findUser(id int, user *models.User) error {

	database.DB.Find(&user, "id =?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

func Login(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var user models.User
	database.DB.Where("email =?", data["email"]).First(&user)

	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{"message": "user not found"})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	hi := "welcome" + user.User_Name

	return c.Status(200).JSON(hi)

}

func LoginWithCookie(c *fiber.Ctx) error {
	var user models.User
	Login(c)

	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Issuer:    strconv.Itoa(int(user.ID)),
	}

	claims_c := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := claims_c.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"message": "could not log in"})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "loged out",
	})

}

func GetUserByCookie(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)
	var user models.User

	database.DB.Where("id =?", claims.Issuer).First(&user)

	return c.JSON(user)

}

func GetUsers(c *fiber.Ctx) error {

	users := []models.User{}

	database.DB.Find(&users)

	//	ResponseUsers := []User{}

	//for _, user := range users {
	//	ResponseUser := CreateResponseUser(user)
	//ResponseUsers = append(ResponseUsers, ResponseUser)

	//}
	return c.Status(200).JSON(users)

}

func GetUserByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please insure that userid is integer")

	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).JSON(user)

}
func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please insure that id is integer")

	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type UpdateUser struct {
		User_Name string `json:"user_name" `
		Email     string `gorm:"unique" json:"email"`
		Phone     string `gorm:"unique" `
		Password  string
	}
	var updateData UpdateUser
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.User_Name = updateData.User_Name
	user.Email = updateData.Email
	user.Phone = updateData.Phone
	user.Password = []byte(updateData.Password)

	database.DB.Save(&user)

	return c.Status(200).JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please insure that id is integer")

	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		return c.Status(400).JSON(err.Error())

	}
	//deleting corresponding orders,if are;
	var order models.Order
	refer := order.UserRefer
	if int(user.ID) == refer {
		database.DB.Delete(&order, "user_refer=? ", user.ID)
		return c.Status(400).JSON("related orders are deleted too")
	}

	return c.Status(200).SendString("Successfully deleted")
}
