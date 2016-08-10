package loadgen

import "go_multi_thread_programming/logging"

var logger logging.Loggering

func init() {
	logger=logging.NewSimpleLogger()
}
