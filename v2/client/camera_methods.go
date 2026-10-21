package client

import (
	"fmt"
	"strings"
)

// Returns every operational dashboard in the organization, regardless of which users it is shared with, along with the widgets on each dashboard.
//
// [Verkada API Docs - Get All Operational Dashboards]
//
// [Verkada API Docs - Get All Operational Dashboards]: https://apidocs.verkada.com/v2/reference/getlistdashboardsviewv2
func (c *CameraClient) GetAllDashboards(options *GetAllDashboardsOptions) (*GetAllDashboardsResponse, error) {
	if options == nil {
		options = &GetAllDashboardsOptions{}
	}
	// page_size must be between 1 and 200
	if options.Limit != nil && (*options.Limit < 1 || *options.Limit > 200) {
		return nil, fmt.Errorf("parameter limit (%d) is not between 1 and 200", *options.Limit)
	}
	var ret GetAllDashboardsResponse
	url := c.client.baseURL + "/v2/analytics/operational_dashboard"
	err := c.client.MakeVerkadaRequest("GET", url, options, nil, &ret, 0)
	if err != nil {
		return nil, err
	}
	if c.client.AutoPaginate {
		for ret.Cursor != "" {
			var nextRet GetAllDashboardsResponse
			options.Cursor = ret.Cursor
			err = c.client.MakeVerkadaRequest("GET", url, options, nil, &nextRet, 0)
			if err != nil {
				return &ret, err
			}
			ret.Items = append(ret.Items, nextRet.Items...)
			ret.Cursor = nextRet.Cursor
		}
	}
	return &ret, err
}

