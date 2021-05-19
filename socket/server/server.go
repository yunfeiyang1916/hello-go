//服务端
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	//tcpServer()
	udpServer()
}

//tcp服务端
func tcpServer() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleTcpClient(conn)
	}
}

//处理Tcp客户端
func handleTcpClient(conn net.Conn) {
	//设置读取超时时间，两分钟
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	defer conn.Close()
	for {
		//一次最多读取128个字节
		request := make([]byte, 128)
		_, err := conn.Read(request)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}

		fmt.Println(string(request))
		//链接是否关闭
		// if readLen == 0 {
		// 	break
		// } else if strings.TrimSpace(string(request[:readLen])) == "timestamp" { //是否打印时间戳
		// 	dayTime := strconv.FormatInt(time.Now().Unix(), 10)
		// 	conn.Write([]byte(dayTime))
		// } else {
		// 	daytime := time.Now().String()
		// 	conn.Write([]byte(daytime))
		// }
	}
}

//udp服务端
func udpServer() {
	service := ":7777"
	udpAddress, err := net.ResolveUDPAddr("udp", service)
	checkError(err)
	conn, err := net.ListenUDP("udp", udpAddress)
	checkError(err)
	for {
		handleTcpClient(conn)
	}
}

func httpServer() {

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "致命错误：%s", err.Error())
		os.Exit(1)
	}
}
