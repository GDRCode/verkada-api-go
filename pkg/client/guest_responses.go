package client

type GetGuestEventsResponse struct {
	Cursor string `json:"cursor"`
	Items  []struct {
		End_time       string `json:"end_time"`
		Event_name     string `json:"event_name"`
		Event_part_id  string `json:"event_part_id"`
		Guest_event_id string `json:"guest_event_id"`
		Is_multipart   bool   `json:"is_multipart"`
		Start_time     string `json:"start_time"`
	} `json:"items"`
}

type CreateGuestEventResponse struct {
	GuestEventDetails
}

type DeleteGuestEventResponse struct {
}

type GetGuestEventByIdResponse struct {
	GuestEventDetails
}

type GuestEventDetails struct {
	End_time          string      `json:"end_time"`
	Event_address     string      `json:"event_address"`
	Event_description string      `json:"event_description"`
	Event_name        string      `json:"event_name"`
	Event_parts       []EventPart `json:"event_parts"`
	Event_source      string      `json:"event_source"`
	Guest_event_id    string      `json:"guest_event_id"`
	Guest_type_id     string      `json:"guest_type_id"`
	Host_id           string      `json:"host_id"`
	Invitees          []Invitee   `json:"invitees"`
	Rsvp_link         string      `json:"rsvp_link"`
	Site_id           string      `json:"site_id"`
	Start_time        string      `jsonL:"start_time"`
	Status            string      `json:"status"`
	Walk_in_enabled   bool        `json:"walk_in_enabled"`
}

type Invitee struct {
	Guest_email        string `json:"guest_email"`
	Guest_full_name    string `json:"guest_full_name"`
	Guest_phone_number string `json:"guest_phone_number"`
	Invited_guest_id   string `json:"invited_guest_id"`
	Notes              string `json:"notes"`
	Registered_time    string `json:"registered_time"`
	Visit_id           string `json:"visit_id"`
}

type EventPart struct {
	End_time      string `json:"end_time"`
	Event_part_id string `json:"event_part_id"`
	Start_time    string `json:"start_time"`
}

type GetApprovedListsResponse struct {
	Cursor string         `json:"cursor"`
	Items  []ApprovedList `json:"items"`
}

type ApprovedList struct {
	Approved_list_id string `json:"approved_list_id"`
	Name             string `json:"name"`
	Site_id          string `json:"site_id"`
}

type AddToApprovedListsResponse struct {
	Count int `json:"count"`
}

type RemoveFromApprovedListsResponse struct {
	Count int `json:"count"`
}

type GetApprovedListMembersResponse struct {
	Cursor string               `json:"cursor"`
	People []ApprovedListMember `json:"people"`
}

type ApprovedListMember struct {
	Address              string `json:"address"`
	Date_of_birth        string `json:"date_of_birth"`
	Email                string `json:"email"`
	Expiration_timestamp string `json:"expiration_timestamp"`
	External_id          string `json:"external_id"`
	Full_name            string `json:"full_name"`
	Person_id            string `json:"person_id"`
	Phone_number         string `json:"phone_number"`
}

type ResetApprovedListResponse struct {
	Count int `json:"count"`
}

type DeleteDenyListResponse struct {
}

type PostDenyListResponse struct {
	Base64_ascii_error_csv string `json:"base64_ascii_error_csv"`
	Org_id                 string `json:"org_id"`
	Site_id                string `json:"site_id"`
}

type GetGuestSitesResponse struct {
	Guest_sites []struct {
		Org_id    string `json:"org_id"`
		Site_id   string `json:"site_id"`
		Site_name string `json:"site_name"`
	} `json:"guest_sites"`
}

type GetGuestVisitsResponse struct {
	Visits []struct {
		Approval_status        string              `json:"approval_status"`
		Check_in_time          int                 `json:"check_in_time"`
		Deleted                bool                `json:"deleted"`
		Device_name            string              `json:"device_name"`
		Guest                  guestInfo           `json:"guest"`
		Host                   hostInfo            `json:"host"`
		Host_approval_status   string              `json:"host_approval_status"`
		Hosts                  []hostInfo          `json:"hosts"`
		Is_contactless         bool                `json:"is_contactless"`
		Next_page_token        string              `json:"next_page_token"`
		Open_ended_responses   []openEndedResponse `json:"open_ended_responses"`
		Questionnaires         []questionnaire     `json:"questionnaires"`
		Security_screen_status string              `json:"security_screen_status"`
		Sign_out_time          int                 `json:"sign_out_time"`
		Signatures             []signature         `json:"signatures"`
		Site_id                string              `json:"site_id"`
		Visit_id               string              `json:"visit_id"`
		Visit_type             string              `json:"visit_type"`
	} `json:"visits"`
}

type GetGuestTypesResponse struct {
	Cursor string `json:"cursor"`
	Items  []struct {
		Enabled_for_invites bool   `json:"enabled_for_invites"`
		Guest_type_id       string `json:"guest_type_id"`
		Name                string `json:"name"`
	} `json:"items"`
}

type GetHostsResponse struct {
	Cursor string `json:"cursor"`
	Items  []struct {
		Email                  string `json:"email"`
		First_name             string `json:"first_name"`
		Full_name              string `json:"full_name"`
		Has_delegate           bool   `json:"has_delegate"`
		Host_id                string `json:"host_id"`
		Last_name              string `json:"last_name"`
		Original_first_name    string `json:"original_first_name"`
		Phone_number           string `json:"phone_number"`
		Requires_host_approval bool   `json:"requires_host_approval"`
	}
}

type guestInfo struct {
	Created      int    `json:"created"`
	Email        string `json:"email"`
	Full_name    string `json:"full_name"`
	Guest_id     string `json:"guest_id"`
	Has_photo    bool   `json:"has_photo"`
	Image_url    string `json:"image_url"`
	Is_sign_urls bool   `json:"is_sign_urls"`
	Person_id    string `json:"person_id"`
	Phone_number string `json:"phone_number"`
}

type hostInfo struct {
	Delivery_location      string `json:"delivery_location"`
	First_name             string `json:"first_name"`
	Full_name              string `json:"full_name"`
	Has_delegate           bool   `json:"has_delegate"`
	Has_email              bool   `json:"has_email"`
	Has_phone_number       bool   `json:"has_phone_number"`
	Host_id                string `json:"host_id"`
	Last_name              string `json:"last_name"`
	Original_first_name    string `json:"original_first_name"`
	Requires_host_approval bool   `json:"requires_host_approval"`
	Student_id             string `json:"student_id"`
	Student_number         string `json:"student_number"`
}

type openEndedResponse struct {
	ChosenAnswers []string `json:"chosen_answers"`
	Question      string   `json:"question"`
	Response      string   `json:"response"`
}

type questionnaire struct {
	Passed        bool   `json:"passed"`
	Questionnaire string `json:"questionnaire"`
}

type signature struct {
	DocSignatureID string `json:"doc_signature_id"`
	DocumentID     string `json:"document_id"`
	PdfURL         string `json:"pdf_url"`
	VisitID        string `json:"visit_id"`
}
