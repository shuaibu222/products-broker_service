package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/shuaibu222/p-broker/reviews"
)

type ReviewsRequestPayload struct {
	Action  string         `json:"action"`
	Reviews ReviewsPayload `json:"reviews,omitempty"`
}

type ReviewsPayload struct {
	Id        string `json:"id"`
	Message   string `json:"message"`
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
}

// HandleSubmission is the main point of entry into the broker. It accepts a JSON
// payload and performs an action based on the value of "action" in that JSON.
func (app *Config) HandleReviews(w http.ResponseWriter, r *http.Request) {
	var requestPayload ReviewsRequestPayload

	err := json.NewDecoder(r.Body).Decode(&requestPayload)
	if err != nil {
		log.Println("Error reading request payload")
		return
	}

	switch requestPayload.Action {

	case "create_review":
		app.createReviewsGrpc(w, requestPayload.Reviews)
	case "get_product_reviews":
		app.getAllReviewsGrpc(w, requestPayload.Reviews)
	case "update_review":
		app.updateReviewGrpc(w, requestPayload.Reviews)
	case "delete_review":
		app.deleteReviewGrpc(w, requestPayload.Reviews)

	default:
		json.NewEncoder(w).Encode("Unknown action")
	}
}

func (app *Config) createReviewsGrpc(w http.ResponseWriter, payload ReviewsPayload) {
	conn, ctx, cancel, err := ReviewsGrpcConnection()
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	connection := reviews.NewReviewsServiceClient(conn)

	defer cancel()

	res, err := connection.CreateReview(ctx, &reviews.ReviewRequest{
		ReviewEntry: &reviews.Review{
			Msg:       payload.Message,
			UserId:    payload.UserId,
			ProductId: payload.ProductId,
		},
	})

	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode("error while creating review")
		return
	}

	json.NewEncoder(w).Encode(res.Response)

}

func (app *Config) getAllReviewsGrpc(w http.ResponseWriter, p ReviewsPayload) {
	conn, ctx, cancel, err := ReviewsGrpcConnection()
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	connection := reviews.NewReviewsServiceClient(conn)
	defer cancel()

	stream, err := connection.GetReviews(ctx, &reviews.ProductId{
		Id: p.ProductId,
	})

	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode("error while getting all reviews")
		return
	}

	var reviews []reviews.Review
	for {
		review, err := stream.Recv()
		if err == io.EOF {
			break // end of stream
		}
		if err != nil {
			log.Println("Error receiving review:", err)
			json.NewEncoder(w).Encode("error while receiving reviews")
			return
		}

		reviews = append(reviews, *review)

	}

	json.NewEncoder(w).Encode(reviews)
}

func (app *Config) updateReviewGrpc(w http.ResponseWriter, p ReviewsPayload) {
	conn, ctx, cancel, err := ReviewsGrpcConnection()
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	connection := reviews.NewReviewsServiceClient(conn)
	defer cancel()

	res, err := connection.UpdateReview(ctx, &reviews.Review{
		Id:        p.Id,
		Msg:       p.Message,
		UserId:    p.UserId,
		ProductId: p.ProductId,
	})

	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode("error while updating review")
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (app *Config) deleteReviewGrpc(w http.ResponseWriter, p ReviewsPayload) {
	conn, ctx, cancel, err := ReviewsGrpcConnection()
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	connection := reviews.NewReviewsServiceClient(conn)
	defer cancel()

	res, err := connection.DeleteReview(ctx, &reviews.ReviewId{
		Id: p.Id,
	})

	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode("error while deleting review")
		return
	}

	json.NewEncoder(w).Encode(res)
}
