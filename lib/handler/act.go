package handler

/*
	Abstrai os handlers das Act do AWS Lambda chamados pelo APIGateway
	uma Act Relaciona uma atividade de qualquer tipo a um RoadMap,
	dando a ela um estágio (ou, stage) ao qual pertence
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
)

type ActHandler struct {
	rps *repository.Repository
}

func NewActHandler(rps *repository.Repository) *ActHandler {
	return &ActHandler{rps: rps}
}

func (h *ActHandler) HandleCreateAct(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	var claims utils.UserClaims
	var ok bool
	if claims, ok = ctx.Value("claims").(utils.UserClaims); !ok {
		r := "Invalid token claims"
		return utils.ErrorResponse(http.StatusUnauthorized, &r), nil
	}

	act := &models.Act{}
	err := json.Unmarshal([]byte(req.Body), act)
	if err != nil {
		r := fmt.Sprintf("Invalid request body: %v", err)
		return utils.ErrorResponse(http.StatusUnprocessableEntity, &r), nil
	}

	srv := services.NewCreateActService(act, &claims.User, h.rps)

	if err = srv.Run(); err != nil {
		r := fmt.Sprintf("Failed to create act: %v", err)
		return utils.ErrorResponse(http.StatusInternalServerError, &r), nil
	}
	return utils.SuccessResponse(http.StatusCreated, act, map[string]string{"Content-Type": "application/json"}), nil
}

func (h *ActHandler) HandleCreateActs(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	var claims utils.UserClaims
	var ok bool
	if claims, ok = ctx.Value("claims").(utils.UserClaims); !ok {
		r := "Invalid token claims"
		return utils.ErrorResponse(http.StatusUnauthorized, &r), nil
	}

	acts := &[]models.Act{}
	err := json.Unmarshal([]byte(req.Body), acts)
	if err != nil {
		r := fmt.Sprintf("Invalid request body: %v", err)
		return utils.ErrorResponse(http.StatusUnprocessableEntity, &r), nil
	}

	srv := services.NewCreateActsService(acts, &claims.User, h.rps)

	if err = srv.Run(); err != nil {
		r := fmt.Sprintf("Failed to create act: %v", err)
		return utils.ErrorResponse(http.StatusInternalServerError, &r), nil
	}
	return utils.SuccessResponse(http.StatusCreated, acts, map[string]string{"Content-Type": "application/json"}), nil
}

func (h *ActHandler) HandleDeleteActs(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	var claims utils.UserClaims
	var ok bool
	if claims, ok = ctx.Value("claims").(utils.UserClaims); !ok {
		r := "Invalid token claims"
		return utils.ErrorResponse(http.StatusUnauthorized, &r), nil
	}
	rdm := &models.RoadMap{}
	err := json.Unmarshal([]byte(req.Body), rdm)
	if err != nil {
		r := fmt.Sprintf("Invalid request body: %v", err)
		return utils.ErrorResponse(http.StatusUnprocessableEntity, &r), nil
	}

	srv := services.NewDeleteActsService(rdm, &claims.User, h.rps)

	if err = srv.Run(); err != nil {
		r := fmt.Sprintf("Failed to delete acts: %v", err)
		return utils.ErrorResponse(http.StatusInternalServerError, &r), nil
	}
	return utils.SuccessResponse(http.StatusNoContent, "", map[string]string{"Content-Type": "application/json"}), nil
}

func (h *ActHandler) HandleReadActs(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	var claims utils.UserClaims
	var ok bool
	if claims, ok = ctx.Value("claims").(utils.UserClaims); !ok {
		r := "Invalid token claims"
		return utils.ErrorResponse(http.StatusUnauthorized, &r), nil
	}
	rdm := &models.RoadMap{}
	err := json.Unmarshal([]byte(req.Body), rdm)
	if err != nil {
		r := fmt.Sprintf("Invalid request body: %v", err)
		return utils.ErrorResponse(http.StatusUnprocessableEntity, &r), nil
	}
	acts := &[]models.Act{}

	srv := services.NewReadActService(acts, rdm, &claims.User, h.rps)

	if err = srv.Run(); err != nil {
		r := fmt.Sprintf("Failed to read acts: %v", err)
		return utils.ErrorResponse(http.StatusInternalServerError, &r), nil
	}
	return utils.SuccessResponse(http.StatusOK, acts, map[string]string{"Content-Type": "application/json"}), nil
}
