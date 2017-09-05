package cron

import (
	"encoding/json"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"
	cmodel "github.com/open-falcon/falcon-plus/common/model"
	"github.com/open-falcon/falcon-plus/modules/alarm/g"
	eventmodel "github.com/open-falcon/falcon-plus/modules/alarm/model/event"
)

func ReadHighEvent() {
	/*event:p0 event:p1*/
	queues := g.Config().Redis.HighQueues
	if len(queues) == 0 {
		return
	}

	for {
		event, err := popEvent(queues)
		if err != nil {
			time.Sleep(time.Second)
			continue
		}
		consume(event, true)
	}
}

func ReadLowEvent() {
	queues := g.Config().Redis.LowQueues
	if len(queues) == 0 {
		return
	}

	for {
		event, err := popEvent(queues)
		if err != nil {
			time.Sleep(time.Second)
			continue
		}
		consume(event, false)
	}
}

/*弹出一个event*/
func popEvent(queues []string) (*cmodel.Event, error) {

	count := len(queues)

	/*params 是为了组装redis命令使用，params[count]=0, queues不能被改变，因此make一个*/
	params := make([]interface{}, count+1)
	for i := 0; i < count; i++ {
		params[i] = queues[i]
	}
	// set timeout 0
	params[count] = 0

	rc := g.RedisConnPool.Get() /*从redis中拿一个conn*/
	defer rc.Close()

	/*BRPOP event:p0 event:p1 0
		BRPOP 按照给出的key顺序查看list，找到第一个非空的list的尾部弹出一个元素
	    redis> RPUSH list1 a b c
	            (integer) 3
	    redis> BRPOP list1 list2 0
	       1) "list1"
	       2) "c"
	*/

	reply, err := redis.Strings(rc.Do("BRPOP", params...))
	if err != nil {
		log.Errorf("get alarm event from redis fail: %v", err)
		return nil, err
	}

	var event cmodel.Event
	/*json encode*/
	err = json.Unmarshal([]byte(reply[1]), &event)
	if err != nil {
		log.Errorf("parse alarm event fail: %v", err)
		return nil, err
	}

	log.Debugf("pop event: %s", event.String())

	//insert event into database
	eventmodel.InsertEvent(&event)
	// events no longer saved in memory

	return &event, nil
}
