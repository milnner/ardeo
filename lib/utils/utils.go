package utils

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

// Função auxiliar para criar uma resposta de erro
func ErrorResponse(status int, message *string) events.APIGatewayV2HTTPResponse {
	return events.APIGatewayV2HTTPResponse{
		StatusCode: status,
		Body:       fmt.Sprintf(`{"error": "%s"}`, *message),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}

// Função auxiliar para criar uma resposta de sucesso
func SuccessResponse(status int, body interface{}, headers map[string]string) events.APIGatewayV2HTTPResponse {
	responseBody, _ := json.Marshal(body)
	return events.APIGatewayV2HTTPResponse{
		StatusCode: status,
		Body:       string(responseBody),
		Headers:    headers,
	}
}
