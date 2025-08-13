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
