package main

import (
	"fmt"
	"kube-opt/action"
	"kube-opt/constant"
	"os"
)

var actions map[string]func(string)

func main() {
	namespace := os.Getenv("INPUT_NAMESPACE")
	_action := os.Getenv("INPUT_ACTION")
	action_list := make([]string, 0, len(actions))
	for _act := range actions {
		action_list = append(action_list, _act)
	}
	if len(_action) == 0 {
		println(fmt.Sprintf("Must specific one action in list:%+v", action_list))
		os.Exit(1)
	}

	if len(namespace) == 0 {
		namespace = "default"
	}
	println(fmt.Sprintf("Use namespace:%s", namespace))

	println(fmt.Sprintf("Use action:%s", _action))
	if act, ok := actions[_action]; !ok {
		println(fmt.Sprintf("Invalid action:%s,please use action in list:%+v", _action, action_list))
		os.Exit(1)
	} else {
		act(namespace)
	}

	os.Exit(0)
}

func init() {
	actions = make(map[string]func(string))
	actions[constant.DelPod] = action.DelPod
}
