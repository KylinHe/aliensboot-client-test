/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2017/11/10
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package msg

import (
	. "github.com/KylinHe/aliensboot-server/protocol"
)

//构建注册消息
//登录服务器注册账号
//message login_register{
//optional string username = 1;	    //用户名
//optional string password = 2;		//密码
//}
//
////登录服务器注册账号返回
//message login_register_ret {
//optional register_Result result = 1;    //注册结果
//optional int64 uid = 2;                 //用户id 注册成功返回此字段
//optional string token = 3;              //登录令牌 注册成功返回此字段
//optional string gameServer = 4;         //返回游戏服务器地址
//optional string msg = 5;                //反馈消息 登录失败返回
//}

func BuildLoginRequest(username string, pwd string, sessionID int32) *Request {
	message := &Request{
		Passport:&Request_UserLogin{
			UserLogin:&UserLogin{
				Username: username,
				Password: pwd,
			},
		},
	}
	return message
}

//func BuildRequest1(session int32, uid string, content string) *service1.Request1 {
//	message := &service1.Request1{
//		Session:  proto.Int32(session),
//		No: proto.Int(1),
//		Request:proto.String(uid + content),
//	}
//
//	return message
//}
//
//func BuildRequest2(session int32, uid string, content string) *service2.Request2 {
//	message := &service2.Request2{
//		Session:  proto.Int32(session),
//		No: proto.Int(1),
//		Request:proto.String(uid + content),
//	}
//
//	return message
//}
