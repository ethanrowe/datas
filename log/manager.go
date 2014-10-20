package log

import (
    "log"
)

type LogResultLike interface {
    Code() int
    Error() error
}

type LogResult struct {
    code int
    err  error
}

func (r LogResult) Code() int {
    return r.code
}

func (r LogResult) Error() error {
    return r.err
}

type PostResult struct {
    LogResult
    LogID int
}

type LogOperation func(*Store, chan<- LogResultLike)

type LogOperator struct {
    Operation LogOperation
    Chan chan<- LogResultLike
}

type Manager struct {
    store *Store
    requests chan *LogOperator
    done chan int
}

func New(log *Store) *Manager {
    return &Manager{log, make(chan *LogOperator), make(chan int)}
}

func (m *Manager) Activate() {
    for operator := range m.requests {
        log.Print("Manager performing next operation.")
        operator.Operation(m.store, operator.Chan)
    }
    log.Print("Manager request queue closed.")
    close(m.done)
}

func (m *Manager) Close() {
    close(m.requests)
}

func (m *Manager) Wait() {
    <-m.done
}

func (m *Manager) LogWindow(window string) (logid int, err error) {
    var result PostResult
    resultChan := make(chan LogResultLike)
    fn := func (s *Store, r chan<- LogResultLike) {
        windowId := s.Next()
        log.Printf("Logged window id %d for window: %s", windowId, window)
        r <- PostResult{LogResult{1, nil}, windowId}
    }
    m.requests <- &LogOperator{fn, resultChan}
    result = (<-resultChan).(PostResult)
    return result.LogID, result.Error()
}

func (m *Manager) PrintSequence() (err error) {
    resultChan := make(chan LogResultLike)
    m.requests <- &LogOperator{
        func (s *Store, r chan<- LogResultLike) {
            log.Printf("Current window id sequence value is: %d", s.Current())
            r <- new(LogResult)
        },
        resultChan,
    }
    return (<-resultChan).Error()
}

