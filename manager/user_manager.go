package manager

import (
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/addione/New/models"
	"github.com/addione/New/repository"
)

type UserManager struct {
	userRepo *repository.UserRepo
}

func newUserManager(mdi *ManagerDIContainer) *UserManager {
	return &UserManager{
		userRepo: mdi.repositoryDIContainer.GetUserRepo(),
	}
}

func (um *UserManager) CreateNewUser() {
	um.userRepo.CreateNewUser(um.getUser())

}

func (um *UserManager) CleanDatabase() {
	um.userRepo.CleanDatabase()
}

func (um *UserManager) getUser() *models.User {

	var name string
	salutation := randomdata.Title(randomdata.RandomGender)

	switch salutation {
	case "Mr":
		name = randomdata.FirstName(randomdata.Male)
	default:
		name = randomdata.FirstName(randomdata.Female)
	}

	lastname := randomdata.LastName()

	user := models.User{
		Name:        name,
		Email:       name + lastname + `@gmail.com`,
		Pass:        "pass",
		PhoneNumber: randomdata.PhoneNumber(),
		Balance:     float64(randomdata.Number(1000, 2000000)),
		Salutation:  salutation,
		Address:     randomdata.Address(),
		Details: models.Details{
			FirstName: name,
			LastName:  lastname,
		},
		CreatedAt: time.Now(),
	}

	return &user
}
