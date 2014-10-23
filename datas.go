package main

import (
    store "github.com/ethanrowe/datas/log"
    "github.com/ethanrowe/datas/service"
    "log"
)

func main() {
    mgr := store.New(new(store.Store))
    log.Print("Activating the log manager.")
    go mgr.Activate()

    log.Print("Activating the server on port 8080.")
    service.Serve(":8080", mgr)

    log.Print("Stopped server; closing manager.")
    mgr.Close()
    mgr.Wait()
}

