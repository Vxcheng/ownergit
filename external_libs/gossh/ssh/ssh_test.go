package ssh

import (
	"strings"
	"testing"
)

func TestSwitch(t *testing.T) {
	t.Run("", func(t *testing.T) {
		ex := NewNationalSSH("\n")
		err := ex.Dial( "admin", "Enmo@123456", "192.168.13.70",22)
		if err != nil{
			t.Error()
		}
		defer ex.Close()
		//
		//out, err := ex.RunCmd("system-view")
		//if err != nil{
		//	t.Error()
		//}

		cmd := []string{
			"dis clock",
			"system-view",
			"dis interface brief",
			"interface HGE1/0/1",
			"dis this",
			//"quit",
			//"quit",
		}
		out, err := ex.RunCmds(true, cmd...)
		if err != nil{
			t.Error()
		}


		t.Log(strings.Join(out, "\n"))
	})
}
