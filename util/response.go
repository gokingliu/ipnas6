package util

import (
	"encoding/json"
	"fmt"
	"ipnas6/entity"
	"net/http"
)

func JSONResponse(res http.ResponseWriter, status entity.Status, result interface{}) error {
	// 返回 JSON 文本
	res.Header().Set("Content-Type", "application/json")

	// 组装返回体
	resp := make(map[string]interface{})
	resp["code"] = status.Code
	resp["msg"] = status.Msg
	resp["result"] = result

	// 将数据 JSON 编码
	rsp, jsonErr := json.Marshal(resp)
	if jsonErr != nil {
		_, _ = fmt.Fprintln(res, jsonErr)
		resp["code"] = entity.JsonMarshalErr.Code
		resp["msg"] = entity.JsonMarshalErr.Msg
		resp["result"] = nil
		rsp, _ = json.Marshal(resp)
	}

	// 写入 HTTP Response
	_, writeErr := res.Write(rsp)
	if writeErr != nil {
		_, _ = fmt.Fprintln(res, writeErr)
	}

	return nil
}
