package tencent

type ConvParams struct {
	Actions []*Action `json:"actions"`
}

type Action struct {
	OuterActionID string  `json:"outer_action_id" structs:"outer_action_id"`
	ActionTime    int64   `json:"action_time" structs:"action_time"`
	UserID        *UserID `json:"user_id" structs:"user_id"`
	ActionType    string  `json:"action_type" structs:"action_type"`
}

type UserID struct {
	HashIMEI      string `json:"hash_imei" structs:"hash_imei"`
	HashOAID      string `json:"hash_oaid" structs:"hash_oaid"`
	HashAndroidID string `json:"hash_android_id" structs:"hash_android_id"`
}
