package i18n

import (
	"encoding/json"
	"flag"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
	"os"
)

type locale struct {
	Name     string
	Langs    map[string]interface{}
	TimeZone string
}

// 所有的 locales.
var Locales = map[string]locale{}

func init() {
	flag.Set("logtostderr", "true")

	flag.Parse()

	// TODO: 加载所有语言配置
	bytes, _ := ioutil.ReadFile("i18n/zh_CN.json")

	zhCN := locale{Name: "zh_CN"}

	// TODO: 时区

	err := json.Unmarshal(bytes, &zhCN.Langs)
	if err != nil {
		glog.Error(err)

		os.Exit(-1)
	}

	Locales["zh_CN"] = zhCN
	glog.Info("Loaded [zh_CN] locale configuration")
}

func GetLangs(r *http.Request) map[string]interface{} {
	locale := getLocale(r)

	return Locales[locale].Langs
}

func getLocale(r *http.Request) string {
	// TODO: 从请求中获取 locale
	return "zh_CN"
}
