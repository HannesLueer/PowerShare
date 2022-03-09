package locations

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

type Location struct {
	DisplayName string
	Name string
	Coordinate Coordinate
}

type Coordinate struct {
	Lat float64
	Lng float64
}


func LocationsHandler(w http.ResponseWriter, r *http.Request) {
	/*allLocations := [...]Location{
		{
			DisplayName: "Test1",
			Name:        "t1",
			Coordinate:  Coordinate{Lat: 50, Lng: 10},
		},
		{
			DisplayName: "Test2",
			Name:        "t2",
			Coordinate:  Coordinate{Lat: 51, Lng: 11},
		},
	}*/

	var allLocations []Location

	for lng := -180.0; lng < 180.0; lng+=1 {
		for lat := -90.0; lat < 90.0; lat+=5 {
			for j := 1.0; j < 3.0; j+=1 {
				lng_ := lng + j * rand.Float64()
				lat_ := lat + j * rand.Float64()
				for i := 1.0; i < 3.0; i+=1 {
					lng_ += i * 0.003 * rand.Float64()
					lat_ += i * -0.003 * rand.Float64()
					allLocations = append(allLocations, Location{
						DisplayName: "lng: " + strconv.FormatFloat(lng_, 'f', 5, 64) + " lat: " + strconv.FormatFloat(lat_, 'f', 5, 64),
						Name:        "lng: " + strconv.FormatFloat(lng_, 'f', 5, 64) + " lat: " + strconv.FormatFloat(lat_, 'f', 5, 64),
						Coordinate:  Coordinate{Lng: lng_, Lat: lat_},
					})
				}
			}
		}
	}
	fmt.Println(len(allLocations))

	jsonResp, err := json.Marshal(allLocations)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(allLocations))

	w.Write(jsonResp)
	return
}
