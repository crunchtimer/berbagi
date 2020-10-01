package value_object

import "errors"

type AccessToken struct {
	Token string
}

var InvalidAccessTokenException = errors.New("invalid access token")

func NewAccessToken(token string) (t *AccessToken, err error) {
	if token == "" {
		err = InvalidAccessTokenException
		return
	}
	t = new(AccessToken)
	t.Token = token
	return
}
