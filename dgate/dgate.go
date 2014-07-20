package dgate

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmoiron/jsonq"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

const apiUrl = "https://deploygate.com/api/users"

/***********************
* Public Methods
************************/

func GetSettings() (string, string) {
	settingFile := getSettingFilePath()
	fileByte, err := ioutil.ReadFile(settingFile)
	if err != nil {
		return "", ""
	}
	json := stringToJsonq(string(fileByte))
	name, _ := json.String("name")
	token, _ := json.String("token")

	return name, token
}

func IsLogin() bool {
	name, token := GetSettings()
	return name != "" && token != ""
}

func Login() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Owner Name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Your Token: ")
	scanner.Scan()
	token := scanner.Text()

	settings := `{"name":"` + name + `","token":"` + token + `"}`
	writeSettingFile(settings)
	fmt.Println("Login Success!")
}

func Logout(name string) {
	settings := `{"name":"` + name + `","token":""}`
	writeSettingFile(settings)
	fmt.Println("Logout Success!")
}

func InviteGet(ownerName string, packageName string, token string) string {
	uri := getUri(ownerName, packageName, token)
	res, err := httpGet(uri, map[string]string{
		"token": token,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

func InvitePost(ownerName string, packageName string, token string, userNames []string) string {
	uri := getUri(ownerName, packageName, token)
	res, err := httpPost(uri, map[string]string{
		"token": token,
		"users": getUserNamesString(userNames),
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

func InviteDelete(ownerName string, packageName string, token string, userNames []string) string {
	uri := getUri(ownerName, packageName, token)
	res, err := httpDelete(uri, map[string]string{
		"token": token,
		"users": getUserNamesString(userNames),
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

/***********************
* Private Methods
************************/

func stringToJsonq(jsonString string) *jsonq.JsonQuery {
	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(jsonString))
	dec.Decode(&data)
	json := jsonq.NewQuery(data)
	return json
}

func getSettingFilePath() string {
	return os.Getenv("HOME") + "/.dgate"
}

func writeSettingFile(settings string) {
	settingFile := getSettingFilePath()
	ioutil.WriteFile(settingFile, []byte(settings), 0644)
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

func PrintUsersName(jsonString string) {
	users := getUsersName(jsonString)
	for _, user := range users {
		fmt.Println(user)
	}
}

func PrintResult(jsonString string) {
	err := checkError(jsonString)
	if err != nil {
		fmt.Println(err)
	} else {
		json := stringToJsonq(jsonString)
		invite, _ := json.String("results", "invite")

		fmt.Println("Success Message: " + invite)
	}
}

func getUserNamesString(userNames []string) string {
	var nameBuffer bytes.Buffer
	for _, name := range userNames {
		nameBuffer.WriteString("," + name)
	}
	names := nameBuffer.String()
	re, _ := regexp.Compile("^,")
	names = re.ReplaceAllString(names, "")

	return "[" + names + "]"
}

func getUri(ownerName string, packageName string, token string) string {
	return apiUrl + "/" + ownerName + "/apps/" + packageName + "/members"
}
