package client

import (
	"fmt"
	"github.com/KylinHe/aliensboot-client-test/base"
	"github.com/KylinHe/aliensboot-client-test/msg"
	"github.com/KylinHe/aliensboot-core/log"
	"os"

	"github.com/gogo/protobuf/proto"
	"sync"
	"time"
)

type ClientBot struct {
	sync.RWMutex
	m_Idx int32 //索引

	syncStartTime time.Time

	network string //TCP KCP

	seq      int32
	Token    string
	Username string
	Password string

	synctime  int64

	secretkey string

	m_SrvInfo string //游戏服务器信息

	agent Agent

	m_bReadyDisCon bool //是否准备断开连接

	m_Channel        chan interface{}
	m_Channel_Closed bool //通道是否 已经关闭了

	m_State int32 //玩家状态
}

/**
*	@brief 获取角色索引
 */
func (this *ClientBot) GetIdx() int32 { return this.m_Idx }

/**
*	@brief 初始化
 */
func (this *ClientBot) Init(idx int32, network string, address string, syncTime int64, secretkey string) bool {
	this.m_Idx = idx
	this.synctime = syncTime
	this.m_SrvInfo = address
	this.network = network
	this.secretkey = secretkey
	this.m_State = base.PLAYER_STATE_NONE //默认设置为 无状态
	this._open_channel()             //打开通道

	if !this.IsConnect(){
		//fmt.Printf(">>> try to connect \n");
		if !this.Connect() { //连接失败
			fmt.Printf(">>> connect failed %v\n", this.m_Idx)
			return false
		}
		fmt.Printf(">>> connect success %v\n", this.m_Idx)
	}
	return true
}

func (this *ClientBot) Reconnect() bool {
	//已经连接了这个地址不需要重复连接
	if this.IsConnect() {
		return true
	}
	//this.m_SrvInfo = this.passport
	this.m_State = base.PLAYER_STATE_NONE //默认设置为 无状态
	this._open_channel()             //打开通道

	if this.IsConnect() {
		this.DisConnect()
		fmt.Printf(">>> reconnect success %v \n", this.m_Idx)
	}
	fmt.Printf(">>> try to connect %v\n", this.m_Idx)
	if !this.Connect() { //连接失败
		fmt.Printf(">>> reconnect failed %v\n", this.m_Idx)
		return false
	} else {
		fmt.Printf(">>> reconnect success %v\n", this.m_Idx)
	}
	return true
}

/**
*	@brief 插入操作码
 */
func (this *ClientBot) AcceptOp(op int) {
	if this.m_Channel == nil {
		return
	}
	select {
	case this.m_Channel <- op:
	default:
		fmt.Printf("%d message channel full\n", this.m_Idx)
		//TODO 消息管道满了需要通知客户端消息请求太过频繁
	}
}

func (this *ClientBot) isCipher() bool {
	return this.secretkey != ""
}

/**
*	@brief 请求连接
 */
func (this *ClientBot) Connect() bool {
	this.Lock()
	defer this.Unlock()
	if this.m_SrvInfo == "" {
		return false
	}

	this.agent = NewAgent(this.network)
	this.agent.OnMsg(this.handleMsg)
	err := this.agent.Run(this.m_SrvInfo)
	if err != nil {
		log.Errorf("connect err : %v", err)
		return false
	}
	return true
}

/**
*	@brief 目前是否处于连接中
 */
func (this *ClientBot) IsConnect() bool {
	this.RLock()
	defer this.RUnlock()
	return this.agent != nil
}

/**
*	@brief 断开连接
 */
func (this *ClientBot) DisConnect() bool {
	this.Lock()
	defer this.Unlock()
	if this.agent == nil {
		return false
	}
	err := this.agent.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Idx=%d, %s", this.m_Idx, err.Error())
	}
	this.agent = nil
	return true
}

func (this *ClientBot) handleMsg(msg interface{}) {
	this._receive_Message(msg.(proto.Message))

}

func (this *ClientBot) _open_channel() { //打开消息管道
	if this.m_Channel != nil {
		return
	}
	this.m_Channel_Closed = false
	this.m_Channel = make(chan interface{}, 10) //10个通道的最大数量
	go func() {
		for {
			//只要消息管道没有关闭，就一直等待消息
			v, open := <-this.m_Channel
			if !open {
				this.m_Channel = nil
				break
			}
			opType, _ := v.(int)
			this._accept_op(opType)
		}
		this._close_channel()
	}()
}

func (this *ClientBot) _close_channel() { //关闭通道
	if this.m_Channel_Closed == true {
		return
	} //已经准备要关闭了
	if this.m_Channel == nil {
		return
	}
	close(this.m_Channel)
	this.m_Channel_Closed = true
}

func (this *ClientBot) _accept_op(opType int) {
	switch opType {
	case base.OP_SYNC:
		this.syncData()
		break
	}
}

func (this *ClientBot) syncData() {
	this.seq++
	sessionID := this.seq*10000 + this.m_Idx
	this.syncStartTime = time.Now()
	this._send_Message(1, msg.BuildLoginRequest(this.Username, "111111", sessionID))
}


func (this *ClientBot) _send_Message(id uint16, message proto.Message) { //发送消息
	this.RLock()
	defer this.RUnlock()
	if this.agent == nil {
		return
	}
	this.agent.WriteMsg([]interface{}{id, message})
}

func (this *ClientBot) _receive_Message(message proto.Message) { //接收消息
	fmt.Printf("%v => receive: %v %v\n", this.Username, message, time.Now())
	this.syncData()
	//code := message.GetSequence()
}
