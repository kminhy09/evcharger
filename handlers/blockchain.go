package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

func GetChain(c echo.Context) error {

	// GET 호출
	resp, err := http.Get("http://52.78.185.234:7050/chain")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// 결과 출력
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	return c.JSON(http.StatusCreated, string(data))
}
