package gossh

import (
	"fmt"
	"github.com/shenbowei/switch-ssh-go"
	"regexp"
	"strings"
	"unicode"
)

func Demo1() {
	user := "admin"
	password := "Enmo@123456"
	ipPort := "192.168.13.70:22"

	//get the switch brand(vendor), include h3c,huawei and cisco
	//brand, err := ssh.GetSSHBrand(user, password, ipPort)
	//if err != nil {
	//	fmt.Println("GetSSHBrand err:\n", err.Error())
	//}
	//fmt.Println("Device brand is:\n", brand)

	//run the cmds in the switch, and get the execution results
	cmds := make([]string, 0)
	cmds = append(cmds, "dis clock")
	//cmds = append(cmds, "dis vlan")
	result, err := ssh.RunCommands(user, password, ipPort, cmds...)
	if err != nil {
		fmt.Println("RunCommands err:\n", err.Error())
	}
	fmt.Println("RunCommands result:\n", result)
}

func RunSwitchCommands(user, password, ipPort string, cmds ...string) (values []string, err error) {
	out, err := ssh.RunCommands(user, password, ipPort, cmds...)
	if err != nil {
		return
	}

	re := regexp.MustCompile(`<.+>|\[.+\]`)
	outter:
	for _, line := range strings.Split(out, "\n") {
		line = strings.Trim(line, "\r")
		line = strings.TrimFunc(line, unicode.IsSpace)
		if line == "" ||
			re.MatchString(line) {
			continue
		}

		for _, cmd :=  range cmds{
			if  strings.Contains(line, cmd){
				continue outter
			}
		}


		values = append(values, line)
	}
	return
}

func init()  {
	ssh.IsLogDebug = false
}