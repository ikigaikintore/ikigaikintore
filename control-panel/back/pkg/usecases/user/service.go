package user

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/ikigaikintore/ikigaikintore/control-panel/back/pkg/domain"
	"github.com/ikigaikintore/ikigaikintore/control-panel/back/pkg/ports"
)

type service struct {
}

func (s service) Filter(ctx context.Context, filter domain.Filter) []domain.User {
	//TODO implement me
	panic("implement me")
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
