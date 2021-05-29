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

func webResponse(writer http.ResponseWriter, code int, status string, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)

	response, err := json.Marshal(model.WebResponse{
		Code: code,
		Status: status,
		Data: data,
	})

	if err != nil {
		log.Println(err)
	}

	writer.Write(response)
	return
}

func (c *UserController) listUsers(writer http.ResponseWriter, request *http.Request) {
	var data []*entity.User
	var err error
	name := request.URL.Query().Get("name")
	if name == "" {
		data, err = c.UserService.ListUsers()
	} else {
		data, err = c.UserService.SearchUsers(name)
	}

	if err != nil {
		webResponse(writer, http.StatusInternalServerError, model.StatusInternalServerError, nil)
		return
	}

	if data == nil {
		webResponse(writer, http.StatusNotFound, model.StatusNotFound, nil)
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

	webResponse(writer, http.StatusOK, model.StatusOK, dataUsers)
	return
}

func (c *UserController) createUser(writer http.ResponseWriter, request *http.Request) {
	var input model.CreateUserRequest
	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		log.Println(err.Error())

		webResponse(writer, http.StatusInternalServerError, model.StatusInternalServerError, nil)
		return
	}

	id, err := c.UserService.CreateUser(input.Email, input.Password, input.Name)
	if err != nil {
		log.Println(err)

		webResponse(writer, http.StatusInternalServerError, model.StatusInternalServerError, nil)
		return
	}

	data := model.CreateUserResponse{
		ID: id,
		Email: input.Email,
		Name: input.Name,
	}

	webResponse(writer, http.StatusCreated, model.StatusCreated, data)
	return
}