// Returns the MQTT broker configuration currently set for a camera's Object Position Events.
// The configured password is never returned in the response.
//
// [Verkada API Docs - Get Object Position MQTT Config]
//
// [Verkada API Docs - Get Object Position MQTT Config]: https://apidocs.verkada.com/v2/reference/getoccupancytrendsmqttconfigviewv2-1
func (c *CameraClient) GetMQTTConfig(camera_id string) (*GetMQTTConfigResponse, error) {
	var ret GetMQTTConfigResponse
	url := c.client.baseURL + "/v2/analytics/cameras/" + camera_id + "/mqtt_configs"
	err := c.client.MakeVerkadaRequest("GET", url, nil, nil, &ret, 0)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

// Replaces the MQTT broker configuration for a camera's Object Position Events.
// Object Position Events from the camera will be published to the MQTT broker specified in the config.
// The camera is identified by the camera_id URL path parameter.
//
// The configured password is never returned in the response.
//
// [Verkada API Docs - Set Object Position MQTT Config]
//
// [Verkada API Docs - Set Object Position MQTT Config]: https://apidocs.verkada.com/v2/reference/putoccupancytrendsmqttconfigviewv2
func (c *CameraClient) SetMQTTConfig(camera_id string, broker_cert string, broker_host_port string, options *SetMQTTConfigOptions) (*SetMQTTConfigResponse, error) {
	if options == nil {
		options = &SetMQTTConfigOptions{}
	}
	body := setMQTTConfigBody{
		Broker_cert:      broker_cert,
		Broker_host_port: broker_host_port,
		Client_password:  options.Client_password,
		Client_username:  options.Client_username,
	}
	var ret SetMQTTConfigResponse
	url := c.client.baseURL + "/v2/analytics/cameras/" + camera_id + "/mqtt_configs"
	err := c.client.MakeVerkadaRequest("PUT", url, nil, &body, &ret, 0)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

// Returns occupancy trends data for a particular dashboard over a specified time range.
// Responses include the start and end time of the data requested and dashboard information (ID, name) alongside per-bucket in counts, out counts, and net occupancy.
//
// [Verkada API Docs - Get Dashboard Occupancy Trend Data]
//
// [Verkada API Docs - Get Dashboard Occupancy Trend Data]: https://apidocs.verkada.com/v2/reference/getdashboardoccupancytrendsviewv2
func (c *CameraClient) GetDashboardOTData(dashboard_id string, options *GetDashboardOTDataOptions) (*GetDashboardOTDataResponse, error) {
	if options == nil {
		options = &GetDashboardOTDataOptions{}
	}
	intervalValidation := map[string]bool{
		"":      true,
		"PT15M": true,
		"PT1H":  true,
		"PT1D":  true,
	}
	if !intervalValidation[options.Interval] {
		return nil, fmt.Errorf("could not validate interval parameter: %s", options.Interval)
	}
	var ret GetDashboardOTDataResponse
	url := c.client.baseURL + "/v2/analytics/operational_dashboard/" + dashboard_id + "/occupancy_trends"
	err := c.client.MakeVerkadaRequest("GET", url, options, nil, &ret, 0)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

// Returns all data for a particular dashboard's widgets over a specified time range.
// Responses will include the start and end time of the data requested and dashboard information (ID, name) alongside the readings.
//
// [Verkada API Docs - Get Dashboard Widget Trend Data]
//
// [Verkada API Docs - Get Dashboard Widget Trend Data]: https://apidocs.verkada.com/v2/reference/postdashboardwidgettrendsview
func (c *CameraClient) GetDashBoardWidgetTrendData(dashboard_id string, options *GetDashboardWidgetDataOptions) (*GetDashboardWidgetDataResponse, error) {
	if options == nil {
		options = &GetDashboardWidgetDataOptions{}
	}
	intervalValidation := map[string]bool{
		"":      true,
		"PT15M": true,
		"PT1H":  true,
		"PT1D":  true,
	}
	if !intervalValidation[options.Interval] {
		return nil, fmt.Errorf("could not validate interval parameter: %s", options.Interval)
	}
	widgetValidation := map[string]bool{
		"occupancy":    true,
		"helix":        true,
		"conversion":   true,
		"queue":        true,
		"deviceAlerts": true,
	}
	for _, widget_type := range options.Widget_types {
		if !widgetValidation[widget_type] {
			return nil, fmt.Errorf("parameter widget_types should only contain \"occupancy\", \"helix\", \"conversion\", and/or \"queue\" - received %s", widget_type)
		}
	}
	body := getDashboardWidgetDataBody{
		End_time:     options.End_time,
		Interval:     options.Interval,
		Site_ids:     options.Site_ids,
		Start_time:   options.Start_time,
		Widget_ids:   options.Widget_ids,
		Widget_types: options.Widget_types,
	}
	var ret GetDashboardWidgetDataResponse
	url := c.client.baseURL + "/v2/analytics/operational_dashboard/" + dashboard_id + "/widget_trends/query"
	err := c.client.MakeVerkadaRequest("POST", url, nil, body, &ret, 0)
	return &ret, err
}

// Returns the max count of people and vehicles within a specified time range.
// Timestamps (start_time, end_time) are provided as ISO 8601 strings (e.g. 2025-07-17T21:06:20Z).
//
// [Verkada API Docs - Get Max People/Vehicle Counts]
//
// [Verkada API Docs - Get Max People/Vehicle Counts]: https://apidocs.verkada.com/v2/reference/getmaxobjectcountsviewv2
func (c *CameraClient) GetMaxCounts(camera_id string, options *GetMaxCountsOptions) (*GetMaxCountsResponse, error) {
	if options == nil {
		options = &GetMaxCountsOptions{}
	}
	// translating search zones to valid query
	var b strings.Builder
	for _, pair := range options.Search_zones {
		if len(pair) != 2 {
			return nil, fmt.Errorf("failer to parse GetMaxCountsOptions search_zones: inner arrays must have length 2")
		}
		fmt.Fprintf(&b, "%d.%d.", pair[0], pair[0])
	}
	options.search_zones = strings.TrimSuffix(b.String(), ".")
	// options.Search_zones = nil
	var ret GetMaxCountsResponse
	url := c.client.baseURL + "/v2/cameras/" + camera_id + "/analytics/max_object_counts"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	return &ret, err
}

// Returns the count of people and vehicles detected on a camera within a specified time range, paginated.
// Timestamps in the request are ISO 8601 strings; timestamps in the response are RFC 3339 strings.
// Continue paginating while the response includes a cursor value.
//
// [Verkada API Docs - Get People/Vehicle Counts]
//
// [Verkada API Docs - Get People/Vehicle Counts]: https://apidocs.verkada.com/v2/reference/getobjectcountsviewv2
func (c *CameraClient) GetObjectCounts(camera_id string, options *GetObjectCountsOptions) (*GetObjectCountsResponse, error) {
	if options == nil {
		options = &GetObjectCountsOptions{}
	}
	// page_size must be between 1 and 200
	if options.Limit != nil && (*options.Limit < 1 || *options.Limit > 200) {
		return nil, fmt.Errorf("parameter limit (%d) is not between 1 and 200", *options.Limit)
	}
	var ret GetObjectCountsResponse
	url := c.client.baseURL + "/v2/cameras/" + camera_id + "/analytics/object_counts"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	if err != nil {
		return nil, err
	}
	if c.client.AutoPaginate {
		for ret.Cursor != "" {
			var nextRet GetObjectCountsResponse
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
