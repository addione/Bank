package manager

import (
	"errors"
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

func (um *UserManager) CreateNewUser(ur *models.UserRequest) (*models.User, error) {
	user, _ := um.userRepo.GetUserByEmail(ur.Email)
	if user.ID != 0 {
		return nil, errors.New("user Already Exists")
	}

	return um.userRepo.CreateNewUser(um.PrepareUser(ur)), nil

}

func (um *UserManager) GetUserById(userId int64) (*models.UserMysql, error) {
	user, err := um.userRepo.GetUserById(userId)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (um *UserManager) CleanDatabase() {
	um.userRepo.CleanDatabase()
}

func (um *UserManager) ListUsers() []*models.User {
	return um.userRepo.GetAllUsers()
}

func (um *UserManager) PrepareUser(u *models.UserRequest) *models.User {
	return &models.User{
		Name:        u.Name,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		Balance:     0,
		Address:     u.Address,
		Details:     u.Details,
	}
}

func (um *UserManager) CreateNewUserBO() *models.User {
	return um.userRepo.CreateNewUser(um.getUser())

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
