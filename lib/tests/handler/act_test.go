package handler_test

// testar o handler de CreateAct
import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"ardeolib.sapions.com/handler"
	"ardeolib.sapions.com/models"
	"ardeolib.sapions.com/repository"
	"github.com/aws/aws-lambda-go/events"
	"github.com/gocql/gocql"
	"github.com/stretchr/testify/assert"
)

var (
	Act1 = models.Act{
		RoadMapUUID: Roadmap1.UUID,
		ActUUID:     gocql.UUID{0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		ActType:     models.OQNAActT,
		Stage:       1,
	}
	Act2 = models.Act{
		RoadMapUUID: Roadmap1.UUID,
		ActUUID:     gocql.UUID{0xFF, 0xFF, 0x0F, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		ActType:     models.OQNAActT,
		Stage:       1,
	}
	Act3 = models.Act{
		RoadMapUUID: Roadmap2.UUID,
		ActUUID:     gocql.UUID{0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		ActType:     models.OQNAActT,
		Stage:       1,
	}
)

func TestHandleCreateAct(t *testing.T) {

	repo := repository.NewRepository(CassandraSession)
	actHandler := handler.NewActHandler(repo)
	var reqBody models.Act

	reqBody = Act1

	reqBodyBytes, _ := json.Marshal(reqBody)
	req := events.APIGatewayV2HTTPRequest{
		Body: string(reqBodyBytes),
	}

	resp, err := actHandler.HandleCreateAct(context.Background(), req)
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestHandleCreateActs(t *testing.T) {
	repo := repository.NewRepository(CassandraSession)
	actHandler := handler.NewActHandler(repo)
	var reqBody []models.Act

	reqBody = append(reqBody, Act1)

	reqBody = append(reqBody, Act2)

	reqBody = append(reqBody, Act3)

	reqBodyBytes, _ := json.Marshal(reqBody)
	req := events.APIGatewayV2HTTPRequest{
		Body: string(reqBodyBytes),
	}

	resp, err := actHandler.HandleCreateActs(context.Background(), req)
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestDeleteActs(t *testing.T) {
	repo := repository.NewRepository(CassandraSession)
	actHandler := handler.NewActHandler(repo)
	var reqBody = Roadmap1

	reqBodyBytes, _ := json.Marshal(reqBody)
	req := events.APIGatewayV2HTTPRequest{
		Body: string(reqBodyBytes),
	}

	resp, err := actHandler.HandleCreateActs(context.Background(), req)
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	resp, err = actHandler.HandleDeleteActs(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestReadAct(t *testing.T) {
	repo := repository.NewRepository(CassandraSession)
	actHandler := handler.NewActHandler(repo)
	var reqBody = Roadmap1

	reqBodyBytes, _ := json.Marshal(reqBody)
	req := events.APIGatewayV2HTTPRequest{
		Body: string(reqBodyBytes),
	}

	resp, err := actHandler.HandleReadActs(context.Background(), req)
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
