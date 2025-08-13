package client

type DeleteHelixEventOptions struct {
	camera_id      string `name:"camera_id"`
	time_ms        int64  `name:"time_ms"`
	event_type_uid string `name:"event_type_uid"`
}

type GetHelixEventOptions struct {
	camera_id      string `name:"camera_id"`
	time_ms        int64  `name:"time_ms"`
	event_type_uid string `name:"event_type_uid"`
}

type UpdateHelixEventOptions struct {
	camera_id      string `name:"camera_id"`
	time_ms        int64  `name:"time_ms"`
	event_type_uid string `name:"event_type_uid"`
}

type UpdateHelixEventBody struct {
	Attributes any  `json:"attributes"`
	Flagged    bool `json:"flagged"`
}

type CreateHelixEventBody struct {
	Attributes any
	Flagged    bool
}

type SearchHelixEventBody struct {
	Attribute_filters []struct {
		Attribute_key   string `json:"attribute_key"`
		Attribute_value string `json:"attribute_value"`
		Op              string `json:"op"`
	} `json:"attribute_filters,omitempty"`
	Camera_ids     []string `json:"camera_ids,omitempty"`
	End_time_ms    int64    `json:"end_time_ms,omitempty"`
	Event_type_uid string   `json:"event_type_uid,omitempty"`
	Flagged        bool     `json:"flagged,omitempty"`
	Keywords       []string `json:"keywords,omitempty"`
	Start_time_ms  int64    `json:"start_time_ms,omitempty"`
}

type DeleteHelixEventTypeOptions struct {
	event_type_uid string `name:"event_type_uid"`
}

type GetHelixEventTypesOptions struct {
	Event_type_uid string `name:"event_type_uid"`
	Name           string `name:"name"`
}

type UpdateHelixEventTypeOptions struct {
	event_type_uid string `name:"event_type_uid"`
}
