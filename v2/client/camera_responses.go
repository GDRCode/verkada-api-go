package client

type GetAllDashboardsResponse struct {
	Cursor string      `json:"cursor"`
	Items  []Dashboard `json:"items"`
}

type Dashboard struct {
	Created_at     string   `json:"created_at"`
	Dashboard_id   string   `json:"dashboard_id"`
	Dashboard_name string   `json:"dashboard_name"`
	Site_ids       []string `json:"site_ids"`
	Timezone       string   `json:"timezone"`
	Widgets        []struct {
		Widget_id   string `json:"widget_id"`
		Widget_name string `json:"widget_name"`
		Widget_type string `json:"widget_type"`
	} `json:"widgets"`
}

type GetMQTTConfigResponse struct {
	Broker_cert      string `json:"broker_cert"`
	Broker_host_port string `json:"broker_host_port"`
	Camera_id        string `json:"camera_id"`
	Client_username  string `json:"client_username"`
}

type SetMQTTConfigResponse struct {
	Broker_cert      string `json:"broker_cert"`
	Broker_host_port string `json:"broker_host_port"`
	Camera_id        string `json:"camera_id"`
	Client_username  string `json:"client_username"`
}

type GetDashboardOTDataResponse struct {
	Dashboard_id   string `json:"dashboard_id"`
	Dashboard_name string `json:"dashboard_name"`
	End_time       string `json:"end_time"`
	In_counts      []struct {
		End_time   string  `json:"end_time"`
		Start_time string  `json:"start_time"`
		Value      float64 `json:"value"`
	} `json:"in_counts"`
	Interval      string `json:"interval"`
	Net_occupancy []struct {
		End_time   string  `json:"end_time"`
		Start_time string  `json:"start_time"`
		Value      float64 `json:"value"`
	} `json:"net_occupancy"`
	Out_counts []struct {
		End_time   string  `json:"end_time"`
		Start_time string  `json:"start_time"`
		Value      float64 `json:"value"`
	} `json:"out_counts"`
	Start_time string `json:"start_time"`
}

type GetDashboardWidgetDataResponse struct {
	Dashboard_id   string `json:"dashboard_id"`
	Dashboard_name string `json:"dashboard_name"`
	End_time       string `json:"end_time"`
	Interval       string `json:"interval"`
	Start_time     string `json:"start_time"`
	Widgets        []struct {
		Conversion_data struct {
			Conversion_rates              map[string]float64 `json:"conversion_rates"`
			Helix_cameras_used            []string           `json:"helix_cameras_used"`
			Helix_counts                  map[string]float64 `json:"helix_counts"`
			Occupancy_camera_presets_used []struct {
				Camera_id string `json:"camera_id"`
				Preset_id string `json:"preset_id"`
			} `json:"occupancy_camera_presets_used"`
			Occupancy_in map[string]float64 `json:"occupancy_in"`
		} `json:"conversion_data"`
		Device_alerts_data struct {
			Alert_count map[string]float64 `json:"alert_count"`
		} `json:"device_alerts_data"`
		Helix_data struct {
			Cameras_used []string       `json:"cameras_used"`
			Helix_stats  map[string]any `json:"helix_stats"`
		} `json:"helix_data"`
		Occupancy_data struct {
			Camera_presets_used []struct {
				Camera_id string `json:"camera_id"`
				Preset_id string `json:"preset_id"`
			} `json:"camera_presets_used"`
			In_counts     map[string]float64 `json:"in_counts"`
			Net_occupancy map[string]float64 `json:"net_occupancy"`
			Out_counts    map[string]float64 `json:"out_counts"`
		} `json:"occupancy_data"`
		Queue_data struct {
			Avg_queue_length    map[string]float64 `json:"avg_queue_length"`
			Avg_wait_time       map[string]float64 `json:"avg_wait_time"`
			Camera_presets_used []struct {
				Camera_id string `json:"camera_id"`
				Preset_id string `json:"preset_id"`
			} `json:"camera_presets_used"`
			Max_queue_length map[string]float64 `json:"max_queue_length"`
			Max_wait_time    map[string]float64 `json:"max_wait_time"`
		} `json:"queue_data"`
		Widget_id   string `json:"widget_id"`
		Widget_name string `json:"widget_name"`
		Widget_type string `json:"widget_type"`
	} `json:"widgets"`
}

type GetMaxCountsResponse struct {
	People_count  int `json:"people_count"`
	Vehicle_count int `json:"vehicle_count"`
}

type GetObjectCountsResponse struct {
	Cursor string `json:"cursor"`
	Items  []struct {
		Detected_at   string `json:"detected_at"`
		People_count  int    `json:"people_count"`
		Vehicle_count int    `json:"vehicle_count"`
	} `json:"items"`
}

type GetOTDataResponse struct {
	Camera_id  string `json:"camera_id"`
	End_time   string `json:"end_time"`
	Interval   string `json:"interval"`
	Preset_id  string `json:"preset_id"`
	Start_time string `json:"start_time"`
	Trends     struct {
		Trend_in []struct {
			Count      int    `json:"count"`
			End_time   string `json:"end_time"`
			Start_time string `json:"start_time"`
		} `json:"trend_in"`
		Trend_out []struct {
			Count      int    `json:"count"`
			End_time   string `json:"end_time"`
			Start_time string `json:"start_time"`
		} `json:"trend_out"`
	} `json:"trends"`
	Type string `json:"type"`
}
