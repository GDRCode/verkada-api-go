package client

type GetAllDashboardsOptions struct {
	Limit  *int   `name:"limit"`
	Cursor string `name:"cursor"`
}

type SetMQTTConfigOptions struct {
	Client_password string
	Client_username string
}

type setMQTTConfigBody struct {
	Broker_cert      string `json:"broker_cert"`
	Broker_host_port string `json:"broker_host_port"`
	Client_password  string `json:"client_password"`
	Client_username  string `json:"client_username"`
}

type GetDashboardOTDataOptions struct {
	Start_time string `name:"start_time"`
	End_time   string `name:"end_time"`
	Interval   string `name:"interval"`
}

type GetDashboardWidgetDataOptions struct {
	End_time     string
	Interval     string
	Site_ids     []string
	Start_time   string
	Widget_ids   []string
	Widget_types []string
}

type getDashboardWidgetDataBody struct {
	End_time     string   `json:"end_time"`
	Interval     string   `json:"interval"`
	Site_ids     []string `json:"site_ids"`
	Start_time   string   `json:"start_time"`
	Widget_ids   []string `json:"widget_ids"`
	Widget_types []string `json:"widget_types"`
}

type GetMaxCountsOptions struct {
	Start_time   string `name:"start_time"`
	End_time     string `name:"end_time"`
	Search_zones [][]int64
	search_zones string `name:"search#zones"`
}

type GetObjectCountsOptions struct {
	Start_time *string `name:"start_time"`
	End_time   *string `name:"end_time"`
	Cursor     string  `name:"cursor"`
	Limit      *int    `name:"limit"`
}

type GetOTDataOptions struct {
	Start_time string `name:"start_time"`
	End_time   string `name:"end_time"`
	Interval   string `name:"interval"`
	Type       string `name:"type"`
	Preset_id  string `name:"preset_id"`
}
