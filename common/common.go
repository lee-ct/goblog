package common

import (
	"encoding/json"
	"go-blog/config"
	"go-blog/model"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var Template model.HtmlTemplate

// 加载模版
func LoadTemplate() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	//耗时
	go func() {
		var err error
		Template, err = model.InitTemplate(config.Cfg.SystemConfig.CurrentDir + "/template/")
		if err != nil {
			panic(err)
		}

		wg.Done()
	}()
	wg.Wait()
}

func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var parms map[string]interface{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("读取body失败", err)
	}

	_ = json.Unmarshal(body, &parms)

	return parms
}

func Success(w http.ResponseWriter, data interface{}) {
	var result model.Result
	result.Data = data
	result.Error = ""
	result.Code = 200
	resJson, _ := json.Marshal(result)
	_, err := w.Write(resJson)
	if err != nil {
		log.Println(err)
	}
}

func Error(w http.ResponseWriter, data interface{}) {
	var result model.Result
	result.Data = data
	result.Error = ""
	result.Code = -999
	resJson, _ := json.Marshal(result)
	w.Write(resJson)
}
