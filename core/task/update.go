package task

import (
	"time"

	"github.com/zu1k/coolq-pushbot/core/bot"
	"github.com/zu1k/coolq-pushbot/core/config"
	"github.com/zu1k/coolq-pushbot/core/model"
)

func Update() {
	for {
		sources := model.GetSubscribedNormalSources()
		for _, source := range sources {
			c, err := source.GetNewContents()
			if err == nil {
				subs := model.GetSubscriberBySource(&source)
				bot.BroadNews(&source, subs, c)
			}
			if source.ErrorCount >= config.ErrorThreshold {
				bot.BroadSourceError(&source)
			}
		}
		time.Sleep(time.Duration(1) * time.Minute)
	}
}
