package client

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
