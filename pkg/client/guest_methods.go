package client

import "fmt"

// Returns a paginated list of Guest events and their event parts for a site within a time range.
//
// [Verkada API Docs - Get Guest Events for a Site]
//
// [Verkada API Docs - Get Guest Events for a Site]: https://apidocs.verkada.com/reference/getguesteventviewv2
func (c *GuestClient) GetGuestEvents(site_id string, start_time string, end_time string, options *GetGuestEventsOptions) (*GetGuestEventsResponse, error) {
	if options == nil {
		options = &GetGuestEventsOptions{}
	}
	if options.Limit != nil && (*options.Limit < 0 || *options.Limit > 200) {
		return nil, fmt.Errorf("parameter limit (%d) is not between 0 and 200", *options.Limit)
	}
	var ret GetGuestEventsResponse
	url := c.client.baseURL + "/v2/guest/guest_events"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	if err != nil {
		return nil, err
	}
	if c.client.AutoPaginate {
		for ret.Cursor != "" {
			var nextRet GetGuestEventsResponse
			options.Cursor = ret.Cursor
			err = c.client.MakeVerkadaRequest("GET", url, *options, nil, &nextRet, 0)
			if err != nil {
				return &ret, err
			}
			ret.Items = append(ret.Items, nextRet.Items...)
			ret.Cursor = nextRet.Cursor
		}
	}
	return &ret, err
}

// Creates a Guest event based on a host, guest type, invitees, and specified start and end times.
// Supports multi-part events by providing multiple time ranges.
//
// [Verkada API Docs - Create a Guest Event]
//
// [Verkada API Docs - Create a Guest Event]: https://apidocs.verkada.com/reference/postguesteventviewv2
func (c *GuestClient) CreateGuestEvent(event_times []EventTime, guest_type_id string, host_id string, site_id string, body *CreateGuestEventBody) (*CreateGuestEventResponse, error) {
	if body == nil {
		body = &CreateGuestEventBody{}
	}
	fullBody := struct {
		CreateGuestEventBody
		Event_times   []EventTime `json:"event_times"`
		Guest_type_id string      `json:"guest_type_id"`
		Host_id       string      `json:"host_id"`
		Site_id       string      `json:"site_id"`
	}{
		CreateGuestEventBody: *body,
		Event_times:          event_times,
		Guest_type_id:        guest_type_id,
		Host_id:              host_id,
		Site_id:              site_id,
	}
	// check for valid number of event times
	if len(event_times) < 1 || len(event_times) > 30 {
		return nil, fmt.Errorf("number of event times must be between 1 and 30 - received %v", len(event_times))
	}
	// check if each invitee notes is under the character limit
	for _, invitee := range body.Invitees {
		if len(invitee.Notes) > 100 {
			return nil, fmt.Errorf("exceeded 100 character limit for notes on invitee: %s", invitee.Full_name)
		}
	}
	var ret CreateGuestEventResponse
	url := c.client.baseURL + "/v2/guest/guest_events"
	err := c.client.MakeVerkadaRequest("POST", url, nil, fullBody, &ret, 0)
	return &ret, err
}

// Returns list of approved lists for the organization.
// Optionally filter by site_id to get lists for a specific site.
//
// [Verkada API Docs - Get Approved Lists]
//
// [Verkada API Docs - Get Approved Lists]: https://apidocs.verkada.com/reference/getapprovedlistsviewv2
func (c *GuestClient) GetApprovedLists(options *GetApprovedListsOptions) (*GetApprovedListsResponse, error) {
	if options == nil {
		options = &GetApprovedListsOptions{}
	}
	if options.Limit != nil && (*options.Limit < 0 || *options.Limit > 1000) {
		return nil, fmt.Errorf("parameter limit (%d) is not between 0 and 1000", *options.Limit)
	}
	var ret GetApprovedListsResponse
	url := c.client.baseURL + "v2/guest/approved_lists"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	if err != nil {
		return nil, err
	}
	if c.client.AutoPaginate {
		for ret.Cursor != "" {
			var nextRet GetApprovedListsResponse
			options.Cursor = ret.Cursor
			err = c.client.MakeVerkadaRequest("GET", url, *options, nil, &nextRet, 0)
			if err != nil {
				return &ret, err
			}
			ret.Items = append(ret.Items, nextRet.Items...)
			ret.Cursor = nextRet.Cursor
		}
	}
	return &ret, err
}

// Get specific details for a Guest event with its event parts and invitees by guest_event_id.
//
// [Verkada API Docs - Get Guest Event by ID]
//
// [Verkada API Docs - Get Guest Event by ID]: https://apidocs.verkada.com/reference/getguestsingleeventviewv2
func (c *GuestClient) DeleteGuestEvent(guest_event_id string) (*DeleteGuestEventResponse, error) {
	var ret DeleteGuestEventResponse
	url := c.client.baseURL + "/v2/guest/guest_events/" + guest_event_id
	err := c.client.MakeVerkadaRequest("DELETE", url, nil, nil, &ret, 0)
	return &ret, err
}

