package handlers

import (
	"log"
	"net/http"

	"github.com/byuoitav/cgi-microservice/helpers"
	"github.com/labstack/echo"
)

func VolumeUp(context echo.Context) error {

	log.Printf("Raising volume on %s...", context.Param("address"))

	err := helpers.BuildAndSendPayload(context.Param("address"), "send", "control", "4054", "12")
	if err != nil {
		log.Printf("ERROR: %v", err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	log.Printf("Done.")
	return context.JSON(http.StatusOK, "success")

}

func VolumeDown(context echo.Context) error {
	log.Printf("Lowering volume on %s...", context.Param("address"))

	err := helpers.BuildAndSendPayload(context.Param("address"), "send", "control", "4054", "13")
	if err != nil {
		log.Printf("ERROR: %v", err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	log.Printf("Done.")
	return context.JSON(http.StatusOK, "success")
}

func Mute(context echo.Context) error {
	log.Printf("Sending mute command to %s", context.Param("address"))

	err := helpers.BuildAndSendPayload(context.Param("address"), "send", "control", "4054", "14")
	if err != nil {
		log.Printf("ERROR: %v", err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	log.Printf("Done.")
	return context.JSON(http.StatusOK, "success")
}
