package userservice

import (
	"github.com/yukihiratype2/cookbook-server/internal/db"
	m "github.com/yukihiratype2/cookbook-server/internal/model/app"
)

// UserService struct
type UserService struct {
	UserDB *db.UserHandler
}

// New func
func New(uh *db.UserHandler) *UserService {
	us := UserService{}
	us.UserDB = uh
	return &us
}

func (us *UserService) Create(newUser *m.User) {
	newUser.Password = hashAndSaltPassword(newUser.Password)
	us.UserDB.Create(newUser)
}

func (us *UserService) Authentication(authUser *m.User) (err error) {
	findedUser, err := us.UserDB.FindByEmail(authUser.Email)
	if err != nil {
		return
	}
	err = verifyPassword(findedUser.Password, authUser.Password)
	return
}

func (us *UserService) GetUserByEmail(email string) (findedUser m.User, err error) {
	findedUser, err = us.UserDB.FindByEmail(email)
	return
}
