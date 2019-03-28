/**
 * @author Jose Nidhin
 */
package handlers

import (
	"net/http"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	res := &apiSuccess{}
	res.Status = http.StatusOK
	res.Data = []string{"hello", "world"}
	renderResponse(w, http.StatusOK, res)
}
