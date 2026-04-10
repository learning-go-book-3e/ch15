package stub

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type getPetNamesStub struct {
	Entities
}

func (ps getPetNamesStub) GetPets(userID string) ([]Pet, error) {
	switch userID {
	case "1":
		return []Pet{{Name: "Bubbles"}}, nil
	case "2":
		return []Pet{{Name: "Stampy"}, {Name: "Snowball II"}}, nil
	default:
		return nil, &InvalidIDError{ID: userID}
	}
}

func TestLogicGetPetNames(t *testing.T) {
	data := []struct {
		name     string
		userID   string
		petNames []string
		err      error
	}{
		{"case1", "1", []string{"Bubbles"}, nil},
		{"case2", "2", []string{"Stampy", "Snowball II"}, nil},
		{"case3", "3", nil, &InvalidIDError{ID: "3"}},
	}
	l := Logic{getPetNamesStub{}}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			petNames, err := l.GetPetNames(d.userID)
			if !errors.Is(err, d.err) {
				t.Errorf("expected error %v, got %v", d.err, err)
			}
			if diff := cmp.Diff(d.petNames, petNames); diff != "" {
				t.Error(diff)
			}
		})
	}
}

type entitiesStub struct {
	getUser     func(id string) (User, error)
	getPets     func(userID string) ([]Pet, error)
	getChildren func(userID string) ([]Person, error)
	getFriends  func(userID string) ([]Person, error)
	saveUser    func(user User) error
}

func (es entitiesStub) GetUser(id string) (User, error) {
	return es.getUser(id)
}

func (es entitiesStub) GetPets(userID string) ([]Pet, error) {
	return es.getPets(userID)
}

func (es entitiesStub) GetChildren(userID string) ([]Person, error) {
	return es.getChildren(userID)
}

func (es entitiesStub) GetFriends(userID string) ([]Person, error) {
	return es.getFriends(userID)
}

func (es entitiesStub) SaveUser(user User) error {
	return es.saveUser(user)
}

func TestLogicGetPetNames2(t *testing.T) {
	data := []struct {
		name     string
		getPets  func(userID string) ([]Pet, error)
		userID   string
		petNames []string
		err      error
	}{
		{"case1", func(userID string) ([]Pet, error) {
			return []Pet{{Name: "Bubbles"}}, nil
		}, "1", []string{"Bubbles"}, nil},
		{"case2", func(userID string) ([]Pet, error) {
			return []Pet{{Name: "Stampy"}, {Name: "Snowball II"}}, nil
		}, "2", []string{"Stampy", "Snowball II"}, nil},
		{"case3", func(userID string) ([]Pet, error) {
			return nil, &InvalidIDError{ID: "3"}
		}, "3", nil, &InvalidIDError{ID: "3"}},
	}
	l := Logic{}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			l.Entities = entitiesStub{getPets: d.getPets}
			petNames, err := l.GetPetNames(d.userID)
			if diff := cmp.Diff(d.petNames, petNames); diff != "" {
				t.Error(diff)
			}
			if !errors.Is(err, d.err) {
				t.Errorf("Expected error %v, got %v", d.err, err)
			}
		})
	}
}
