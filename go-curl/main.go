package main

import (
	"fmt"
	"net"
	"time"
)

var ch = make(chan int)

func main() {
	conn, err := net.DialTimeout("tcp", "localhost:8080", time.Second)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	req := "GET / HTTP/1.1\r\nHost: localhost:8080\r\n\r\n"
	conn.Write([]byte(req))

	conn.SetReadDeadline(time.Now().Add(2 * time.Second))

	// 没有SetReadDeadline 不会断开,不会打印响应信息
	// var buf bytes.Buffer
	// len, err := io.Copy(&buf, conn)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("len%d:%s", len, buf.String())

	// 在没有 SetReadDeadline 的情况下 会自动断开，并不会报出超时异常，
	// 会打印响应信息
	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", string(buf))

	// 没有SetReadDeadline 不会断开,会打印响应信息
	// buf := make([]byte, 1024)
	// len := 0
	// for {
	// 	fmt.Printf("%d %s", len, string(buf))
	// 	n, err := conn.Read(buf[len:])
	// 	if n > 0 {
	// 		fmt.Printf("n:%d\n", n)
	// 		len += n
	// 	}
	// 	if err != nil {
	// 		if err != io.EOF {
	// 			fmt.Println(err)
	// 			break
	// 		}
	// 		break
	// 	}
	// }

	// 没有SetReadDeadline 不会断开,不会打印响应信息
	// buf1, err := ioutil.ReadAll(conn)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(buf1))
}
