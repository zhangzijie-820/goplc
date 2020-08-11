package goplc

import (
	"github.com/MiguelValentine/goplc/enip/cip/epath/segment"
	"log"
	"testing"
	"time"
)

func TestEncapsulation(t *testing.T) {
	//go tcpServer()
	r := segment.PortBuild(1, []byte{0x01})
	log.Printf("%#x\n", r)

	a, b := NewOriginator("10.211.55.7", 1, nil)
	err := a.Connect()
	log.Println(a, b, err)
	time.Sleep(time.Second * 5)
}

//func tcpServer() {
//	host := ":10809"
//	tcpAddr, err := net.ResolveTCPAddr("tcp4", host)
//	if err != nil {
//		log.Printf("resolve tcp addr failed: %v\n", err)
//		return
//	}
//	listener, err := net.ListenTCP("tcp", tcpAddr)
//	if err != nil {
//		log.Printf("listen tcp port failed: %v\n", err)
//		return
//	}
//	conn, err := listener.AcceptTCP()
//	if err != nil {
//		log.Printf("Accept failed:%v\n", err)
//	}
//
//	data := &encapsulation.Encapsulation{}
//	data.Command = encapsulation.CommandRegisterSession
//	_buf := data.Buffer()
//	log.Println(_buf)
//
//	_, _ = conn.Write(_buf[0:3])
//	time.Sleep(time.Millisecond * 200)
//	_, _ = conn.Write(_buf[3:])
//	conn.Close()
//}
