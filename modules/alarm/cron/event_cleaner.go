package cron

import (
	"github.com/open-falcon/falcon-plus/modules/alarm/g"
	eventmodel "github.com/open-falcon/falcon-plus/modules/alarm/model/event"
	"time"
)

func CleanExpiredEvent() {
	for {

		/*最多保留七天*/
		retention_days := g.Config().Housekeeper.EventRetentionDays
		/*delete batch : 100*/
		delete_batch := g.Config().Housekeeper.EventDeleteBatch

		now := time.Now()
		before := now.Add(time.Duration(-retention_days*24) * time.Hour)
		eventmodel.DeleteEventOlder(before, delete_batch)

		time.Sleep(time.Second * 60)
	}
}
