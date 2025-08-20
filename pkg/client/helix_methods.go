package client

import "fmt"

// This method can be used to delete a Helix event from Command.
// The required parameters to successfully delete a Helix event are the associated Camera ID, Event Type UID, and the exact event epoch time in milliseconds.
//
// [Verkada API Docs - Delete a Helix Event]
//
// [Verkada API Docs - Delete a Helix Event]: https://apidocs.verkada.com/reference/deletevideotaggingeventviewv1
func (c *HelixClient) DeleteHelixEvent(camera_id string, time_ms int64, event_type_uid string) (*DeleteHelixEventResponse, error) {
	options := &DeleteHelixEventOptions{camera_id: camera_id, time_ms: Int64(time_ms), event_type_uid: event_type_uid}
	var ret DeleteHelixEventResponse
	url := c.client.baseURL + "/cameras/v1/video_tagging/event"
	err := c.client.MakeVerkadaRequest("DELETE", url, *options, nil, &ret, 0)
	return &ret, err
}

// This method can be used to retrieve a Helix Event that has already been posted to Command.
// In the return message, the users will be able to see the corresponding attribute values for that unique event.
// To successfully retrieve a Helix Event, users will need to input the associated Camera ID, Event Type UID, and the exact event epoch time in milliseconds.
//
// [Verkada API Docs - Get a Helix Event]
//
// [Verkada API Docs - Get a Helix Event]: https://apidocs.verkada.com/reference/getvideotaggingeventviewv1
func (c *HelixClient) GetHelixEvent(camera_id string, time_ms int64, event_type_uid string) (*GetHelixEventResponse, error) {
	options := &GetHelixEventOptions{camera_id: camera_id, time_ms: Int64(time_ms), event_type_uid: event_type_uid}
	var ret GetHelixEventResponse
	url := c.client.baseURL + "/cameras/v1/video_tagging/event"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	return &ret, err
}

// This method can be used to update a Helix Event that has already been posted to Command.
// This is especially useful if a user needs to add an additional attribute key to the existing event, along with its new corresponding value.
// To successfully update a Helix Event, users will need to input the associated Camera ID, Event Type UID, exact event epoch time in milliseconds, as well as the new attribute key and attribute value that is being updated.
//
// [Verkada API Docs - Update a Helix Event]
//
// [Verkada API Docs - Update a Helix Event]: https://apidocs.verkada.com/reference/patchvideotaggingeventviewv1
func (c *HelixClient) UpdateHelixEvent(camera_id string, time_ms int64, event_type_uid string, body *UpdateHelixEventBody) (*UpdateHelixEventResponse, error) {
	attributes := make(map[string]any, len(body.Attributes))
	for _, item := range body.Attributes {
		attributes[item.Key] = item.Value
	}
	fullBody := struct {
		Attributes any  `json:"attributes,omitempty"`
		Flagged    bool `json:"flagged,omitempty"`
	}{
		Attributes: attributes,
		Flagged:    body.Flagged,
	}
	options := &UpdateHelixEventOptions{camera_id: camera_id, time_ms: Int64(time_ms), event_type_uid: event_type_uid}
	var ret UpdateHelixEventResponse
	url := c.client.baseURL + "/cameras/v1/video_tagging/event"
	err := c.client.MakeVerkadaRequest("PATCH", url, *options, fullBody, &ret, 0)
	return &ret, err
}

