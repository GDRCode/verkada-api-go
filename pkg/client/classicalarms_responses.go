package client

type GetAlarmDevicesResponse struct {
	Devices []struct {
		Device_id   string `json:"device_id"`
		Device_type string `json:"device_type"`
		Site_id     string `json:"site_id"`
	} `json:"devices"`
}

type GetAlarmSitesResponse struct {
	Sites []struct {
		Site_id             string `json:"site_id"`
		Site_name           string `json:"site_name"`
		Site_security_level string `json:"site_security_level"`
		Site_state          string `json:"site_state"`
	} `json:"sites"`
}
