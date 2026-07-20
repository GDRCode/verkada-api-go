package client

import "fmt"

// Returns all alerts for all (or subset of) sensors in an org over a specified time range.
//
// [Verkada API Docs - Get Sensor Alerts]
//
// [Verkada API Docs - Get Sensor Alerts]: https://apidocs.verkada.com/reference/getsensoralertsviewv1
func (c *SensorClient) GetSensorAlerts(device_ids []string, options *GetSensorAlertsOptions) (*GetSensorAlertsResponse, error) {
	if options == nil {
		options = &GetSensorAlertsOptions{}
	}
	options.device_ids = device_ids
	// page_size must be between 1 and 200
	if options.Page_size != nil && (*options.Page_size < 1 || *options.Page_size > 200) {
		return nil, fmt.Errorf("parameter page_size (%d) is not between 1 and 200", *options.Page_size)
	}
	// Notification type must be one of the following:
	fields_validation := map[string]bool{
		"humidity":              true,
		"motion":                true,
		"noise_level":           true,
		"pm_2_5":                true,
		"pm_4_0":                true,
		"pm_1_0_0":              true,
		"tamper":                true,
		"temperature":           true,
		"tvoc":                  true,
		"usa_air_quality_index": true,
		"vape_index":            true,
		"carbon_dioxide":        true,
		"carbon_monoxide":       true,
		"barometric_pressure":   true,
		"formaldehyde":          true,
		"ambient_light":         true,
		"tvoc_index":            true,
		"heat_index":            true,
	}
	for _, param := range options.Fields {
		if ok := fields_validation[param]; !ok {
			return nil, fmt.Errorf("could not validate parameter in fields: %s", param)
		}
	}
	var ret GetSensorAlertsResponse
	url := c.client.baseURL + "/environment/v1/alerts"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	if err != nil {
		return nil, err
	}
	if c.client.AutoPaginate {
		for ret.Next_page_token != "" {
			var nextRet GetSensorAlertsResponse
			options.Page_token = ret.Next_page_token
			err = c.client.MakeVerkadaRequest("GET", url, *options, nil, &nextRet, 0)
			if err != nil {
				return &ret, err
			}
			ret.Alert_events = append(ret.Alert_events, nextRet.Alert_events...)
			ret.Next_page_token = nextRet.Next_page_token
		}
	}
	return &ret, err
}

// Returns all sensor readings for a particular sensor over a specified time range.
//
// [Verkada API Docs - Get Sensor Data]
//
// [Verkada API Docs - Get Sensor Data]: https://apidocs.verkada.com/reference/getsensordataviewv1
func (c *SensorClient) GetSensorData(device_id string, options *GetSensorDataOptions) (*GetSensorDataResponse, error) {
	if options == nil {
		options = &GetSensorDataOptions{}
	}
	options.device_id = device_id
	// page_size must be between 1 and 200
	if options.Page_size != nil && (*options.Page_size < 1 || *options.Page_size > 200) {
		return nil, fmt.Errorf("parameter page_size (%d) is not between 1 and 200", *options.Page_size)
	}
	// Notification type must be one of the following:
	fields_validation := map[string]bool{
		"humidity":              true,
		"motion":                true,
		"noise_level":           true,
		"pm_2_5":                true,
		"pm_4_0":                true,
		"pm_1_0_0":              true,
		"tamper":                true,
		"temperature":           true,
		"tvoc":                  true,
		"usa_air_quality_index": true,
		"vape_index":            true,
		"carbon_dioxide":        true,
		"carbon_monoxide":       true,
		"barometric_pressure":   true,
		"formaldehyde":          true,
		"ambient_light":         true,
		"tvoc_index":            true,
		"heat_index":            true,
	}
	for _, param := range options.Fields {
		if ok := fields_validation[param]; !ok {
			return nil, fmt.Errorf("could not validate parameter in fields: %s", param)
		}
	}
	var ret GetSensorDataResponse
	url := c.client.baseURL + "/environment/v1/data"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	if err != nil {
		return nil, err
	}
	if c.client.AutoPaginate {
		for ret.Next_page_token != "" {
			var nextRet GetSensorDataResponse
			options.Page_token = ret.Next_page_token
			err = c.client.MakeVerkadaRequest("GET", url, *options, nil, &nextRet, 0)
			if err != nil {
				return &ret, err
			}
			ret.Data = append(ret.Data, nextRet.Data...)
			ret.Next_page_token = nextRet.Next_page_token
		}
	}
	return &ret, err
}
