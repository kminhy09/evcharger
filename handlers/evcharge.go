package handlers

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/kminhy09/evcharger/conf"
	"github.com/kminhy09/evcharger/models"

	"github.com/labstack/echo"
)

func GetEvChargers(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetEvChargers(db))
	}
}

func ListEvChargerApiData(c echo.Context) error {

	evcharger := GetEvChargerApiData()

	return c.JSON(http.StatusCreated, evcharger)
}

func GetEvChargerApiData() models.EvCharger {
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
	//fmt.Println(string(data))

	return response.EvCharger
}

func SyncEvChargers(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		// 공공포털 api 전기자동차 충전소 조회
		evcItems := GetEvChargerApiData().EvCharger.EvcItems

		fmt.Println("##### Delete Start")
		// 기존 데이터 삭제
		count, delerr := models.DeleteEvChargers(db)
		if delerr == nil {
			fmt.Printf("##### DeleteEvChargers Success : Count >>> %d\n", count)
		} else {
			return delerr
		}
		fmt.Println("##### Delete End\n")

		fmt.Println("##### Insert Start\n")
		// 다시 입력
		inserterr := models.PutEvChargers(db, evcItems)
		if inserterr == nil {
			fmt.Fprintln(os.Stdout, "##### PutEvChargers Success")
		} else {
			return inserterr
		}
		fmt.Println("##### Insert End")

		return c.JSON(http.StatusCreated, "SyncEvChargers Success")
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fatal error %s", err.Error())
	}
}
