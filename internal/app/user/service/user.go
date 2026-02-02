package service

import (
	entity "Concurrencia/internal/app/user/entities"
	"Concurrencia/internal/app/user/repository"
	"fmt"
	"sync"
)

type Service struct {
	db UserRepository
}

func NewService(db UserRepository) *Service {
	return &Service{db: db}
}

func (s *Service) CreateUser(user *entity.User) error {

	var wg sync.WaitGroup

	wg.Add(2)

	var result1, result2 string

	go func() {
		defer wg.Done()
		result1 = repository.GetRepositoryOne()
	}()

	go func() {
		defer wg.Done()
		result2 = repository.GetRepositoryTwo()
	}()

	wg.Wait()
	/*result1 = repository.GetRepositoryOne()
	result2 = repository.GetRepositoryTwo()*/

	fmt.Println(result1)
	fmt.Println(result2)

	/*err := s.db.CreateUser(user)
	if err != nil {
		return err
	}*/
	return nil
}

func (s *Service) GetUser(id string) (*entity.User, error) {
	user, err := s.db.FindUserById(id)
	if err != nil {
		return nil, err
	}
	return &entity.User{
		Id:   user.Id,
		Name: user.Name,
	}, nil
}

func (s *Service) PublishTopic(msg *entity.Msg) {
	fmt.Println("El mensaje se publico correctamente", msg.Value)
}
