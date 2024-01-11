package gossh

import (
	"fmt"
	"strings"
	"testing"
	"time"
)
// GOPROXY="https://mirrors.aliyun.com/goproxy/,direct";
func TestSSH(t *testing.T)  {
	t.Run("", func(t *testing.T) {
		ex := NewMellanoxSSH("\n")
		err := ex.Dial( "admin", "admin", "192.168.59.51",22)
		if err != nil{
			t.Error()
		}

		out, err := ex.RunCmd("show user")
		if err != nil{
			t.Error()
		}

		t.Log(out)
	})

	t.Run("", func(t *testing.T) {
		ex := NewMellanoxSSH("\n")
		err := ex.Dial( "admin", "Enmo@123456", "192.168.13.70",22)
		if err != nil{
			t.Error()
		}

		out, err := ex.RunCmd("system-view")
		if err != nil{
			t.Error()
		}

		t.Log(out)
	})
}

func TestSwitch(t *testing.T)  {
	t.Run("", func(t *testing.T) {
		Demo1()
	})

// same session
	t.Run("", func(t *testing.T) {
		cmds := []string{
			//"system-view",
			"dis interface brief",
			//"interface HGE1/0/1",
			//"dis this",
			//"quit",
			//"quit",
		}
		out, err := RunSwitchCommands("admin", "Enmo@123456", fmt.Sprintf("%s:%d", "192.168.13.70",22),cmds...)
		if err != nil{
			t.Error()
		}

		cmds = []string{
			"interface HGE1/0/1",
			"dis this",
		}
		out, err = RunSwitchCommands("admin", "Enmo@123456", fmt.Sprintf("%s:%d", "192.168.13.70",22),cmds...)
		if err != nil{
			t.Error()
		}
		t.Log(strings.Join(out, "\n"))
	})
}

func TestSSHRunnerMultiple(t *testing.T) {
	user := "admin"
	password := "Enmo@123456"
	ipPort := "192.168.13.70:22"
	cmds := make([]string, 0)
	cmds = append(cmds, "dis clock")
	//cmds = append(cmds, "dis vlan")

	for i := 0; i < 2; i++ {
		go func(i int) {
			result, err := RunSwitchCommands(user, password, ipPort, cmds...)
			if err != nil {
				t.Logf("RunCommands<%d> err:%s", i, err.Error())
			}
			t.Logf("RunCommands<%d> result:\n%s", i, result)
		}(i)
	}

	time.Sleep(15 * time.Second)
}