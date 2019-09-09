package service

import (
	"CleanArchitecture_SampleApp/domain"
	"CleanArchitecture_SampleApp/usecase/repository"
	"log"

	"github.com/google/uuid"
)

type authService struct {
	UserRepository repository.UserRepository
}

type AuthService interface {
	CreateUser(userName *string) (*string, error)
}

func NewAuthService(ur repository.UserRepository) AuthService {
	return &authService{UserRepository: ur}
}

func (authService *authService) CreateUser(userName *string) (*string, error) {
	// UUIDでユーザIDを生成する
	userID, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	userIDString := userID.String()

	// UUIDで認証トークンを生成する
	authToken, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	authTokenString := authToken.String()

	user := domain.User{
		UserID:    userIDString,
		AuthToken: authTokenString,
		Name:      *userName,
	}

	// データベースにユーザデータを登録する
	err = authService.UserRepository.Insert(user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &authTokenString, nil
}
