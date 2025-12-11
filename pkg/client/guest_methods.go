package client

import "fmt"

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
	err := c.client.MakeVerkadaRequest("GET", url, options, nil, &ret, 0)
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
	err := c.client.MakeVerkadaRequest("GET", url, options, nil, &ret, 0)
	return &ret, err
}
