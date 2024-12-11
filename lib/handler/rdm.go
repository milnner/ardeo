package handler

/*
	Abstrai os handlers dos RoadMap do AWS Lambda chamados pelo APIGateway
	um RoadMap é um roteiro de estudos tutorados, que possui etapas, cada etapa (ou, stage)
	possui atividades associadas, a associação acontece de acordo com quem monta o roteiro
*/

import (
	"ardeolib.sapions.com/repository"
)

type RoadMapHandler struct {
	rps *repository.Repository
}

func NewRoadMapHandler(rps *repository.Repository) *RoadMapHandler {
	return &RoadMapHandler{rps: rps}
}

// func (h *RoadMapHandler) HandleCreateRoadMap(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
// 	// var claims utils.UserClaims
// 	// var ok bool
// 	// if claims, ok = ctx.Value("claims").(utils.UserClaims); !ok {
// 	// 	r := "Invalid token claims"
// 	// 	return utils.ErrorResponse(http.StatusUnauthorized, &r), nil
// 	// }

// 	roadMap := &models.RoadMap{}
// 	err := json.Unmarshal([]byte(req.Body), roadMap)
// 	if err != nil {
// 		r := fmt.Sprintf("Invalid request body: %v", err)
// 		return utils.ErrorResponse(http.StatusUnprocessableEntity, &r), nil
// 	}

// 	roadMap.UUID = uuid.New()

// 	srv := services.NewCreateRoadMapService(roadMap, h.rps)

// 	if err = srv.Run(); err != nil {
// 		r := fmt.Sprintf("Failed to create roadmap: %v", err)
// 		return utils.ErrorResponse(http.StatusInternalServerError, &r), nil
// 	}
// 	return utils.SuccessResponse(http.StatusCreated, roadMap, map[string]string{"Content-Type": "application/json"}), nil
// }
