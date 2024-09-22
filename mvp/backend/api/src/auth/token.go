package auth

import "time"

type TockenModel struct {
	AccessToken  string        `bson:"access_token"`
	TokenType    string        `bson:"token_type,omitempty"`
	RefreshToken string        `bson:"refresh_token,omitempty"`
	Expiry       time.Time     `bson:"expiry,omitempty"`
	ExpiresIn    int64         `bson:"expires_in,omitempty"`
	ExpiryDelta  time.Duration `bson:"expiry_delta,omitempty"`
	Raw          interface{}   `bson:"raw,omitempty"`
}
