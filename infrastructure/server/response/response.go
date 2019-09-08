package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// Success HTTPコード:200 正常終了を処理する
func Success(writer http.ResponseWriter, res interface{}) {
	data, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		httpError(writer, http.StatusInternalServerError, "Failed Json Marshal")
		return
	}
	writer.Write(data)
}

// BadRequest HTTPコード:400 BadRequestを処理する
func BadRequest(writer http.ResponseWriter, message string) {
	httpError(writer, http.StatusBadRequest, message)
}

// InternalServerError HTTPコード:500 InternalServerErrorを処理する
func InternalServerError(writer http.ResponseWriter, message string) {
	httpError(writer, http.StatusInternalServerError, message)
}

// httpError エラー用のレスポンス出力を行う
func httpError(writer http.ResponseWriter, code int, message string) {
	data, _ := json.Marshal(errorResponse{
		Code:    code,
		Message: message,
	})
	writer.WriteHeader(code)
	if data != nil {
		writer.Write(data)
	}
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
