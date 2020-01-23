package bot

import (
	"fmt"
	"log"

	"github.com/Tnze/CoolQ-Golang-SDK/v2/cqp"

	"github.com/zu1k/coolq-pushbot/core/config"
	"github.com/zu1k/coolq-pushbot/core/model"
)

func registFeed(qqnum int64, url string) {
	cqp.SendPrivateMsg(qqnum, "处理中...")
	source, err := model.FindOrNewSourceByUrl(url)
	if err != nil {
		cqp.SendPrivateMsg(qqnum, fmt.Sprintf("%s，订阅失败", err))
		return
	}

	err = model.RegistFeed(qqnum, source.ID)
	log.Printf("%d subscribe [%d]%s %s", qqnum, source.ID, source.Title, source.Link)

	if err == nil {
		cqp.SendPrivateMsg(qqnum, fmt.Sprintf("[%s](%s) 订阅成功", source.Title, source.Link))
	} else {
		cqp.SendPrivateMsg(qqnum, "订阅失败")
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
				log.Println("BroadNews tpldata.Render err ", err)
				cqp.SendPrivateMsg(1393385930, "Render error")
				return
			}
			cqp.SendPrivateMsg(sub.UserID, msg)
		}
	}
}

func BroadSourceError(source *model.Source) {
	subs := model.GetSubscriberBySource(source)
	for _, sub := range subs {
		message := fmt.Sprintf("[%s](%s) 已经累计连续%d次更新失败，暂时停止更新", source.Title, source.Link, config.ErrorThreshold)
		cqp.SendPrivateMsg(sub.UserID, message)
	}
}
