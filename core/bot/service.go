package bot

import (
	"fmt"

	"github.com/zu1k/coolq-pushbot/core/config"
	"github.com/zu1k/coolq-pushbot/core/model"
)

func registFeed(qqnum int64, url string, isGroup bool) {
	SendMsg(qqnum, "处理中...", isGroup)
	source, err := model.FindOrNewSourceByUrl(url)
	if err != nil {
		SendMsg(qqnum, fmt.Sprintf("%s，订阅失败", err), isGroup)
		return
	}
	err = model.RegistFeed(qqnum, source.ID, isGroup)
	if err == nil {
		SendMsg(qqnum, fmt.Sprintf("[%s](%s) 订阅成功", source.Title, source.Link), isGroup)
	} else {
		SendMsg(qqnum, "订阅失败", isGroup)
	}
}

//BroadNews send new contents message to subscriber
func BroadNews(source *model.Source, subs []model.Subscribe, contents []model.Content) {
	for _, content := range contents {
		previewText := trimDescription(content.Description, config.PreviewText)
		for _, sub := range subs {
			tpldata := &config.TplData{
				SourceTitle:  source.Title,
				ContentTitle: content.Title,
				RawLink:      content.RawLink,
				PreviewText:  previewText,
			}
			msg, err := tpldata.Render()
			if err != nil {
				return
			}
			SendMsg(sub.UserID, msg, sub.IsGroup)
		}
	}
}
