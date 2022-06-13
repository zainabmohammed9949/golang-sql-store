package controllers

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/zainabmohammed9949/golang-sql-store/database"
	"github.com/zainabmohammed9949/golang-sql-store/models"
)

type Order struct {
	ID         uint      `json:"id"`
	User       User      `json:"user"`
	Product    Product   `json:"product"`
	Ordered_At time.Time `json:"order_date"`
}

func CreateResponseOrder(o models.Order, user User, product Product) Order {
	return Order{ID: o.ID, User: user, Product: product, Ordered_At: o.Ordered_At}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	if err := findUser(order.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var product models.Product

	if err := findproduct(order.ProductRefer, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var seller models.Seller
	database.DB.Create(&order)

	responseUser := CreateResponseUser(user)
	responserSeller := CreateRespSeller(seller)
	responseProduct := CreateResponseProduct(product, responserSeller)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)

}
func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}
	database.DB.Find(&orders)

	allOrders := []Order{}

	for _, order := range orders {
		var user models.User
		var seller models.Seller
		var product models.Product
		database.DB.Find(&user, "id = ?", order.UserRefer)
		database.DB.Find(&product, "id = ?", order.ProductRefer)
		database.DB.Find(&seller, "id = ?", product.SellerRefer)

		allOrder := CreateResponseOrder(order, CreateResponseUser(user), CreateResponseProduct(product, CreateRespSeller(seller)))
		allOrders = append(allOrders, allOrder)

	}
	return c.Status(200).JSON(allOrders)

}

func FindOrder(id int, order *models.Order) error {

	database.DB.Find(&order, "id =?", id)
	if order.ID == 0 {
		return errors.New("order does not exist")
	}
	return nil
}

func GetOrderByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var order models.Order

	if err != nil {
		return c.Status(400).JSON("Please insure that id is integer")

	}

	if err := FindOrder(id, &order); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var user models.User
	var product models.Product
	var seller models.Seller

	if err := findproduct(order.ProductRefer, &product); err != nil {
		return c.Status(400).JSON("check  product  ")
	}

	if err := findUser(order.UserRefer, &user); err != nil {
		return c.Status(400).JSON("check  user  ")
	}
	if err := findSeller(product.SellerRefer, &seller); err != nil {
		return c.Status(400).JSON("check seller  ")
	}

	database.DB.First(&user, order.UserRefer)
	database.DB.First(&product, order.ProductRefer)
	database.DB.First(&seller, product.SellerRefer)

	responseOrder := CreateResponseOrder(order, CreateResponseUser(user), CreateResponseProduct(product, CreateRespSeller(seller)))

	return c.Status(200).JSON(responseOrder)
}
func CancelOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var order models.Order

	if err != nil {
		return c.Status(400).JSON("Please insure that id is integer")

	}

	if err := FindOrder(id, &order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.DB.Delete(&order).Error; err != nil {
		return c.Status(400).JSON(err.Error())

	}

	return c.Status(200).SendString("Successfully deleted")
}
