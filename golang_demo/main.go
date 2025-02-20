package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

// 接口请求域名
const mainURL = "https://api.salesmartly.com"

// 项目API Token
const apiToken = "DXPvaTUacQJi2ym"

// 项目Id
const projectID = "dnu1b5"

// 生成 MD5 签名
func generateMD5Signature(apiToken string, params map[string]string) string {
	// 将参数按字典序排序
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 拼接参数字符串
	var paramStr strings.Builder
	paramStr.WriteString(apiToken)
	for _, k := range keys {
		paramStr.WriteString("&")
		paramStr.WriteString(k)
		paramStr.WriteString("=")
		paramStr.WriteString(params[k])
	}

	// 拼接 API Token

	fmt.Println("paramStr:", paramStr.String())

	// 计算 MD5
	hash := md5.Sum([]byte(paramStr.String()))
	return hex.EncodeToString(hash[:])
}

// 发送 GET 请求
func sendGetRequest(apiURL string, params map[string]string, headers map[string]string) {
	// 构建完整的 URL
	u, err := url.Parse(mainURL + apiURL)
	if err != nil {
		fmt.Printf("URL 解析失败: %v\n", err)
		return
	}

	// 添加查询参数
	q := u.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()

	// 创建请求
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		fmt.Printf("创建 GET 请求失败: %v\n", err)
		return
	}

	// 添加请求头
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("GET 请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return
	}

	// 解析 JSON 响应
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Printf("解析 JSON 响应失败: %v\n", err)
		return
	}

	fmt.Println("GET 请求成功，响应数据：")
	fmt.Println(data)
}

// 发送 POST 请求
func sendPostRequest(apiURL string, payload map[string]string, headers map[string]string) {
	// 构建完整的 URL
	u, err := url.Parse(mainURL + apiURL)
	if err != nil {
		fmt.Printf("URL 解析失败: %v\n", err)
		return
	}

	// 编码表单数据
	formData := url.Values{}
	for k, v := range payload {
		formData.Add(k, v)
	}

	// 创建请求
	req, err := http.NewRequest("POST", u.String(), bytes.NewBufferString(formData.Encode()))
	if err != nil {
		fmt.Printf("创建 POST 请求失败: %v\n", err)
		return
	}

	// 添加请求头
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("POST 请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return
	}

	// 解析 JSON 响应
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Printf("解析 JSON 响应失败: %v\n", err)
		return
	}

	fmt.Println("POST 请求成功，响应数据：")
	fmt.Println(data)
}

func main() {
	// 定义请求的接口及参数
	getURL := "/api/chat-user/get-contact-list"
	getParams := map[string]string{
		"project_id":   projectID,
		"updated_time": `{"start":1680000000,"end":1814027206}`,
	}

	// 参数加密出签名
	getSign := generateMD5Signature(apiToken, getParams)
	getHeader := map[string]string{
		"external-sign": getSign,
	}

	// 发送 GET 请求
	sendGetRequest(getURL, getParams, getHeader)
}
