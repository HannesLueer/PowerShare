package charging

import (
	"PowerShare/handler/charging/helper/shelly"
	"net/http"
)

func DebugHandler(w http.ResponseWriter, r *http.Request) {

	shelly.TurnPowerOn()

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("hallo"))
}
