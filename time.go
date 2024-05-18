package constant

import "time"

const (
	Time30Min       = time.Minute * 30
	TimePing        = time.Second * 5 /* 多久 ping 一次 */
	TimePingRound   = time.Minute * 2 /* 一轮 ping 持续多久 */
	TimeKeepAlive   = time.Minute * 1
	TimeScanSession = time.Minute * 10
)
