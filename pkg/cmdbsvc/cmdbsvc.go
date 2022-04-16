package cmdbsvc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kubook/pkg/auth"
	"log"
	"net/http"
	"strings"
)

const (
	getService = "/api/v1/dynamic/bt_service?page_size=200&page_index="
)

type BtInfoResponse struct {
	Code     int 	   `json:"code"`
	Mseeage  string    `json:"mseeage"`
	Data     []*BtData `json:"data"`
	Page_info struct{
		Pages int  `json:"pages"`
		Total int  `json:"total"`
	} `json:"page_info"`
}

type BtData struct {
	Workload_name  string  `json:"workload_name"`
	Workload_type  string  `json:"workload_type"`
	Service_giturl string  `json:"service_giturl"`
	Service_project string `json:"service_project"`
	Metadata struct{
		Status  string  `json:"status"`
	}  `json:"metadata"`
}

type CmdbClient struct {
	*auth.Client
}

type Config struct {
	CmdbHost string
	CmdbUsername string
	CmdbPassword string
}

func New(conf Config) *CmdbClient{
	return &CmdbClient{
		auth.New(conf.CmdbHost,conf.CmdbUsername,conf.CmdbPassword),
	}
}

func (c *CmdbClient) GetServices() []*BtData{
	token := c.Client.GetToken()
	var res_page bool = true
	var page int =1
	var btdata,data []*BtData
	for res_page {
		data,res_page = c.GetCmdb(token,page)
		btdata = append(btdata,data...)
		page++
	}
	return btdata
}

// GetCmdb 获取服务列表
func (c *CmdbClient) GetCmdb(s string,num int) ([]*BtData,bool){
	url := fmt.Sprintf("https://cmdb.btpoc.com%s%d",getService,num)
	payload := strings.NewReader("payOrderNo=")
	req , _ := http.NewRequest("GET",url,payload)
	req.Header.Add("Content-Type","application/json; charset=utf-8")
	req.Header.Add("Authorization",s)
	res,err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	defer res.Body.Close()
	body,_ := ioutil.ReadAll(res.Body)
	var resp BtInfoResponse
	if err := json.Unmarshal(body,&resp); err != nil{
		log.Printf("getservice json 反序列化错误: %v",err)
		return nil,false
	}
	if num*200 <= resp.Page_info.Total{
		return resp.Data,true
	}else {
		return resp.Data,false
	}

}