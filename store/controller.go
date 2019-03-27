package store

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Controller struct {
	Repository Repository
}

// Index GET /
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	products := c.Repository.GetProducts() // list of all products
	render.JSON(w, r, products)
}

// AddProduct POST /
func (c *Controller) AddProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request

	log.Println(body)

	if err != nil {
		log.Fatalln("Error AddProduct", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddProduct", err)
	}

	if err := json.Unmarshal(body, &product); err != nil { // unmarshall body contents as a type Candidate
		w.WriteHeader(422) // unprocessable entity
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error AddProduct unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	log.Println(product)
	success := c.Repository.AddProduct(product) // adds the product to the DB
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}

// SearchProduct GET /
func (c *Controller) SearchProduct(w http.ResponseWriter, r *http.Request) {
	query := chi.URLParam(r, "query")

	log.Println("Search Query - " + query)
	products := c.Repository.GetProductsByString(query)
	data, _ := json.Marshal(products)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// UpdateProduct PUT /
func (c *Controller) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var product Product
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		log.Fatalln("Error UpdateProduct", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error UpdateProduct", err)
	}

	if err := json.Unmarshal(body, &product); err != nil { // unmarshall body contents as a type Candidate
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error UpdateProduct unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	productId, err := strconv.Atoi(id)
	product.ID = productId
	log.Println(product.ID)
	success := c.Repository.UpdateProduct(product) // updates the product in the DB

	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}

// GetProduct GET - Gets a single product by ID /
func (c *Controller) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	log.Println(id)
	productid, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalln("Error GetProduct", err)
	}

	product := c.Repository.GetProductById(productid)

	render.JSON(w, r, product)
}

// DeleteProduct DELETE /
func (c *Controller) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	log.Println(id)
	productid, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalln("Error GetProduct", err)
	}

	if err := c.Repository.DeleteProduct(productid); err != "" { // delete a product by id
		log.Println(err)
		if strings.Contains(err, "404") {
			w.WriteHeader(http.StatusNotFound)
		} else if strings.Contains(err, "500") {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}
