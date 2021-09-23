package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	clientv3 "go.etcd.io/etcd/client/v3"
	"reflect"
	"strconv"
	"strings"
)

type Login struct {
	*User
}
type User struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	salt     int64
}
type LoginResponse struct {
	*User
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}

func (l *Login) Verify(val interface{}) bool {
	v := reflect.ValueOf(val).Elem()
	t := v.Type()
	var user = new(*User)
	u := reflect.ValueOf(user).Elem()
	for i := 0; i < v.NumField(); i++ {
		tv := u.FieldByName(t.Field(i).Name)
		if tv.IsValid() == false {
			continue
		}
		tv.Set(v.Field(i))
	}
	if user != nil {
		h := md5.New()
		h.Write([]byte(l.Password))
		h.Write([]byte(strconv.FormatInt(l.salt, 10)))
		newPass := hex.EncodeToString(h.Sum(nil))
		return reflect.DeepEqual(newPass, val)

	}
	return false
}

type ListOptions struct {
	PageInfo
}

type PageInfo struct {
	Page     int64  `json:"page" form:"page"`
	PageSize int64  `json:"pageSize" form:"pageSize"`
	Key      string `json:"key" form:"key"`
}

type Page struct {
	Total    int64 `json:"total"`
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}

type App struct {
	Name    string   `json:"name" form:"name"`       //name
	AppId   string   `json:"appId" form:"appId"`     //appId
	GroupId []string `json:"groupId" form:"groupId"` //groupId
}

func MarshalJSON(a interface{}) ([]byte, error) {
	var maps = make(map[string]interface{})
	v := reflect.ValueOf(a).Elem()
	st := v.Type()
	for i := 0; i < v.NumField(); i++ {
		key := strings.Split(st.Field(i).Tag.Get("json"), ",")[0]
		maps[key] = v.Field(i).Interface()
	}
	return json.Marshal(maps)
}

type Group struct {
}

type PageList struct {
	List interface{} `json:"list"`
	Page
}
type Result struct {
	Result interface{} `json:"result"`
}

func (p *PageList) Unmarshal(total int64, response *clientv3.GetResponse) {
	if len(response.Kvs) > 0 {
		list := make([]interface{}, 0)
		for _, kv := range response.Kvs {
			list = append(list, string(kv.Value))
		}
		p.List = &list
		p.Total = total
		//p.PageSize

	}
	//response.Kvs
	//json.Unmarshal()
}
