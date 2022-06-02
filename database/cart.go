package database

import (

	//"github.com/jinzhu/gorm"
	"errors"
	"time"

	"github.com/zainabmohammed9949/golang-sql-store/models"
)

var (
	ErrCantFindProduct    = errors.New("can't find the product")
	ErrCantDecodeProduct  = errors.New("can't find the product")
	ErrCantUpdateUser     = errors.New("this user is not valid")
	ErrCantRemoveItemCart = errors.New("cannot remove this item from the cart")
	ErrCantGetItem        = errors.New("was unable to get the item from the cart")
	ErrCantBuyCartItem    = errors.New("cannot update the purches")
	ErrUserIdIsNotFound   = errors.New("the user is not found")
	ErrUserIdIsNotValid   = errors.New("the user is not valid")
)

func InstantBuyer() error {

	var product_details models.ProductUser
	var orders_detail models.Order

	orders_detail.ID = orders_detail.ID
	orders_detail.Ordered_At = time.Now()
	orders_detail.Order_Cart = make([]models.ProductUser, 0)
	orders_detail.Payment_Method.COD = true
	orders_detail.Price = int(product_details.Price)

	return nil

}
