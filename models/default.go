package models

var defaultSettings = []Setting{
	{Name: SiteURLKey, Value: "http://localhost:8080", Type: "basic"},
	{Name: SecretKey, Value: "secret_key", Type: "basic"},
}
