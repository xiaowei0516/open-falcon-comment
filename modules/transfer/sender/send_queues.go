package sender

import (
	"github.com/open-falcon/falcon-plus/modules/transfer/g"
	nlist "github.com/toolkits/container/list"
)

/*为每个发送的对象创建一个定长的queue
目的：
1. 起缓冲作用，减缓后端graph、judge的压力，尤其是尖峰流量
2. 避免流量的丢失，起缓冲作用，当连接池异常的时候，重连可以保障流量不丢失
*/
func initSendQueues() {
	cfg := g.Config()
	for node := range cfg.Judge.Cluster {
		Q := nlist.NewSafeListLimited(DefaultSendQueueMaxSize)
		JudgeQueues[node] = Q
	}

	/*
	   一个graph可以配置多个addr，可以保障数据的安全性，当一个宕机，另外一个照常工作。
	*/
	for node, nitem := range cfg.Graph.ClusterList {
		for _, addr := range nitem.Addrs {
			Q := nlist.NewSafeListLimited(DefaultSendQueueMaxSize)
			/*graph-0011.11.11.11:8000*/
			GraphQueues[node+addr] = Q
		}
	}

	if cfg.Tsdb.Enabled {
		TsdbQueue = nlist.NewSafeListLimited(DefaultSendQueueMaxSize)
	}
}
