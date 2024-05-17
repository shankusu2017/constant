package constant

import "time"

const (
	Time30Min       = time.Minute * 30
	TimePing        = time.Second * 1
	TimeKeepAlive   = time.Second * 7
	TimeScanSession = time.Minute * 10
)
