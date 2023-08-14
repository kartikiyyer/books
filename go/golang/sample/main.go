package main

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	xj "github.com/basgys/goxml2json"
	yaml "gopkg.in/yaml.v2"
)

func parseHtmlTemplate() {
	fmt.Println(os.Getwd())
	_, err := template.ParseFiles("index.html")

	fmt.Println(err)
}

func postinfluxdbapi() {
	phoneCount := 1
	websiteCount := 1
	i := 1
	for {
		phoneCount += i + 10
		websiteCount += i + 12
		data := fmt.Sprintf("orders phone=%d,website=%d", phoneCount, websiteCount)
		resp, _ := http.Post("http://localhost:8086/write?db=food_data", "text/plain", bytes.NewBuffer([]byte(data)))
		fmt.Println(resp.StatusCode)
		time.Sleep(time.Minute / 2)
		i++
	}
}

func duration() {
	duration := 10 * time.Minute
	fmt.Println(int64(duration / time.Second))

	fmt.Println(int(3 * 30 * 24 * time.Hour / time.Second))
}

func formatMultilineString() {
	downsampleFluxQuery := fmt.Sprintf(`
from(bucket:"%s")
|> range(start: -%ds)
|> filter(fn:(r) => r._measurement == "session" and r._field == "forward_sampled_bytes" or r._field == "forward_sampled_bytes" or r._field == "forward_sampled_pkts" or r._field == "reverse_sampled_bytes" or r._field == "reverse_sampled_pkts")
|> keep(columns:["_time", "_field", "_value", "local_ip", "remote_ip", "forward_sampled_bytes", "forward_sampled_bytes", "forward_sampled_pkts", "reverse_sampled_bytes", "reverse_sampled_pkts"])
|> pivot(rowKey:["_time"],columnKey: ["_field"],valueColumn: "_value")
|> reduce( identity:{forward_sampled_bytes_sum:0,forward_sampled_pkts_sum:0,reverse_sampled_bytes_sum:0,reverse_sampled_pkts_sum:0}, fn: (r, accumulator) => ({forward_sampled_bytes_sum: accumulator.forward_sampled_bytes_sum + r.forward_sampled_bytes,forward_sampled_pkts_sum: accumulator.forward_sampled_pkts_sum + r.forward_sampled_pkts,reverse_sampled_bytes_sum: accumulator.reverse_sampled_bytes_sum + r.reverse_sampled_bytes,reverse_sampled_pkts_sum: accumulator.reverse_sampled_pkts_sum + r.reverse_sampled_pkts}))
|> rename(columns: {forward_sampled_bytes_sum:"forward_sampled_bytes",forward_sampled_pkts_sum:"forward_sampled_pkts",reverse_sampled_bytes_sum:"reverse_sampled_bytes",reverse_sampled_pkts_sum:"reverse_sampled_pkts"})
|> to(bucket: "CN2_downsample", org: "contrail")
`, "CN2", 600)

	fmt.Println(downsampleFluxQuery)
}

func httpMethods() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	// r, err := client.Get("https://10.84.3.27:6443/apis/introspect.k8s.io/v1beta1/")
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://10.84.3.27:6443/apis/introspect.k8s.io/v1beta1/10.84.3.27:8085/"), nil)
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6ImtSenZucll0VHJLVjhGSHhfejZxelRkV1ByV29IMVVnVEg4U0ZVTnlnclEifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJjb250cmFpbCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJpbnRyb3NwZWN0LXRva2VuLXp4ZGRzIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6ImludHJvc3BlY3QiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiJlMmY0N2NlZS1hYWE4LTQ5ZTMtYTY4NC02NzE4NjU0ZDUzNzEiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6Y29udHJhaWw6aW50cm9zcGVjdCJ9.hsrsq7E9HsC6_AroCYpIfT9NH679MyBx5_xSxNz2jRDsS0LZqkib-8hkoO2NYXT7eqJQEKaGdxffmpM-_P8vcTWOUNFOFVR1wCsP5yj8G5nz2AGmuS6AXnxn-t_d19Twmv0H3YcCyIcMUxr4B5F4scvtjPrKWsRvTPqktJxHYw6n7R6AYwEWVOx6FzuIhNtl7phakv1o_kkOqlZLnVIctQHcOsl4S3G6VUe3buaGfmUOXt8fRdz7QN7wFx4n5PmTpyC2_RBz8Np7yRo6mydpZYGnB-5Ai7iQZCc1MRIJUVq-nakzJn6xGfs6VWY_wW7mZlb34KFkCGkvjCMgJ0G09A")
	r, err := client.Do(req)
	fmt.Println(r.StatusCode)
	fmt.Println(r.Header.Get("content-type"))
	data, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	fmt.Println(string(data))
	fmt.Println(err)
}

