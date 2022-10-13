package main

import (
	"github.com/thb-cmyk/aliyum-demo/databasic"
)

func main() {

	databasic.All_Init()

	databasic.Broker()

	databasic.ProceNode_register(dataProccessor, "LightStatus")

	MysqlInit()

	defer MysqlDeInit()

	go Aliyun_Connect()

	databasic.ProceNode_register(voltageProccesser, "voltage")
	databasic.ProceNode_register(checkmodeProccesser, "check_mode")
	databasic.ProceNode_register(errorinfoProccesser, "error_info")

	IntrefaceInit()

}
