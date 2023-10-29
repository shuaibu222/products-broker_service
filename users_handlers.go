package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/shuaibu222/p-broker/users"
)

type UsersRequestPayload struct {
	Action string       `json:"action"`
	Users  UsersPayload `json:"users,omitempty"`
}

type UsersPayload struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

// HandleSubmission is the main point of entry into the broker. It accepts a JSON
// payload and performs an action based on the value of "action" in that JSON.
func (app *Config) HandleUsers(w http.ResponseWriter, r *http.Request) {
	var requestPayload UsersRequestPayload

	err := json.NewDecoder(r.Body).Decode(&requestPayload)
	if err != nil {
		log.Println("Error reading request payload")
		return
	}

	switch requestPayload.Action {

	case "create_user":
		app.createUsersGrpc(w, requestPayload.Users)
	case "get_all_users":
		app.getAllUsersGrpc(w)
	case "get_user":
		app.getUserByIdGrpc(w, requestPayload.Users)
	case "update_user":
		app.updateUserGrpc(w, requestPayload.Users)
	case "delete_user":
		app.deleteUserGrpc(w, requestPayload.Users)
	default:
		json.NewEncoder(w).Encode("Unknown action")
	}
}

func (app *Config) createUsersGrpc(w http.ResponseWriter, payload UsersPayload) {
	conn, ctx, cancel, err := UsersGrpcConnection()
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	connection := users.NewUsersServiceClient(conn)
	defer cancel()

	res, err := connection.CreateUser(ctx, &users.UserRequest{
		UserEntry: &users.User{
			Username: payload.Username,
			Password: payload.Password,
		},
	})

	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode("error while creating user")
		return
	}

	json.NewEncoder(w).Encode(res.Response)
}

func (app *Config) getAllUsersGrpc(w http.ResponseWriter) {
	conn, ctx, cancel, err := UsersGrpcConnection()
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	connection := users.NewUsersServiceClient(conn)
	defer cancel()

	res, err := connection.GetAllUsers(ctx, &users.NoParams{})

	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode("error while getting all users")
		return
	}

	json.NewEncoder(w).Encode(res.Response)
}

func (app *Config) getUserByIdGrpc(w http.ResponseWriter, p UsersPayload) {
	conn, ctx, cancel, err := UsersGrpcConnection()
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	connection := users.NewUsersServiceClient(conn)
	defer cancel()

	res, err := connection.GetUserById(ctx, &users.UserId{
		Id: p.Id,
	})

	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode("error while getting user by his id")
		return
	}

	json.NewEncoder(w).Encode(res.Response)
}

func (app *Config) getUserByUsernameGrpc(w http.ResponseWriter, p UsersPayload) (users.User, error) {
	conn, ctx, cancel, err := UsersGrpcConnection()
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	connection := users.NewUsersServiceClient(conn)
	defer cancel()

	res, err := connection.GetUserByUsername(ctx, &users.User{
		Username: p.Username,
	})

	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode("error while getting user by his username")
	}

	return *res.Response, nil
}

func (app *Config) updateUserGrpc(w http.ResponseWriter, p UsersPayload) {
	conn, ctx, cancel, err := UsersGrpcConnection()
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	connection := users.NewUsersServiceClient(conn)
	defer cancel()

	res, err := connection.UpdateUser(ctx, &users.User{
		Id:       p.Id,
		Username: p.Username,
		Password: p.Password,
	})

	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode("error while updating user")
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (app *Config) deleteUserGrpc(w http.ResponseWriter, p UsersPayload) {
	conn, ctx, cancel, err := UsersGrpcConnection()
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	connection := users.NewUsersServiceClient(conn)
	defer cancel()

	res, err := connection.DeleteUser(ctx, &users.UserId{
		Id: p.Id,
	})

	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode("error while deleting user")
		return
	}

	json.NewEncoder(w).Encode(res)
}
