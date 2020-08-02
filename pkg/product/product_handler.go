package product

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Response data
type Response struct {
	Status int      `json:"status"`
	Data   []string `json:"data"`
}

var pm ProductManager

func init() {
	pm = NewProductManager()
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Adding new Product")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var product Product
	json.Unmarshal(reqBody, &product)
	err := pm.Add(product)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		log.Println("Product added successfully")
		json.NewEncoder(w).Encode(product)
	}

}

func LoadProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Loading new Product")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var product Product
	json.Unmarshal(reqBody, &product)
	err := pm.Load(product)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		log.Println("Product loaded successfully")
		json.NewEncoder(w).Encode(product)
	}
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting all Product")
	products, err := pm.GetAll()
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		log.Println("Products retrieved successfully")
		json.NewEncoder(w).Encode(products)
	}

}

func EmailSchdedule(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting all Product")
	products, err := pm.GetEmailSchedule()
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		log.Println("Email Schedule List retrieved successfully")
		json.NewEncoder(w).Encode(products)
	}

}
