package main

/**
 * @Author elastic·H
 * @Date 2024-10-22
 * @File: smppTestServer.go
 * @Description:
 */

import (
	"github.com/fiorix/go-smpp/smpp/pdu"
	"github.com/fiorix/go-smpp/smpp/pdu/pdufield"
	"github.com/fiorix/go-smpp/smpp/smpptest"
	"log"
)

// 自定义处理函数，处理从客户端接收到的消息
func handlePDU(c smpptest.Conn, m pdu.Body) {
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

func main() {
	// 创建并启动服务器
	// 注意：端口是随机的！
	server := smpptest.NewUnstartedServer()

	// 设置自定义处理函数
	server.Handler = handlePDU

	// 启动服务器
	server.Start()

	// 打印服务器地址，用于客户端连接
	log.Printf("SMPP test server is running on %s\n", server.Addr())

	// 阻塞以保持服务器运行
	select {}
}
