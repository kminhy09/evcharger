package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

type (
	Auth struct {
		EnrollId     string `json:"enrollId"`
		EnrollSecret string `json:"enrollSecret"`
	}

	ChinCodeReq struct {
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
		Args []string `json:"args"`
	}
	PBCResponse struct {
		OK      string `json:"OK"`
		Jsonrpc string `json:"jsonrpc"`
		Result  Result `json:"result"`
		ID      int    `json:"id"`
	}
	Result struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
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

	reqdata := c.Param("params")
	resultdata := strings.Split(reqdata, ",")
	auth := Auth{}
	auth.EnrollId = resultdata[0]
	auth.EnrollSecret = resultdata[1]

	pbytes, _ := json.Marshal(auth)
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

	response := PBCResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusCreated, response)
}

func PostDeploy(c echo.Context) error {
	reqdata := c.Param("params")
	resultdata := strings.Split(reqdata, ",")

	deploy := ChinCodeReq{}
	deploy.Jsonrpc = "2.0"
	deploy.Method = "deploy"
	params := Params{}
	params.Type = 1
	chaincodeID := ChaincodeID{}
	chaincodeID.Name = "mycc"
	params.ChaincodeID = chaincodeID
	ctorMsg := CtorMsg{}
	ctorMsg.Args = []string{resultdata[0], resultdata[1], resultdata[2], resultdata[3], resultdata[4]}
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

	response := PBCResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusCreated, response)
}

func PostInvoke(c echo.Context) error {
	reqdata := c.Param("params")
	resultdata := strings.Split(reqdata, ",")

	invoke := ChinCodeReq{}
	invoke.Jsonrpc = "2.0"
	invoke.Method = "invoke"
	params := Params{}
	params.Type = 1
	chaincodeID := ChaincodeID{}
	chaincodeID.Name = "mycc"
	params.ChaincodeID = chaincodeID
	ctorMsg := CtorMsg{}
	ctorMsg.Args = []string{resultdata[0], resultdata[1], resultdata[2], resultdata[3]}
	params.CtorMsg = ctorMsg
	params.SecureContext = "admin"
	invoke.Params = params
	invoke.ID = 3

	pbytes, _ := json.Marshal(invoke)
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

	response := PBCResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusCreated, response)
}

func PostQuery(c echo.Context) error {
	reqdata := c.Param("params")
	resultdata := strings.Split(reqdata, ",")

	query := ChinCodeReq{}
	query.Jsonrpc = "2.0"
	query.Method = "query"
	params := Params{}
	params.Type = 1
	chaincodeID := ChaincodeID{}
	chaincodeID.Name = "mycc"
	params.ChaincodeID = chaincodeID
	ctorMsg := CtorMsg{}
	ctorMsg.Args = []string{resultdata[0], resultdata[1]}
	params.CtorMsg = ctorMsg
	params.SecureContext = "admin"
	query.Params = params
	query.ID = 5

	pbytes, _ := json.Marshal(query)
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

	response := PBCResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusCreated, response)
}
