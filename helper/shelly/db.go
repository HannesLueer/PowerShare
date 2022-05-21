package shelly

import (
	"PowerShare/database"
	"PowerShare/models"
	"fmt"
	"log"
	"net/http"
)

func getHost(deviceId int64) (host string, err error) {
	sqlStatement := `SELECT host FROM shelly_devices WHERE device_id=$6`

	err = database.DB.QueryRow(sqlStatement, deviceId).Scan(&host)
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
		return "", fmt.Errorf("internal error")
	}

	return host, nil
}

func AddDeviceToDB(body models.ShellyAddRemoveDeviceBody) (errCode int, err error) {
	sqlStatement := `INSERT INTO shelly_devices (device_id, user_id, device_type, device_code, host) VALUES ($1, $2, $3, $4, $5) RETURNING device_id`

	_, err = database.DB.Exec(sqlStatement, body.DeviceId, body.UserId, body.DeviceType, body.DeviceCode, body.Host)
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
		return http.StatusInternalServerError, fmt.Errorf("an error has occurred")
	}

	return http.StatusOK, nil
}

func RemoveDeviceFromDB(body models.ShellyAddRemoveDeviceBody) (errCode int, err error) {
	sqlStatement := `DELETE FROM shelly_devices WHERE device_id = $1 AND user_id=$2`

	_, err = database.DB.Exec(sqlStatement, body.DeviceId, body.UserId)
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
		return http.StatusInternalServerError, fmt.Errorf("an error has occurred")
	}

	return http.StatusOK, nil
}
