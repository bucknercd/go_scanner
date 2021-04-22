package base
import (
    //"fmt"
    "net"
    "time"
    "strconv"
)

// TODO: implement a timeout for the Recv func! conn.SetReadDeadline : https://stackoverflow.com/questions/34892507/reading-from-a-tcp-connection-in-golang


type BaseSocket struct {
    host string
    port int
    transport string
    timeout int
    readTimeout int
    conn net.Conn
}

func NewSocket(host string, port int, transport string, timeout int) (*BaseSocket, error) {
    sock := BaseSocket{host: host, port: port, transport: transport, timeout: timeout, readTimeout: 15}
    conn, err := net.DialTimeout(transport,  host + ":" + strconv.Itoa(port), time.Duration(timeout) * time.Millisecond)
    if err != nil {
        return nil, err
    }
    sock.conn = conn
    return &sock, nil
}

func (s *BaseSocket) Send(data []byte) (int, error) {
    n, err := s.conn.Write(data)
    return n, err
}

func (s *BaseSocket) Recv() ([]byte, error) {
    s.conn.SetReadDeadline(time.Now().Add(time.Duration(s.readTimeout) * time.Second))
    buf := make([]byte, 4096)
    n, err := s.conn.Read(buf)
    return buf[:n], err
}

func (s *BaseSocket) Close() {
    s.conn.Close()
}

func (s *BaseSocket) SetReadTimeout(timeout int) {
    s.conn.SetReadDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
}