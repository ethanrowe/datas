package service

import (
    "net/http"
    "io"
    "log"
    store "github.com/ethanrowe/datas/log"
)

func toString(r io.Reader) string {
    buf := make([]byte, 1000)
    n, err := r.Read(buf)
    if n == 0 && err != nil {
        log.Fatal(err)
    }
    return string(buf[:])
}

func handleWindowPost(m *store.Manager, w http.ResponseWriter, req *http.Request) {
    if req.Method == "POST" {
        _, _ = m.LogWindow(toString(req.Body))
        io.WriteString(w, "Okay.\n")
    } else {
        io.WriteString(w, "Nope.\n")
    }
}

func Serve(addr string, mgr *store.Manager) error {
    http.HandleFunc("/windows", func(w http.ResponseWriter, req *http.Request) {
        handleWindowPost(mgr, w, req)
    })
    return http.ListenAndServe(addr, nil)
}


