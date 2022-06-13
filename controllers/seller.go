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

type Seller struct {
	ID          uint      `gorm:"unique;  primaryKey json :seller_id"`
	Store_Name  string    `json:"store_name" `
	Seller_Name string    `json:"seller_name"`
	Password    []byte    `json:"password" `
	Email       string    `gorm:"unique" `
	Phone       string    ` json:"phone" `
	Sub_Fees    uint      `json:"fees"`
	Joined_At   time.Time `json:"joiend_at"`
}

func CreateResponseSeller(s models.Seller) Seller {
	return Seller{ID: s.ID, Joined_At: s.Joined_At, Seller_Name: s.Seller_Name, Store_Name: s.Store_Name, Email: s.Email, Phone: s.Phone}
}
func CreateRespSeller(s models.Seller) models.Seller {
	return models.Seller{ID: s.ID, Joined_At: s.Joined_At, Seller_Name: s.Seller_Name, Store_Name: s.Store_Name, Email: s.Email, Phone: s.Phone}
}
func SellerSignup(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var seller models.Seller
	database.DB.Where("email =?", data["email"]).First(&seller)

	if seller.ID == 0 {

		password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
		seller := models.Seller{
			Seller_Name: data["seller_name"],
			Email:       data["email"],
			Password:    password,
			Store_Name:  data["store_name"],
			Phone:       data["phone"],
			Joined_At:   time.Now(),
		}
		database.DB.Create(&seller)

		return c.Status(200).JSON(seller)
	}

	c.Status(fiber.StatusBadRequest)
	return c.JSON(fiber.Map{"message": " try another email"})
}

func SellerLogin(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var seller models.Seller
	database.DB.Where("email =?", data["email"]).First(&seller)

	if seller.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{"message": "seller not found"})
	}

	if err := bcrypt.CompareHashAndPassword(seller.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	hi := "welcome" + seller.Seller_Name

	return c.Status(200).JSON(hi)

}

func SellerLoginWithCookie(c *fiber.Ctx) error {
	var seller models.Seller
	Login(c)

	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Issuer:    strconv.Itoa(int(seller.ID)),
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

	return c.JSON(seller)
}

func SellerLogout(c *fiber.Ctx) error {

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

func GetSellerByCookie(c *fiber.Ctx) error {
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
	var seller models.Seller

	database.DB.Where("id =?", claims.Issuer).First(&seller)

	return c.JSON(seller)

}

func findSeller(id int, seller *models.Seller) error {

	database.DB.Find(&seller, "id =?", id)
	if seller.ID == 0 {
		return errors.New("seller does not exist")
	}
	return nil
}
func GetSellers(c *fiber.Ctx) error {

	sellers := []models.Seller{}

	database.DB.Find(&sellers)
	//ResponseSellers := []Seller{}

	//for _, seller := range sellers {
	//	ResponseSeller := CreateResponseSeller(seller)
	//		ResponseSellers = append(ResponseSellers, ResponseSeller)

	//	}
	return c.Status(200).JSON(sellers)

}

func GetSellerByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var seller models.Seller

	if err != nil {
		return c.Status(400).JSON("Please insure that id is integer")

	}

	if err := findSeller(id, &seller); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(seller)
}

func UpdateSeller(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var seller models.Seller

	if err != nil {
		return c.Status(400).JSON("Please insure that id is integer")

	}

	if err := findSeller(id, &seller); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type UpdateSeller struct {
		Seller_Name string `json:"seller_name"`
		Password    []byte `json:"password" `
		Email       string `gorm:"unique" `
	}
	var updateData UpdateSeller
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	seller.Seller_Name = updateData.Seller_Name
	seller.Email = updateData.Email
	seller.Password = updateData.Password

	database.DB.Save(&seller)

	res := CreateResponseSeller(seller)
	return c.Status(200).JSON(res)
}

func DeleteSeller(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var seller models.Seller

	if err != nil {
		return c.Status(400).JSON("Please insure that id is integer")

	}

	if err := findSeller(id, &seller); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.DB.Delete(&seller).Error; err != nil {
		return c.Status(400).JSON(err.Error())

	}

	return c.Status(200).SendString("Successfully deleted")
}
