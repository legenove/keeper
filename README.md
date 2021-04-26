# keeper

## Introduce

Keeper is a plug-in of schedule task and delay task, which is based on time wheel.

## Quick start

Add a schedule task

* the `SetKeeper` first param should be unique.
```go
packet example


import (
    "fmt"
    "time"
)


func main() {
   SetKeeper("example1", exampleFunc, 1*time.Second, true)
   time.Sleep(10 *time.Second)
   DelKeeper("example1")
}

func exampleFunc() error {
    fmt.Println("example, func", time.Now().Unix())
    return nil
}

```

Add a delay task

```go
packet example


import (
    "fmt"
    "time"
)


func main() {
   SetDelay("deley_example", exampleFunc, 2*time.Second)
   time.Sleep(10 *time.Second)
}

func exampleFunc() error {
    fmt.Println("example, func", time.Now().Unix())
    return nil
}

```
