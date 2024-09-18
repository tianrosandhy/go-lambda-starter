package requestparser

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type APIResp struct {
	Type    string      `json:"type"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func HTTPErrorResp(err error, httpCode int, errorCode ...string) events.APIGatewayProxyResponse {
	if len(errorCode) == 0 {
		errorCode = []string{"ERR_GENERIC"}
	}

	bodyResp := APIResp{
		Type:    "error",
		Code:    errorCode[0],
		Message: err.Error(),
	}

	return events.APIGatewayProxyResponse{
		StatusCode: httpCode,
		Body:       bodyParser(bodyResp),
	}
}

func HttpSuccessResp(message string, data ...interface{}) events.APIGatewayProxyResponse {
	bodyResp := APIResp{
		Type:    "success",
		Code:    "SUCCESS",
		Message: message,
	}
	if len(data) > 0 {
		bodyResp.Data = data[0]
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       bodyParser(bodyResp),
	}
}

func bodyParser(body APIResp) string {
	bodyByte, _ := json.Marshal(body)
	return string(bodyByte)
}
