package loadgen

import "go_multi_thread_programming/logging"

var logger logging.Logger

func init() {
	logger=logging.NewSimpleLogger()
}
