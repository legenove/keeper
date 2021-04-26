package keeper

import (
    "github.com/RussellLuo/timingwheel"
    "sync"
    "time"
)

type keeperSyncFunc func() error

type keeperJob struct {
    syncFunc keeperSyncFunc
    interval time.Duration
}

var keeperMapper = make(map[string]*keeperJob)
var keeperLock sync.Mutex
var timeWheel *timingwheel.TimingWheel

func init() {
    timeWheel = timingwheel.NewTimingWheel(500*time.Millisecond, 1000)
    timeWheel.Start()
}

// syncFunc is a func for run return an error
// interval is repeat interval for sync function
// execNow is execute immediate
func SetKeeper(keeperName string, syncFunc keeperSyncFunc, intervalDuration time.Duration, execNow bool) {
    defer recoverError(keeperName)
    if intervalDuration < 500*time.Millisecond {
        intervalDuration = 500 * time.Millisecond
    }
    kj := &keeperJob{
        syncFunc: syncFunc,
        interval: intervalDuration,
    }
    keeperLock.Lock()
    defer keeperLock.Unlock()
    keeperMapper[keeperName] = kj
    setKeeperToTimeWheel(keeperName, intervalDuration, execNow)
}

func DelKeeper(keeperName string) {
    keeperLock.Lock()
    defer keeperLock.Unlock()
    delete(keeperMapper, keeperName)
}

func SetDelay(delayName string, syncFunc keeperSyncFunc, intervalDuration time.Duration) {
    defer recoverError(delayName)
    timeWheel.AfterFunc(intervalDuration, func() {
        syncDelay(delayName, syncFunc)
    })
}

func setKeeperToTimeWheel(keeperName string, intervalDuration time.Duration, execNow bool) {
    var d time.Duration
    if execNow {
        d = 0
    } else {
        d = intervalDuration
    }
    timeWheel.AfterFunc(d, func() {
        syncKeeper(keeperName)
    })
}

func syncKeeper(keeperName string) {
    keeperLock.Lock()
    if ke, ok := keeperMapper[keeperName]; ok {
        keeperLock.Unlock()
        doTask(keeperName, ke.syncFunc)
        d := ke.interval
        timeWheel.AfterFunc(d, func() {
            syncKeeper(keeperName)
        })
        return
    } else {
        keeperLock.Unlock()
    }
}

func syncDelay(keeperName string, syncFunc keeperSyncFunc) {
    doTask(keeperName, syncFunc)
}

func doTask(keeperName string, f keeperSyncFunc) {
    defer recoverError(keeperName)
    err := f()
    if err != nil {
        GetLogger()(err.Error(), keeperName)
    }
}

func recoverError(keeperName string) {
    if err := recover(); err != nil {
        var reason string
        switch err.(type) {
        case error:
            reason = err.(error).Error()
        default:
            reason = "unknow error"
        }
        GetLogger()(reason, keeperName)
    }
}
