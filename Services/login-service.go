package Services

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"gg-cms/DTOs"
	"gg-cms/DataRepos"
	"gg-cms/Models"
	"hash"
	"os"
	"time"
)

type LoginService interface {
	Login(username string, password string) (bool, bool)
	Register(registration DTOs.Registration) (Models.User, error)
}

type loginService struct {
	userRepo DataRepos.UserRepo
}

func NewLoginService(repo DataRepos.UserRepo) LoginService {
	return &loginService{
		userRepo: repo,
	}
}

func (service *loginService) Login(username string, password string) (bool, bool) {
	exists, err := service.userRepo.ExistsAny(true)
	isAdmin := false
	if especialUser(username, password) && !exists && err == nil {
		return true, true
	}
	encodedPassword := EncodePassword(username, password)
	exists, isAdmin, err = service.userRepo.ExistsCredential(username, encodedPassword)
	if exists && err == nil {
		return true, isAdmin
	}

	return false, false
}

func (service *loginService) Register(registration DTOs.Registration) (Models.User, error) {
	realPassword := EncodePassword(registration.UserName, registration.Password)
	var user = Models.User{
		UserName: registration.UserName,
		Password: realPassword,
		Status: "P",
		Email: registration.Email,
		CreatedBy:  registration.CreatedBy,
		CreatedDate: time.Now(),
	}

	//TODO: add module of send mail to confirm account
	user, err := service.userRepo.Insert(user)

	return user, err
}


func especialUser(username string, password string) bool {
	secret := os.Getenv("ESPECIAL_USER")
	if secret != "" && secret == username + ":" + password {
		return true
	}
	return false
}

func EncodePassword(username string, password string) string {
	secret := getSecretForPassword() + username
	data := password
	strHash := generateHMAC(sha512.New, secret, data)
	return strHash
}

func getSecretForPassword() string {
	secret := os.Getenv("PASSWORD_SECRET")
	if secret == "" {
		return "$ecr3ct"
	}
	return secret
}

func generateHMAC(hash func()hash.Hash, secret string, data string) string{
	h := hmac.New(hash, []byte(secret))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}