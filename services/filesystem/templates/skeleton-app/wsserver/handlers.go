package wsserver

import (
	"skeleton-app/types"
)

type WsHandler func(msg UserMessage, userConn *UserConnection) (answer UserResponse)

type RpcMessage struct {
	Message UserMessage
}

var defaultHandler = func(msg UserMessage, userConn *UserConnection) (answer UserResponse) {
	return UserResponse{
		Result:  "Unsupported message type",
		Success: false,
	}
}

var handlers = map[string]WsHandler{}

func SetMessageHandler(messageType string, h WsHandler) {
	handlers[messageType] = h
}

func GetMessageHandler(messageType string) (h WsHandler) {
	h, ok := handlers[messageType]
	if !ok {
		h = defaultHandler
	}
	return
}

func SetTokenHandler(msg UserMessage, userConn *UserConnection) (answer UserResponse) {
	if msg.Data != nil {
		token := msg.Data.(string)

		auth := types.Authenticator{Token: token}
		if auth.IsAuthorized() {
			userConn.HttpToken = token
			answer = UserResponse{
				Type:    msg.Type,
				Result:  "token was set",
				Success: true,
			}
			return
		}
	}
	answer = UserResponse{
		Type:    msg.Type,
		Result:  "invalid token",
		Success: false,
	}
	return
}

func HelloHandler(msg UserMessage, userConn *UserConnection) (answer UserResponse) {
	name := ""
	if msg.Data == nil {
		name = "Anonymous"
	} else {
		name = msg.Data.(string)
	}

	answer = UserResponse{
		Type:    msg.Type,
		Result:  "Hello " + name,
		Success: true,
	}
	return
}
