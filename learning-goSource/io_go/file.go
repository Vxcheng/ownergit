package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const (
	IbSwitchPortInfoPath = "./zmanager-oracle/ibcard/ibSwitchNetPortNumbersInfo.sh"
	ZdataOraclePath      = "./zmanager-oracle/ibcard"
)
const (
	currentTimeStamp = `currentTimeStamp=$(date "+%Y%m%d%H%M%S")`
	contentEcho      = `
content=%s
echo -e ${content}`
	contentHead = `"
----------------------------------------------------------\n
Disabled IBSwitchNetPortNumbers Info:`
	contentIBInfo = `\n\t
	CA Name: %s\n\t
	Port: %s\n\t
	SwitchNumber: %s\n\t
	SwitchNetPortNumber: %s\n\t
	Host: %s\n\t
	UpdateTime: %s\n
	`
	contentTips = `
----------------------------------------------------------\n\n
	
Tips: If you want enable IBSwitchNetPortNumbers, you can execute those commands in other nodes connecting this IB Switch.`
	contentEnable = `\n\t
	ibportstate -C %s  %s %s enable`
	contentTail = `\n\n

And please after confirm enable IBSwitchNetPortNumbers success, then delete file by excute the follow command:\n\t
	mv -f /opt/zdata/zmanager/zmanager-oracle/ibSwitchNetPortNumbersInfo.sh /opt/zdata/zmanager/zmanager-oracle/ibSwitchNetPortNumbersInfo_${currentTimeStamp}.sh\n
	"`
	timeFormat = "2006-01-02 15:04:05"
)

func IsEixst(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func ChmodFile(filename string) error {
	return os.Chmod(filename, os.ModePerm)
}

type FileHandle struct{}

func NewFileHandle() *FileHandle {
	return &FileHandle{}
}
func (h *FileHandle) generateShellFile(iBName, iBPort, switchNumber, switchNetPortNumber, hostIp string, updateTime int64) error {
	var (
		err error
		f   *os.File
	)
	if !IsEixst(ZdataOraclePath) {
		if err = os.Mkdir(ZdataOraclePath, 0555); err != nil {
			return err
		}
	}
	if !IsEixst(IbSwitchPortInfoPath) {
		if _, err = os.Create(IbSwitchPortInfoPath); err != nil {
			return err
		}
	}
	if f, err = os.OpenFile(IbSwitchPortInfoPath, os.O_RDWR, 0644); err != nil {
		return err
	}
	defer f.Close()

	fByte, err := ioutil.ReadFile(IbSwitchPortInfoPath)
	if err != nil {
		return err
	}
	content := string(fByte)
	ibInfo := fmt.Sprintf(contentIBInfo, iBName, iBPort, switchNumber, switchNetPortNumber, hostIp, time.Unix(updateTime, 0).Format(timeFormat))
	enableInfo := fmt.Sprintf(contentEnable, iBName, switchNumber, switchNetPortNumber)
	if content == "" {
		if _, err = f.Write([]byte(currentTimeStamp + fmt.Sprintf(contentEcho, contentHead+ibInfo+contentTips+enableInfo+contentTail))); err != nil {
			return err
		}
		return ChmodFile(IbSwitchPortInfoPath)
	}

	if strings.Contains(content, ibInfo) {
		return nil // exist
	}
	content = strings.ReplaceAll(content, "Disabled IBSwitchNetPortNumbers Info:", fmt.Sprintf("Disabled IBSwitchNetPortNumbers Info:%s", ibInfo))
	content = strings.ReplaceAll(content, "you can execute those commands in other nodes connecting this IB Switch.", fmt.Sprintf("you can execute those commands in other nodes connecting this IB Switch.%s", enableInfo))
	if _, err = f.WriteString(content); err != nil {
		return err
	}
	return ChmodFile(IbSwitchPortInfoPath)
}
