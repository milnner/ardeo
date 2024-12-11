package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"ardeolib.sapions.com/utils"
	"github.com/aws/aws-lambda-go/events"
)

type LambdaHandler func(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error)

type JWTMiddleware struct {
	jwtService utils.JWTUtil
}

func NewJWTMiddleware(jwtService utils.JWTUtil) *JWTMiddleware {
	return &JWTMiddleware{jwtService: jwtService}
}

func (m *JWTMiddleware) AWSLambdaHandler(next LambdaHandler) LambdaHandler {
	return func(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
		token := req.Headers["Authorization"]
		if token == "" {
			r := "Authorization token is required"
			return utils.ErrorResponse(http.StatusUnauthorized, &r), nil
		}

		token = strings.ReplaceAll(token, "Bearer ", "")

		userClaims, err := m.jwtService.ValidateToken(token)
		if err != nil {
			r := fmt.Sprintf("Invalid token: %v", err)
			return utils.ErrorResponse(http.StatusUnauthorized, &r), nil
		}

		if expirationTime, err := userClaims.GetExpirationTime(); err != nil {
			r := fmt.Sprintf("Invalid token: %v", err)
			return utils.ErrorResponse(http.StatusUnauthorized, &r), nil
		} else if expirationTime.Time.Before(time.Now()) {
			r := fmt.Sprintf("Expired token")
			return utils.ErrorResponse(http.StatusUnauthorized, &r), nil
		}

		ctx = context.WithValue(ctx, "claims", *userClaims)

		return next(ctx, req)
	}
}
