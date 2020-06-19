package main

import (
	"encoding/base64"
	"encoding/json"
	"log"

	"mysql-agent/rowsql"

	"github.com/mitchellh/mapstructure"
)

type User struct {
	Name string
	Addr `mapstructure:",squash"`
}

type Addr struct {
	Dong string
}

func main() {
	convert()
}

func convert() {
	var max int64 = 10
	log.Print(int(max))
}

func unmarshal() {
	str := `
	{
		"name": "ming"
	}
	`
	// var u User
	err := json.Unmarshal([]byte(str), &struct{}{})
	if err != nil {
		log.Fatal(err)
	}
}

type AsmGroupStat struct {
	rowsql.ASM_DISKGROUP_STAT `mapstructure:",squash"`
	Key                       string `json:"key"`
}

func unmarshal2() {
	str := `
	[
    {
        "group_number": 1,
        "diskgroup_name": "DATA",
        "ausize": 4,
        "state": "CONNECTED",
        "type": "NORMAL",
        "total_mb": 153600,
        "free_mb": 142496,
        "hot_used_mb": 0,
        "cold_used_mb": 11104,
        "usable_file_mb": 45648,
        "required_mirror_free_mb": 51200,
        "offline_disks": 0,
        "compatible_asm": "12.2.0.1.0",
        "compatible_rdbms": "12.2.0.1.0",
        "compatible_advm": "",
        "update_datetime": "2020-04-16 17:03:53"
    }
	]
	`
	container := make([]map[string]interface{}, 0)
	err := json.Unmarshal([]byte(str), &container)
	if err != nil {
		log.Panicf("json.Unmarshal err: %s", err.Error())
	}
	list := make([]AsmGroupStat, 0)
	for _, c := range container {
		var u AsmGroupStat
		if err = mapstructure.Decode(c, &u); err != nil {
			log.Panicf("mapstructure.Decode err: %s", err.Error())
		}
		list = append(list, u)
	}

	byteL, err := json.Marshal(list)
	if err != nil {
		log.Panicf("json.Marshal err: %s", err.Error())
	}
	log.Printf("byteL: %v\n", byteL)
	log.Printf("len %d,%s", len(list), string(byteL))

}
func baseDecode() {
	password, _ := base64.StdEncoding.DecodeString("cm9vdDEyMw==")
	log.Println("pwd, ", string(password))
}

func unmarshal3() {
	str := `
	[
    {
        "group_number": 1,
        "diskgroup_name": "DATA",
        "ausize": 4,
        "state": "CONNECTED",
        "type": "NORMAL",
        "total_mb": 153600,
        "free_mb": 142496,
        "hot_used_mb": 0,
        "cold_used_mb": 11104,
        "usable_file_mb": 45648,
        "required_mirror_free_mb": 51200,
        "offline_disks": 0,
        "compatible_asm": "12.2.0.1.0",
        "compatible_rdbms": "12.2.0.1.0",
        "compatible_advm": "",
        "update_datetime": "2020-04-16 17:03:53"
    }
	]
	`
	var list []*AsmGroupStat
	err := json.Unmarshal([]byte(str), &list)
	if err != nil {
		log.Panicf("json.Unmarshal err: %s", err.Error())
	}

	byteL, err := json.Marshal(list)
	if err != nil {
		log.Panicf("json.Marshal err: %s", err.Error())
	}
	log.Printf("byteL: %v\n", byteL)
	log.Printf("len %d,%s", len(list), string(byteL))
}
