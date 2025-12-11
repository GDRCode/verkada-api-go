package client

type GetSensorAlertsOptions struct {
	device_ids []string `name:"device_ids"`
	Start_time *int     `name:"start_time"`
	End_time   *int     `name:"end_time"`
	Page_size  *int     `name:"page_size"`
	Page_token string   `name:"page_token"`
	Fields     []string `name:"fields"`
}

type GetSensorDataOptions struct {
	device_id  string   `name:"device_id"`
	Start_time *int     `name:"start_time"`
	End_time   *int     `name:"end_time"`
	Page_token string   `name:"page_token"`
	Page_size  *int     `name:"page_size"`
	Fields     []string `name:"fields"`
	Interval   string   `name:"interval"`
}
