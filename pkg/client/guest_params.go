package client

type GetGuestEventsOptions struct {
	site_id    string `name:"site_id"`
	start_time string `name:"start_time"`
	end_time   string `name:"end_time"`
	Limit      *int   `name:"limit"`
	Cursor     string `name:"cursor"`
}

type CreateGuestEventBody struct {
	Event_address     string           `json:"event_address,omitempty"`
	Event_description string           `json:"event_description,omitempty"`
	Event_name        string           `json:"event_name,omitempty"`
	Invitees          []InviteeContact `json:"invitees,omitempty"`
	Rsvp_enabled      bool             `json:"rsvp_enabled,omitempty"`
	Walk_in_enabled   bool             `json:"walk-_in_enabled,omitempty"`
}

type EventTime struct {
	End_time   string `json:"end_time"`
	Start_time string `json:"start_time"`
}

type InviteeContact struct {
	Email        string `json:"email,omitempty"`
	Full_name    string `json:"full_name"`
	Notes        string `json:"notes,omitempty"`
	Phone_number string `json:"phone_number,omitempty"`
}

type GetApprovedListsOptions struct {
	Site_id string `json:"site_id"`
	Cursor  string `json:"cursor"`
	Limit   *int   `json:"limit"`
}

type ApprovedListAddition struct {
	Address              string   `json:"address,omitempty"`
	Approved_list_ids    []string `json:"approved_list_ids"`
	Date_of_birth        string   `json:"date_of_birth,omitempty"`
	Email                string   `json:"email"`
	Expiration_timestamp string   `json:"expiration_timestamp,omitempty"`
	External_id          string   `json:"external_id,omitempty"`
	Full_name            string   `json:"full_name"`
	Phone_number         string   `json:"phone_number,omitempty"`
}

type ApprovedListRemoval struct {
	Approved_list_ids []string `json:"approved_list_ids,omitempty"`
	External_id       string   `json:"external_id,omitempty"`
	Person_id         string   `json:"person_id,omitempty"`
}

type GetApprovedListMembersOptions struct {
	Cursor string `json:"cursor"`
	Limit  *int   `json:"limit"`
}

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
