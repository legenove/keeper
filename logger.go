package keeper

import (
    "sync"
)

var Logger = logger
var lgmu sync.RWMutex

func logger(reason, keeperName string) {
    //fmt.Println(keeperName, ":", reason)
}

func GetLogger() func(reason, keeperName string) {
    lgmu.RLock()
    defer lgmu.RUnlock()
    return Logger
}
func SetLogger(l func(reason, keeperName string)) {
    lgmu.Lock()
    defer lgmu.Unlock()
    Logger = l
}
