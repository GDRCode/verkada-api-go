package client

import "fmt"

func (c *HelixClient) DeleteHelixEvent(camera_id string, time_ms int64, event_type_uid string) (*DeleteHelixEventResponse, error) {
	options := &DeleteHelixEventOptions{camera_id: camera_id, time_ms: time_ms, event_type_uid: event_type_uid}
	var ret DeleteHelixEventResponse
	url := c.client.baseURL + "/cameras/v1/video_tagging/event"
	err := c.client.MakeVerkadaRequest("DELETE", url, *options, nil, &ret, 0)
	return &ret, err
}

func (c *HelixClient) GetHelixEvent(camera_id string, time_ms int64, event_type_uid string) (*GetHelixEventResponse, error) {
	options := &GetHelixEventOptions{camera_id: camera_id, time_ms: time_ms, event_type_uid: event_type_uid}
	var ret GetHelixEventResponse
	url := c.client.baseURL + "/cameras/v1/video_tagging/event"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	return &ret, err
}

func (c *HelixClient) UpdateHelixEvent(camera_id string, time_ms int64, event_type_uid string, body *UpdateHelixEventBody) (*UpdateHelixEventResponse, error) {
	options := &UpdateHelixEventOptions{camera_id: camera_id, time_ms: time_ms, event_type_uid: event_type_uid}
	var ret UpdateHelixEventResponse
	url := c.client.baseURL + "/cameras/v1/video_tagging/event"
	err := c.client.MakeVerkadaRequest("PATCH", url, *options, body, &ret, 0)
	return &ret, err
}

func (c *HelixClient) CreateHelixEvent(camera_id string, time_ms int64, event_type_uid string, body *CreateHelixEventBody) (*CreateHelixEventResponse, error) {
	fullBody := struct {
		Attributes     any    `json:"attributes,omitempty"`
		Camera_id      string `json:"camera_id"`
		Event_type_uid string `json:"event_type_uid"`
		Flagged        bool   `json:"flagged,omitempty"`
		Time_ms        int64  `json:"time_ms"`
	}{
		Attributes:     body.Attributes,
		Camera_id:      camera_id,
		Event_type_uid: event_type_uid,
		Flagged:        body.Flagged,
		Time_ms:        time_ms,
	}
	var ret CreateHelixEventResponse
	url := c.client.baseURL + "/cameras/v1/video_tagging/event"
	err := c.client.MakeVerkadaRequest("POST", url, nil, fullBody, &ret, 0)
	return &ret, err
}

func (c *HelixClient) SearchHelixEvent(body *SearchHelixEventBody) (*SearchHelixEventResponse, error) {
	var ret SearchHelixEventResponse
	url := c.client.baseURL + "/cameras/v1/video_tagging/event"
	err := c.client.MakeVerkadaRequest("POST", url, nil, body, &ret, 0)
	return &ret, err
}

func (c *HelixClient) DeleteHelixEventType(event_type_uid string) (*DeleteHelixEventTypeResponse, error) {
	options := &DeleteHelixEventTypeOptions{event_type_uid: event_type_uid}
	var ret DeleteHelixEventTypeResponse
	url := c.client.baseURL + "/cameras/v1/video_tagging/event_type"
	err := c.client.MakeVerkadaRequest("DELETE", url, *options, nil, &ret, 0)
	return &ret, err
}

func (c *HelixClient) GetHelixEventTypes(options *GetHelixEventTypesOptions) (*GetHelixEventTypesResponse, error) {
	if options == nil {
		options = &GetHelixEventTypesOptions{}
	}
	// shouldn't filter by both uid and name
	if options.Event_type_uid != "" && options.Name != "" {
		return nil, fmt.Errorf("should not filter by both uid and name - received uid %s and name %s", options.Event_type_uid, options.Name)
	}
	var ret GetHelixEventTypesResponse
	url := c.client.baseURL + "/cameras/v1/video_tagging/event_type"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	return &ret, err
}

func (c *HelixClient) UpdateHelixEventType(event_type_uid string, event_schema any, name string) (*UpdateHelixEventTypeResponse, error) {
	options := &UpdateHelixEventTypeOptions{event_type_uid: event_type_uid}
	fullBody := struct {
		Event_schema any    `json:"event_schema"`
		Name         string `json:"name"`
	}{
		Event_schema: event_schema,
		Name:         name,
	}
	var ret UpdateHelixEventTypeResponse
	url := c.client.baseURL + "/cameras/v1/video_tagging/event_type"
	err := c.client.MakeVerkadaRequest("PATCH", url, *options, fullBody, &ret, 0)
	return &ret, err
}

func (c *HelixClient) CreateHelixEventType(event_schema any, name string) (*CreateHelixEventTypeResponse, error) {
	fullBody := struct {
		Event_schema any    `json:"event_schema"`
		Name         string `json:"name"`
	}{
		Event_schema: event_schema,
		Name:         name,
	}
	var ret CreateHelixEventTypeResponse
	url := c.client.baseURL + "/cameras/v1/video_tagging/event_type"
	err := c.client.MakeVerkadaRequest("POST", url, nil, fullBody, &ret, 0)
	return &ret, err
}
