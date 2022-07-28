package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-social-media-user-profile-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToSocialMediaUserProfileCollection(raw []byte, l *logger.Logger) ([]SocialMediaUserProfileCollection, error) {
	pm := &responses.SocialMediaUserProfileCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to SocialMediaUserProfileCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	socialMediaUserProfileCollection := make([]SocialMediaUserProfileCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		socialMediaUserProfileCollection = append(socialMediaUserProfileCollection, SocialMediaUserProfileCollection{
			ObjectID:                        data.ObjectID,
			ETag:                            data.ETag,
			SocialMediaUserCategoryCode:     data.SocialMediaUserCategoryCode,
			SocialMediaUserCategoryCodeText: data.SocialMediaUserCategoryCodeText,
			ID:                              data.ID,
			AccountUUID:                     data.AccountUUID,
			EntityLastChangedOn:             data.EntityLastChangedOn,
			AccountInternalID:               data.AccountInternalID,
		})
	}

	return socialMediaUserProfileCollection, nil
}
