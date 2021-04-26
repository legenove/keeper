package keeper

import (
    "fmt"
    "testing"
    "time"
)

func TestSetKeeper(t *testing.T) {
    SetKeeper("test1", test1Func, 1*time.Second, true)
    SetKeeper("test2", test2Func, 2*time.Second, true)
    SetKeeper("test3", test3Func, 3*time.Second, true)
    time.Sleep(10 * time.Second)
    DelKeeper("test1")
    time.Sleep(5 * time.Second)
    SetDelay("deley_test1", test1Func, 2*time.Second)
    time.Sleep(10 * time.Second)
}

func test1Func() error {
    fmt.Println("test1, func", time.Now().Unix())
    return nil
}

func test2Func() error {
    fmt.Println("test2, func", time.Now().Unix())
    return nil
}

func test3Func() error {
    fmt.Println("test3, func", time.Now().Unix())
    return nil
}
