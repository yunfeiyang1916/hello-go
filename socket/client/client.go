//客户端
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//tcpClient()
	udpClient()
}

//tcp客户端
func tcpClient() {
	service := "127.0.0.1:7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte("HEAD /HTTP/1.0\r\n\r\n"))
	checkError(err)
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("客户端名称：")
	clientName, _ := inputReader.ReadString('\n')
	//去除\r\n
	trimmedClient := strings.Trim(clientName, "\r\n")
	_, err = conn.Write([]byte(trimmedClient))
	checkError(err)
	for {
		fmt.Println("请输入：")
		input, _ := inputReader.ReadString('\n')
		//去除\r\n
		trimmedInput := strings.Trim(input, "\r\n")
		if strings.ToUpper(trimmedInput) == "Q" {
			return
		}
		conn.Write([]byte(trimmedClient + " 发送消息：" + trimmedInput))
	}
}

//udp客户端
func udpClient() {
	service := "127.0.0.1:7777"
	updAddress, err := net.ResolveUDPAddr("udp", service)
	checkError(err)
	conn, err := net.DialUDP("udp", nil, updAddress)
	checkError(err)
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Printf("请输入客户端名称：")
	clientName, _ := inputReader.ReadString('\n')
	//去除\r\n
	trimmedClient := strings.Trim(clientName, "\r\n")
	conn.Write([]byte(trimmedClient))
	for {
		input, _ := inputReader.ReadString('\n')
		//去除\r\n
		trimmedInput := strings.Trim(input, "\r\n")
		conn.Write([]byte(fmt.Sprintf("%s 发送消息：%s", trimmedClient, trimmedInput)))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "致命错误：%s", err.Error())
		os.Exit(1)
	}
}
