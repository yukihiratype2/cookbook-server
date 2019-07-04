package model

import (
	"time"
)

// User struct
type User struct {
	ID         string     `bson:"_id"`
	Email      string     `bson:"email"`
	Password   string     `bson:"password,omitempty"`
	UserInfo   UserInfo   `bson:"user_info,omitempty"`
	UserStatus UserStatus `bson:"user_status,omitempty"`
}

// UserInfo struct
type UserInfo struct {
	UserName    string    `bson:"user_name,omitempty"`
	CreatedAt   time.Time `bson:"created_at,omitempty"`
	LastLogin   time.Time `bson:"last_login,omitempty"`
	LastLoginAt string    `bson:"last_login_at,omitempty"`
}

// UserStatus struct
type UserStatus struct {
	Suspend         bool      `bson:"suspend,omitempty"`
	SuspendReason   string    `bson:"suspend_reason,omitempty"`
	Premium         bool      `bson:"premium,omitempty"`
	PremiumExpireAt time.Time `bson:"premium_expire_at,omitempty"`
}
