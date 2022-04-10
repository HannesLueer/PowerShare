package testdata

import (
	"PowerShare/handler/chargingStation"
	"PowerShare/handler/user"
	"PowerShare/models"
	"fmt"
	"log"
	"math/rand"
)

func FillDB() {
	fillUsers()
	fillChargers()

	log.Println("DB filled")
}

const numberUsers = 50
const numberChargers = 5_000

func fillUsers() {
	for userCount := 0; userCount < numberUsers; userCount++ {
		_, _, err := user.SignUp(models.User{
			ID:           0,
			Name:         fmt.Sprintf("User%d", userCount),
			Email:        fmt.Sprintf("user%d@test.com", userCount),
			Password:     "123",
			PasswordHash: "",
			Role:         0,
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func fillChargers() {
	for userCount := 0; userCount < numberUsers; userCount++ {
		for chargerCount := 0; chargerCount < numberChargers/numberUsers; chargerCount++ {
			pos, title := getRandomPosition()
			chargingStation.CreateCharger(models.Charger{
				Title:      title,
				Position:   pos,
				Cost:       3,
				IsOccupied: false,
			},
				fmt.Sprintf("user%d@test.com", userCount),
			)
		}
	}
}

func getRandomPosition() (models.Coordinate, string) {
	lat := rand.Float64()*360 - 180
	lng := rand.Float64()*180 - 90

	return models.Coordinate{
			Lat: lat,
			Lng: lng,
		},
		fmt.Sprintf("Lng: %f Lat: %f", lng, lat)
}
