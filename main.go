package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ddliu/go-httpclient"
	"github.com/jmoiron/jsonq"
	"io/ioutil"
	"strings"
)

const (
	USERAGENT       = "golang deploygate cl tool"
	TIMEOUT         = 30
	CONNECT_TIMEOUT = 5
)

const apiUrl = "https://deploygate.com/api/users"

var client = httpclient.NewHttpClient(map[int]interface{}{
	httpclient.OPT_USERAGENT: USERAGENT,
	httpclient.OPT_TIMEOUT:   TIMEOUT,
})

func printUserName(jsonString string) {
	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(jsonString))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)

	apiError, _ := jq.Bool("error")
	if apiError {
		message, _ := jq.String("message")
		fmt.Println("Api Error Message: " + message)
	}
	users, _ := jq.ArrayOfObjects("results", "users")

	for _, value := range users {
		fmt.Println(value["name"])
	}
}

func getJsonString(ownerName string, packageName string, token string) string {
	res, _ := client.Get(apiUrl+"/"+ownerName+"/apps/"+packageName+"/members", map[string]string{
		"token": token,
	})
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

func invitePost(ownerName string, packageName string, token string, userName string) {
	res, _ := client.Post(apiUrl+"/"+ownerName+"/apps/"+packageName+"/members", map[string]string{
		"token": token,
		"users": "[" + userName + "]",
	})
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

var (
	ownerName   string
	packageName string
	token       string
	userName    string
	get         bool
	invite      bool
	delete      bool
)

func flagInit() {
	flag.StringVar(&ownerName, "ownername", "", "Your DeployGate Owner Name")
	flag.StringVar(&ownerName, "o", "", "Your DeployGate Owner Name")
	flag.StringVar(&packageName, "packageName", "", "Your DeployGate App Package Name")
	flag.StringVar(&packageName, "p", "", "Your DeployGate App Package Name")
	flag.StringVar(&token, "token", "", "Your DeployGate App Package Name")
	flag.StringVar(&token, "t", "", "Your DeployGate App Package Name")
	flag.StringVar(&userName, "username", "", "Your DeployGate App Package Name")
	flag.StringVar(&userName, "u", "", "Your DeployGate App Package Name")
	flag.BoolVar(&get, "get", false, "Your DeployGate App Package Name")
	flag.BoolVar(&get, "g", false, "Your DeployGate App Package Name")
	flag.BoolVar(&invite, "invite", false, "Your DeployGate App Package Name")
	flag.BoolVar(&invite, "i", false, "Your DeployGate App Package Name")
	flag.BoolVar(&delete, "delete", false, "Your DeployGate App Package Name")
	flag.BoolVar(&delete, "d", false, "Your DeployGate App Package Name")

	flag.Parse()
}

func main() {
	flagInit()

	if get {
		printUserName(getJsonString(ownerName, packageName, token))
	}
	if invite {
		invitePost(ownerName, packageName, token, userName)
	}
}
