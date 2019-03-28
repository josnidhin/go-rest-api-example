/**
 * @author Jose Nidhin
 */
package handlers

import (
	"encoding/json"
	"net/http"
)

type apiResponse struct {
	Status int `json:"status"`
}

type apiSuccess struct {
	apiResponse
	Data interface{} `json:"data"`
}

type apiError struct {
	apiResponse
	Message string `json:"message"`
}

func render(w http.ResponseWriter, status int, data interface{}) {
	jsonData, err := json.Marshal(data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	w.Write([]byte(jsonData))
}

func renderResponse(w http.ResponseWriter, status int, data *apiSuccess) {
	render(w, status, data)
}

func renderError(w http.ResponseWriter, status int, data *apiError) {
	render(w, status, data)
}

func Default404(w http.ResponseWriter, r *http.Request) {
	res := &apiError{}
	res.Status = http.StatusNotFound
	renderError(w, http.StatusNotFound, res)
}
