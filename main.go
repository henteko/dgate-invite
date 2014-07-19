package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		packageName = flag.String("p", "", "Yout DeployGate App Package Name")
		userName    = flag.String("u", "", "Invite Target User Name")
		getFlag     = flag.Bool("g", false, "Get App member")
		inviteFlag  = flag.Bool("i", false, "Invite Target User for App member")
		deleteFlag  = flag.Bool("d", false, "Delete Target User for App member")
		loginFlag   = flag.Bool("login", false, "DeployGate Login")
		logoutFlag  = flag.Bool("logout", false, "DeployGate Logout")
	)
	flag.Parse()

	if checkLogin() {
		name, loginToken := getSettings()
		if *loginFlag {
			dgateLogin()
		} else if *logoutFlag {
			dgateLogout(name)
		} else if *getFlag {
			printUsersName(inviteGet(name, *packageName, loginToken))
		} else if *inviteFlag {
			printResult(invitePost(name, *packageName, loginToken, *userName))
		} else if *deleteFlag {
			printResult(inviteDelete(name, *packageName, loginToken, *userName))
		}
	} else {
		fmt.Println("Please Login")
		dgateLogin()
	}
}
