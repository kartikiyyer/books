package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func main4() {
	fmt.Println("hi")
	base, _ := url.Parse("http://proxy/cn2-introspect/Snh_SandeshSendingParamsSet?system_logs_rate_limit=&disable_object_logs=&disable_all_logs=&disable_flows=")

	fmt.Println(base.RawQuery)

	type vrouter struct {
		Name string
		Url  string
	}
	vrouters := []vrouter{{"v1", "http://10.84.18.9:8085/"}, {"v2", "http://10.84.18.9:8085/"}}
	data, _ := json.Marshal(vrouters)
	// v1 := &vrouter{name: "v1", url: "http://10.84.18.9:8085/"}
	// data, _ := json.Marshal(v1)
	fmt.Println(string(data))

	type response1 struct {
		Page   int
		Fruits []string
	}
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

}
