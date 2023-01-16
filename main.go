package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/no-src/golongman/json"
	"github.com/no-src/golongman/provider"
	"github.com/no-src/log"
)

func main() {
	log.Info("hello, welcome to golongman !")
	recommendTodayProvider := provider.NewRecommendTodayProvider()
	recommend, err := recommendTodayProvider.GetRecommendToday()
	if err != nil {
		log.Error(err, "获取今日推荐信息异常")
		return
	}
	jsonStr := _json.GetJsonString(recommend)
	fileName := "./data/" + time.Now().Format("2006-01-02") + ".today.json"
	err = os.MkdirAll(filepath.Dir(fileName), os.ModePerm)
	if err != nil {
		log.Error(err, "创建数据存储目录错误")
		return
	}
	err = ioutil.WriteFile(fileName, []byte(jsonStr), os.ModePerm)
	if err != nil {
		log.Error(err, "写入json文件错误")
	}
	log.Info("今日推荐信息：%s", jsonStr)
}