func esSearch() {
	// distCluster1Host := "test1-distributed-cluster1.englab.juniper.net"
	// distCluster2Host := "test1-distributed-cluster2.englab.juniper.net"
	// masterClusterHost := "k8s-master.englab.juniper.net"
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	body := `
	{
		"query": {
			"bool": {
				"must": [
					{"match": { "kubernetes.host": "k8s-master.englab.juniper.net" }},
					{"match": { "kubernetes.container_name": "contrail-vrouter-agent" }}
				]
			}
		}
	}
	`
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://%s/elasticsearch/contrail.kubernetes-*/_search/", "10.84.2.234"), bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	r, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	if r.StatusCode != 200 {
		fmt.Println("Expected status code: %d but found: %d", 200, r.StatusCode)
	}
	respBody, _ := ioutil.ReadAll(r.Body)
	// respBodyStr := string(respBody)

	type EsOutputTotal struct {
		Value int `json:"value"`
	}

	type EsOutputHits struct {
		Total EsOutputTotal `json:"total"`
	}

	type EsOutput struct {
		Hits EsOutputHits `json:"hits"`
	}

	var esOutput EsOutput
	err = json.Unmarshal(respBody, &esOutput)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(esOutput)
	if esOutput.Hits.Total.Value == 0 {
		fmt.Println("No value are present")
	}
}

