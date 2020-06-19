currentTimeStamp=$(date "+%Y%m%d%H%M%S")
content="
----------------------------------------------------------\n
Disabled IBSwitchNetPortNumbers Info:\n\t
	CA Name: mlx4_0\n\t
	Port: 1\n\t
	SwitchNumber: 1\n\t
	SwitchNetPortNumber: 1\n\t
	Host: 192.168.10.61\n\t
	UpdateTime: 2020-03-09 01:27:57\n
	\n\t
	CA Name: mlx4_0\n\t
	Port: 0\n\t
	SwitchNumber: 0\n\t
	SwitchNetPortNumber: 0\n\t
	Host: 192.168.10.62\n\t
	UpdateTime: 2020-03-09 01:27:57\n
	
----------------------------------------------------------\n\n
	
Tips: If you want enable IBSwitchNetPortNumbers, you can execute those commands in other nodes connecting this IB Switch.\n\t
	ibportstate -C mlx4_0  1 1 enable\n\t
	ibportstate -C mlx4_0  0 0 enable\n\n

And please after confirm enable IBSwitchNetPortNumbers success, then delete file by excute the follow command:\n\t
	mv -f /opt/zdata/zmanager/zmanager-oracle/ibSwitchNetPortNumbersInfo.sh /opt/zdata/zmanager/zmanager-oracle/ibSwitchNetPortNumbersInfo_${currentTimeStamp}.sh\n
	"
echo -e ${content}