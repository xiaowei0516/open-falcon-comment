package plugins

import (
	"github.com/open-falcon/falcon-plus/modules/agent/g"
	"github.com/toolkits/file"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

// key: sys/ntp/60_ntp.py
//传入 sys/ntp ,然后进行拼接 ./plugin/sys/ntp,搜索 该文件下的所有文件，找到 $cycle_$xx 类型
//的文件，然后拼接成最终的文件./plugin/sys/ntp/600_ntp.py
//return
// ret["./plugin/sys/ntp/600_ntp.py"] = Plugin
// plugin 包含文件的名称、修改时间、Cycle
func ListPlugins(relativePath string) map[string]*Plugin {
	ret := make(map[string]*Plugin)
	if relativePath == "" {
		return ret
	}

	dir := filepath.Join(g.Config().Plugin.Dir, relativePath)

	if !file.IsExist(dir) || file.IsFile(dir) {
		return ret
	}

	fs, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println("can not list files under", dir)
		return ret
	}

	for _, f := range fs {
		if f.IsDir() {
			continue
		}

		filename := f.Name()
		arr := strings.Split(filename, "_")
		if len(arr) < 2 {
			continue
		}

		// filename should be: $cycle_$xx
		var cycle int
		cycle, err = strconv.Atoi(arr[0])
		if err != nil {
			continue
		}

		fpath := filepath.Join(relativePath, filename)
		plugin := &Plugin{FilePath: fpath, MTime: f.ModTime().Unix(), Cycle: cycle}
		ret[fpath] = plugin
	}

	return ret
}
