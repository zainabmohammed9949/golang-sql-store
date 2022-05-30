package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/zainabmohammed9949/golang-mysql-store/models"

	//"github.com/jinzhu/gorm"

	"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerfyPassward(userPassword string, givenPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(givenPassword), []byte(userPassword))
	valid := true
	msg := ""
	if err != nil {
		msg = "Login or Password is not correct"
		valid = false
	}
	return valid, msg
}
func signup(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "./templates/signup.html")
		return
	}
	email := req.FormValue("email")
	password := req.FormValue("password")

	slct, err := DB.Prepare("SELECT email FROM users WHERE email = ? ")
	if err != nil {
		http.Error(res, "server error, unable to create your account.", 500)
		return
	}
	_, err = slct.Exec(email)
	switch {
	case err == sql.ErrNoRows:
		HashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		if err != nil {
			http.Error(res, "server error, unable to create your account.", 500)
			return
		}
		_, err = DB.Exec("INSERT INTO users (email,password) VALUES (?,?)", email, HashedPassword)
		if err != nil {
			http.Error(res, "server error, unable to create your account.", 500)
			return
		}
		res.Write([]byte("user created!"))
		return
	case err != nil:
		http.Error(res, "server error, unable to create your account.", 500)
		return
	default:
		http.Redirect(res, req, "/", 301)
	}

}

//func ProductViewerAdmins() gin.HandlerFunc {

//}
//
//}
func searchProduct(DB *sql.DB, res http.ResponseWriter, r *http.Request, p models.Product) ([]models.Product, error) {
	if r.Method != "POST" {
		http.ServeFile(res, r, "home.html")

	}
	product := []models.Product{}
	rowstext := DB.QueryRow("SELECT * FROM products WHERE prod_name =?", *p.Product_Name)
	for rowstext.Next() {
		var id uint
		var product_name *string
		var price *uint64
		var rating *uint8
		var image *string
		err2 := rowstext.Scan(&id, &product_name, &price, &image)
		if err2 != nil {
			return nil, err2
		}
		product := models.Product{id, product_name, price, rating, image}
		products := append(products, product)
	}
	return products, nil

}

func login(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "login.html")
		return
	}
	email := req.FormValue("email")
	password := req.FormValue("password")

	var dbemail, psw string
	err := DB.QueryRow("SELECT email,password FROM users WHERE email =?", email).Scan(&dbemail, &psw)

	if err != nil {
		http.Redirect(res, req, "/login", 301)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(psw), []byte(password))

	if err != nil {
		http.Redirect(res, req, "/login", 301)
		return
	}

	res.Write([]byte("Hello" + email))

}
