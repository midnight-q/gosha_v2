package router

import (
    "skeleton-app/wsserver"
)

func HandleWss(msg wsserver.UserMessage, con *wsserver.UserConnection) (answer wsserver.UserResponse) {

    answer = wsserver.UserResponse{
        Type:    msg.Type,
        Result:  "response from home",
        Success: true,
    }

    return
}
