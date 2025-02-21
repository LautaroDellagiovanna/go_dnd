package handlers

import (
	"database/sql"
	"go_dnd/internal/models"
	"go_dnd/internal/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// En golando se crean structs para los datos a enviar

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{
		userService: services.NewUserService(db),
	}
}

func (h *UserHandler) GetUsers(context *gin.Context) {
	var u models.User
	var err error
	var users []models.User

	u.ID, err = strconv.Atoi(context.Query("id"))
	u.Email = context.Query("email")
	u.Name = context.Query("name")

	if err != nil {
		context.IndentedJSON(http.StatusForbidden, gin.H{"message": "id must be a number"})
	}

	if u.ID > 0 && (u.Email != "" || u.Name != "") {
		log.Printf("Excecuting find user by query.\n")
		users, _ = h.userService.GetUsers(&u)
	} else {
		users, _ = h.userService.GetAllUsers()
	}

	context.IndentedJSON(http.StatusOK, users)
}

func (h *UserHandler) AddUser(context *gin.Context) {
	var newUser models.User

	if err := context.BindJSON(&newUser); err != nil {
		return
	}

	h.userService.AddUser(&newUser)

	context.IndentedJSON(http.StatusCreated, newUser)
}

// gets a path parameter
func (h *UserHandler) GetUser(context *gin.Context) {
	param := context.Param("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		context.IndentedJSON(http.StatusForbidden, gin.H{"message": "id must be a number"})
	}

	user, err := h.userService.GetUserByID(id)

	if err != nil {
		log.Printf("Error finding user: %v", err)
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	context.IndentedJSON(http.StatusFound, user)

}
