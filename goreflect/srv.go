package main

import "encoding/json"

type LoginArgs struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

type Data struct {
	Uid   int    `json:"uid"`
	Token string `json:"token"`
}

type LoginReply struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data *Data  `json:"data"`
}

func NewRequest(js string) *LoginArgs {
	var arg *LoginArgs
	err := json.Unmarshal([]byte(js), &arg)
	if err != nil {
		return nil
	}

	return arg
}

func NewReply(js interface{}) string {
	if _, ok := js.(*LoginReply); !ok {
		return ""
	}

	r := js.(*LoginReply)
	s, err := json.Marshal(r)
	if err != nil {
		return ""
	}

	return string(s)
}
