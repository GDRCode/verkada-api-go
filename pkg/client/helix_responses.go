package client

type DeleteHelixEventResponse struct {
}

type GetHelixEventResponse struct {
	Attributes     any    `json:"attributes"`
	Camera_id      string `json:"camera_id"`
	Event_type_uid string `json:"event_type_uid"`
	Flagged        bool   `json:"flagged"`
	Org_id         string `json:"org_id"`
	Time_ms        int    `json:"time_ms"`
}

type UpdateHelixEventResponse struct {
}

type CreateHelixEventResponse struct {
}

type SearchHelixEventResponse struct {
	Attributes     any    `json:"attributes"`
	Camera_id      string `json:"camera_id"`
	Event_type_uid string `json:"event_type_uid"`
	Flagged        bool   `json:"flagged"`
	Org_id         string `json:"org_id"`
	Time_ms        int    `json:"time_ms"`
}

type DeleteHelixEventTypeResponse struct {
}

type GetHelixEventTypesResponse struct {
	Event_types []struct {
		Event_schema   any    `json:"event_schema"`
		Event_type_uid string `json:"event_type_uid"`
		Name           string `json:"name"`
		Org_id         string `json:"org_id"`
	} `json:"event_types"`
}

type UpdateHelixEventTypeResponse struct {
}

type CreateHelixEventTypeResponse struct {
	Event_schema   any    `json:"event_schema"`
	Event_type_uid string `json:"event_type_uid"`
	Name           string `json:"name"`
	Org_id         string `json:"org_id"`
}

type BatchCreateHelixEventResponse struct {
	Job_id      string `json:"job_id"`
	Status      string `json:"status"`
	Status_url  string `json:"status_url"`
	Total_items int    `json:"total_items"`
}

type GetHelixBatchJobStatusResponse struct {
	Completed_at     string `json:"completed_at"`
	Created_at       string `json:"created_at"`
	Error_message    string `json:"error_message"`
	Failed_items     int    `json:"failed_items"`
	Failure_file_url string `json:"failure_file_url"`
	Job_id           string `json:"job_id"`
	Processed_items  int    `json:"processed_items"`
	Status           string `json:"status"`
	Successful_items int    `json:"successful_items"`
	Total_items      int    `json:"total_items"`
	Updated_at       string `json:"updated_at"`
}