// This method can be used to generate a Helix Event in Command.
// Users will be able to specify the attribute values for each attribute key that was previously defined in the Event Type creation process.
// To successfully create a Helix Event, users will need to input the associated Camera ID, Event Type UID, and the exact event epoch timestamp in milliseconds.
//
// Users are not required to provide an attribute value for all of the attribute keys tied to that specific Event Type.
// If an attribute value is not available or does not exist, users can simply disregard the attribute altogether when making the POST request.
//
// [Verkada API Docs - Create a Helix Event]
//
// [Verkada API Docs - Create a Helix Event]: https://apidocs.verkada.com/reference/postvideotaggingeventviewv1
func (c *HelixClient) CreateHelixEvent(camera_id string, time_ms int64, event_type_uid string, body *CreateHelixEventBody) (*CreateHelixEventResponse, error) {
	attributes := make(map[string]any, len(body.Attributes))
	for _, item := range body.Attributes {
		attributes[item.Key] = item.Value
	}
	fullBody := struct {
		Attributes     any    `json:"attributes,omitempty"`
		Camera_id      string `json:"camera_id"`
		Event_type_uid string `json:"event_type_uid"`
		Flagged        bool   `json:"flagged,omitempty"`
		Time_ms        int64  `json:"time_ms"`
	}{
		Attributes:     attributes,
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

// This method can be used to search for either a single or multiple Helix Events that have already been posted to Command.
// In the return message, the users will be able to see the corresponding attribute keys and attribute values for those specific Helix Events.
//
// The only required parameters to search for Helix Events is a Verkada API Token with Helix permissions.
// Users will be returned a complete list of all Helix Events that are currently available in Command.
// Users can further narrow down their search by adding:
//   - Camera ID: returns all Helix Events linked to that specific camera or list of cameras.
//   - Event Type UID: returns all Helix Events that share that specific Event Type UID.
//   - Start and End Times: returns all Helix Events that have occurred during that time range.
//   - Attributes Keys and Values: returns all Helix Events that have attributes keys and values matching the user's entered parameters.
//
// [Verkada API Docs - Search a Helix Event]
//
// [Verkada API Docs - Search a Helix Event]: https://apidocs.verkada.com/reference/postvideotaggingeventsearchviewv1
func (c *HelixClient) SearchHelixEvent(body *SearchHelixEventBody) (*SearchHelixEventResponse, error) {
	var ret SearchHelixEventResponse
	url := c.client.baseURL + "/cameras/v1/video_tagging/event"
	err := c.client.MakeVerkadaRequest("POST", url, nil, body, &ret, 0)
	return &ret, err
}

// This method can be used to delete an Event Type from Command.
// The required parameter to successfully delete an Event Type is the Event Type UID.
//
// Note: Once you have deleted an Event Type UID, you will no longer be able to use that specific Event Type UID to create any future Helix Events.
// Additionally, all Helix Events linked to this Event Type UID will be deleted from Command.
//
// [Verkada API Docs - Delete a Helix Event Type]
//
// [Verkada API Docs - Delete a Helix Event Type]: https://apidocs.verkada.com/reference/deletevideotaggingeventtypeviewv1
func (c *HelixClient) DeleteHelixEventType(event_type_uid string) (*DeleteHelixEventTypeResponse, error) {
	options := &DeleteHelixEventTypeOptions{event_type_uid: event_type_uid}
	var ret DeleteHelixEventTypeResponse
	url := c.client.baseURL + "/cameras/v1/video_tagging/event_type"
	err := c.client.MakeVerkadaRequest("DELETE", url, *options, nil, &ret, 0)
	return &ret, err
}

// This method can be used to retrieve a complete list of all Event Types from a specific Command organization, along with its associated names and schemas.
// Users can specify either an Event Type UID or an Event Type name to only retrieve the information tied to that specific Event Type.
//
// [Verkada API Docs - Get List of Helix Event Types]
//
// [Verkada API Docs - Get List of Helix Event Types]: https://apidocs.verkada.com/reference/getvideotaggingeventtypeviewv1
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

// This method can be used to update an Event Type.
// This is especially useful if a user needs to add an additional attribute key to the existing Event Type or change the name of the Event Type.
// To successfully update an Event Type, users will need to input the Event Type UID as well as the new attribute keys or Event Type name that is being updated.
//
// If either event_schema or name are left empty, that aspect will not be affected.
// Importantly, if event_schema is not empty, only the included keys and data types will be in the updated Helix Event Type.
//
// The event_schema is a map with strings for both keys and values.
//
// Each key will be the field name for a Helix event type attributes.
// The corresponding value must be a string representing the desired data type for that field:
// "string", "integer", "float", or "boolean"
//
// [Verkada API Docs - Update a Helix Event Type]
//
// [Verkada API Docs - Update a Helix Event Type]: https://apidocs.verkada.com/reference/patchvideotaggingeventtypeviewv1
func (c *HelixClient) UpdateHelixEventType(event_type_uid string, event_schema map[string]string, name string) (*UpdateHelixEventTypeResponse, error) {
	// validate data types in event_schema
	data_type_validation := map[string]bool{
		"string":  true,
		"integer": true,
		"float":   true,
		"boolean": true,
	}
	for key, value := range event_schema {
		if ok := data_type_validation[value]; !ok {
			return nil, fmt.Errorf("could not validate field type - received key: %s and value: %s", key, value)
		}
	}
	options := &UpdateHelixEventTypeOptions{event_type_uid: event_type_uid}
	fullBody := struct {
		Event_schema map[string]string `json:"event_schema"`
		Name         string            `json:"name"`
	}{
		Event_schema: event_schema,
		Name:         name,
	}
	var ret UpdateHelixEventTypeResponse
	url := c.client.baseURL + "/cameras/v1/video_tagging/event_type"
	err := c.client.MakeVerkadaRequest("PATCH", url, *options, fullBody, &ret, 0)
	return &ret, err
}

// This method can be used to generate an Event Type by defining the Event Type schema.
// The event_schema is a map with strings for both keys and values.
//
// Each key will be the field name for a Helix event type attributes.
// The corresponding value must be a string representing the desired data type for that field:
// "string", "integer", "float", or "boolean"
//
// [Verkada API Docs - Create a Helix Event Type]
//
// [Verkada API Docs - Create a Helix Event Type]: https://apidocs.verkada.com/reference/postvideotaggingeventtypeviewv1
func (c *HelixClient) CreateHelixEventType(event_schema map[string]string, name string) (*CreateHelixEventTypeResponse, error) {
	// validate data types in event_schema
	data_type_validation := map[string]bool{
		"string":  true,
		"integer": true,
		"float":   true,
		"boolean": true,
	}
	for key, value := range event_schema {
		if ok := data_type_validation[value]; !ok {
			return nil, fmt.Errorf("could not validate field type - received key: %s and value: %s", key, value)
		}
	}
	fullBody := struct {
		Event_schema map[string]string `json:"event_schema"`
		Name         string            `json:"name"`
	}{
		Event_schema: event_schema,
		Name:         name,
	}
	var ret CreateHelixEventTypeResponse
	url := c.client.baseURL + "/cameras/v1/video_tagging/event_type"
	err := c.client.MakeVerkadaRequest("POST", url, nil, fullBody, &ret, 0)
	return &ret, err
}
