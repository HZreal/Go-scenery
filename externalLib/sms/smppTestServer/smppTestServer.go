package main

/**
 * @Author elastic·H
 * @Date 2024-10-22
 * @File: smppTestServer.go
 * @Description: 基于 go-smpp 的测试服务端
 */

import (
	"fmt"
	"github.com/fiorix/go-smpp/smpp/pdu"
	"github.com/fiorix/go-smpp/smpp/pdu/pdufield"
	"github.com/fiorix/go-smpp/smpp/smpptest"
	"log"
	"time"
)

// 自定义处理函数，处理从客户端接收到的消息
func handlePDU1(c smpptest.Conn, m pdu.Body) {
	log.Printf("Received PDU from client: %#v\n", m)

	// 示例：如果收到的是 submit_sm PDU，则返回一个 submit_sm_resp
	if m.Header().ID == pdu.SubmitSMID {
		// 提取 short_message 字段
		shortMessage, ok := m.Fields()[pdufield.ShortMessage]
		if ok {
			// 打印短信内容
			log.Printf("Received Short Message: %s\n", shortMessage.String())
		} else {
			log.Println("No short message found in PDU.")
		}

		resp := pdu.NewSubmitSMResp()
		resp.Header().Seq = m.Header().Seq
		resp.Fields().Set(pdufield.MessageID, "12345") // 短信的标识符 MessageID
		c.Write(resp)
		log.Println("Sent SubmitSMResp to client.")
	}
}

// 源码 transceiver_test.go 的测试处理器，可参考
func handlerInSourceCode(c smpptest.Conn, p pdu.Body) {
	switch p.Header().ID {
	case pdu.SubmitSMID:
		r := pdu.NewSubmitSMResp()
		r.Header().Seq = p.Header().Seq
		r.Fields().Set(pdufield.MessageID, "foobar")
		c.Write(r)
		pf := p.Fields()
		rd := pf[pdufield.RegisteredDelivery]
		if rd.Bytes()[0] == 0 {
			return
		}
		r = pdu.NewDeliverSM()
		f := r.Fields()
		f.Set(pdufield.SourceAddr, pf[pdufield.SourceAddr])
		f.Set(pdufield.DestinationAddr, pf[pdufield.DestinationAddr])
		f.Set(pdufield.ShortMessage, pf[pdufield.ShortMessage])
		c.Write(r)
	default:
		smpptest.EchoHandler(c, p)
	}
}

// 基于 handlerInSource 的改动，模拟 5 秒后的状态发送
func handlePDU2(c smpptest.Conn, p pdu.Body) {
	switch p.Header().ID {
	case pdu.SubmitSMID:
		messageID := generateMessageID()

		r := pdu.NewSubmitSMResp()
		r.Header().Seq = p.Header().Seq
		r.Fields().Set(pdufield.MessageID, messageID)
		c.Write(r)

		go func() {
			time.Sleep(5 * time.Second)

			pf := p.Fields()
			rd := pf[pdufield.RegisteredDelivery]
			if rd.Bytes()[0] == 0 {
				return
			}
			r = pdu.NewDeliverSM()
			f := r.Fields()
			f.Set(pdufield.SourceAddr, pf[pdufield.SourceAddr])
			f.Set(pdufield.DestinationAddr, pf[pdufield.DestinationAddr])
			f.Set(pdufield.ShortMessage, pf[pdufield.ShortMessage])
			// f.Set(pdufield.MessageID, messageID) // 设置 MessageID
			// f.Set(pdufield.MessageState, uint8(2)) // 设置状态为成功递送
			c.Write(r)

			log.Printf("Sent deliver_sm (status report) to client with MessageID: %s\n", messageID)
		}()

	default:
		log.Println("Received PDU from client，default handler for the other PDU types")
		smpptest.EchoHandler(c, p)
	}
}

func generateMessageID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func main() {
	// 创建并启动服务器
	// ！！！注意：端口是随机的
	server := smpptest.NewUnstartedServer()

	// server.Handler = handlePDU1
	// server.Handler = handlerInSourceCode
	server.Handler = handlePDU2

	// 启动
	server.Start()
	log.Printf("SMPP test server is running on %s\n", server.Addr())

	//
	select {}
}
