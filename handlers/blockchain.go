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

type (
	Deploy struct {
		Jsonrpc string `json:"jsonrpc"`
		Method  string `json:"method"`
		Params  Params `json:"params"`
		ID      int    `json:"id"`
	}

	Params struct {
		Type          int         `json:"type"`
		ChaincodeID   ChaincodeID `json:"chaincodeID"`
		CtorMsg       CtorMsg     `json:"ctorMsg"`
		SecureContext string      `json:"secureContext"`
	}

	ChaincodeID struct {
		Name string `json:"name"`
	}
	CtorMsg struct {
		Args [5]string `json:"args"`
	}
)

func PostDeploy(c echo.Context) error {

	deploy := Deploy{}
	deploy.Jsonrpc = "2.0"
	deploy.Method = "deploy"
	params := Params{}
	params.Type = 1
	chaincodeID := ChaincodeID{}
	chaincodeID.Name = "mycc"
	params.ChaincodeID = chaincodeID
	ctorMsg := CtorMsg{}
	ctorMsg.Args = [5]string{"init", "a", "100", "b", "200"}
	params.CtorMsg = ctorMsg
	params.SecureContext = "admin"
	deploy.Params = params
	deploy.ID = 1

	pbytes, _ := json.Marshal(deploy)
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
