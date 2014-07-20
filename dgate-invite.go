package main

import (
	"flag"
	"fmt"
	"github.com/henteko/dgate-invite/dgate"
)

func main() {
	var (
		packageName = flag.String("p", "", "Your DeployGate App Package Name")
		getFlag     = flag.Bool("g", false, "Get App member")
		inviteFlag  = flag.Bool("i", false, "Invite Target User for App member")
		deleteFlag  = flag.Bool("d", false, "Delete Target User for App member")

		loginFlag  = flag.Bool("login", false, "DeployGate Login")
		logoutFlag = flag.Bool("logout", false, "DeployGate Logout")
	)
	flag.Parse()

	if dgate.IsLogin() {
		name, loginToken := dgate.GetSettings()
		if *loginFlag {
			dgate.Login()
		} else if *logoutFlag {
			dgate.Logout(name)
		} else if *getFlag {
			dgate.PrintUsersName(dgate.InviteGet(name, *packageName, loginToken))
		} else if *inviteFlag {
			if len(flag.Args()) == 0 {
				fmt.Println("Please input target user name")
				return
			}
			dgate.PrintResult(dgate.InvitePost(name, *packageName, loginToken, flag.Args()))
		} else if *deleteFlag {
			if len(flag.Args()) == 0 {
				fmt.Println("Please input target user name")
				return
			}
			dgate.PrintResult(dgate.InviteDelete(name, *packageName, loginToken, flag.Args()))
		}
	} else {
		fmt.Println("Please Login")
		dgate.Login()
	}
}
