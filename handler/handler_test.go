package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	"testing"
	"todo/mock"
	"todo/model"
)

func TestHandler_CreateTodo(t *testing.T) {

	response := model.Todo{ID: "123", Text: "test"}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockITodoService(ctrl)
	mockService.EXPECT().CreateTodo("test").Return(response)

	handler := NewTodoHandler(mockService)
	req := httptest.NewRequest("POST", "/todo", nil)
	req.Form = map[string][]string{"text": {"test"}}

	app := fiber.New()
	app.Post("/todo", handler.CreateTodo)

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("expected status code to be 200, got %d", resp.StatusCode)
	}
}
