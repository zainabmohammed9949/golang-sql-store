package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/zainabmohammed9949/golang-sql-store/database"
)

var db *gorm.DB

type Seller struct {
	gorm.Model
	ID              uint            `gorm:"unique;autoincrement;json :"buyer_id""`
	Store_Name      *string         `json:"store_name" `
	Seller_Name     *string         `json:"seller_name"`
	Password        *string         `json:"password" `
	Email           *string         `gorm:"unique;json:"email""`
	Phone           *string         `gorm:"unique;json:"phone""`
	Token           *string         `json:"token"`
	Refresh_Token   *string         `json:"refresh_token"`
	Joined_At       time.Time       `json:"joiend_at"`
	Deleted_At      time.Time       `json:"deleted_at"`
	Seller_products []Product       `gorm :"many2many:sellerproducts;json:"sellerprod"`
	Sub_Fees        uint            `json :"fees"`
	Address_Details []SellerAddress `gorm :"many2many:seller_address;" `
}

type User struct {
	gorm.Model
	ID              uint          `gorm:"unique;autoincrement;default:uuid_generate_v3; json :"user_id`
	User_Name       *string       `json:"user_name" `
	Password        *string       `json:"password" `
	Email           *string       `gorm:"unique;json:"email""`
	Phone           *string       `gorm:"unique;json:"phone""`
	UserCart        []ProductUser `gorm :"many2many:userproducts;"`
	Address_Details []UserAddress `gorm :"many2many:useraddress;" `
	Order_Status    []Order       `gorm :"many2many:orders" ;foreignKey:Refer;joinForeignKey:UserReferID;References:UserRefer;joinReferences:OrderRefer"`
	Refer           uint          `gorm:"index:unique"`
}

type Product struct {
	gorm.Model
	ID           uint `gorm:"json:prod_id;primary_key; unique; autoinecrement"`
	Product_Name *string
	Price        *string
	Image        *string
	Sellers      []*Seller `gorm :"many2many:sellerproducts" ;`
}
type ProductUser struct {
	gorm.Model
	ID           uint `gorm "json:prod_id;unique"`
	Product_Name *string
	Price        uint64 `json:"price" bson:"price"`
	Rating       *uint
	Image        *string `json:"image" bson:"image"`
	Users        []*User `gorm :"many2many:userproducts" ;`
}

type UserAddress struct {
	House *string `json:"house_name" bson:"house_name"`
	City  *string `json:"city_name" bson:"city_name"`
	Users []*User `gorm :"many2many:useraddress" ;`
}
type SellerAddress struct {
	Store   *string   `json:"store_name" bson:"store_name"`
	City    *string   `json:"city_name" bson:"city_name"`
	Sellers []*Seller `gorm :"many2many:selleraddress" ;`
}

type Order struct {
	gorm.Model
	ID             uint          `gorm "json:order_id;unique"`
	Order_Cart     []ProductUser `json:"order_list" bson:"order_list"`
	Ordered_At     time.Time     `json:"order_at" bson:"order_at"`
	Price          int           `json:"total_price" bson:"total_price"`
	Discount       *int          `json:"discount" bson:"discount"`
	Payment_Method Payment       `json:"payment_method" `
	Users          []*User       `gorm :"many2many:orders" ;`
}

type Payment struct {
	Digital bool
	COD     bool
}

func init() {
	database.Connect()
	db = database.GetDB()
	db.AutoMigrate(&Seller{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Product{})

}
func (u *User) CreateUser() *User {
	db.NewRecord(u)
	db.Create(&u)
	return u
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	db.Find(&users)
	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)
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
