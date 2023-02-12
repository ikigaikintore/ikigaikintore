package main

import (
	"encoding/json"
	apigen "github.com/ervitis/control-panel/api/v1"
	"github.com/ervitis/control-panel/common/global/mock"
	"github.com/ervitis/control-panel/pkg/domain"
	"github.com/ervitis/control-panel/pkg/ports"
	"github.com/ervitis/control-panel/pkg/services/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

type requestBody struct {
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
}

type usersResponse struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
}

/*
{"name": "victor", "birthday": "2020-10-04"}
*/

var (
	svc ports.IUserService
)

func init() {
	svc = user.NewUserService()
	mock.Users = user.LoadMockList()
}

type Server struct{}

func (s Server) Login(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) SignIn(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) CreateUser(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) DeleteUser(ctx echo.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) GetUser(ctx echo.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) UpdateUser(ctx echo.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) GetUsers(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s Server) FilterUsers(ctx echo.Context) error {
	var body requestBody
	if err := json.NewDecoder(ctx.Request().Body).Decode(&body); err != nil {
		log.Println(err)
		if err := ctx.JSON(http.StatusBadRequest, nil); err != nil {
			return err
		}
	}

	filters := domain.NewFilter(domain.WithBirthday(body.Birthday), domain.WithName(body.Name))

	users := svc.Filter(ctx.Request().Context(), filters)
	return ctx.JSON(http.StatusOK, toUsersResponse(users))
}

func toUsersResponse(users []domain.User) []usersResponse {
	r := make([]usersResponse, len(users), len(users))
	for i, v := range users {
		r[i] = usersResponse{
			ID:       uint64(v.ID),
			Name:     v.Name,
			Birthday: v.Birthday,
		}
	}
	return r
}

func main() {
	s := Server{}
	r := echo.New()

	r.Use(middleware.CORS())

	apigen.RegisterHandlersWithBaseURL(r, s, "/v1")

	r.Logger.Fatal(r.Start(":8080"))
}
