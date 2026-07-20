package client

type GetVXDevicesResponse struct {
	Devices []struct {
		App_version           string `json:"app_version"`
		Claimed_serial_number string `json:"claimed_serial_number"`
		Device_id             string `json:"device_id"`
		Ip_address            string `json:"ip_address"`
		Last_seen_at          string `json:"last_seen_at"`
		Last_status           string `json:"last_status"`
		Location              string `json:"location"`
		Location_lat          int    `json:"location_lat"`
		Location_lon          int    `json:"location_lon"`
		Name                  string `json:"name"`
		Site_id               string `json:"site_id"`
		Timezone              string `json:"timezone"`
	} `json:"devices"`
}
