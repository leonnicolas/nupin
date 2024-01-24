package config

import "github.com/redis/go-redis/v9"

type Config struct {
	Oidc struct {
		IssuerURL    string `yaml:"issuerUrl"`
		ClientID     string `yaml:"clientId"`
		ClientSecret string `yaml:"clientSecret"`
		RedirectUrl  string `yaml:"redirectUrl"`
		CookieDomain string `yaml:"cookieDomain"`
	} `yaml:"oidc"`
	Web struct {
		Address string `yaml:"address"`
	} `yaml:"web"`
	Nuki struct {
		APIKey          string `yaml:"apiKey"`
		SmartLockDevice int64  `yaml:"smartLockDevice"`
	} `yaml:"nuki"`
	StorageType string `yaml:"storageType"`
	Redis       struct {
		UseSentinel     bool                   `yaml:"sentinel"`
		FailoverOptions *redis.FailoverOptions `yaml:"failoverOptions,omitempty"`
		Options         *redis.Options         `yaml:"options,omitempty"`
	} `yaml:"redis"`
}
