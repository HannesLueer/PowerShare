package testdata

import (
	"PowerShare/handler/chargingStation"
	"PowerShare/models"
	"fmt"
	"math/rand"
	"strconv"
)

func FillDB(){
	fillChargers()

	fmt.Println("DB filled")
}

func fillChargers() {
	for lng := -180.0; lng < 180.0; lng+=3 {
		for lat := -90.0; lat < 90.0; lat+=3 {
			for j := 1.0; j < 4.0; j+=1 {
				lng_ := lng + j * rand.Float64()
				lat_ := lat + j * rand.Float64()
				for i := 1.0; i < 4.0; i+=1 {
					lng_ += i * 0.003 * rand.Float64()
					lat_ += i * -0.003 * rand.Float64()

					chargingStation.CreateCharger(models.Charger{
						Title: "lng: " + strconv.FormatFloat(lng_, 'f', 5, 64) + " lat: " + strconv.FormatFloat(lat_, 'f', 5, 64),
						Position: models.Coordinate{Lat: lng_, Lng: lat_},
						Cost: 3,
						IsOccupied: false,
					})
				}
			}
		}
	}
}
