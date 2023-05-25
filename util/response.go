package util

import (
	"encoding/json"
	"fmt"
	"ipnas6/entity"
	"net/http"
)

func JSONResponse(res http.ResponseWriter, status entity.Status, result interface{}) error {
	res.Header().Set("Content-Type", "application/json")

	resp := make(map[string]interface{})
	resp["code"] = status.Code
	resp["msg"] = status.Msg
	resp["result"] = result

	rsp, jsonErr := json.Marshal(resp)
	if jsonErr != nil {
		_, _ = fmt.Fprintln(res, jsonErr)
		resp["code"] = entity.JsonMarshalErr.Code
		resp["msg"] = entity.JsonMarshalErr.Msg
		resp["result"] = nil
		rsp, _ = json.Marshal(resp)
	}

	if _, err := res.Write(rsp); err != nil {
		_, _ = fmt.Fprintln(res, err)
	}

	return nil
}
