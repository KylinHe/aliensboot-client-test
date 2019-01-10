package main

import (
	"flag"
	"fmt"
	"github.com/KylinHe/aliensboot-client-test/base"
	"github.com/KylinHe/aliensboot-client-test/client"
	"github.com/KylinHe/aliensboot-client-test/constant"
	"strconv"
	"time"
)

var g_Players map[int32]*client.ClientBot

var (
	gate    	  string
	network       string
	preaccount    string
	password      string
	accountnum    int
	synctime      int
	exist         bool
	secretkey     string
)

//stresstest -accountserver "127.0.0.1:18811" -gate "127.0.0.1:18812" -preaccount "test_" -password "11111111" -accountnum 3000 -synctime 5 -exist=true
func main() {
	//形参
	flag.StringVar(&gate, "gate", "", "gate address")
	flag.StringVar(&preaccount, "preaccount", "", "pre account")

	flag.StringVar(&password, "password", "11111111", "password")
	flag.StringVar(&network, "network", constant.NetworkTCP, "tcp, kcp, ws")
	flag.StringVar(&secretkey, "secretkey", "", "cipher key")
	flag.IntVar(&accountnum, "accountnum", 1, "account number")
	flag.IntVar(&synctime, "synctime", 5, "synctime=5s")
	flag.BoolVar(&exist, "exist", false, "account is exist?")

	flag.Parse()

	if gate == "" || preaccount == "" || password == "" {
		println("Please input correct params => stresstest -h")
		println("stresstest -gate \"127.0.0.1:18812\" -preaccount \"test\" -password \"11111111\" -secretkey \"abcd\" -accountnum 3000 -synctime 5 -exist=true")
		return
	}

	println("secript key: " + secretkey)
	g_Players = make(map[int32]*client.ClientBot)

	fmt.Printf("##Create %v connections\n", accountnum)
	time.Sleep(3 * time.Second)

	if exist {
		fmt.Println("##Login account")
	} else {
		fmt.Println("##Register account")
	}

	for i := 1; i <= accountnum; i++ {
		p := &client.ClientBot{}
		p.Username = preaccount + "_" + strconv.Itoa(i)
		p.Password = password
		//p.Init(int32(i), "127.0.0.1:3568")
		p.Init(int32(i), network, gate, int64(synctime), secretkey)
		g_Players[p.GetIdx()] = p
	}
	
	fmt.Println("##Sync data")

	for _, p := range g_Players {
		p.AcceptOp(base.OP_SYNC)
	}
	time.Sleep(time.Hour)
}
