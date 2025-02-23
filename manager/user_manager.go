package manager

import (
	"errors"
	"strconv"
	"time"

	"github.com/addione/New/helpers"
	"github.com/addione/New/src/request"

	"github.com/Pallinder/go-randomdata"
	"github.com/addione/New/models"
	"github.com/addione/New/repository"
)

type UserManager struct {
	userRepo *repository.UserRepo
	hash     *helpers.Hashing
	jwt      *helpers.JwtHelper
}

func newUserManager(mdi *ManagerDIContainer) *UserManager {

	return &UserManager{
		userRepo: mdi.repositoryDIContainer.GetUserRepo(),
		hash:     mdi.helpers.GetHashing(),
		jwt:      mdi.helpers.GetJwtTokenHelper(),
	}
}

func (um *UserManager) UpdateUser(userId int64, ur *models.UserUpdateRequest) error {
	return um.userRepo.UpdateUserByID(userId, ur)
}

func (um *UserManager) CreateNewUser(ur *request.CreateUserParams) (*models.User, error) {
	ur.Email = strconv.FormatInt(time.Now().Unix(), 10) + ur.Email
	// ur.Email = fmt.Sprintf("%f", time.Now().Unix()) + ur.Email

	user, _ := um.userRepo.GetUserByEmail(ur.Email)
	if user.ID != 0 {
		return nil, errors.New("user Already Exists")
	}
	preparedUser, err := um.PrepareUser(ur)
	if err != nil {
		return nil, err
	}
	return um.userRepo.CreateNewUser(preparedUser), nil

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

func (um *UserManager) PrepareUser(u *request.CreateUserParams) (*models.User, error) {
	password, err := um.hash.HashPassword(u.Password)
	if err != nil {
		return &models.User{}, err
	}

	return &models.User{
		Name:        u.Name,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		Balance:     0,
		Address:     u.Address,
		Details:     u.Details,
		Pass:        password,
	}, nil
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

	password, _ := um.hash.HashPassword("test")

	user := models.User{
		Name:        name,
		Email:       name + lastname + `@gmail.com`,
		Pass:        password,
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

func (um *UserManager) ValidateCredentialsAndGetToken(loginParams *models.UserLoginRequest) (string, error) {
	user, err := um.userRepo.ValidateAndGetCredentials(loginParams)
	if err != nil {
		return "", err
	}
	if !um.hash.CheckPassword(loginParams.Password, user.Password) {
		return " ", errors.New("invalid username or password")
	}

	token, err := um.jwt.GenerateToken(user.Email, user.ID)

	if err != nil {
		return token, err
	}

	return token, nil
}
