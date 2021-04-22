package tcpscan
import (
    "base"
    "fmt"
    //"time"
    "regexp"
    "sync"
)

// what to do when a syn is sent but we get no resp (aka RST) ??

func Scan(host string) []int {
    tcpConns := make(chan int, 1000)
    var waitgroup sync.WaitGroup
    var mutex sync.Mutex

    tcpPorts := []int{}
    timeout := 20
    for port := 1; port <= 10000; port++ {
        waitgroup.Add(1)
        go connect(host, port, timeout, tcpConns, &waitgroup, &mutex)
    }
    waitgroup.Wait()
    close(tcpConns)
    for tcpConn := range tcpConns {
        tcpPorts = append(tcpPorts, tcpConn)
    }
    //fmt.Println(tcpPorts)
    return tcpPorts
}

func connect(host string, port int, timeout int, tcpConns chan int, wg *sync.WaitGroup, m *sync.Mutex) {
    m.Lock()
    sock, err := base.NewSocket(host, port, "tcp", timeout)
    m.Unlock()
    if err != nil {
        if matched, _ := regexp.MatchString(`files`, err.Error()); matched {
            fmt.Println(err)
        } else {
            //fmt.Println(err)
        }

    }
    if sock != nil && err == nil {
        m.Lock()
        tcpConns<-port
        //fmt.Printf("PORT %d detected ++++++++++++++++\n", port)
        sock.Close()
        m.Unlock()
    }
    wg.Done()
}
