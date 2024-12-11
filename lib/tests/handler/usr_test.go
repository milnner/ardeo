package handler_test

import (
	"ardeolib.sapions.com/handler"
)

var (
	user1 = handler.UserBody{
		"test1@test.com",
		"test",
		true,
		"Ab1aaasdsdsds0@#$%",
	}
	user2 = handler.UserBody{
		"test2@test.com",
		"test",
		true,
		"Ab1aaasdsdsds0@#$%",
	}
	user3 = handler.UserBody{
		"test3@test.com",
		"test",
		true,
		"Ab1aaasdsdsds0@#$%",
	}
)

// func TestCreateUser(t *testing.T) {
// 	repo := repository.NewRepository(CassandraSession)
// 	usrHandler := handler.NewUserHandler(repo)
// 	var reqBody handler.UserBody

// 	reqBodyBytes, _ := json.Marshal(reqBody)
// 	req := events.APIGatewayV2HTTPRequest{
// 		Body: string(reqBodyBytes),
// 	}

// 	resp, err := usrHandler.HandleCreateUser(context.Background(), req)
// 	// Assert
// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusCreated, resp.StatusCode)
// }
