/**
 * @author Jose Nidhin
 */
package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func parseJsonRequest(r *http.Request, reqData interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024*1024))

	if err != nil {
		return err
	}

	err = json.Unmarshal(body, reqData)

	if err != nil {
		return err
	}

	err = validate.Struct(reqData)

	if err != nil {
		return err
	}

	return nil
}
