package test_3_net

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"testing"
	"time"
)

//TCP listen
func TestTcp(t *testing.T) {
	//定义一个tcp断点
	var tcpAddr *net.TCPAddr
	//通过ResolveTCPAddr实例一个具体的tcp断点
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	//打开一个tcp断点监听
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()
	fmt.Println("Server ready to read ...")
	//循环接收客户端的连接，创建一个协程具体去处理连接
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("A client connected :" + tcpConn.RemoteAddr().String())
		go tcpPipe(tcpConn)
	}
}

func tcpPipe(conn *net.TCPConn) {
	//tcp连接的地址
	ipStr := conn.RemoteAddr().String()

	defer func() {
		fmt.Println(" Disconnected : " + ipStr)
		_ = conn.Close()
	}()

	//获取一个连接的reader读取流
	reader := bufio.NewReader(conn)
	i := 0
	//接收并返回消息
	for {
		message, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		fmt.Println(string(message))

		time.Sleep(time.Second * 3)

		msg := time.Now().String() + conn.RemoteAddr().String() + " Server Say hello! \n"

		b := []byte(msg)

		_, _ = conn.Write(b)

		i++

		if i > 10 {
			break
		}
	}
}

//TCP write
func TestWriteTCP(t *testing.T) {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")

	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		fmt.Println("Client connect error ! " + err.Error())
		return
	}

	defer conn.Close()

	fmt.Println(conn.LocalAddr().String() + " : Client connected!")

	onMessageReceived(conn)
	return
}

func onMessageReceived(conn *net.TCPConn) {

	reader := bufio.NewReader(conn)
	b := []byte(conn.LocalAddr().String() + " Say hello to Server... \n")
	_, _ = conn.Write(b)
	for {
		msg, err := reader.ReadString('\n')
		fmt.Println("ReadString")
		fmt.Println(msg)

		if err != nil || err == io.EOF {
			fmt.Println(err)
			break
		}
		time.Sleep(time.Second * 2)

		fmt.Println("writing...")

		b := []byte(conn.LocalAddr().String() + " write data to Server... \n")
		_, err = conn.Write(b)

		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func TestUdp(t *testing.T) {

	//定义一个udp断点
	var udpAddr *net.UDPAddr
	//通过ResolveUDPAddr实例一个具体的tcp断点
	udpAddr, _ = net.ResolveUDPAddr("udp", "127.0.0.1:9999")
	//打开一个udp断点监听
	udpListener, _ := net.ListenUDP("tcp", udpAddr)
	defer udpListener.Close()
	fmt.Println("Server ready to read ...")
	//循环接收客户端的连接，创建一个协程具体去处理连接
	for {
		var buf []byte
		_, _, err := udpListener.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(string(buf))
		//fmt.Println("A client connected :" + tcpConn.RemoteAddr().String())
		//go tcpPipe(tcpConn)
	}

}
