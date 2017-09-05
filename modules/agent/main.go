package main

import (
	"flag"
	"fmt"
	"github.com/open-falcon/falcon-plus/modules/agent/cron"
	"github.com/open-falcon/falcon-plus/modules/agent/funcs"
	"github.com/open-falcon/falcon-plus/modules/agent/g"
	"github.com/open-falcon/falcon-plus/modules/agent/http"
	"os"
)

func main() {

	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	check := flag.Bool("check", false, "check collector")

	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	if *check {
		funcs.CheckCollector()
		os.Exit(0)
	}

	//解析cfg.json
	fmt.Println(*cfg)
	//存储cfgjson内容到全局变量中
	g.ParseConfig(*cfg)

	//struct
	if g.Config().Debug {
		g.InitLog("debug")
	} else {
		g.InitLog("info")
	}

	g.InitRootDir()
	g.InitLocalIp()
	g.InitRpcClients()

	funcs.BuildMappers()

	go cron.InitDataHistory() //每秒钟更新cpu和磁盘的信息

	cron.ReportAgentStatus() //指定时间向HBS汇报信息,hostname plugin_version IP  agent_version

	//同步插件
	cron.SyncMinePlugins()
	//同步内置metric，ｐｒｏｃ\ports\net\du
	cron.SyncBuiltinMetrics()
	cron.SyncTrustableIps()
	cron.Collect()

	go http.Start()

	select {}

}
