package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainHandlerSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?city=moscow&count=2", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler вернул неправильный код состояния: получили %v, ожидали %v",
			status, http.StatusOK)
	}

	if responseRecorder.Body.String() == "" {
		t.Errorf("handler вернул пустое тело ответа")
	}
}

func TestMainHandlerWrongCity(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?city=paris&count=2", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler вернул неправильный код состояния: получили %v, ожидали %v",
			status, http.StatusBadRequest)
	}

	expected := "wrong city value"
	if responseRecorder.Body.String() != expected {
		t.Errorf("handler вернул неожиданное тело ответа: получили %v, ожидали %v",
			responseRecorder.Body.String(), expected)
	}
}

func TestMainHandlerCountMoreThanTotal(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?city=moscow&count=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler вернул неправильный код состояния: получили %v, ожидали %v",
			status, http.StatusOK)
	}

	expected := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"
	if responseRecorder.Body.String() != expected {
		t.Errorf("handler вернул неожиданное тело ответа: получили %v, ожидали %v",
			responseRecorder.Body.String(), expected)
	}
}
