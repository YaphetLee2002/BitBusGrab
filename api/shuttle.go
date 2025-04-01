package api

import (
	"awesomeProject/config"
	"awesomeProject/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func makeRequest(url string, headers map[string]string) ([]byte, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	return body, nil
}

func getDefaultHeaders() map[string]string {
	return map[string]string{
		"Host":     config.ApiHost,
		"Accept":   "application/json",
		"apitoken": config.ApiToken,
		"apitime":  config.ApiTime,
	}
}

// GetShuttleList 获取班车列表
func GetShuttleList(date string, address string) (*models.ShuttleList, error) {
	url := fmt.Sprintf("http://%s/vehicle/get-list?page=1&limit=20&date=%s&address=%s",
		config.ApiHost, date, address)

	body, err := makeRequest(url, getDefaultHeaders())
	if err != nil {
		return nil, fmt.Errorf("请求班车列表失败: %w", err)
	}

	var shuttleList models.ShuttleList
	err = json.Unmarshal(body, &shuttleList)
	if err != nil {
		return nil, fmt.Errorf("解析班车列表失败: %w", err)
	}

	return &shuttleList, nil
}

// GetShuttleInfo 获取班车详情
func GetShuttleInfo(id string, userid string) (*models.ShuttleInfo, error) {
	url := fmt.Sprintf("http://%s/vehicle/get-info?id=%s&userid=%s",
		config.ApiHost, id, userid)

	body, err := makeRequest(url, getDefaultHeaders())
	if err != nil {
		return nil, fmt.Errorf("请求班车详情失败: %w", err)
	}

	var shuttleInfo models.ShuttleInfo
	err = json.Unmarshal(body, &shuttleInfo)
	if err != nil {
		return nil, fmt.Errorf("解析班车详情失败: %w", err)
	}

	return &shuttleInfo, nil
}

// GetReservedSeats 获取座位预定状态
func GetReservedSeats(id string, date string, userid string) (*models.ReservedSeats, error) {
	url := fmt.Sprintf("http://%s/vehicle/get-reserved-seats?id=%s&date=%s&userid=%s",
		config.ApiHost, id, date, userid)

	body, err := makeRequest(url, getDefaultHeaders())
	if err != nil {
		return nil, fmt.Errorf("请求座位预定状态失败: %w", err)
	}

	var reservedSeats models.ReservedSeats
	err = json.Unmarshal(body, &reservedSeats)
	if err != nil {
		return nil, fmt.Errorf("解析座位预定状态失败: %w", err)
	}

	return &reservedSeats, nil
}

// CreateOrder 下单
func CreateOrder(id string, date string, userid string, seatNumber int64) (*string, error) {
	data := fmt.Sprintf("id=%s&date=%s&seat_number=%d&userid=%s", id, date, seatNumber, userid)
	url := fmt.Sprintf("http://%s/vehicle/create-order", config.ApiHost)

	headers := getDefaultHeaders()
	headers["Content-Type"] = "application/x-www-form-urlencoded"

	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}
	var resp models.Response
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}
	if resp.Message != "ok" {
		return nil, fmt.Errorf(resp.Message)
	}
	return &resp.Message, nil
}
