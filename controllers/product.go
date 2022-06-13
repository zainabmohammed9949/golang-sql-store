package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/zainabmohammed9949/golang-sql-store/database"
	"github.com/zainabmohammed9949/golang-sql-store/models"
)

type Product struct {
	ID           uint          `json:"prod_id"`
	Product_Name string        `json:"prod_name"`
	Price        int           `json:"prod_price"`
	Image        string        `json:"prod_image"`
	Seller       models.Seller `json:"seller"`
}

func CreateResponseProduct(p models.Product, seller models.Seller) Product {
	return Product{ID: p.ID, Seller: seller, Product_Name: p.Product_Name, Image: p.Image, Price: p.Price}
}
func findproduct(id int, p *models.Product) error {

	database.DB.Find(&p, "id =?", id)
	if p.ID == 0 {
		return errors.New("product does not exist")
	}
	return nil
}

func AddProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var seller models.Seller
	if err := findSeller(product.SellerRefer, &seller); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Create(&product)
	responseSeller := CreateRespSeller(seller)
	responseProduct := CreateResponseProduct(product, responseSeller)
	return c.Status(200).JSON(responseProduct)
}
func GetProducts(c *fiber.Ctx) error {

	products := []models.Product{}
	var seller models.Seller
	//var product models.Product

	database.DB.Find(&products)
	database.DB.Find(&seller)

	ResponseProducts := []Product{}

	for _, product := range products {

		ResponseProduct := CreateResponseProduct(product, CreateRespSeller(seller))
		ResponseProducts = append(ResponseProducts, ResponseProduct)

	}
	return c.Status(200).JSON(ResponseProducts)

}

func GetProductByID(c *fiber.Ctx) error {
	var product models.Product
	var seller models.Seller
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Please insure that id is integer")

	}

	if err := findproduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())

	}
	if err := findSeller(product.SellerRefer, &seller); err != nil {
		return c.Status(400).JSON(err.Error())

	}
	res := CreateResponseProduct(product, CreateRespSeller(seller))
	return c.Status(200).JSON(res)

}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product models.Product

	if err != nil {
		return c.Status(400).JSON("Please insure that id is integer")

	}

	if err := findproduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.DB.Delete(&product).Error; err != nil {
		return c.Status(400).JSON(err.Error())

	}

	return c.Status(200).SendString("Successfully deleted")
}
