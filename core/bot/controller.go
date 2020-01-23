package bot

import (
	"fmt"
	"strings"

	"github.com/Tnze/CoolQ-Golang-SDK/v2/cqp"
	"github.com/zu1k/coolq-pushbot/core/model"
)

func StartCmdCtr(qqnum int64) {
	cqp.SendPrivateMsg(qqnum, fmt.Sprintf("欢迎使用RSS订阅机器人。"))
}

func SubCmdCtr(qqnum int64, msg string) {
	msgs := strings.SplitN(msg, " ", 2)
	if len(msgs) == 2 {
		url := msgs[1]
		if CheckUrl(url) {
			registFeed(qqnum, url)
		} else {
			cqp.SendPrivateMsg(qqnum, "请在命令后跟上合法的url")
		}
	} else {
		cqp.SendPrivateMsg(qqnum, "请发送 /sub RSS URL")
	}
}

func ListCmdCtr(qqnum int64) {
	sources, _ := model.GetSourcesByUserID(qqnum)
	message := "当前订阅列表："
	if len(sources) == 0 {
		message = "订阅列表为空"
	} else {
		for index, source := range sources {
			message = message + fmt.Sprintf("\n\n[%d] %s\n%s", index+1, source.Title, source.Link)
		}
	}
	cqp.SendPrivateMsg(qqnum, message)
}

func UnsubCmdCtr(qqnum int64, msg string) {
	msgs := strings.SplitN(msg, " ", 2)
	if len(msgs) == 2 {
		url := msgs[1]
		if CheckUrl(url) {
			source, _ := model.GetSourceByUrl(url)
			sub, err := model.GetSubByUserIDAndURL(qqnum, url)
			if err != nil {
				if err.Error() == "record not found" {
					cqp.SendPrivateMsg(qqnum, "未订阅该RSS源")
				} else {
					cqp.SendPrivateMsg(qqnum, "退订失败")
				}
			} else {
				err := sub.Unsub()
				if err == nil {
					cqp.SendPrivateMsg(qqnum, fmt.Sprintf("退订 [%s](%s) 成功", source.Title, source.Link))
				} else {
					cqp.SendPrivateMsg(qqnum, err.Error())
				}
			}
		}
	}
}

func HelpCmdCtr(qqnum int64) {
	message := `
命令：
/sub 订阅源
/unsub  取消订阅
/list 查看当前订阅源
/help 帮助
`
	cqp.SendPrivateMsg(qqnum, message)
}
