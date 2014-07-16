package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmoiron/jsonq"
	"io/ioutil"
	"strings"
)

const apiUrl = "https://deploygate.com/api/users"

func stringToJsonq(jsonString string) *jsonq.JsonQuery {
	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(jsonString))
	dec.Decode(&data)
	return jsonq.NewQuery(data)
}

func checkError(jsonString string) error {
	json := stringToJsonq(jsonString)

	apiError, _ := json.Bool("error")
	if apiError {
		message, _ := json.String("message")
		return errors.New("Api Error Message: " + message)
	}
	return nil
}

func getUsersName(jsonString string) []string {
	err := checkError(jsonString)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	json := stringToJsonq(jsonString)
	users, _ := json.ArrayOfObjects("results", "users")

	result := make([]string, 0)
	for _, value := range users {
		result = append(result, fmt.Sprintf("%v", value["name"]))
	}
	return result
}

func printUsersName(jsonString string) {
	users := getUsersName(jsonString)
	for _, user := range users {
		fmt.Println(user)
	}
}

func printResult(jsonString string) {
	err := checkError(jsonString)
	if err != nil {
		fmt.Println(err)
	} else {
		json := stringToJsonq(jsonString)
		invite, _ := json.String("results", "invite")

		fmt.Println("Success Message: " + invite)
	}
}

func getUri(ownerName string, packageName string, token string) string {
	return apiUrl + "/" + ownerName + "/apps/" + packageName + "/members"
}

func inviteGet(ownerName string, packageName string, token string) string {
	uri := getUri(ownerName, packageName, token)
	res, _ := httpGet(uri, map[string]string{
		"token": token,
	})
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

func invitePost(ownerName string, packageName string, token string, userName string) string {
	uri := getUri(ownerName, packageName, token)
	res, _ := httpPost(uri, map[string]string{
		"token": token,
		"users": "[" + userName + "]",
	})
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

func inviteDelete(ownerName string, packageName string, token string, userName string) string {
	uri := getUri(ownerName, packageName, token)
	res, _ := httpDelete(uri, map[string]string{
		"token": token,
		"users": "[" + userName + "]",
	})
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}
