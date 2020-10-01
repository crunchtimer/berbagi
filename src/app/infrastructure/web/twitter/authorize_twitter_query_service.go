package twitter

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/hourglasshoro/berbagi/src/app/domain/value_object"
	query_service "github.com/hourglasshoro/berbagi/src/app/usecase/query"
	"net/http"
	"os"
)

type AuthorizeTwitterQueryService struct {
}

func NewAuthorizeTwitterQueryService() *AuthorizeTwitterQueryService {
	qs := new(AuthorizeTwitterQueryService)
	return qs
}

type AuthResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

func (qs *AuthorizeTwitterQueryService) AuthTwitter() (token value_object.AccessToken, err error) {
	url := "https://api.twitter.com/oauth2/token?grant_type=client_credentials"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return
	}
	credential := []byte(fmt.Sprintf("%s:%s", os.Getenv("TWITTER_API_KEY"), os.Getenv("TWITTER_API_SECRET")))
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString(credential)))
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return
	}
	if res.StatusCode != http.StatusOK {
		err = query_service.CannotGetTwitterAccessTokenException
	}
	defer res.Body.Close()

	var auth AuthResponse
	err = json.NewDecoder(res.Body).Decode(&auth)
	if err != nil {
		return
	}
	tok, err := value_object.NewAccessToken(auth.AccessToken)
	if err != nil {
		return
	}
	token = *tok
	return
}
