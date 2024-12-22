package application

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/MoodyShoo/go-http-calculator/pkg/calculation"
)

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

type Request struct {
	Expression string `json:"expression"`
}

type Response interface {
	ToJSON() ([]byte, error)
}

type SuccessResponse struct {
	Result float64 `json:"result"`
}

func (r SuccessResponse) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (r ErrorResponse) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

func sendResponse(w http.ResponseWriter, response Response, status int) {
	w.WriteHeader(status)
	resp, err := response.ToJSON()
	if err != nil {
		http.Error(w, `{"error": "Failed to encode response"}`, http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendResponse(w, ErrorResponse{Error: "Invalid request body"}, http.StatusUnprocessableEntity)
		return
	}

	log.Printf("Request: {\"expression\" : %s}", req.Expression)

	result, err := calculation.Calc(req.Expression)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		sendResponse(w, ErrorResponse{Error: "Expression is not valid"}, http.StatusUnprocessableEntity)
	} else {
		log.Printf("Result: %f", result)
		sendResponse(w, SuccessResponse{Result: result}, http.StatusOK)
	}
}

func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	log.Printf("Server running on: %s", a.config.Addr)
	return http.ListenAndServe(":"+a.config.Addr, nil)
}
