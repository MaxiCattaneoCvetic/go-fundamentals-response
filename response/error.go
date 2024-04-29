package response

/*

	v0.0.1 -> siempre los releases son de este formato


*/
import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func InternalServerError(msg string) Response {
	return error(msg, http.StatusInternalServerError)
}
func NotFound(msg string) Response {
	return error(msg, http.StatusNotFound)
}

func UnAuthorized(msg string) Response {
	return error(msg, http.StatusUnauthorized)
}

func Forbidden(msg string) Response {
	return error(msg, http.StatusForbidden)
}
func BadRequest(msg string) Response {
	return error(msg, http.StatusBadRequest)
}

func error(msg string, code int) Response {
	return &ErrorResponse{
		Message: msg,
		Status:  code,
	}
}

func (e ErrorResponse) Error() string {
	return ""
}
func (e ErrorResponse) StatusCode() int {
	return e.Status
}
func (e ErrorResponse) GetBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e ErrorResponse) GetData() interface{} {
	return nil

}
