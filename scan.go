package main
import (
        //"base"
        "fmt"
        "flag"
        "strings"
        "log"
        "ipaddr"
        "hostalive"
        "sync"
)

var target string

func main() {
    flag.StringVar(&target, "t", "default", "Target. Ip range or a single ip.")
    //flag.IntVar
    flag.Parse()
    if target == "default" {
        log.Fatal("Please provide a target. (valid ip range or a single ip)");
    }

    launched := make(chan int)
    //done := make(chan bool)
    limit := 0

    ipRange := strings.Split(target, "-")
    startIP, _ := ipaddr.NewIP(ipRange[0])
    var endIP *ipaddr.IP
    if len(ipRange) > 1 {
        endIP, _ = ipaddr.NewIP(ipRange[1])
    }
    ip := startIP

    var waitgroup sync.WaitGroup
    var mutex sync.Mutex

    for {
        waitgroup.Add(1)
        go scan(ip.String(), launched, &limit, &waitgroup, &mutex)
        limit += <-launched
        //fmt.Println(limit)
        for limit > 256 {
        }
        if endIP == nil || ip.Equals(endIP) {
            break
        }
        ip = ip.Next()
    }
    waitgroup.Wait()
    fmt.Println("Scan completed.")
}

func scan(ip string, launched chan int, limit *int, wg *sync.WaitGroup, m *sync.Mutex) {
    fmt.Printf("Scanning ip %s ...\n", ip)
    launched<-1
    hostalive.Scan(ip, m)
    m.Lock()
    *limit -= 1
    m.Unlock()
    wg.Done()
}

// create a list of ips to scan!