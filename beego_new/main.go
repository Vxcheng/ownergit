package main

import (
	"fmt"
	_ "ownergit/beego_new/routers"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	initDB()
	m := &InfiniBandPort{
		Name:                 "mlx4_0",
		Port:                 "1",
		SwitchNumber:         "12",
		SwitchNetPortNumber:  "13",
		MetaHostIp:           "127.0.0.1",
		MetaUpdatedTimestamp: 1584427013,
		PortGuid:             "111111111222",
		MetaHostname:         "cell03",
	}
	m.AddOrUpdate()
	m.JudgePortGuidExsit()
	m.UpdateFields("22", "22")
	beego.Run()
}

func init() {
	// orm.RegisterModel(new(InfiniBandPort))
}

type InfiniBandPort struct {
	BaseInfo
	PortGuid             string `json:"port_guid" orm:"column(port_guid);null"`
	MetaHostname         string `json:"meta_hostname" orm:"column(meta_hostname);null"`
	MetaHostIp           string `json:"meta_host_ip" orm:"column(meta_host_ip);null"`
	MetaUpdatedTimestamp int64  `json:"meta_updated_timestamp" orm:"column(meta_updated_timestamp);null"`
	Name                 string `json:"name" orm:"column(name);null"`
	Port                 string `json:"port" orm:"column(port);null"`
	SwitchNumber         string `json:"switch_nmber" orm:"column(switch_number);null"`
	SwitchNetPortNumber  string `json:"switchnet_port_number" orm:"column(switch_net_port_number);null"`
}

type BaseInfo struct {
	ID         int   `orm:"column(id);auto;pk"`
	UpdateTime int64 `orm:"column(update_time);null"`
	IsDel      bool  `orm:"column(is_del);null"`
}

func (m InfiniBandPort) TableName() string {
	return "INFINIBAND_PORT"
}

func (m *InfiniBandPort) JudgePortGuidExsit() (bool, error) {
	o := orm.NewOrm()
	m.IsDel = false
	err := o.Read(m, "PortGuid", "IsDel")
	if err != nil && err != orm.ErrNoRows {
		return false, err
	}
	if err == orm.ErrNoRows {
		return false, nil
	}
	return true, nil
}

func (m *InfiniBandPort) UpdateFields(switchNumber, switchNetPortNumber string) error {
	o := orm.NewOrm()
	m.SwitchNumber = switchNumber
	m.SwitchNetPortNumber = switchNetPortNumber
	m.IsDel = true
	m.UpdateTime = time.Now().Unix()
	if _, err := o.Update(m); err != nil {
		return err
	}
	return nil
}

func (m InfiniBandPort) String() string {
	return fmt.Sprintf("port_guid: %s, ca_name: %s, port: %s, switch_number:%s, switch_net_port_number: %s", m.PortGuid, m.Name, m.Port, m.SwitchNumber, m.SwitchNetPortNumber)
}

func (m *InfiniBandPort) AddOrUpdate() error {
	o := orm.NewOrm()
	b := &InfiniBandPort{
		Name:                 m.Name,
		Port:                 m.Port,
		MetaHostname:         m.MetaHostname,
		MetaHostIp:           m.MetaHostIp,
		MetaUpdatedTimestamp: m.MetaUpdatedTimestamp,
		SwitchNumber:         m.SwitchNumber,
		SwitchNetPortNumber:  m.SwitchNetPortNumber,
	}
	err := o.Read(m, "PortGuid")
	if err != nil && err != orm.ErrNoRows {
		return err
	}
	m.UpdateTime = time.Now().Unix()
	m.IsDel = false
	if err == orm.ErrNoRows {

		if _, err = o.Insert(m); err != nil {
			return err
		}
		return nil
	}

	m.Name = b.Name
	m.Port = b.Port
	m.MetaHostname = b.MetaHostname
	m.MetaHostIp = b.MetaHostIp
	m.MetaUpdatedTimestamp = b.MetaUpdatedTimestamp
	m.SwitchNumber = b.SwitchNumber
	m.SwitchNetPortNumber = b.SwitchNetPortNumber
	if _, err = o.Update(m); err != nil {
		return err
	}
	return nil
}

func initDB() {
	orm.Debug = true
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", "root", "zdata_2019", "192.168.10.67", "3306", "zMon")
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(new(InfiniBandPort))

	orm.RunSyncdb("default", false, true)
}
