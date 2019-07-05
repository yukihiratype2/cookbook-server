package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// User struct
type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	Email      string             `bson:"email"`
	Password   string             `bson:"password"`
	Role       []string           `bson:"role"`
	UserInfo   UserInfo           `bson:"user_info"`
	UserStatus UserStatus         `bson:"user_status"`
}

// UserInfo struct
type UserInfo struct {
	UserName    string    `bson:"user_name"`
	CreatedAt   time.Time `bson:"created_at"`
	LastLogin   time.Time `bson:"last_login"`
	LastLoginAt string    `bson:"last_login_at"`
}

// UserStatus struct
type UserStatus struct {
	Suspend         bool      `bson:"suspend,omitempty"`
	SuspendReason   string    `bson:"suspend_reason,omitempty"`
	Premium         bool      `bson:"premium,omitempty"`
	PremiumExpireAt time.Time `bson:"premium_expire_at,omitempty"`
}
