package g

var Modules map[string]bool
var BinOf map[string]string
var PidOf map[string]string
var AllModulesInOrder []string
var ModuleApps map[string]string
var logpathOf map[string]string


func init() {
	Modules = map[string]bool{
		"agent":      true,
		"aggregator": true,
		"graph":      true,
		"hbs":        true,
		"judge":      true,
		"nodata":     true,
		"transfer":   true,
		"gateway":    true,
		"api":        true,
		"alarm":      true,
	}

	BinOf = map[string]string{
		"agent":      "./agent/bin/opsultra-agent",
		"aggregator": "./aggregator/bin/opsultra-aggregator",
		"graph":      "./graph/bin/opsultra-graph",
		"hbs":        "./hbs/bin/opsultra-hbs",
		"judge":      "./judge/bin/opsultra-judge",
		"nodata":     "./nodata/bin/opsultra-nodata",
		"transfer":   "./transfer/bin/opsultra-transfer",
		"gateway":    "./gateway/bin/opsultra-gateway",
		"api":        "./api/bin/opsultra-api",
		"alarm":      "./alarm/bin/opsultra-alarm",
	}

	cfgOf = map[string]string{
		"agent":      "./agent/config/cfg.json",
		"aggregator": "./aggregator/config/cfg.json",
		"graph":      "./graph/config/cfg.json",
		"hbs":        "./hbs/config/cfg.json",
		"judge":      "./judge/config/cfg.json",
		"nodata":     "./nodata/config/cfg.json",
		"transfer":   "./transfer/config/cfg.json",
		"gateway":    "./gateway/config/cfg.json",
		"api":        "./api/config/cfg.json",
		"alarm":      "./alarm/config/cfg.json",
	}

	ModuleApps = map[string]string{
		"agent":      "opsultra-agent",
		"aggregator": "opsultra-aggregator",
		"graph":      "opsultra-graph",
		"hbs":        "opsultra-hbs",
		"judge":      "opsultra-judge",
		"nodata":     "opsultra-nodata",
		"transfer":   "opsultra-transfer",
		"gateway":    "opsultra-gateway",
		"api":        "opsultra-api",
		"alarm":      "opsultra-alarm",
	}

	logpathOf = map[string]string{
		"agent":      "./agent/logs/agent.log",
		"aggregator": "./aggregator/logs/aggregator.log",
		"graph":      "./graph/logs/graph.log",
		"hbs":        "./hbs/logs/hbs.log",
		"judge":      "./judge/logs/judge.log",
		"nodata":     "./nodata/logs/nodata.log",
		"transfer":   "./transfer/logs/transfer.log",
		"gateway":    "./gateway/logs/gateway.log",
		"api":        "./api/logs/api.log",
		"alarm":      "./alarm/logs/alarm.log",
	}

	PidOf = map[string]string{
		"agent":      "<NOT SET>",
		"aggregator": "<NOT SET>",
		"graph":      "<NOT SET>",
		"hbs":        "<NOT SET>",
		"judge":      "<NOT SET>",
		"nodata":     "<NOT SET>",
		"transfer":   "<NOT SET>",
		"gateway":    "<NOT SET>",
		"api":        "<NOT SET>",
		"alarm":      "<NOT SET>",
	}

	// Modules are deployed in this order
	AllModulesInOrder = []string{
		"graph",
		"hbs",
		"judge",
		"transfer",
		"nodata",
		"aggregator",
		"agent",
		"gateway",
		"api",
		"alarm",
	}
}

func Bin(name string)string{
    p,_ := filepath.Abs(BinOf[name])
    retrun p
}

func Cfg(name string)string{
    p,_ := filepath.Abs(cfgOf[name])
    return p
}

func LogPath(name string)string{
    p,_ := filepath.Abs(logpathOf[name])
    return p
}
func LogDir(name string)string{
    d,_ := filepath.Abs(filepath.Dir(logpathOf[name]))
    return d
}














