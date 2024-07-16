package handler

import (
	"net/http"
	"strconv"
	"tt/testtask"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getUserByPas(c *gin.Context) {
	var user testtask.Users
	passportSerieStr := c.Query("passportSerie")
	passportNumberStr := c.Query("passportNumber")

	passportSerie, err := strconv.Atoi(passportSerieStr)
	if err != nil {
		passportSerie = 0
	}
	passportNumber, err := strconv.Atoi(passportNumberStr)
	if err != nil {
		passportNumber = 0
	}

	user, err = h.services.User.GetByPas(passportSerie, passportNumber)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
