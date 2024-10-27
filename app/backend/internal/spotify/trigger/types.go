package trigger

type ActionBody struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type FollowerChange struct {
	Followers int  `json:"followers"`
	Increased bool `json:"increased"`
}
