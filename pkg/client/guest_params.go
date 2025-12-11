package client

type DeleteDenyListOptions struct {
	site_id string `name:"site_id"`
}

type PostDenyListOptions struct {
	site_id string `name:"site_id"`
}

type PostDenyListBody struct {
	Base64_ascii_deny_list_csv string `json:"base64_ascii_deny_list_csv"`
}

type GetGuestVisitsOptions struct {
	site_id    string `name:"site_id"`
	start_time *int   `name:"start_time"`
	end_time   *int   `name:"end_time"`
	Page_token string `name:"page_token"`
	Page_size  *int   `name:"page_size"`
}

type GetGuestTypesOptions struct {
	site_id string `name:"site_id"`
	Cursor  string `name:"cursor"`
	Limit   *int   `name:"limit"`
}

type GetHostsOptions struct {
	site_id string `name:"site_id"`
	Email   string `name:"email"`
	Cursor  string `name:"cursor"`
	Limit   *int   `name:"limit"`
}
