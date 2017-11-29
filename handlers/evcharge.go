package handlers

import (
	"encoding/xml"
	"evcharger/conf"
	"evcharger/models"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func ListEvCharger(c echo.Context) error {

	evcharger := GetEvChargerData()

	return c.JSON(http.StatusCreated, evcharger)
}

func GetEvChargerData() models.EvCharger {
	// GET 호출
	resp, err := http.Get(conf.URL + conf.ServiceKey)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// 결과 출력
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	response := models.Response{}

	err = xml.Unmarshal(data, &response)
	checkError(err)
	fmt.Println(string(data))

	return response.EvCharger
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}
