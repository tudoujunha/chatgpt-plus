package service

import (
	"chatplus/core/types"
	"chatplus/store"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type AliYunSmsService struct {
	config *types.AppConfig
	db     *store.LevelDB
	client *dysmsapi.Client
}

func NewAliYunSmsService(config *types.AppConfig, db *store.LevelDB) (*AliYunSmsService, error) {
	// 创建阿里云短信客户端
	client, err := dysmsapi.NewClientWithAccessKey(
		"cn-hangzhou",
		config.SmsConfig.AccessKey,
		config.SmsConfig.AccessSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	return &AliYunSmsService{
		config: config,
		db:     db,
		client: client,
	}, nil
}

func (s *AliYunSmsService) SendVerifyCode(mobile string, code int) error {
	// 创建短信请求并设置参数
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.Domain = s.config.SmsConfig.Domain
	request.PhoneNumbers = mobile
	request.SignName = "飞行的蜗牛"
	request.TemplateCode = "SMS_281460317"
	request.TemplateParam = fmt.Sprintf("{\"code\":\"%d\"}", code) // 短信模板中的参数

	// 发送短信
	response, err := s.client.SendSms(request)
	if err != nil {
		return fmt.Errorf("failed to send SMS:%v", err)
	}

	if response.Code != "OK" {
		return fmt.Errorf("failed to send SMS:%v", response.Message)
	}

	return nil
}
