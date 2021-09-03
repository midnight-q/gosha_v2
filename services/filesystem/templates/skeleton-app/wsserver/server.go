package wsserver

import (
	"encoding/json"
	"net"
	"net/http"
	"sync"
	"time"

	"errors"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/google/uuid"
)

var instance *sync.Map
var once sync.Once

func GetConnections() *sync.Map {
	once.Do(func() {
		instance = &sync.Map{}
	})

	return instance
}

type UserMessage struct {
	Type string
	Data interface{}
}

type UserResponse struct {
	Type    string
	Result  interface{}
	Success bool
}

type UserConnection struct {
	Con       net.Conn
	WsToken   string
	HttpToken string
	UpdatedAt time.Time
}

func Run(host, port string) (err error) {
	setHandlers()
	err = http.ListenAndServe(host+":"+port, http.HandlerFunc(wssHandler))

	return
}

func setHandlers() {
	SetMessageHandler("SetToken", SetTokenHandler)
	SetMessageHandler("hello", HelloHandler)
}

func wssHandler(w http.ResponseWriter, r *http.Request) {

	code := uuid.New()

	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		// handle error
		return
	}

	cons := GetConnections()
	cons.Store(code, &UserConnection{
		Con:       conn,
		WsToken:   code.String(),
		HttpToken: "",
		UpdatedAt: time.Now(),
	})

	go runWssReader(cons, code)

}

func runWssReader(cons *sync.Map, code uuid.UUID) {

	defer closeConnection(code)

	for {

		if v, isOk := cons.Load(code); isOk {

			c, isOK := v.(*UserConnection)
			if !isOK {
				return
			}

			msg, op, err := wsutil.ReadClientData(c.Con)

			if err != nil {
				return
			}

			if len(msg) > 0 {

				cons.Store(code, c)
				message, err := ParseMessage(msg)
				if err != nil {
					continue
				}

				answer := GetMessageHandler(message.Type)(message, c)

				answer.Type = answer.Type + "Response"

				bAnswer, err := json.Marshal(answer)
				if err != nil {
					continue
				}
				err = wsutil.WriteServerMessage(c.Con, op, bAnswer)
				if err != nil {
					return
				}
			}

		} else {
			return
		}
	}
}

func ParseMessage(msg []byte) (UserMessage, error) {
	message := UserMessage{}
	err := json.Unmarshal(msg, &message)
	return message, err
}

func SendToUserByToken(token string, msg []byte) (err error) {
	cons := GetConnections()

	sent := false

	cons.Range(func(code interface{}, _ interface{}) bool {
		if v, ok := cons.Load(code); ok == true {
			if c, isOk := v.(*UserConnection); isOk {
				if c.HttpToken == token {
					err = wsutil.WriteServerMessage(c.Con, ws.OpText, msg)
					if err != nil {
						return false
					}
					sent = true
				}
			}
		}

		return true
	})

	if sent {
		return nil
	}

	return errors.New("User not found with token:" + token)
}

func closeConnection(code uuid.UUID) {

	cons := GetConnections()

	c, ok := cons.Load(code)

	if ok {

		if userConnection, isOk := c.(UserConnection); isOk {
			userConnection.Con.Close()
			cons.Delete(code)
		}
	}
}
