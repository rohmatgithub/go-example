package dbsharding

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type productResponse struct {
	Products []productData `json:"products"`
}
type productData struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Brand       string  `json:"brand"`
	SKU         string  `json:"sku"`
	Weight      float64 `json:"weight"`
}

type customerResponse struct {
	Users []cutomerData `json:"users"`
}
type cutomerData struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   struct {
		Address string `json:"address"`
		City    string `json:"city"`
		State   string `json:"state"`
	} `json:"address"`
}

func prepareDataMaster(db *gorm.DB) {
	getCustomer(db)
	getProducts(db)
}

func getProducts(db *gorm.DB) {
	url := "https://dummyjson.com/products?limit=200"
	// insert data product
	body, err := makeRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Failed to get data: %v\n", err)
	}

	// Unmarshal the JSON response into the APIResponse struct
	var apiResponse productResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v\n", err)
	}

	for _, p := range apiResponse.Products {
		product := Product{
			Title:       p.Title,
			Description: p.Description,
			CategoryID:  0,
			Price:       p.Price * 10000, // because from api price in dollar
			Brand:       p.Brand,
			SKU:         p.SKU,
			Weight:      p.Weight,
		}

		var pc ProductCategory
		result := db.First(&pc, "code = ?", p.Category)
		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			log.Fatalf("Error retrieving product category: %v\n", result.Error)
		}

		if pc.ID == 0 {
			pc.Code = p.Category
			pc.Name = p.Category
			err = db.Create(&pc).Error
			if err != nil {
				log.Fatalf("Error inserting product category: %v\n", err)
			}
		}
		product.CategoryID = pc.ID

		err = db.Create(&product).Error
		if err != nil {
			log.Fatalf("Error inserting product: %v\n", err)
		}
	}
}

func getCustomer(db *gorm.DB) {

	url := "https://dummyjson.com/users?limit=250"
	// insert data product
	body, err := makeRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Failed to get data: %v\n", err)
	}

	// Unmarshal the JSON response into the APIResponse struct
	var cr customerResponse
	err = json.Unmarshal(body, &cr)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v\n", err)
	}

	for _, p := range cr.Users {
		customer := Customer{
			Name:    fmt.Sprintf("%s %s", p.FirstName, p.LastName),
			Email:   p.Email,
			Phone:   p.Phone,
			Address: fmt.Sprintf("%s, %s, %s", p.Address.Address, p.Address.City, p.Address.State),
		}
		err = db.Create(&customer).Error
		if err != nil {
			log.Fatalf("Error inserting customer: %v\n", err)
		}
	}
}

func makeRequest(method, url string, requestBody interface{}) ([]byte, error) {
	var reqBody *bytes.Buffer
	if requestBody != nil {
		// Convert request body to JSON
		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %v", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	} else {
		reqBody = bytes.NewBuffer(nil)
	}

	// Create a new HTTP request
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %v", err)
	}

	// If request body exists, set content type as JSON
	if requestBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// Perform the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the status code indicates an error
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK response: %s", resp.Status)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	return body, nil
}
