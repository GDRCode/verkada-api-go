package client

// Gets information about the all the devices in an alarm site specified by site_id.
//
// This method is for Classic Alarms sites and devices ONLY.
//
// [Verkada API Docs - Get Alarm Devices]
//
// [Verkada API Docs - Get Alarm Devices]: https://apidocs.verkada.com/reference/getalarmsdevicesviewv1
func (c *ClassicAlarmsClient) GetAlarmDevices(site_id string, options *GetAlarmDevicesOptions) (*GetAlarmDevicesResponse, error) {
	if options == nil {
		options = &GetAlarmDevicesOptions{}
	}
	options.site_id = site_id
	var ret GetAlarmDevicesResponse
	url := c.client.baseURL + "/alarms/v1/devices"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	return &ret, err
}

// Returns information about a list of alarm sites
//
// This method is for Classic Alarms sites ONLY.
//
// [Verkada API Docs - GGet Site Information]
//
// [Verkada API Docs - Get Site Information]: https://apidocs.verkada.com/reference/getalarmssitesviewv1
func (c *ClassicAlarmsClient) GetAlarmSites(site_ids []string, options *GetAlarmSitesOptions) (*GetAlarmSitesResponse, error) {
	if options == nil {
		options = &GetAlarmSitesOptions{}
	}
	options.Site_ids = site_ids
	var ret GetAlarmSitesResponse
	url := c.client.baseURL + "/alarms/v1/sites"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	return &ret, err
}
