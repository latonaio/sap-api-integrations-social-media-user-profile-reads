package responses

type SocialMediaUserProfileCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                        string `json:"ObjectID"`
			ETag                            string `json:"ETag"`
			SocialMediaUserCategoryCode     string `json:"SocialMediaUserCategoryCode"`
			SocialMediaUserCategoryCodeText string `json:"SocialMediaUserCategoryCodeText"`
			ID                              string `json:"ID"`
			AccountUUID                     string `json:"AccountUUID"`
			EntityLastChangedOn             string `json:"EntityLastChangedOn"`
			AccountInternalID               string `json:"AccountInternalID"`
		} `json:"results"`
	} `json:"d"`
}
