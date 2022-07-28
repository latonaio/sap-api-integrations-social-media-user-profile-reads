package sap_api_output_formatter

type SocialMediaUserProfile struct {
	ConnectionKey              string `json:"connection_key"`
	Result                     bool   `json:"result"`
	RedisKey                   string `json:"redis_key"`
	Filepath                   string `json:"filepath"`
	APISchema                  string `json:"api_schema"`
	SocialMediaUserProfileCode string `json:"social_media_user_profile_code"`
	Deleted                    bool   `json:"deleted"`
}

type SocialMediaUserProfileCollection struct {
	ObjectID                        string `json:"ObjectID"`
	ETag                            string `json:"ETag"`
	SocialMediaUserCategoryCode     string `json:"SocialMediaUserCategoryCode"`
	SocialMediaUserCategoryCodeText string `json:"SocialMediaUserCategoryCodeText"`
	ID                              string `json:"ID"`
	AccountUUID                     string `json:"AccountUUID"`
	EntityLastChangedOn             string `json:"EntityLastChangedOn"`
	AccountInternalID               string `json:"AccountInternalID"`
}
