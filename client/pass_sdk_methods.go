package client

// Creates an SDK token for Pass App access.
//
// [Verkada API Docs - Create SDK Token]
//
// [Verkada API Docs - Create SDK Token]: https://apidocs.verkada.com/reference/post-v2accessuserpasssdk_token
func (c *PassSDKClient) CreateSDKToken(user_id string, code_challenge string) (*CreateSDKTokenResponse, error) {
	body := &createSDKTokenBody{
		User_id:        user_id,
		Code_challenge: code_challenge,
	}
	var ret CreateSDKTokenResponse
	url := c.client.baseURL + "/v2/access/user/pass/sdk_token"
	err := c.client.MakeVerkadaRequest("POST", url, nil, body, &ret, 0)
	return &ret, err
}
