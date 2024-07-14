package config

import "time"

// Don't change this one to const
var Linkvertises = [3]string{"https://link-hub.net/611514/key-test", "https://link-hub.net/611514/key-test-2", "https://link-hub.net/611514/key-test-3"}

const (
	Port        = 8080
	ProUrl      = "https://buy.stripe.com/6oE2ba4rV5zk0jCfYY"
	DownloadUrl = "https://bit.ly/autoblox-download"
	DiscordUrl  = "https://discord.gg/ne3pYRpnSN"

	KeyDuration    = time.Hour * 12
	ProKeyDuration = time.Hour * 744
)
