package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("can't find the product")
	ERrCantDecodeProduct  = errors.New("can't find the product")
	ErrCantUpdateUser     = errors.New("this user is not valid")
	ErrCantRemoveItemCart = errors.New("cannot remove this item from the cart")
	ErrCantGetItem        = errors.New("was unable to get the item from the cart")
	ErrCantBuyCartItem    = errors.New("cannot update the purches")
)

func AddProductToCart() error {

}
func RemoveCartItem() error {

}
func BuyItemFromCart() {

}
func InstantBuyer() {

}
