package main

import (
    store "github.com/ethanrowe/datas/log"
    "log"
    "strconv"
    "sync"
)

func main() {
    mgr := store.New(new(store.Store))

    log.Print("Enqueuing 10 operations.")
    wg := new(sync.WaitGroup)
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(inc int) {
            if inc % 3 != 0 {
                name := strconv.Itoa(inc)
                log.Print("Requesting log for window: ", name)
                logid, err := mgr.LogWindow(name)
                log.Print("Window ", name, " got log id: ", logid, " error ", err)
            } else {
                log.Printf("At %d asking for PrintSequence", inc)
                err := mgr.PrintSequence()
                log.Print("Increment ", inc, " got err: ", err)
            }
            log.Printf("Done with goroutine increment %d", inc)
            wg.Done()
        }(i)
    }

    log.Print("Activating the log manager.")
    go mgr.Activate()
    log.Print("Waiting for input goroutines.")
    wg.Wait()
    log.Print("Closing manager.")
    mgr.Close()
    log.Print("Waiting for manager.")
    mgr.Wait()
}

