package main

import (
	"strings"

	"github.com/zu1k/coolq-pushbot/core/task"

	"github.com/Tnze/CoolQ-Golang-SDK/v2/cqp"
	"github.com/zu1k/coolq-pushbot/core/bot"
)

//go:generate cqcfg -c .
// cqp: 名称: RSSPushBot
// cqp: 版本: 1.0.0:1
// cqp: 作者: zu1k
// cqp: 简介: RSS订阅机器人
func main() { /*此处应当留空*/ }

func init() {
	cqp.AppID = "co.tgbot.rsspush"
	cqp.PrivateMsg = onPrivateMsg
	cqp.GroupMsg = onGroupMsg
	cqp.FriendAdd = friendAdd
	cqp.FriendRequest = onFriendRequest
	cqp.GroupRequest = onGroupRequest
	cqp.Start = Start
}

func onPrivateMsg(subType, msgID int32, fromQQ int64, msg string, font int32) int32 {
	msgs := strings.SplitN(msg, " ", 2)
	order := msgs[0]
	switch order {
	case "/start":
		bot.StartCmdCtr(fromQQ, false)
	case "/sub":
		bot.SubCmdCtr(fromQQ, msg, false)
	case "/list":
		bot.ListCmdCtr(fromQQ, false)
	case "/unsub":
		bot.UnsubCmdCtr(fromQQ, msg, false)
	case "/help":
		bot.HelpCmdCtr(fromQQ, false)
	}
	return 0
}

func onGroupMsg(subType, msgID int32, fromGroup, fromQQ int64, fromAnonymous, msg string, font int32) int32 {
	msgs := strings.SplitN(msg, " ", 2)
	order := msgs[0]
	switch order {
	case "/sub":
		bot.SubCmdCtr(fromGroup, msg, true)
	case "/list":
		bot.ListCmdCtr(fromGroup, true)
	case "/unsub":
		bot.UnsubCmdCtr(fromGroup, msg, true)
	case "/help":
		bot.HelpCmdCtr(fromGroup, true)
	}
	return 0
}

func onGroupRequest(subType, sendTime int32, fromGroup, fromQQ int64, msg, responseFlag string) int32 {
	cqp.SetGroupAddRequest(responseFlag, 2, 1)
	return 0
}

func onFriendRequest(subType, sendTime int32, fromQQ int64, msg, responseFlag string) int32 {
	cqp.SetFriendAddRequest(responseFlag, 1, "")
	return 0
}

func friendAdd(subType, sendTime int32, fromQQ int64) int32 {
	bot.HelpCmdCtr(fromQQ, false)
	return 0
}

func Start() int32 {
	go task.Update()
	return 0
}
