package repository

import (
	"github.com/auth/src/context/security/authentication/domain/interfaces"
)

type MongoAuthRepository struct {
	repo interfaces.IAuthenticationRepository
}

func NewMongoAuthRepository(repo interfaces.IAuthenticationRepository) *MongoAuthRepository {
	return &MongoAuthRepository{
		repo: repo,
	}
}

func (s *MongoAuthRepository) GetById(id string) {

}

func (s *MongoAuthRepository) Create(request string) {

}
