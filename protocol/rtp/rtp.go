package rtp

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	//SERVER_IP       = "127.0.0.1"
	SERVER_IP       = "0.0.0.0"
	SERVER_PORT     = 5082
	SERVER_RECV_LEN = 1024
)

func InitRtp() {
	fmt.Println("=========Start Rtp")
	go StartRtp()
	fmt.Println("=========End Rtp")
}
func StartRtp() {
	address := SERVER_IP + ":" + strconv.Itoa(SERVER_PORT)
	fmt.Println("address:", address)
	addr, err := net.ResolveUDPAddr("udp", address)
	fmt.Println("addr:", addr, "|", "error:", err)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", addr)
	fmt.Println("conn::listen:", conn, "error:", err)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer conn.Close()

	for {
		// Here must use make and give the lenth of buffer
		data := make([]byte, SERVER_RECV_LEN)
		//fmt.Println("data:", data)
		_, rAddr, err := conn.ReadFromUDP(data)
		//fmt.Println("rAddr:", rAddr, "error:", err)
		fmt.Println("=============================================数据来了")
		if err != nil {
			fmt.Println(err)
			continue
		}

		strData := string(data)
		fmt.Println("===============Received:", strData)

		upper := strings.ToUpper(strData)
		_, err = conn.WriteToUDP([]byte(upper), rAddr)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("============Send:", upper)
	}
}
