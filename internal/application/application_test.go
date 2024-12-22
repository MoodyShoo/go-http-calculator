package application_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MoodyShoo/go-http-calculator/internal/application"
)

func TestHandlerLogic(t *testing.T) {
	cases := []struct {
		name       string
		statusCode int
		request    string
		want       string
	}{
		{
			name:       "Valid Sum",
			statusCode: http.StatusOK,
			request:    `{"expression": "5 + 3"}`,
			want:       `{"result":8}`,
		},
		{
			name:       "Invalid Expression",
			statusCode: http.StatusUnprocessableEntity,
			request:    `{"expression": "5 / 0"}`,
			want:       `{"error":"Expression is not valid"}`,
		},
		{
			name:       "Invalid Request Body",
			statusCode: http.StatusUnprocessableEntity,
			request:    `invalid json`,
			want:       `{"error":"Invalid request body"}`,
		},
		{
			name:       "Valid Double Sum",
			statusCode: http.StatusOK,
			request:    `{"expression": "2.5 + 3.412"}`,
			want:       `{"result":5.912}`,
		},
		{
			name:       "Expression Priority Order",
			statusCode: http.StatusOK,
			request:    `{"expression": "(2.5 + 3.412) / 2 + (1 * 2)"}`,
			want:       `{"result":4.9559999999999995}`,
		},
		{
			name:       "Spaces doesn't matter",
			statusCode: http.StatusOK,
			request:    `{"expression": "              3          * 2            + 1        "}`,
			want:       `{"result":7}`,
		},
		{
			name:       "Has letter in expression",
			statusCode: http.StatusUnprocessableEntity,
			request:    `{"expression": "360+5A"}`,
			want:       `{"error":"Expression is not valid"}`,
		},
		{
			name:       "Uncomplete expression",
			statusCode: http.StatusUnprocessableEntity,
			request:    `{"expression": "360+5-"}`,
			want:       `{"error":"Expression is not valid"}`,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewReader([]byte(tt.request)))
			w := httptest.NewRecorder()

			application.CalcHandler(w, req)

			if status := w.Code; status != tt.statusCode {
				t.Errorf("Expected status %d, got %d", tt.statusCode, status)
			}

			if got := w.Body.String(); got != tt.want {
				t.Errorf("Expected body %s, got %s", tt.want, got)
			}
		})
	}
}
