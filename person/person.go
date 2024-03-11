package person

import (
	"github.com/go-faker/faker/v4"
	"fmt"
)


type coordinates struct {
	lat float64
	long float64
}

type person struct {
	firstname string
	lastname string
	coordinates 
	email string
	phone string
	dateOfBirth string
}





func GetFakeUser () person {
	person := person {
		firstname: faker.FirstName(),
		lastname: faker.LastName(),
		coordinates: coordinates {
			lat: faker.Longitude(),
			long: faker.Latitude(),
		},
		email: faker.Email(),
		phone: faker.Phonenumber(),
		dateOfBirth: faker.Date(),
	}

	fmt.Println(person)

	return person

}