package main

import (
	"log"
	"strings"

	"github.com/zu1k/coolq-pushbot/core/task"

	"github.com/Tnze/CoolQ-Golang-SDK/v2/cqp"
	"github.com/zu1k/coolq-pushbot/core/bot"
)

//go:generate cqcfg -c .
// cqp: 名称: RSSPushBot
// cqp: 版本: 1.0.0:0
// cqp: 作者: zu1k
// cqp: 简介: RSS订阅机器人
func main() { /*此处应当留空*/ }

func init() {
	cqp.AppID = "co.tgbot.rsspush"
	cqp.PrivateMsg = onPrivateMsg
	cqp.FriendAdd = friendAdd
	cqp.Start = Start
}

func onPrivateMsg(subType, msgID int32, fromQQ int64, msg string, font int32) int32 {
	log.Printf("QQ: %d\tMsg: %s", fromQQ, msg)
	msgs := strings.SplitN(msg, " ", 2)
	order := msgs[0]
	switch order {
	case "/start":
		bot.StartCmdCtr(fromQQ)
	case "/sub":
		bot.SubCmdCtr(fromQQ, msg)
	case "/list":
		bot.ListCmdCtr(fromQQ)
	case "/unsub":
		bot.UnsubCmdCtr(fromQQ, msg)
	case "/help":
		bot.HelpCmdCtr(fromQQ)
	}
	return 0
}

func friendAdd(subType, sendTime int32, fromQQ int64) int32 {
	bot.HelpCmdCtr(fromQQ)
	return 0
}

func Start() int32 {
	go task.Update()
	return 0
}
