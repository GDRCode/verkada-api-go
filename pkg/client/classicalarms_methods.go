package client

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
