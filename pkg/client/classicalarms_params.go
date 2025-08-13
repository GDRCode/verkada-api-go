package client

type GetAlarmDevicesOptions struct {
	site_id    string   `name:"site_id"`
	Device_ids []string `name:"device_ids"`
}

type GetAlarmSitesOptions struct {
	Site_ids []string `name:"site_ids"`
}
