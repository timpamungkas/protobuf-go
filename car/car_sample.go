package car

import (
	"log"
	"my-protobuf/protogen/car"

	"github.com/google/uuid"
)

func ValidateCar() {
	car := car.Car{
		CarId: uuid.New().String(),
		Brand: "Bmw",
		Model: "735i Sport",
		Price: 155000,
	}

	err := car.ValidateAll()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(&car)
}
