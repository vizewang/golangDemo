package loadgen

import (
	"go_multi_thread_programming/cpt7/loadgen/lib"
	"time"
)

func NewGenerator(caller lib.Caller,timeoutNs time.Duration,lps uint32,durationNs time.Duration,resultCh chan *lib.CallResult)(lib.Generator,error) {
	logger.Infoln("New a load generator...")
	logger.Infoln("Checking the parameters...")
}