package stub

import (
	"errors"
	"fmt"
)

type InvalidIDError struct {
	ID string
}

func (ie *InvalidIDError) Error() string {
	return fmt.Sprintf("invalid id: %s", ie.ID)
}

func (ie *InvalidIDError) Is(err error) bool {
	if e, ok := errors.AsType[*InvalidIDError](err); ok {
		return e.ID == ie.ID
	}
	return false
}

type User struct{}
type Pet struct {
	Name string
}
type Person struct{}

type Entities interface {
	GetUser(id string) (User, error)
	GetPets(userID string) ([]Pet, error)
	GetChildren(userID string) ([]Person, error)
	GetFriends(userID string) ([]Person, error)
	SaveUser(user User) error
}

type Logic struct {
	Entities Entities
}

func (l Logic) GetPetNames(userId string) ([]string, error) {
	pets, err := l.Entities.GetPets(userId)
	if err != nil {
		return nil, err
	}
	out := make([]string, len(pets))
	for _, p := range pets {
		out = append(out, p.Name)
	}
	return out, nil
}
