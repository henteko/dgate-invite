package main

import (
	"flag"
)

var (
	ownerName   string
	packageName string
	token       string
	userName    string
	get         bool
	invite      bool
	delete      bool
	login       bool
	logout      bool
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

	flag.BoolVar(&login, "login", false, "Your DeployGate App Package Name")
	flag.BoolVar(&logout, "logout", false, "Your DeployGate App Package Name")

	flag.Parse()
}

func main() {
	flagInit()

	if checkLogin() {
		name, loginToken := getSettings()
		if login {
			dgateLogin(ownerName, token)
		} else if logout {
			dgateLogout(name)
		} else if get {
			printUsersName(inviteGet(name, packageName, loginToken))
		} else if invite {
			printResult(invitePost(name, packageName, loginToken, userName))
		} else if delete {
			printResult(inviteDelete(name, packageName, loginToken, userName))
		}
	} else {
		dgateLogin(ownerName, token)
	}
}
