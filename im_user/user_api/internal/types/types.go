// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type FriendInfoRequest struct {
	FriendName string `form:"friend_name"`
}

type FriendInfoResponse struct {
	FriendID uint   `json:"friendID"`
	Nickname string `json:"nickname"`
	Abstract string `json:"abstract"`
	Avatar   string `json:"avatar"`
	Notice   string `json:"notice"`
}

type UserInfoRequest struct {
	UserName string `json:"username"`
}

type UserInfoResponse struct {
	UserID               uint                  `json:"userID"`
	Nickname             string                `json:"nickname"`
	Abstract             string                `json:"abstract"`
	Avatar               string                `json:"avatar"`
	RecallMessage        *string               `json:"recallMessage"`
	FriendOnline         bool                  `json:"friendOnline"`
	Sound                bool                  `json:"sound"`
	SecureLink           bool                  `json:"secureLink"`
	SavePwd              bool                  `json:"savePwd"`
	SearchUser           int8                  `json:"searchUser"`
	Verification         int8                  `json:"verification"`
	VerificationQuestion *VerificationQuestion `json:"verificationQuestion"`
}

type UserInfoUpdateRequest struct {
	Nickname             *string               `json:"nickname,optional" user:"nickname"`
	Abstract             *string               `json:"abstract,optional" user:"abstract"`
	Avatar               *string               `json:"avatar,optional" user:"avatar"`
	RecallMessage        *string               `json:"recallMessage,optional" user_conf:"recall_message"`
	FriendOnline         *bool                 `json:"friendOnline,optional" user_conf:"friend_online"`
	Sound                *bool                 `json:"sound,optional" user_conf:"sound"`
	SecureLink           *bool                 `json:"secureLink,optional" user_conf:"secure_link"`
	SavePwd              *bool                 `json:"savePwd,optional" user_conf:"save_pwd"`
	SearchUser           *int8                 `json:"searchUser,optional" user_conf:"search_user"`
	Verification         *int8                 `json:"verification,optional" user_conf:"verification"`
	VerificationQuestion *VerificationQuestion `json:"verificationQuestion,optional" user_conf:"verification_question"`
}

type UserInfoUpdateResponse struct {
}

type VerificationQuestion struct {
	Problem1 *string `json:"problem1,optional" user_conf:"problem1"`
	Problem2 *string `json:"problem2,optional" user_conf:"problem2"`
	Problem3 *string `json:"problem3,optional" user_conf:"problem3"`
	Answer1  *string `json:"answer1,optional" user_conf:"answer1"`
	Answer2  *string `json:"answer2,optional" user_conf:"answer2"`
	Answer3  *string `json:"answer3,optional" user_conf:"answer3"`
}
