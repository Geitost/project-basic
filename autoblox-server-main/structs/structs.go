package structs

import (
	"gorm.io/gorm"
)

// Database Models
type Key struct {
	gorm.Model
	ID         uint
	Ip         string `gorm:"primaryKey"`
	Pro        bool   `gorm:"default:false"`
	Checkpoint uint8
	Value      string `gorm:"uniqueIndex;not null"`
	ExpiresAt  int64
	CreatedAt  int64 `gorm:"autoCreateTime:milli"`
	UpdatedAt  int64 `gorm:"autoUpdateTime:milli"`
}

type CheckoutSession struct {
	gorm.Model
	ID        uint
	Value     string `gorm:"uniqueIndex;not null"`
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli"`
}

// Other Structs
type HCaptchaRes struct {
	Success     bool     `json:"success"`
	ChallengeTs string   `json:"challenge_ts"`
	Hostname    string   `json:"hostname"`
	Credit      bool     `json:"credit"`
	ErrorCode   []string `json:"error-codes"`
}

type HCaptchaReq struct {
	Token string `json:"token"`
}

type ChannelMessage struct {
	Data       map[string]string
	StatusCode int
	Message    string
}
