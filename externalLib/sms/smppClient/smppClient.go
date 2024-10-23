package main

/**
 * @Author elastic·H
 * @Date 2024-10-22
 * @File: smppClient.go
 * @Description: 基于 go-smpp 的客户端发送短信 demo
 */

import (
	"errors"
	"fmt"
	"github.com/fiorix/go-smpp/smpp"
	"github.com/fiorix/go-smpp/smpp/pdu"
	"github.com/fiorix/go-smpp/smpp/pdu/pdufield"
	"github.com/fiorix/go-smpp/smpp/pdu/pdutext"
	"github.com/fiorix/go-smpp/smpp/pdu/pdutlv"
	"github.com/fiorix/go-smpp/smpp/smpptest"
	"log"
)

// 注意：连接 测试服务端 时，端口不固定，看 测试服务端 启动的实际端口
var port = "65130"

// 简单 demo，仅发送，对应测试服务端的 handlePDU1 处理
func test1() {
	tx := &smpp.Transmitter{
		Addr: "127.0.0.1:" + port,
		// Addr: "123.253.140.228:7099",

		// 账户名称
		User: smpptest.DefaultUser,
		// User:   "222001",

		// 密码
		Passwd: smpptest.DefaultPasswd,
		// Passwd: "9?_)YT9K",
	}
	conn := tx.Bind()
	defer tx.Close()

	// 查看连接状态
	var status smpp.ConnStatus
	if status = <-conn; status.Error() != nil {
		log.Fatalln("Unable to connect, aborting:", status.Error())
	}
	log.Println("Connection completed, status:", status.Status().String())

	// 异步检查状态
	go func() {
		for c := range conn {
			log.Println("SMPP connection status:", c.Status())
		}
	}()

	// 发送
	sm, err := tx.Submit(&smpp.ShortMessage{
		// 发送号
		Src: "10690",
		// 目的接收号
		Dst: "5514996740534",
		// 内容
		Text: pdutext.Raw([]byte("发送的短信内容 xxxxxxxxxxxxxxxxxxxxxxxxxx")),
		// 设置短信最终状态的回执
		Register: pdufield.FinalDeliveryReceipt,
		// 传递附加信息
		TLVFields: pdutlv.Fields{
			pdutlv.TagReceiptedMessageID: pdutlv.CString("msgIdIIIIIIIIIIIIIIIII"),
		},
	})
	if errors.Is(err, smpp.ErrNotConnected) {
		log.Println("SMPP connection not connected", err)
		return
	}
	if err != nil {
		log.Println("SMPP error", err)
		return
	}

	fmt.Println("sm.RespID()  ---->  ", sm.RespID())
	resHeader := sm.Resp().Header()
	fmt.Println("resHeader  ---->  ", resHeader)
	resLen := sm.Resp().Len()
	fmt.Println("resLen  ---->  ", resLen)
	resFields := sm.Resp().Fields()
	fmt.Println("resFields  ---->  ", resFields)
	resFieldList := sm.Resp().FieldList()
	fmt.Println("resFieldList  ---->  ", resFieldList)
	resTLVFields := sm.Resp().TLVFields()
	fmt.Println("resTLVFields  ---->  ", resTLVFields)

	// time.Sleep(30 * time.Second)
	select {}
}

// 终端结果：
// 2024/10/22 15:44:37 Connection completed, status: Connected
// sm.RespID()  ---->   AB3EC846FD13102F
// resHeader  ---->   &{33 SubmitSMResp OK 2}
// resLen  ---->   33
// resFields  ---->   map[message_id:AB3EC846FD13102F]
// resFieldList  ---->   [message_id]
// resTLVFields  ---->   map[]

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// 处理状态报告（deliver_sm）
func handleDeliverSM(p pdu.Body) {
	log.Println("handleDeliverSM Received a PDU for status ---->  ")

	// 提取 MessageID
	fields := p.Fields()
	if messageID, ok := fields[pdufield.MessageID]; ok {
		fmt.Printf("Received status report for MessageID: %s\n", messageID.String())

		// 提取 MessageState
		if messageState, ok := fields[pdufield.MessageState]; ok {
			fmt.Printf("Message State: %d\n", messageState)

			// 根据状态更新数据库中的消息状态（如递送成功或失败）
			// updateMessageStateInDatabase(messageID.String(), messageState.String())
		}
	}
}

// 发送后等待接收状态回执
func test2() {
	//
	tc := &smpp.Transceiver{
		Addr: "127.0.0.1:" + port,
		// Addr: "123.253.140.228:7099",
		// 账户名称
		User: smpptest.DefaultUser,
		// User:   "222001",
		Passwd: smpptest.DefaultPasswd,
		// Passwd: "9?_)YT9K",

		Handler: handleDeliverSM, // 设置处理 PDU 的 Handler
	}
	conn := tc.Bind()
	defer tc.Close()

	// 查看连接状态
	var status smpp.ConnStatus
	if status = <-conn; status.Error() != nil {
		log.Fatalln("Unable to connect, aborting:", status.Error())
	}
	log.Println("Connection completed, status:", status.Status().String())

	// 异步检查状态
	go func() {
		for c := range conn {
			log.Println("SMPP connection status:", c.Status())
		}
	}()

	// 发送
	sm, err := tc.Submit(&smpp.ShortMessage{
		Src:  "10690",
		Dst:  "5514996740534",
		Text: pdutext.Raw([]byte("发送的短信内容 xxxxxxxxxxxxxxxxxxxxxxxxxx")),
		// 设置短信最终状态的回执
		Register: pdufield.FinalDeliveryReceipt,
		TLVFields: pdutlv.Fields{
			pdutlv.TagReceiptedMessageID: pdutlv.CString("msgIdIIIIIIIIIIIIIIIII"),
		},
	})
	if errors.Is(err, smpp.ErrNotConnected) {
		log.Println("SMPP connection not connected", err)
		return
	}
	if err != nil {
		log.Println("SMPP error", err)
		return
	}

	fmt.Println("sm.RespID()  ---->  ", sm.RespID())

	// time.Sleep(30 * time.Second)
	select {}
}

func main() {
	// test1()
	test2()
}
