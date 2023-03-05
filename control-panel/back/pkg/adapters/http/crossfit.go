package http

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func (h Handler) StartBookingCrossfit(ctx echo.Context) error {
	if err := h.agendaService.StartBooking(ctx.Request().Context()); err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (h Handler) GetCrossfitStatus(ctx echo.Context) error {
	status, err := h.agendaService.CrossfitStatus(ctx.Request().Context())
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, NewProcessStatus(status))
}

func (h Handler) GetCrossfitToken(ctx echo.Context) error {
	token, err := h.agendaService.GetToken(ctx.Request().Context())
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, NewGoogleCredentials(token))
}

func (h Handler) SetCrossfitToken(ctx echo.Context) error {
	var reqBody ReqGoogleToken
	if err := ctx.Bind(&reqBody); err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusInternalServerError)
	}
	err := h.agendaService.SetToken(ctx.Request().Context(), reqBody.To())
	if err != nil {
		log.Println(err)
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.NoContent(http.StatusNoContent)
}
