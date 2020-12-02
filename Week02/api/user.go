package api

import (
	"database/sql"
	"encoding/json"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"practice/Go训练营/3.第二周作业/service"
	"strconv"
)

func GetUserById(w http.ResponseWriter, r *http.Request) {
	// 简单检验输入值
	uids, ok := r.URL.Query()["uid"]
	if !ok || len(uids[0]) < 1 {
		w.WriteHeader(400)
		w.Write([]byte("缺少参数"))
		return
	}

	uid, err := strconv.Atoi(uids[0])
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("参数错误"))
		return
	}

	user, err := service.GetUserInfo(uid)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(404)
		w.Write([]byte("找不到该用户"))
	} else if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("服务错误"))
	}

	// 返回信息
	if err == nil {
		data, _ := json.Marshal(user)
		w.WriteHeader(200)
		w.Write(data)
	}

	if err != nil {
		log.Printf("original error:%T %v\n", errors.Cause(err), errors.Cause(err))
		log.Printf("stack strace:\n%+v\n", err)
	}

}
