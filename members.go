package googleadmindirectory

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
)

type MembersResponse struct {
	Kind    string   `json:"kind"`
	Etag    string   `json:"etag"`
	Members []Member `json:"members"`
}

type Member struct {
	Kind   string `json:"kind"`
	Etag   string `json:"etag"`
	ID     string `json:"id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Type   string `json:"type"`
	Status string `json:"status"`
}

func (service *Service) Members(groupID string) (*[]Member, *errortools.Error) {
	membersReponse := MembersResponse{}

	requestConfig := oauth2.RequestConfig{
		URL:           service.url(fmt.Sprintf("groups/%s/members", groupID)),
		ResponseModel: &membersReponse,
	}

	_, _, e := service.googleService.Get(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &membersReponse.Members, nil
}
