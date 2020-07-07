package Services

import (
	"gg-cms/DTOs"
	"gg-cms/DataRepos"
	"gg-cms/Models"
	"context"
)

type UserService interface {
	Save(user Models.User) (DTOs.User, error)
	Update(user Models.User) (DTOs.User, error)
	Delete(ID string) error
	Find(username string) (DTOs.User, error)
	FindAll(limit int64, skip int64) ([]DTOs.User, error)
}

type userService struct {
	userRepo DataRepos.UserRepo

}

func NewUserService(repo DataRepos.UserRepo) UserService {
	return &userService{
		userRepo: repo,
	}
}

func (service *userService) Save(user Models.User) (DTOs.User, error) {
	repoUser, err := service.userRepo.Insert(user)
	return DTOs.User{
		ID:          repoUser.ID,
		UserName:    repoUser.UserName,
		Status:      repoUser.Status,
		IsAdmin:     repoUser.IsAdmin,
		Email:       repoUser.Email,
		CreatedDate: repoUser.CreatedDate,
		CreatedBy:   repoUser.CreatedBy,
	}, err
}

func (service *userService) Update(user Models.User) (DTOs.User, error) {
	repoUser, err := service.userRepo.Update(user)
	return DTOs.User{
		ID:          repoUser.ID,
		UserName:    repoUser.UserName,
		Status:      repoUser.Status,
		IsAdmin:     repoUser.IsAdmin,
		Email:       repoUser.Email,
		CreatedDate: repoUser.CreatedDate,
		CreatedBy:   repoUser.CreatedBy,
	}, err
}

func (service *userService) Delete(ID string) error {
	err := service.userRepo.Delete(ID)
	return  err
}

func (service *userService) Find(username string) (DTOs.User, error) {
	repoUser, err := service.userRepo.Get(username)
	return DTOs.User{
		ID:          repoUser.ID,
		UserName:    repoUser.UserName,
		Status:      repoUser.Status,
		IsAdmin:     repoUser.IsAdmin,
		Email:       repoUser.Email,
		CreatedDate: repoUser.CreatedDate,
		CreatedBy:   repoUser.CreatedBy,
	}, err
}

func (service *userService) FindAll(limit int64, skip int64) ([]DTOs.User, error) {
	var results = make([]DTOs.User, 0)
	cur, err := service.userRepo.GetAllUsers(limit, skip)

	if err != nil {
		return results, err
	}
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem DTOs.User
		err := cur.Decode(&elem)
		if err != nil {
			return results, err
		}

		results = append(results, elem)
	}

	return results, nil

}