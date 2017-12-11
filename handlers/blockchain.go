package handlers

import (
	"bytes"
	"encoding/json"
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

	return c.JSON(http.StatusCreated, string(data))
}

func PostRegistrar(c echo.Context) error {
	params := map[string]string{
		"enrollId":     "admin",
		"enrollSecret": "Xurw3yU9zI0l",
	}
	pbytes, _ := json.Marshal(params)
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post("http://52.78.185.234:7050/registrar", "application/json", buff)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// 결과 출력
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusCreated, string(data))
}

func PostDeploy(c echo.Context) error {
	params := map[string]string{
		"enrollId":     "admin",
		"enrollSecret": "Xurw3yU9zI0l",
	}
	pbytes, _ := json.Marshal(params)
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post("http://52.78.185.234:7050/chaincode", "application/json", buff)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// 결과 출력
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusCreated, string(data))
}
