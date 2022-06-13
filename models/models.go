package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Seller struct {
	ID          uint   `gorm:" primaryKey json:seller_id"`
	Store_Name  string `json:"store_name" `
	Seller_Name string `json:"seller_name"`
	Password    []byte `json:"password" `
	Email       string `gorm:"unique" json:"email" `
	Phone       string `json:"phone" gorm:"unique"`
	Joined_At   time.Time
	Deleted_At  time.Time
	Sub_Fees    uint `json:"fees"`
}

type User struct {
	ID         uint
	User_Name  string `json:"user_name" `
	Password   []byte `json:"password" `
	Email      string `gorm:"unique" json:"email"`
	Phone      string `json:"phone"`
	Created_at time.Time
}
type ProductUser struct {
	OrderID      uint
	ID           uint `json:"id"`
	Product_Name string
	Price        uint64 `json:"price"`
	Rating       *uint
	Image        string  `json:"image" `
	OrderRefer   int     `json:"order_id"`
	Order        User    `gorm:"foreignKey:OrderRefer"`
	ProductRefer int     `json:"prod_id"`
	Product      Product `gorm:"foreignKey:ProductRefer"`
}
type Product struct {
	ID           uint   `json:"prod_id" gorm:"primaryKey ; autoincrement" `
	Product_Name string `json:"prod_name"`
	Price        int    `json:"prod_price"`
	Image        string `json:"prod_image"`
	SellerRefer  int    `json:"seller_id"`
	Seller       Seller `gorm:"foreignKey:SellerRefer"`
}

//type SellerAddress struct {
//	SellerID uint
//	Store    *string `json:"store_name" `
//	City     *string `json:"city_name"`
//	Seller   Seller
//}
type UserAddress struct {
	ID        uint   `json:"useraddress_id"`
	City      string `json:"city"`
	Regin     string `json:"regin"`
	UserRefer int    `json:"user_id"`
	User      User   `gorm:"foreignKey:UserRefer"`
}
type Order struct {
	ID         uint      `json:"id" gorm:"primaryKey ; autoincrement " `
	Ordered_At time.Time `json:"order_date"`

	UserRefer    int     `json:"user_id"`
	User         User    `gorm:"foreignKey:UserRefer"`
	ProductRefer int     `json:"prod_id"`
	Product      Product `gorm:"foreignKey:ProductRefer"`
	//Order_Cart     []ProductUser `json:"order_list" `
	//Price          int           `json:"total_price" `
	//Discount       *int    `json:"discount" `
	//Payment_Method Payment `json:"payment_method" `
}

type Payment struct {
	gorm.Model
	OrderID uint
	Digital bool
	COD     bool
}
