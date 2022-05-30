package models

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/zainabmohammed9949/golang-mysql-store/database"
)

var db *gorm.DB

type buyer struct {
	gorm.Model
	ID               uint      `gorm:"unique;autoincrement;json :"buyer_id`
	Store_Name       *string   `json:"store_name" `
	Last_Name        *string   `json:"last_name"`
	Password         *string   `json:"password" `
	Email            *string   `gorm:"unique;json:"email"`
	Phone            *string   `json:"phone"`
	Token            *string   `json:"token"`
	Refresh_Token    *string   `json:"refresh_token"`
	Created_At       time.Time `json:"created_at"`
	Updated_At       time.Time `json:"updated_at"`
	iserted_products []Product `json:"byerprod"`
}

type User struct {
	gorm.Model
	First_Name    *string   `json:"first_name" `
	Last_Name     *string   `json:"last_name"`
	Password      *string   `json:"password" `
	Email         *string   `gorm:"unique;json:"email"`
	Phone         *string   `json:"phone"`
	Token         *string   `json:"token"`
	Refresh_Token *string   `json:"refresh_token"`
	Created_At    time.Time `json:"created_at"`
	Updated_At    time.Time `json:"updated_at"`
	//User_ID         string        `json:"user_id"`
	UserCart        []ProductUser `json:"usercart" bson:"usercart"`
	Address_Details []Address     `json:"address" bson:"address"`
	Order_Status    []Order       `json:"orders" bson:"orders"`
	Notes           *string       `json:user_notes`
}

type Product struct {
	ID           uint    `gorm:"json:prod_id; unique; autoinecrement"`
	Product_Name *string `json:"product_name"`
	Price        *uint64 `json:"price"`
	Rating       *uint8  `json:"rating"`
	Image        *string `json:"image"`
}
type ProductUser struct {
	gorm.Model
	ID           uint    `gorm "json:prod_id;unique"`
	Product_Name *string `json:"product_name" bson:"product_name"`
	Price        uint64  `json:"price" bson:"price"`
	Rating       *uint   `json:"rating" bson:"rating"`
	Image        *string `json:"image" bson:"image"`
}

type Address struct {
	House *string `json:"house_name" bson:"house_name"`
	//Street  *string `json:"street_name" bson:"street_name"`
	City *string `json:"city_name" bson:"city_name"`
	//Pincode *string `json:"pin_code" bson:"pin_code"`
}

type Order struct {
	gorm.Model
	ID             uint          `gorm "json:order_id;unique"`
	Order_Cart     []ProductUser `json:"order_list" bson:"order_list"`
	Ordered_At     time.Time     `json:"order_at" bson:"order_at"`
	Price          int           `json:"total_price" bson:"total_price"`
	Discount       *int          `json:"discount" bson:"discount"`
	Payment_Method Payment       `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	Digital bool
	COD     bool
}

func init() {
	database.Connect()
	db = database.GetDB()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&ProductUser{})
	db.AutoMigrate(&Order{})
	db.AutoMigrate(&Address{})

}

//func GetProductByName(name string) (*Product, *gorm.DB) {
//var getProduct = []Product{}
//db := db.Where("Product_Name", name).Find(&getProduct)
//	return &getProduct, db
//}
func (u *User) CreateUser() *User {
	db.NewRecord(u)
	db.Create(&u)
	return u
}

func GetAllUsers() []User {
	var users []User
	db.Find(&users)
	return users
}
func GetAllProds() []Product {
	var prod []Product
	db.Find(&prod)
	return prod
}
func GetProductUser() []ProductUser {
	var userprod []ProductUser
	db.Find(&userprod)
	return userprod
}
func GetUserOrders() []Order {
	var userord []Order
	db.Find(&userord)
	return userord
}
func GetordersUsersById(Id int64) (*Order, *gorm.DB) {
	var getord Order
	db := db.Where("ID=?", Id).Find(&getord)
	return &getord, db
}
func AddAddress(Db *gorm.DB, prod Product, b buyer) error {

	rows, err := Db.Query("INSERT  INTO address(home,city) VALUES (?,?) WHERE email =?", a.House, a.City, u.Email)
	if err != nil {
		log.Println("Invalid")
		return err
	}
	log.Printf("%d address created", rows)
	return nil

}