func readWriteFile() {
	kubeconfig := []byte("YXBpVmVyc2lvbjogdjEKY2x1c3RlcnM6Ci0gY2x1c3RlcjoKICAgIGNlcnRpZmljYXRlLWF1dGhvcml0eS1kYXRhOiBMUzB0TFMxQ1JVZEpUaUJEUlZKVVNVWkpRMEZVUlMwdExTMHRDazFKU1VNMWVrTkRRV01yWjBGM1NVSkJaMGxDUVVSQlRrSm5hM0ZvYTJsSE9YY3dRa0ZSYzBaQlJFRldUVkpOZDBWUldVUldVVkZFUlhkd2NtUlhTbXdLWTIwMWJHUkhWbnBOUWpSWVJGUkplRTFFWTNkUFZFVXpUbFJuZDAxR2IxaEVWRTE0VFVSamQwNTZSVE5PVkdkM1RVWnZkMFpVUlZSTlFrVkhRVEZWUlFwQmVFMUxZVE5XYVZwWVNuVmFXRkpzWTNwRFEwRlRTWGRFVVZsS1MyOWFTV2gyWTA1QlVVVkNRbEZCUkdkblJWQkJSRU5EUVZGdlEyZG5SVUpCVDNaT0NuWk1hVkJQT1hCdGFtZzRPV1ZSWm1FMFlsRm5jWFZoVEU1S2RtWnNXV2xhTDNKVFZ5dGhNWHBvZVZkV1dGQTRMMVpsVGtKdVNHWk9URFIySzFwblVHb0thWHBHSzIxSFVtOXVNWFJ6WWsxaGRVRk5lblJQWmk4MU1VbFRSWEJJTmxjNVdXNUxRbEl3ZDBsUFJHaERUeXN6WW5veGNVSmFRMEpuT1hkdVdtRjNid3BGZERWWFVXbE1ObVUwVldaUVMwaFFZMlE0ZDBJM1QxbGhRbk51TVRVeWFFRk9VME5tZGpSUFNYUjNNMWQxYWxoSVowRnFXSEJwU2tKV01FTmFRMGd6Q2psbWFGTTNjbkY0ZVhoWFJuRTFTMEpLT0RkT2NGUm9TbVJUVGtKYWIycHliakJpWjBkMmJrOUxVV1ZqVVhJeWMxbERXVWxzWW1kb01UVmxVRE01ZVdNS2QzcHdVRkpJVGpONmFsbFpRbkk0YkZFNVpFZ3ZWVXR2ZVU1SFQzZDVVbXMxYXpCSlRucFpRV3gyWVhBM1NGRnhWMUpoYkV3M1MwUjNPVkZIWmxkYWN3cFFiWFZuTUUxV1ZHaFZMMFZxTUhkTVZrMXJRMEYzUlVGQllVNURUVVZCZDBSbldVUldVakJRUVZGSUwwSkJVVVJCWjB0clRVRTRSMEV4VldSRmQwVkNDaTkzVVVaTlFVMUNRV1k0ZDBoUldVUldVakJQUWtKWlJVWkxTMXBpVVRSMFVHczVTRkJMTjJFMmFIVkJjbFpIYmtvMU5VaE5RVEJIUTFOeFIxTkpZak1LUkZGRlFrTjNWVUZCTkVsQ1FWRkJMekp1Y21Od01XeEhSbnBrVm1ONGJXbDVLMjh4ZWxSeFRua3JRa2hvTlhoUldEbFZjV0ZMY2tSUFowRTFaR0Y0VUFvclNqSTBkazUxUlhsd05qQmplRTUzVWxaSVJXWXdMeTlwV21SUVNqWmFkVE5qVURaa2FIQjBVRkJLTDJKSk5UUTVjRnBUY3pnek5qVmllSEJsT0U1MkNqbFViVXhRWlZoM2JDdGtObGh1TlRVeVNscDNhMEZwWkd4dlJuQXpkVXczUVZWQlkxaGtPR0U0WjBkNU5qVmxhV0YxU25OYVUxcEVVblJOUTBwU1Vub0tVRTVTV1c5S1JYbzFlbWhYVldoT01FY3lTaTlxUW1KUllrWlZOMnhxWkdVemJHeElVVzlGWXpaSGFYWXJSRU5NVG5sRlVqWmFZelpSWnpsdmRtZG1jQXBRV1VsaFRrZzJWbmQzYWpVemNURTJOVFZ0SzBKMFNYUk5hVTFqV25GU2VEWm9lVTlFWlZaV2ExZDFlV1VyZUVKclNVWkJia2x5TlZKamFXSlBhMEZYQ2xKcmVGWjROSFZoVkhVd1ltSkRTRTE0V1U5NUsweGFjRkpQZDAxSE1qYzVaakpSWXdvdExTMHRMVVZPUkNCRFJWSlVTVVpKUTBGVVJTMHRMUzB0Q2c9PQogICAgc2VydmVyOiBodHRwczovLzEwLjg0LjIuMTE3OjY0NDMKICBuYW1lOiBrdWJlcm5ldGVzCmNvbnRleHRzOgotIGNvbnRleHQ6CiAgICBjbHVzdGVyOiBrdWJlcm5ldGVzCiAgICB1c2VyOiBrdWJlcm5ldGVzLWFkbWluCiAgbmFtZToga3ViZXJuZXRlcy1hZG1pbkBrdWJlcm5ldGVzCmN1cnJlbnQtY29udGV4dDoga3ViZXJuZXRlcy1hZG1pbkBrdWJlcm5ldGVzCmtpbmQ6IENvbmZpZwpwcmVmZXJlbmNlczoge30KdXNlcnM6Ci0gbmFtZToga3ViZXJuZXRlcy1hZG1pbgogIHVzZXI6CiAgICBjbGllbnQtY2VydGlmaWNhdGUtZGF0YTogTFMwdExTMUNSVWRKVGlCRFJWSlVTVVpKUTBGVVJTMHRMUzB0Q2sxSlNVUkZla05EUVdaMVowRjNTVUpCWjBsSlZYSnlLMnB2YVZoVVVHZDNSRkZaU2t0dldrbG9kbU5PUVZGRlRFSlJRWGRHVkVWVVRVSkZSMEV4VlVVS1FYaE5TMkV6Vm1sYVdFcDFXbGhTYkdONlFXVkdkekI1VFZSQk0wMUVhM2hPZWxVMFRVUkNZVVozTUhsTmFrRXpUVVJyZUU1NlZUUk5SRTVoVFVSUmVBcEdla0ZXUW1kT1ZrSkJiMVJFYms0MVl6TlNiR0pVY0hSWldFNHdXbGhLZWsxU2EzZEdkMWxFVmxGUlJFVjRRbkprVjBwc1kyMDFiR1JIVm5wTVYwWnJDbUpYYkhWTlNVbENTV3BCVGtKbmEzRm9hMmxIT1hjd1FrRlJSVVpCUVU5RFFWRTRRVTFKU1VKRFowdERRVkZGUVhJdlFqbFdlVWQyY0Rsa2NGY3ZTVWdLZDI1R1IzVmpjVlpKY0V4UWNUSmlhVFZCYUc4dllqRXplaXMzVWtaTlF6WnJWazVFYTJ4Qk1XVkZjMlZRUkROMGIxRmFkVTl3YVdaQmEwVlZWalZUYVFwd1FVNHpkMnhVUkhkWlZITlJZMEY2YUd3elMya3laMUpRVGtrNGFVOXJkbmhOY1VoSlprRXlWbGxpTUM5WVZYZDNNSHB3U25ndlpFOVhjVXBUZWtSdUNuWTBVVlk1ZFZGckswbEVPVTlvUVVGa01UbGhhbkpNTHpSc2FuSldNR0U0VVZKQ2VVVTFiV2RRYUVveUwwWnpOR3RFVDA5aVNGbFJNblJhYmxoMFJqVUtiRzV4UkZGM1ExVjRVRmRTU0hWc2JuWlpSR2x4VWpjck9DdGxaR0V6YWpWelZucFhWazltZFhFd1pVYzBTQzg0V0U1M1FteE9PVkZEU3pReVJUQmxUUXB3UXk4NFlsbEVaa3cwYkdWQ2JUYzFiWFEyY2xKMlZVRm5PR1V6U0ZaYWJHbFFkWFYxTWxSVWJWWnFjRXB0TkZCbVlrMVBla2s0TWpOaFVWRTVVMXBrQ2xKTk0zcFZkMGxFUVZGQlFtOHdaM2RTYWtGUFFtZE9Wa2hST0VKQlpqaEZRa0ZOUTBKaFFYZEZkMWxFVmxJd2JFSkJkM2REWjFsSlMzZFpRa0pSVlVnS1FYZEpkMGgzV1VSV1VqQnFRa0puZDBadlFWVnZjR3gwUkdrd0sxUXdZemh5ZEhKeFJ6UkRkRlZoWTI1dWEyTjNSRkZaU2t0dldrbG9kbU5PUVZGRlRBcENVVUZFWjJkRlFrRkVhVTlLVUdoVlUwRkRORFpyZHpBd2NuQnVWbEJDVmxGMFoxVkZhREZxYVVsbVNXNHdkR2RtWldaeFJWTlFaMlZPTmpOTGN5OTJDa2t2VWtSaVJGSktZaXRWY0dwWWFFTlZXRFZaYWpWRk4xTnlWRUV4U0ZKVFQyNDVTazByYkhkUVFYZERTbUowY0dsaWVVbFRla1JzVkVobGNVWnJPRGtLZDFKQloyMDFSMVozWkdWM09WQmxSRUY0WmpOV2RIbHhkVGd5ZVU1ek56WXdPR1kzWVRVNWEzRlVVV2gxY0V4WFUyUTNRVmhRVUN0RFUzbGxaV2xoWmdwc1owWkVNRXgzV1dObFJVRXdRMEowWVdORk9XUnBjbXROWWxsR1JVUkhNblpRVWl0dGRWbEpaR2x6WjNaQ1l6UjVhMFl2ZFZJdmFrSTJiakp2YjB4VkNqQlhMMWxvYzFOUk5UbFdTRGhsYmpkVk1EWTJjaTgwTjBaa1ZFRlRaM052TUhFMVp5OUpORXR4YWxWRWFFMDNkazltV1hGRFYwWlZSRFp5TkRKa2RuRUtObnBxY25oWFdEZHBOVUZ3V1VOV1kxWXlPWGh3V1dNeUwwVnJSbk15UVQwS0xTMHRMUzFGVGtRZ1EwVlNWRWxHU1VOQlZFVXRMUzB0TFFvPQogICAgY2xpZW50LWtleS1kYXRhOiBMUzB0TFMxQ1JVZEpUaUJTVTBFZ1VGSkpWa0ZVUlNCTFJWa3RMUzB0TFFwTlNVbEZiM2RKUWtGQlMwTkJVVVZCY2k5Q09WWjVSM1p3T1dSd1Z5OUpTSGR1UmtkMVkzRldTWEJNVUhFeVltazFRV2h2TDJJeE0zb3JOMUpHVFVNMkNtdFdUa1JyYkVFeFpVVnpaVkJFTTNSdlVWcDFUM0JwWmtGclJWVldOVk5wY0VGT00zZHNWRVIzV1ZSelVXTkJlbWhzTTB0cE1tZFNVRTVKT0dsUGEzWUtlRTF4U0VsbVFUSldXV0l3TDFoVmQzY3dlbkJLZUM5a1QxZHhTbE42Ukc1Mk5GRldPWFZSYXl0SlJEbFBhRUZCWkRFNVlXcHlUQzgwYkdweVZqQmhPQXBSVWtKNVJUVnRaMUJvU2pJdlJuTTBhMFJQVDJKSVdWRXlkRnB1V0hSR05XeHVjVVJSZDBOVmVGQlhVa2gxYkc1MldVUnBjVkkzS3pnclpXUmhNMm8xQ25OV2VsZFdUMloxY1RCbFJ6UklMemhZVG5kQ2JFNDVVVU5MTkRKRk1HVk5jRU12T0dKWlJHWk1OR3hsUW0wM05XMTBObkpTZGxWQlp6aGxNMGhXV213S2FWQjFkWFV5VkZSdFZtcHdTbTAwVUdaaVRVOTZTVGd5TTJGUlVUbFRXbVJTVFRONlZYZEpSRUZSUVVKQmIwbENRVVphT1U5YVExRjJZMGxpYzJVMlF3cE1NMEpXYzFSbU5rdzFRV05pWW5oTWVYSlFaM052VVhsR1QydENPVWxTYjB4bk4zTjVTMWgyYVdsdFJtVXJjVkJaZG1GeGNUVjBaMnhoU2s1RVkzWXdDa0pJVUhWU1NIZHBiMVY2VEhCRWVUWlhSbTUzVmpOYWNXaFVZa292ZG05eE56bEhhelJuVFRoREsyNUdaWG92ZURCVVVHYzFhV3BYYTBSRUsyWXhiVmtLVWtGVGR6QlFPV0Z5YUZndmFYWjRNVU14VjJwaFYyZEhPVXRXTWpCVFJXdFJlbU5hUjNCVlNXSjBjakpYTml0SFNGcFRPRzkzY0hreWRWWm1TbE5qTlFvclFXMUZWWFJ5WmswNE56UlZVSEpOUVhGRE5rNDFlaXQ0WVV0eVVEaFdTVlZyVlhSaUwzQlhaSHBKYVZCNmJuaDVNVXhLSzJwRldGZFZWa2swWjJkaUNubHNaRWhoUjFWRVRUZFhVbFZZZURselYyUlNXVGgxU2xwcGJtaFBPVFJsTTNRNVdIUllhbGhQWnpCc1FXUTFPREphZFhOck5WRlNOblpTTUZwbmRVRUtabGRCYW05dVJVTm5XVVZCTm0xNVIxTmtVWE5CZW5OWkswMWFaMmxwV0UxT1VGTjFjMEZaUlZaSVIxVkJNMGx1VmxKWFpYRnZVWHBMVm5ZeGJESmtOUXBLVVVZeU5FMXFPVU5RT0ZWeEswdEhWemxKYm1WMWFHSjNTRFJUV1VoU1FYcFhaa1JMU0dZeVZVcE9OelZJVlRCTGIyTmtjams1VDJGUE1saEhURUpPQ2tKTlVtWlNUMFExWlc5clRYZEdXVWMwZFhWYVlrOXllV1p1YW5kWWRFbERNR3htYm10bFlVaGpaVzU2VDFoR09EUkJNMmRhYzNORFoxbEZRWGREU0RFS1ZEQlVWbGR3VEROVGJVa3dMemhhVUdRNWJ6SlVjVzFIZEdkT1JGQnZNRXRST0VKdGNEUmhTa1lyTVVwTVQzVm5PRmhMYlZwdGNVRXhVbHA2TnpZemJBcG1SREYyYWpadVpsSjBRVkJUY25kMFRuaE1aVm8xU1VJdlJYaFZSbXhxTUUxUVJrVlFRM0pRWTBrellTdDJhR3RGWVZJMFpHczJZMmcyYkhGSlZrYzFDak51YUdSMmMwMHhPVmhRVW1WblZUUTBha1V5VWtGdU5XZFpRMWx6U1hkbE1VcFBPVVJLYTBObldVVkJOR0ZTTkZab1VXVkZWRmxRVVZaYWFHdHNabm9LYVZneVJISXZZalp5TDFaeVZGTXpValJxYjJaTlZ6VjVaVWh6Um0xRGFVaEhWRWRvY0ZwcFVYaDBiMUZ2WVc4elVHRmhOVlJpY21Wd2FHZDZlRUptYVFwVlJsSjVTM292WkVoWk16UnJUbmx1ZVRncmJYRmhVeXRMYnpKWVFreExaVlE1Y0dSQlZWZDZibFpxVW5Bd2FXMDVSVVp0ZFhOdmFtYzRNRkpSV2xCUkNrWXpkRFZqU25nckwzZDZOMmwxTVUxRU4xbE1jVFpWUTJkWlFYY3pOMHRIU1c4eFpsSXJiRTQyTms1MGRVOVdORXRJVFU5SFdVOXBRWHA0TUhsbE0yd0thRU5xV1hwd2JFeEZPVzg0WVdkbU4zY3hNV1Y1YUV4U2IyRXJSMDVoWTJSeWIyTldaRmhJVG10d2JVSlBiRTFMU1cxMVpYcDJOV2RVYnpWS1ZGRkdOQW92WW1GTk1EVnpTRVkxTm5adWMxaGtNSGRsZEcxS2RYRjVjMkZ5YVZJeVRqUlFTaTh4ZVhBdmMwdEhSblExVW1KMldXSXZWV05aT1ZWb1owWlBUeXRUQ25oaE5HVmhVVXRDWjBoQ2VGaDFjelZSU1VOT1dUWjRjSFZxV1UxdFdGaFlSelZHZVhCVWVHVjZhSGw1Wm1Ob1VVWnVkMVJITkZKNFUxRk1ObUpOWTFJS2VDdExORXRVV1dSSVUxaDRNazlPUld0MVRVSmlkMHR5WjFOYWNFZFFSMWM1TlVweGFFbHpRVkZQVkZWaVpuQlhiMDkxUlZKbVppOWlLMUpSZEN0V2F3cFZjbkZITjJZd2FVbERUVzVVUzNONVJpOXJNR1J3U1U1Nk9EVTNXWEUyZW1wNEsyVk1UVE4zZHpOSVpEbFpPRUoyVVZVM0NpMHRMUzB0UlU1RUlGSlRRU0JRVWtsV1FWUkZJRXRGV1MwdExTMHRDZz09Cg==")

	fmt.Println(string(kubeconfig))
	kubeconfig, _ = base64.StdEncoding.DecodeString(string(kubeconfig))
	fmt.Println(string(kubeconfig))
	tmpfile, err := ioutil.TempFile(".", "kubeconfig")
	if err != nil {
		fmt.Printf("Unable to create kubeconfig file: %v", err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(kubeconfig); err != nil {
		fmt.Printf("Unable to write kubeconfig file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		fmt.Printf("Unable to close kubeconfig file: %v", err)
	}
	fmt.Printf("Kubeconfig file path: %s", tmpfile.Name())
}

func xmlTojson() {
	// xml is an io.Reader
	xml := strings.NewReader(`<?xml version="1.0" encoding="UTF-8"?><hello>world</hello>`)
	json, err := xj.Convert(xml)
	if err != nil {
		panic("That's embarrassing...")
	}

	fmt.Println(json.String())
	// {"hello": "world"}
}

// Get certificates from kube config file
func getCertificate() (string, error) {
	kubeconfig := "/Users/kiyyer/Documents/Projects/CN2/kubeconfig/config_117"
	var certificate string
	yamlFileContent, err := ioutil.ReadFile(kubeconfig)
	if err != nil {
		return certificate, err
	}

	type Certificate struct {
		ClientCertificateData string `yaml:"client-certificate-data"`
		ClientKeyData         string `yaml:"client-key-data"`
	}

	type User struct {
		Certificate Certificate `yaml:"user"`
	}
	type KubeConfig struct {
		Users []User `yaml:"users"`
	}
	var kc KubeConfig

	err = yaml.Unmarshal(yamlFileContent, &kc)
	if err != nil {
		return certificate, err
	}

	if len(kc.Users) > 0 {
		clientCertificateBytes, _ := base64.URLEncoding.DecodeString(kc.Users[0].Certificate.ClientCertificateData)
		clientCertificate := string(clientCertificateBytes)
		// fmt.Println(clientCertificate)
		clientKeyBytes, _ := base64.URLEncoding.DecodeString(kc.Users[0].Certificate.ClientKeyData)
		clientKey := string(clientKeyBytes)
		// fmt.Println(clientKey)
		clientCertificateAndKey := fmt.Sprintf("%s%s", url.PathEscape(clientCertificate), url.PathEscape(clientKey))
		certificate = fmt.Sprintf("Cert=\"%s\"", clientCertificateAndKey)
	}

	return certificate, nil
}

func populateYaml() {
	kubeconfigContents, _ := os.ReadFile("kubeconfig.yaml")
	t, _ := template.ParseFiles("userbinding.yaml")

	values := map[string]string{
		"Clusteruuid": "1234",
		"Kubeconfig":  string(kubeconfigContents),
	}

	var s strings.Builder
	t.Execute(&s, values)

	fmt.Println(s.String())

}

func main() {
	fmt.Println("hi")
	// parseHtmlTemplate()
	//postinfluxdbapi()
	//duration()
	//formatMultilineString()
	// httpMethods()
	// readWriteFile()
	// esSearch()
	//xmlTojson()
	// fmt.Println(getCertificate())
	populateYaml()
}
