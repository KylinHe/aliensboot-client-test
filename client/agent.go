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
	"github.com/KylinHe/aliensboot-client-test/constant"
	"github.com/KylinHe/aliensboot-client-test/msg"
)

type Agent interface {

	Run(address string) error

	Close() error

	OnMsg(handler MsgHandler)

	WriteMsg(data interface{}) error

}

type MsgProcessor interface {

	Marshal(message interface{}) ([]byte, error)

	UNMarshal(data []byte) (interface{}, error)
}

type MsgHandler func(msg interface{})


func NewAgent(network string) Agent {
	if network == constant.NetworkTCP {
		return &TCPAgent{processor:&msg.ProtoProcessor{}}
	} else {
		return &KCPAgent{processor:&msg.ProtoProcessor{}}
	}
}