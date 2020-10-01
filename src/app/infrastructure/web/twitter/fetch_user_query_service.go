package twitter

import (
	"encoding/json"
	"fmt"
	"github.com/hourglasshoro/berbagi/src/app/domain/value_object"
	query_service "github.com/hourglasshoro/berbagi/src/app/usecase/query"
	"github.com/hourglasshoro/berbagi/src/app/usecase/query/read_model"
	"net/http"
)

type FetchUserQueryService struct {
}

func NewFetchUserQueryService() *FetchUserQueryService {
	qs := new(FetchUserQueryService)
	return qs
}

type UserResponse struct {
	Data struct {
		Id              string `json:"id"`
		ProfileImageUrl string `json:"profile_image_url"`
		Name            string `json:"name"`
		UserName        string `json:"username"`
	} `json:"data"`
}

func (qs *FetchUserQueryService) FetchUser(userId uint64, token value_object.AccessToken) (readModel read_model.FetchUserReadModel, err error) {
	u := fmt.Sprintf("https://api.twitter.com/2/users/%d?user.fields=profile_image_url", userId)
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.Token))
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return
	}
	if res.StatusCode != http.StatusOK {
		err = query_service.CannotFetchTweetsException
		return
	}
	defer res.Body.Close()

	var userRes UserResponse
	err = json.NewDecoder(res.Body).Decode(&userRes)
	if err != nil {
		return
	}

	readModel = read_model.FetchUserReadModel{
		DisplayName: userRes.Data.Name,
		UserName:    userRes.Data.UserName,
		ImageUrl:    userRes.Data.ProfileImageUrl,
	}
	return
}
