package main

import (
	"Go_learn/server/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
)

func setHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                           //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,Authorization") //header的类型
	w.Header().Set("content-type", "application/json")                           //返回数据格式是json
}

func insertData(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	if r.Method != "POST" {
		logs.Warning("insertData method is not POST")
		w.Write([]byte("false"))
		return
	}

	// 解析表单 主体 r.body
	r.ParseForm()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logs.Error("read r.Body error: \n", err)
		return
	}

	var newPeople models.Person // 返回数据时用的Marshal 解析数据 Unmarshal
	err = json.Unmarshal(body, &newPeople)
	if err != nil {
		logs.Error("unmarshal r.Body error: \n", err)
		return
	}
	logs.Debug(newPeople)

	_, err = models.DB.NamedExec(`INSERT INTO person (first_name,last_name,email) VALUES (:first,:last,:email)`,
		map[string]interface{}{
			"first": newPeople.FirstName,
			"last":  newPeople.LastName,
			"email": newPeople.Email,
		})
	if err != nil {
		logs.Error(err)
		w.Write([]byte("false"))
	}

	logs.Debug("插入数据成功")
	w.Write([]byte("true"))
}

func getData(w http.ResponseWriter, r *http.Request) {
	setHeader(w) // 跨域问题
	if r.Method != "GET" {
		logs.Warning("getData method is not GET")
		return
	}

	var people []models.Person
	err := models.DB.Select(&people, "SELECT * FROM person ORDER BY first_name ASC")
	if err != nil {
		logs.Error(err)
		return
	}
	logs.Debug(people)
	//fmt.Printf("%+v\n",people)

	// 单行查询 QueryRow
	jason := models.Person{}
	err = models.DB.Get(&jason, "SELECT * FROM person WHERE first_name=$1", "Jason")
	fmt.Printf("%#v\n", jason)

	data, err := json.Marshal(people)
	if err != nil {
		logs.Error("marshal people error,\n", err)
		return
	}

	w.Write(data)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/getData", getData)
	mux.HandleFunc("/insertData", insertData)

	err := http.ListenAndServe(":9800", mux)
	if err != nil {
		logs.Error("listen and serve error:\n", err)
	}
}
