/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2019/1/9
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package msg

import (
	"encoding/binary"
	"github.com/KylinHe/aliensboot-server/protocol"
	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
)

type ProtoProcessor struct {

}

func (*ProtoProcessor) UNMarshal(data []byte) (interface{}, error) {
	recv := &protocol.Response{}
	err := proto.Unmarshal(data, recv)
	return recv, err
}


func (*ProtoProcessor) Marshal(msg interface{}) ([]byte, error) {
	ret, ok := msg.([]interface{})
	if !ok {
		return nil, errors.New("invalid msg content")
	}
	id := ret[0].(uint16)
	message := ret[1].(proto.Message)
	buff, err := proto.Marshal(message)

	if err != nil { //序列化失败
		return nil, err
	}
	//fmt.Printf("%v => send: %v - %v %v\n", this.Username, id, message, time.Now())
	//if this.isCipher() {
	//	buff = xxtea.Encrypt(buff, []byte(this.secretkey))
	//	//fmt.Printf("key %v", this.secretkey)
	//}

	m := make([]byte, len(buff)+4)
	binary.LittleEndian.PutUint16(m, uint16(len(buff))+2) //
	binary.LittleEndian.PutUint16(m[2:], id)

	copy(m[4:], buff)
	return m, nil
}

