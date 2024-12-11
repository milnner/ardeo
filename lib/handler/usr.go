package handler

/*
	Abstrai os handlers para User do AWS Lambda chamados pelo APIGateway
	um User pode ser professor ou aluno, caso professor ele tem outras permissões
	como criar atividades e roteiros.
*/

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"ardeolib.sapions.com/models"
	"ardeolib.sapions.com/repository"
	"ardeolib.sapions.com/services"
	"ardeolib.sapions.com/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/gocql/gocql"
)

type UserBody struct {
	Email       string     `json:"email"`
	UUID        gocql.UUID `json:"uuid"`
	IsProfessor bool       `json:"isProfessor"`
	Name        string     `json:"name"`
	Password    string     `json:"password"`
}

func (u *UserBody) ToUser() *models.User {
	return &models.User{
		Email:       u.Email,
		UUID:        u.UUID,
		IsProfessor: u.IsProfessor,
		Name:        u.Name,
	}
}

type UserHandler struct {
	rps *repository.Repository
}

func NewUserHandler(rps *repository.Repository) *UserHandler {
	return &UserHandler{rps: rps}
}

func (h *UserHandler) HandleCreateUser(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	user := &UserBody{}
	err := json.Unmarshal([]byte(req.Body), user)
	if err != nil {
		r := fmt.Sprintf("Invalid request body: %v", err)
		return utils.ErrorResponse(http.StatusUnprocessableEntity, &r), nil
	}

	user.UUID = gocql.TimeUUID()
	srv := services.NewCreateUserService(user.ToUser(), &user.Password, h.rps)
	if err = srv.Run(); err != nil {
		r := fmt.Sprintf("Failed to create user: %v", err)
		return utils.ErrorResponse(http.StatusInternalServerError, &r), nil
	}
	return utils.SuccessResponse(http.StatusCreated, user, map[string]string{"Content-Type": "application/json"}), nil
}
