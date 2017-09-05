package cache

import (
	"log"
	"time"
)

func Init() {
	log.Println("cache begin")

	log.Println("#1 GroupPlugins...")
	GroupPlugins.Init()

	//得到所有的组模板列表，一个组ID可以对应多个模板
	log.Println("#2 GroupTemplates...")
	GroupTemplates.Init()

	//得到所有的hostgroup列表
	log.Println("#3 HostGroupsMap...")
	HostGroupsMap.Init()

	//得到所有的host信息
	log.Println("#4 HostMap...")
	HostMap.Init()

	log.Println("#5 TemplateCache...")
	//得到所有的template列表，其实就是执行select × from template；
	TemplateCache.Init()

	log.Println("#6 Strategies...")
	//得到所有的strategy列表，就是执行select × from strategy；
	Strategies.Init(TemplateCache.GetMap())

	log.Println("#7 HostTemplateIds...")
	//主机组中同一个host对应多个模板， select a.tpl_id, b.host_id from grp_tpl as a inner join grp_host as b on a.grp_id=b.grp_id;
	HostTemplateIds.Init()

	log.Println("#8 ExpressionCache...")
	//得到所有的express信息
	ExpressionCache.Init()

	log.Println("#9 MonitoredHosts...")
	//得到所有的host列表
	MonitoredHosts.Init()

	log.Println("cache done")

	go LoopInit()

}

//每分钟刷新一次cache列表
func LoopInit() {
	for {
		time.Sleep(time.Minute)
		GroupPlugins.Init()
		GroupTemplates.Init()
		HostGroupsMap.Init()
		HostMap.Init()
		TemplateCache.Init()
		Strategies.Init(TemplateCache.GetMap())
		HostTemplateIds.Init()
		ExpressionCache.Init()
		MonitoredHosts.Init()
	}
}
