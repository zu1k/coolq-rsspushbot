package bot

import (
	"fmt"
	"strings"

	"github.com/zu1k/coolq-pushbot/core/model"
)

func StartCmdCtr(qqnum int64, isGroup bool) {
	SendMsg(qqnum, fmt.Sprintf("欢迎使用RSS订阅机器人。"), isGroup)
}

func SubCmdCtr(qqnum int64, msg string, isGroup bool) {
	msgs := strings.SplitN(msg, " ", 2)
	if len(msgs) == 2 {
		url := msgs[1]
		if CheckUrl(url) {
			registFeed(qqnum, url, isGroup)
		} else {
			SendMsg(qqnum, "请在命令后跟上合法的url", isGroup)
		}
	} else {
		SendMsg(qqnum, "请发送 /sub RSS URL", isGroup)
	}
}

func ListCmdCtr(qqnum int64, isGroup bool) {
	sources, _ := model.GetSourcesByUserID(qqnum)
	message := "当前订阅列表："
	if len(sources) == 0 {
		message = "订阅列表为空"
	} else {
		i := 0
		for index, source := range sources {
			i++
			if i == 5 {
				SendMsg(qqnum, message, isGroup)
				message = fmt.Sprintf("[%d] %s\n%s", index+1, source.Title, source.Link)
				i = 0
			} else {
				message = message + fmt.Sprintf("\n\n[%d] %s\n%s", index+1, source.Title, source.Link)
			}
		}
	}
	SendMsg(qqnum, message, isGroup)
}

func UnsubCmdCtr(qqnum int64, msg string, isGroup bool) {
	msgs := strings.SplitN(msg, " ", 2)
	if len(msgs) == 2 {
		url := msgs[1]
		if CheckUrl(url) {
			source, _ := model.GetSourceByUrl(url)
			sub, err := model.GetSubByUserIDAndURL(qqnum, url)
			if err != nil {
				if err.Error() == "record not found" {
					SendMsg(qqnum, "未订阅该RSS源", isGroup)
				} else {
					SendMsg(qqnum, "退订失败", isGroup)
				}
			} else {
				err := sub.Unsub()
				if err == nil {
					SendMsg(qqnum, fmt.Sprintf("退订 [%s](%s) 成功", source.Title, source.Link), isGroup)
				} else {
					SendMsg(qqnum, err.Error(), isGroup)
				}
			}
		}
	}
}

func HelpCmdCtr(qqnum int64, isGroup bool) {
	message := `
命令：
/sub 订阅源
/unsub  取消订阅
/list 查看当前订阅源
/help 帮助
`
	SendMsg(qqnum, message, isGroup)
}
