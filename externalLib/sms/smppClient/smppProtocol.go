package main

/**
 * @Author elastic·H
 * @Date 2024-10-22
 * @File: smppProtocol.go
 * @Description:
 */

import (
	"errors"
	"fmt"
	"github.com/fiorix/go-smpp/smpp"
	"github.com/fiorix/go-smpp/smpp/pdu/pdufield"
	"github.com/fiorix/go-smpp/smpp/pdu/pdutext"
	"github.com/fiorix/go-smpp/smpp/pdu/pdutlv"
	"github.com/fiorix/go-smpp/smpp/smpptest"
	"log"
	"time"
)

func main() {
	// make persistent connection
	tx := &smpp.Transmitter{
		Addr:   "127.0.0.1:54280",
		User:   smpptest.DefaultUser,
		Passwd: smpptest.DefaultPasswd,
	}
	conn := tx.Bind()
	defer tx.Close()

	// check initial connection status
	var status smpp.ConnStatus
	if status = <-conn; status.Error() != nil {
		log.Fatalln("Unable to connect, aborting:", status.Error())
	}
	log.Println("Connection completed, status:", status.Status().String())

	// example of connection checker goroutine
	go func() {
		for c := range conn {
			log.Println("SMPP connection status:", c.Status())
		}
	}()

	sm, err := tx.Submit(&smpp.ShortMessage{
		Src:  "10690",
		Dst:  "5514996740534",
		Text: pdutext.Raw([]byte("发送的短信内容 xxxxxxxxxxxxxxxxxxxxxxxxxx")),
		// 状态回执
		Register: pdufield.NoDeliveryReceipt,
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

	time.Sleep(1 * time.Minute)
}

// 2024/10/22 15:44:37 Connection completed, status: Connected
// sm.RespID()  ---->   AB3EC846FD13102F
// resHeader  ---->   &{33 SubmitSMResp OK 2}
// resLen  ---->   33
// resFields  ---->   map[message_id:AB3EC846FD13102F]
// resFieldList  ---->   [message_id]
// resTLVFields  ---->   map[]
