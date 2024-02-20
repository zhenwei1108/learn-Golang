package dir

import (
	"fmt"
	"os/user"
)

func GetUserMsg() {
	user, err := user.Current()
	if err != nil {
		fmt.Errorf("失败：", err)
	}
	fmt.Println("user.name=", user.Name)
	fmt.Println("user.HomeDir=", user.HomeDir)
	fmt.Println("user.Uid=", user.Uid)
	fmt.Println("user.Username=", user.Username)
}
