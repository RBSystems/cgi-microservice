package handlers

import (
	"log"
	"net/http"

	"github.com/byuoitav/cgi-microservice/helpers"
	"github.com/jessemillar/jsonresp"
	"github.com/labstack/echo"
)

func VolumeUp(context echo.Context) error {

	log.Printf("Raising volume on %s...", context.Param("address"))

	err := helpers.BuildAndSendPayload(context.Param("address"), "send", "control", "4054", "12")
	if err != nil {
		log.Printf("ERROR: %v", err.Error())
		jsonresp.New(context.Response(), http.StatusBadRequest, err.Error())
		return nil
	}

	log.Printf("Done.")
	return context.JSON(http.StatusOK, err)

}

func VolumeDown(context echo.Context) error {
	log.Printf("Lowering volume on %s...", context.Param("address"))

	return nil
}
