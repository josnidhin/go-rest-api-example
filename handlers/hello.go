/**
 * @author Jose Nidhin
 */
package handlers

import (
	"net/http"
)

type customHelloInput struct {
	Name string `json:"name" validate:"required"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	res := &apiSuccess{}
	res.Status = http.StatusOK
	res.Data = []string{"hello", "world"}
	renderResponse(w, http.StatusOK, res)
	return
}

func CustomHello(w http.ResponseWriter, r *http.Request) {
	reqData := &customHelloInput{}
	err := parseJsonRequest(r, reqData)

	if err != nil {
		res := &apiError{}
		res.Status = http.StatusBadRequest
		res.Message = "Bad request"
		renderError(w, http.StatusBadRequest, res)
		return
	}

	res := &apiSuccess{}
	res.Status = http.StatusOK
	res.Data = []string{"hello", reqData.Name}
	renderResponse(w, http.StatusOK, res)
	return
}
