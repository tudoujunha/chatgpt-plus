package core

import (
	"bytes"
	"chatplus/core/types"
	logger2 "chatplus/logger"
	"chatplus/utils"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
)

var logger = logger2.GetLogger()

func NewDefaultConfig() *types.AppConfig {
	return &types.AppConfig{
		Listen:        "0.0.0.0:5678",
		ProxyURL:      "",
		Manager:       types.Manager{Username: "admin", Password: "admin123"},
		StaticDir:     "./static",
		StaticUrl:     "http://localhost/5678/static",
		Redis:         types.RedisConfig{Host: "localhost", Port: 6379, Password: ""},
		AesEncryptKey: utils.RandString(24),
		Session: types.Session{
			Driver:    types.SessionDriverCookie,
			SecretKey: utils.RandString(64),
			Name:      "CHAT_PLUS_SESSION",
			Domain:    "",
			Path:      "/",
			MaxAge:    86400,
			Secure:    true,
			HttpOnly:  false,
			SameSite:  http.SameSiteLaxMode,
		},
		ApiConfig: types.ChatPlusApiConfig{},
		ExtConfig: types.ChatPlusExtConfig{Token: utils.RandString(32)},
		OSS: types.OSSConfig{
			Active: "local",
			Local: types.LocalStorageConfig{
				BaseURL:  "http://localhost/5678/static/upload",
				BasePath: "./static/upload",
			},
		},
	}
}

func LoadConfig(configFile string) (*types.AppConfig, error) {
	var config *types.AppConfig
	_, err := os.Stat(configFile)
	if err != nil {
		logger.Info("creating new config file: ", configFile)
		config = NewDefaultConfig()
		config.Path = configFile
		// save config
		err := SaveConfig(config)
		if err != nil {
			return nil, err
		}

		return config, nil
	}
	_, err = toml.DecodeFile(configFile, &config)
	if err != nil {
		return nil, err
	}

	return config, err
}

func SaveConfig(config *types.AppConfig) error {
	buf := new(bytes.Buffer)
	encoder := toml.NewEncoder(buf)
	if err := encoder.Encode(&config); err != nil {
		return err
	}

	return os.WriteFile(config.Path, buf.Bytes(), 0644)
}
