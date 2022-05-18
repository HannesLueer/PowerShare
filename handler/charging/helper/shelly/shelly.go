package shelly

import (
	"PowerShare/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"math/rand"
	"net/http"
	"os"
)

type Mode string

const (
	on  Mode = "on"
	off Mode = "off"
)

func TurnPowerOn() {
	turnPower("", -1, on) //TODO
}

func TurnPowerOff() {
	turnPower("", -1, off) //TODO
}

func turnPower(host string, deviceId int, mode Mode) (err error, httpStatusCode int) {
	println(mode)
	println(mode == on)
	println(os.Getenv("SHELLY_INTEGRATOR_TAG"))
	println(getAccessToken())

	// open web socket
	socketUrl := fmt.Sprintf("wss://%s:6113/shelly/wss/hk_sock?t=%s", host, getAccessToken())
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		return fmt.Errorf("error connecting to websocket server: %s", err), http.StatusInternalServerError
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
		return fmt.Errorf("error while setting up the request: %s", err), http.StatusInternalServerError
	}

	// write to web socket
	err = conn.WriteMessage(websocket.TextMessage, requestJson)
	if err != nil {
		return fmt.Errorf("error during writing to websocket: %s", err), http.StatusInternalServerError
	}

	return nil, http.StatusOK
}
