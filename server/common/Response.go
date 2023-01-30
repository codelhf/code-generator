package common

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"` //omitempty忽略空值，陷阱：结构体(对象)和带默认值类型的使用*指针类型
}

const (
	SUCCESS = 200
	FAIL    = 999
	ERROR   = 500
)

func Success(w http.ResponseWriter) {
	response(w, SUCCESS, "", nil)
}

func SuccessMsg(w http.ResponseWriter, msg string) {
	response(w, SUCCESS, msg, nil)
}

func SuccessData(w http.ResponseWriter, data interface{}) {
	response(w, SUCCESS, "", data)
}

func SuccessMsgData(w http.ResponseWriter, msg string, data interface{}) {
	response(w, SUCCESS, msg, data)
}

func FailMsg(w http.ResponseWriter, msg string) {
	response(w, FAIL, msg, nil)
}

func response(w http.ResponseWriter, code int, msg string, data interface{}) {
	//定义一个结构体
	res := Response{Code: code, Msg: msg, Data: data}
	//将结构体转化成JSON字符串
	ret, err := json.Marshal(res)
	if err != nil {
		log.Println(err.Error())
	}
	//设置响应头
	w.Header().Set("Content-Type", "application/json")
	//设置200状态
	w.WriteHeader(http.StatusOK)
	//输出
	w.Write(ret)
}
