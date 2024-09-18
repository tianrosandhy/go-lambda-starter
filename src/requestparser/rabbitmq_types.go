package requestparser

type RabbitMQEventRequest struct {
	EventSource        string                     `json:"eventSource"`
	EventSourceArn     string                     `json:"eventSourceArn"`
	RmqMessagesByQueue map[string][]MessageDetail `json:"rmqMessagesByQueue"`
}

type MessageDetail struct {
	BasicProperties BasicProperty `json:"basicProperties"`
	Data            []byte        `json:"data"`
	Redelivered     bool          `json:"redelivered"`
}

type BasicProperty struct {
	AppID           interface{}       `json:"appID"`
	BodySize        int               `json:"bodySize"`
	ClusterID       interface{}       `json:"clusterId"`
	ContentEncoding interface{}       `json:"contentEncoding"`
	ContentType     interface{}       `json:"contentType"`
	CorrelationID   interface{}       `json:"correlationId"`
	DeliveryMode    int               `json:"deliveryMode"`
	Expiration      interface{}       `json:"expiration"`
	Headers         map[string]string `json:"headers"`
	MessageID       interface{}       `json:"messageId"`
	Priority        interface{}       `json:"priority"`
	ReplyTo         interface{}       `json:"replyTo"`
	Timestamp       interface{}       `json:"timestamp"`
	Type            interface{}       `json:"type"`
	UserID          interface{}       `json:"userId"`
}
