package snowflake

import (
	"project/setting"
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}

func InitSnowflake() (err error) {
	sConfig := setting.Conf.SnowflakeConfig
	var st time.Time
	st, err = time.Parse("2006-01-02", sConfig.StartTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(sConfig.MachineID)

	return
}
