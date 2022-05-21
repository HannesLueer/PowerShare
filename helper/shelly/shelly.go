package shelly

import (
	"PowerShare/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"math/rand"
	"net/http"
)

type Mode string

const (
	On  Mode = "on"
	Off Mode = "off"
)

func TurnPower(deviceId int64, mode Mode) (httpStatusCode int, err error) {
	// get host
	host, err := getHost(deviceId)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("unable to find the host of the shelly device")
	}

	// open web socket
	socketUrl := fmt.Sprintf("wss://%s:6113/shelly/wss/hk_sock?t=%s", host, getAccessToken())
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error connecting to websocket server: %s", err)
	}
	defer conn.Close()

	// define request
	request := models.ShellyOnOffCommandRequest{
		Event:    "Shelly:CommandRequest",
		TrId:     rand.Intn(9999),
		DeviceId: deviceId,
		Data: models.ShellyOnOffCommandRequestData{
			Cmd: "relay",
			Params: models.ShellyOnOffCommandRequestDataParams{
				Id:   0,
				Turn: string(mode),
			},
		},
	}
	requestJson, err := json.Marshal(request)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error while setting up the request: %s", err)
	}

	// write to web socket
	err = conn.WriteMessage(websocket.TextMessage, requestJson)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error during writing to websocket: %s", err)
	}

	return http.StatusOK, nil
}