// Deletes a Guest event by guest_event_id.
//
// [Verkada API Docs - Delete a Guest Event]
//
// [Verkada API Docs - Delete a Guest Event]: https://apidocs.verkada.com/reference/deleteguestsingleeventviewv2
func (c *GuestClient) GetGuestEventById(guest_event_id string) (*GetGuestEventByIdResponse, error) {
	var ret GetGuestEventByIdResponse
	url := c.client.baseURL + "/v2/guest/guest_events/" + guest_event_id
	err := c.client.MakeVerkadaRequest("GET", url, nil, nil, &ret, 0)
	return &ret, err
}

// Adds one or more people to one or more approved lists.
// This operation is idempotent - calling it multiple times with the same data won't create duplicates.
// People are matched by external_id first, then by email if no external_id is provided.
//
// [Verkada API Docs - Add People to Approved Lists]
//
// [Verkada API Docs - Add People to Approved Lists]: https://apidocs.verkada.com/reference/patchapprovedlistaddviewv2
func (c *GuestClient) AddToApprovedLists(people []ApprovedListAddition) (*AddToApprovedListsResponse, error) {
	if len(people) > 100 {
		return nil, fmt.Errorf("cannot add more than 100 people at a time to approved list")
	}
	fullBody := struct {
		People []ApprovedListAddition `json:"people"`
	}{
		People: people,
	}
	var ret AddToApprovedListsResponse
	url := c.client.baseURL + "/v2/guest/approved_lists/add"
	err := c.client.MakeVerkadaRequest("PATCH", url, nil, fullBody, &ret, 0)
	return &ret, err
}

// Removes one or more people from one or more approved lists.
// If approved_list_ids is omitted, removes the person from all lists.
// People must be identified by either person_id or external_id.
//
// [Verkada API Docs - Remove People from Approved Lists]
//
// [Verkada API Docs - Remove People from Approved Lists]: https://apidocs.verkada.com/reference/patchapprovedlistremoveviewv2
func (c *GuestClient) RemoveFromApprovedLists(people []ApprovedListRemoval) (*RemoveFromApprovedListsResponse, error) {
	if len(people) < 1 || len(people) > 100 {
		return nil, fmt.Errorf("list of people must be between 1 and 100")
	}
	for _, person := range people {
		if (person.Person_id == "") == (person.External_id == "") {
			return nil, fmt.Errorf("should use one of external_id and person_id - received external_id: %s and person_id: %s", person.External_id, person.Person_id)
		}
	}
	fullBody := struct {
		People []ApprovedListRemoval `json:"people"`
	}{
		People: people,
	}
	var ret RemoveFromApprovedListsResponse
	url := c.client.baseURL + "/v2/guest/approved_lists/remove"
	err := c.client.MakeVerkadaRequest("PATCH", url, nil, fullBody, &ret, 0)
	return &ret, err
}

// Returns a list of people on a specific approved list excluding expired members.
//
// [Verkada API Docs - Get Approved List Members]
//
// [Verkada API Docs - Get Approved List Members]: https://apidocs.verkada.com/reference/getapprovedlistmembersviewv2
func (c *GuestClient) GetApprovedListMembers(approved_list_id string, options *GetApprovedListMembersOptions) (*GetApprovedListMembersResponse, error) {
	if options == nil {
		options = &GetApprovedListMembersOptions{}
	}
	if options.Limit != nil && (*options.Limit < 0 || *options.Limit > 1000) {
		return nil, fmt.Errorf("parameter limit (%d) is not between 0 and 1000", *options.Limit)
	}
	var ret GetApprovedListMembersResponse
	url := c.client.baseURL + "v2/guest/approved_lists/" + approved_list_id
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	if err != nil {
		return nil, err
	}
	if c.client.AutoPaginate {
		for ret.Cursor != "" {
			var nextRet GetApprovedListMembersResponse
			options.Cursor = ret.Cursor
			err = c.client.MakeVerkadaRequest("GET", url, *options, nil, &nextRet, 0)
			if err != nil {
				return &ret, err
			}
			ret.People = append(ret.People, nextRet.People...)
			ret.Cursor = nextRet.Cursor
		}
	}
	return &ret, err
}

// Removes all members from an approved list, effectively clearing the list.
// Returns the count of members removed.
//
// [Verkada API Docs - Reset Approved List]
//
// [Verkada API Docs - Reset Approved List]: https://apidocs.verkada.com/reference/patchapprovedlistresetviewv2
func (c *GuestClient) ResetApprovedList(approved_list_id string) (*ResetApprovedListResponse, error) {
	var ret ResetApprovedListResponse
	url := c.client.baseURL + "/v2/guest/approved_lists/" + approved_list_id + "/reset"
	err := c.client.MakeVerkadaRequest("PATCH", url, nil, nil, &ret, 0)
	return &ret, err
}

