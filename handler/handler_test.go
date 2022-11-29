package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
	"todo/mock"
	"todo/model"
)

var buf bytes.Buffer

func TestHandler_CreateTodo(t *testing.T) {

	var responseObj model.Todo
	mockResponse := model.Todo{ID: "123", Text: "buy some milk"}

	json.NewEncoder(&buf).Encode(
		CreatTodoRequest{Text: "buy some milk"},
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockITodoService(ctrl)
	mockService.
		EXPECT().
		CreateTodo("buy some milk").
		Return(mockResponse).
		Times(1)

	handler := NewTodoHandler(mockService)

	app := fiber.New()
	app.Post("/todo", handler.CreateTodo)

	req := httptest.NewRequest("POST", "/todo", &buf)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}
	asd, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(asd, &responseObj)
	if err != nil {
		t.Errorf("Cannot unmarshal response: %v", err)
	}

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, mockResponse, responseObj)
}
