package handlers

import (
	"bytes"
	"encoding/json"
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

func PostRegistrar(c echo.Context) error {

	jsonData := map[string]string{"enrollId": "admin", "enrollSecret": "Xurw3yU9zI0l"}
	jsonValue, _ := json.Marshal(jsonData)

	// POST 호출
	resp, err = http.Post("http://52.78.185.234:7050/chaincode", "application/json", bytes.NewBuffer(jsonValue))
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
