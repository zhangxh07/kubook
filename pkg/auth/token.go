package auth

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const getTokenUri = "/api/v1/login"

type Client struct {
	host string
	username string
	password string
}

type BtToken struct {
	Msg string       		`json:"msg"`
	Data struct{
		Token string 		`json:"token"`
		Bt_username string  `json:"bt_username"`
	}                		`json:"data"`
	Code int         		`json:"code"`
}

type GetTokenRequest struct {
	Domain int   `json:"domain"`
	Username string  `json:"username"`
	Password string  `json:"password"`
}

func New(host,username,password string) *Client {
	return &Client{
		host: host,
		username: username,
		password: password,
	}
}
// GetToken 获取认证token
func (c *Client)GetToken() string {
	url := c.host + getTokenUri
	request := GetTokenRequest{
		Domain: 1,
		Username: c.username,
		Password: c.password,
	}

	payload, _ := json.Marshal(&request)
	req,_ := http.NewRequest("POST",url, bytes.NewBuffer(payload))
	req.Header.Add("Content-Type","application/json; charset=utf-8")

	res,err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
		return ""
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body,_:= ioutil.ReadAll(res.Body)
	var cmdb_token BtToken

	if err := json.Unmarshal(body, &cmdb_token); err != nil {
		return cmdb_token.Data.Token
	}
	return cmdb_token.Data.Token
}

