/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2019/1/9
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package client

import (
	"fmt"
	"github.com/KylinHe/aliensboot-core/log"
	"github.com/xtaci/kcp-go"
	"sync"
)

type KCPAgent struct {
	con *kcp.UDPSession

	sync.RWMutex

	processor MsgProcessor

	handler MsgHandler
}

func (this *KCPAgent) OnMsg(handler MsgHandler) {
	this.handler = handler
}

func (this *KCPAgent) WriteMsg(obj interface{}) error {
	data, err := this.processor.Marshal(obj)
	if err != nil {
		return err
	}
	_, err1 := this.con.Write(data)
	return err1
}


func (this *KCPAgent) Close() error {
	this.RLock()
	defer this.RUnlock()
	return this.con.Close()
}



func (this *KCPAgent) Run(address string) error {
	conn, err := kcp.DialWithOptions(address, nil, 10, 3)
	if err != nil {
		return err
	}
	conn.SetReadBuffer(64 * 1024)
	conn.SetWriteBuffer(64 * 1024)
	conn.SetNoDelay(1, 10, 2, 1)
	conn.SetStreamMode(true)
	conn.SetWriteDelay(false)
	conn.SetACKNoDelay(true)
	this.con = conn

	buf := make([]byte, 8182)
	go func() {
		for {
			len, err := this.con.Read(buf)
			if err != nil {
				this.Lock()
				defer this.Unlock()
				this.con = nil
				fmt.Printf("read error: %v\n", err.Error())
				break
			}


			data := buf[4:len]
			//if this.isCipher() {
			//	data = xxtea.Decrypt(data, []byte(this.secretkey))
			//}
			msg, err := this.processor.UNMarshal(data)
			if err != nil {
				log.Errorf("unmarshal msg err :", err)
				continue
			}

			if this.handler != nil {
				this.handler(msg)
			}
		}
	}()
	return nil
}
