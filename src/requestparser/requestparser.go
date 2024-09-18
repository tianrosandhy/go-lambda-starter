package requestparser

import (
	"encoding/base64"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

// ParseEventRequest will transform either RabbitMQ event or Raw Event to payload
func ParseEventRequest(request map[string]interface{}, payload interface{}) {
	rawByte, _ := json.Marshal(request)

	// try to parse to RabbitMQEventRequest first
	rabbitRequest := RabbitMQEventRequest{}
	json.Unmarshal(rawByte, &rabbitRequest)
	if rabbitRequest.RmqMessagesByQueue != nil {
		log.Printf("ParseEventRequest : TRY MODE RABBIT MQ EVENT")
		parseAsRabbitMQEvent(rabbitRequest, payload)
		return
	}

	// try to parse as HTTP request
	httpRequest := events.APIGatewayProxyRequest{}
	json.Unmarshal(rawByte, &httpRequest)
	if httpRequest.Body != "" && httpRequest.HTTPMethod != "" {
		log.Printf("ParseEventRequest : TRY MODE HTTP EVENT")
		parseAsHTTPEvent(httpRequest, payload)
		return
	}

	// default: try to parse raw request as JSON
	log.Printf("ParseEventRequest : TRY MODE RAW EVENT")
	json.Unmarshal(rawByte, payload)
}

func parseAsRabbitMQEvent(rabbitRequest RabbitMQEventRequest, payload interface{}) {
	for rabbitQueue, rabbitData := range rabbitRequest.RmqMessagesByQueue {
		if len(rabbitData) > 0 {
			log.Printf("ParseRequest Try parsing QUEUE %s data", rabbitQueue)
			rabbitPayload := rabbitData[0].Data

			// try decode base64 first
			decodedPayload, err := base64.StdEncoding.DecodeString(string(rabbitPayload))
			if err != nil {
				// fail to decode base64, try to parse as JSON instead
				json.Unmarshal(rabbitPayload, payload)
			} else {
				json.Unmarshal(decodedPayload, payload)
			}
		}
		break
	}
}

func parseAsHTTPEvent(httpRequest events.APIGatewayProxyRequest, payload interface{}) {
	body := httpRequest.Body
	if httpRequest.IsBase64Encoded {
		bodyBytes, err := base64.StdEncoding.DecodeString(body)
		if err != nil {
			log.Printf("FAIL parseAsHTTPEvent : %+v", err)
			return
		}
		body = string(bodyBytes)
	}

	err := json.Unmarshal([]byte(body), payload)
	if err != nil {
		log.Printf("FAIL parseAsHTTPEvent UNMARSHAL : %+v", err)
	}
}
