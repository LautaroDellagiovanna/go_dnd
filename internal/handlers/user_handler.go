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
	var users []models.User
	var u models.User

	u.Name = context.Query("name")
	u.Email = context.Query("email")

	users, _ = h.userService.GetUsers(&u)

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
