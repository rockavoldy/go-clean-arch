package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-clean-arch/api/model"
	"go-clean-arch/entity"
	"go-clean-arch/service/user"
	"log"
	"net/http"
)

type UserController struct {
	UserService user.Service
}

func NewUserController(userService *user.Service) UserController {
	return UserController{
		UserService: *userService,
	}
}

func (c *UserController) Route(r *mux.Router) {
	r.HandleFunc("/user", c.listUsers).Methods("GET", "OPTIONS").Name("listUsers")
	r.HandleFunc("/user", c.createUser).Methods("POST", "OPTIONS").Name("createUser")
	r.HandleFunc("/user/{id}", c.getUser).Methods("GET", "OPTIONS").Name("getUser")
	r.HandleFunc("/user/{id}", c.deleteUser).Methods("DELETE", "OPTIONS").Name("deleteUser")
	
}

const errorGetUser = "Error reading users"
const errorCreateUser = "Error adding user"

func (c *UserController) listUsers(writer http.ResponseWriter, request *http.Request) {
	var data []*entity.User
	var err error
	name := request.URL.Query().Get("name")
	if name == "" {
		data, err = c.UserService.ListUsers()
	} else {
		data, err = c.UserService.SearchUsers(name)
	}

	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(errorGetUser))
		return
	}

	if data == nil {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte(errorGetUser))
		return
	}

	var dataUsers []*model.GetUserResponse
	for _, d := range data {
		dataUsers = append(dataUsers, &model.GetUserResponse{
			ID:    d.ID,
			Email: d.Email,
			Name:  d.Name,
		})
	}

	response := &model.WebResponse{Code: 200, Status: "OK", Data: dataUsers}

	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(errorGetUser))
	}
}

func (c *UserController) createUser(writer http.ResponseWriter, request *http.Request) {
	var input model.CreateUserRequest
	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		res, _ := json.Marshal(
			model.WebResponse{
				Code: http.StatusInternalServerError,
				Status: errorCreateUser,
				Data: nil,
			})
		writer.Write(res)
		return
	}

	id, err := c.UserService.CreateUser(input.Email, input.Password, input.Name)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		res, _ := json.Marshal(
			model.WebResponse{
				Code: http.StatusInternalServerError,
				Status: errorCreateUser,
				Data: nil,
			})
		writer.Write(res)
		return
	}

	response, _ := json.Marshal(
		model.WebResponse{
			Code: http.StatusCreated,
			Status: "Created",
			Data: model.CreateUserResponse{
				ID: id,
				Email: input.Email,
				Name: input.Name,
			},
		})

	writer.WriteHeader(http.StatusCreated)
	writer.Write(response)
}

func (c *UserController) getUser(writer http.ResponseWriter, request *http.Request) {
	return
}

func (*UserController) deleteUser(writer http.ResponseWriter, request *http.Request) {
	return
}

