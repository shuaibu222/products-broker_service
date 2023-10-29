package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/shuaibu222/p-broker/products"
)

type RequestPayload struct {
	Action   string          `json:"action"`
	Products ProductsPayload `json:"products,omitempty"`
}

type ProductsPayload struct {
	Id          string `json:"id"`
	Name        string `bson:"name" json:"name"`
	Motto       string `bson:"motto" json:"motto"`
	Link        string `bson:"link" json:"link"`
	Category    string `bson:"category" json:"category"`
	Description string `bson:"description" json:"description"`
	ImageLink   string `bson:"image_link" json:"image_link"`
}

// HandleSubmission is the main point of entry into the broker. It accepts a JSON
// payload and performs an action based on the value of "action" in that JSON.
func (app *Config) HandleProducts(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := json.NewDecoder(r.Body).Decode(&requestPayload)
	if err != nil {
		log.Println("Error reading request payload")
		return
	}

	switch requestPayload.Action {
	case "create_product":
		app.createProductGrpc(w, requestPayload.Products)
	case "get_all_products":
		app.getAllProductsGrpc(w)
	case "get_product":
		app.getProductByIdGrpc(w, requestPayload.Products)
	case "update_product":
		app.updateProductGrpc(w, requestPayload.Products)
	case "delete_product":
		app.deleteProductGrpc(w, requestPayload.Products)
	default:
		json.NewEncoder(w).Encode("Unknown action")
	}
}

func (app *Config) createProductGrpc(w http.ResponseWriter, p ProductsPayload) {
	conn, ctx, cancel, err := ProductsGrpcConnection()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	connection := products.NewProductsServiceClient(conn)

	defer cancel()

	res, err := connection.CreateProducts(ctx, &products.ProductRequest{
		ProductEntry: &products.Product{
			Name:        p.Name,
			Motto:       p.Motto,
			Link:        p.Link,
			Category:    p.Category,
			Description: p.Description,
			ImageLink:   p.ImageLink,
		},
	})
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode("error while creating product")
		return
	}

	json.NewEncoder(w).Encode(res.Response)
}

func (app *Config) getAllProductsGrpc(w http.ResponseWriter) {
	conn, ctx, cancel, err := ProductsGrpcConnection()
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	connection := products.NewProductsServiceClient(conn)
	defer cancel()

	stream, err := connection.GetAllProducts(ctx, &products.NoParams{})

	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode("error while getting all products")
		return
	}

	var products []products.Product

	for {
		product, err := stream.Recv()
		if err == io.EOF {
			break // end of stream
		}
		if err != nil {
			log.Println("Error receiving product:", err)
			json.NewEncoder(w).Encode("error while receiving products")
			return
		}

		products = append(products, *product)
	}

	json.NewEncoder(w).Encode(products)
}

func (app *Config) getProductByIdGrpc(w http.ResponseWriter, p ProductsPayload) {
	conn, ctx, cancel, err := ProductsGrpcConnection()
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	connection := products.NewProductsServiceClient(conn)
	defer cancel()

	res, err := connection.GetProductById(ctx, &products.ProductId{
		Id: p.Id,
	})

	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode("error while getting user by his id")
		return
	}

	json.NewEncoder(w).Encode(res.Response)
}

func (app *Config) updateProductGrpc(w http.ResponseWriter, p ProductsPayload) {
	conn, ctx, cancel, err := ProductsGrpcConnection()
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	connection := products.NewProductsServiceClient(conn)
	defer cancel()

	res, err := connection.UpdateProduct(ctx, &products.Product{
		Id:          p.Id,
		Name:        p.Name,
		Motto:       p.Motto,
		Link:        p.Link,
		Category:    p.Category,
		Description: p.Description,
		ImageLink:   p.ImageLink,
	})

	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode("error while updating product")
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (app *Config) deleteProductGrpc(w http.ResponseWriter, p ProductsPayload) {
	conn, ctx, cancel, err := ProductsGrpcConnection()
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	connection := products.NewProductsServiceClient(conn)
	defer cancel()

	res, err := connection.DeleteProduct(ctx, &products.ProductId{
		Id: p.Id,
	})

	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode("error while deleting product")
		return
	}

	json.NewEncoder(w).Encode(res)
}
