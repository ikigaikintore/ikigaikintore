package user

import (
	"context"
	"encoding/json"
	"github.com/ervitis/control-panel/common/global/mock"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/ervitis/control-panel/pkg/domain"
	"github.com/ervitis/control-panel/pkg/ports"

	"github.com/TeaEntityLab/fpGo/v2"
)

type service struct {
}

func NewUserService() ports.IUserService {
	return service{}
}

func LoadMockList() []domain.User {
	return mockListPeople()
}

func mockListPeople() []domain.User {
	absPath, _ := filepath.Abs(".")
	f, err := os.Open(absPath + "/mock/data/persons.json")
	if err != nil {
		log.Panicln(err)
	}

	b, _ := io.ReadAll(f)
	var data []domain.User
	if err := json.Unmarshal(b, &data); err != nil {
		log.Panicln(err)
	}
	return data
}

func (s service) Filter(_ context.Context, criteria domain.Filter) []domain.User {
	var f1, f2 []domain.User

	f1 = func(criteria domain.Filter) []domain.User {
		return fpgo.Filter(func(user domain.User, i int) bool {
			return criteria.Name == user.Name
		}, mock.Users...)
	}(criteria)

	f2 = func(criteria domain.Filter) []domain.User {
		return fpgo.Filter(func(user domain.User, i int) bool {
			return criteria.Birthday == user.Birthday
		}, mock.Users...)
	}(criteria)

	if len(f1) == 0 && len(f2) == 0 {
		return mock.Users
	}

	if len(f1) > 0 && len(f2) > 0 {
		return fpgo.Concat(f1, f2)
	}

	if len(f1) > 0 {
		return f1
	}

	return f2
}