// Deletes all deny list entries, including the CSV, POI entries, and photos from the specified site.
//
// [Verkada API Docs - Delete a Guest Deny List]
//
// [Verkada API Docs - Delete a Guest Deny List]: https://apidocs.verkada.com/reference/deletedenylistview
func (c *GuestClient) DeleteDenyList(site_id string) (*DeleteDenyListResponse, error) {
	options := &DeleteDenyListOptions{site_id: site_id}
	var ret DeleteDenyListResponse
	url := c.client.baseURL + "/guest/v1/deny_list"
	err := c.client.MakeVerkadaRequest("DELETE", url, *options, nil, &ret, 0)
	return &ret, err
}

// Uploads a deny list in csv format to the specified site.
// Will overwrite any existing site deny list.
//
// [Verkada API Docs - Post Guest Deny List]
//
// [Verkada API Docs - Post Guest Deny List]: https://apidocs.verkada.com/reference/postdenylistview
func (c *GuestClient) PostDenyList(site_id string, uploadFilename string) (*PostDenyListResponse, error) {
	options := &PostDenyListOptions{site_id: site_id}
	var ret PostDenyListResponse
	url := c.client.baseURL + "/guest/v1/deny_list"
	err := c.client.MakeVerkadaRequestWithFile("POST", url, *options, uploadFilename, "text/csv", &ret, 0)
	return &ret, err
}

// Returns a list of Guest sites in an organization.
//
// [Verkada API Docs - Get Guest Sites]
//
// [Verkada API Docs - Get Guest Sites]: https://apidocs.verkada.com/reference/getguestsiteview
func (c *GuestClient) GetGuestSites() (*GetGuestSitesResponse, error) {
	var ret GetGuestSitesResponse
	url := c.client.baseURL + "/guest/v1/sites"
	err := c.client.MakeVerkadaRequest("GET", url, nil, nil, &ret, 0)
	return &ret, err
}

// Returns a list of visits in a site.
// Maximum timeframe is one day (86,400 second difference between start_time and end_time).
//
// [Verkada API Docs - Get Guest Visits]
//
// [Verkada API Docs - Get Guest Visits]: https://apidocs.verkada.com/reference/getvisitview
func (c *GuestClient) GetGuestVisits(site_id string, start_time *int, end_time *int, options *GetGuestVisitsOptions) (*GetGuestVisitsResponse, error) {
	if options == nil {
		options = &GetGuestVisitsOptions{}
	}
	options.site_id, options.start_time, options.end_time = site_id, start_time, end_time
	// max timeframe of one day (86,400 seconds)
	if *options.end_time-*options.start_time > 86400 {
		return nil, fmt.Errorf("difference between start_time and end_time is too large: %d - %d = %d", *options.end_time, *options.start_time, (*options.end_time - *options.start_time))
	}
	// page_size must be between 1 and 200
	if options.Page_size != nil && (*options.Page_size < 1 || *options.Page_size > 200) {
		return nil, fmt.Errorf("parameter page_size (%d) is not between 1 and 200", *options.Page_size)
	}
	var ret GetGuestVisitsResponse
	url := c.client.baseURL + "/guest/v1/visits"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	if err != nil {
		return nil, err
	}
	if c.client.AutoPaginate {
		for len(ret.Visits) > 0 && ret.Visits[len(ret.Visits)-1].Next_page_token != "" {
			var nextRet GetGuestVisitsResponse
			options.Page_token = ret.Visits[len(ret.Visits)-1].Next_page_token
			err = c.client.MakeVerkadaRequest("GET", url, *options, nil, &nextRet, 0)
			if err != nil {
				return &ret, err
			}
			ret.Visits = append(ret.Visits, nextRet.Visits...)
			ret.Visits[len(ret.Visits)-1].Next_page_token = nextRet.Visits[len(nextRet.Visits)-1].Next_page_token
		}
	}
	return &ret, err
}

// Returns a list of Guest types applied to a site.
//
// [Verkada API Docs - Get Guest Types]
//
// [Verkada API Docs - Get Guest Types]: https://apidocs.verkada.com/reference/getguesttypeviewv2
func (c *GuestClient) GetGuestTypes(site_id string, options *GetGuestTypesOptions) (*GetGuestTypesResponse, error) {
	if options == nil {
		options = &GetGuestTypesOptions{}
	}
	options.site_id = site_id
	var ret GetGuestTypesResponse
	url := c.client.baseURL + "/v2/guest/guest_types"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	return &ret, err
}

// Returns a list of hosts in a site.
//
// [Verkada API Docs - Get Hosts]
//
// [Verkada API Docs - Get Hosts]: https://apidocs.verkada.com/reference/gethostviewv2
func (c *GuestClient) GetHosts(site_id string, options *GetHostsOptions) (*GetHostsResponse, error) {
	if options == nil {
		options = &GetHostsOptions{}
	}
	options.site_id = site_id
	var ret GetHostsResponse
	url := c.client.baseURL + "/v2/guest/hosts"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	return &ret, err
}
