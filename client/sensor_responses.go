package client

type GetSensorAlertsResponse struct {
	Alert_events []struct {
		Alert_event_id     string `json:"alert_event_id"`
		Device_id          string `json:"device_id"`
		Device_name        string `json:"device_name"`
		Device_serial      string `json:"device_serial"`
		End_time           int    `json:"end_time"`
		Is_above_max_event bool   `json:"is_above_max_event"`
		Most_extreme_value int    `json:"most_extreme_value"`
		Reading            string `json:"reading"`
		Start_time         int    `json:"start_time"`
		Threshold          int    `json:"threshold"`
	} `json:"alert_events"`
	Next_page_token string `json:"next_page_token"`
}

type GetSensorDataResponse struct {
	Data []struct {
		Heat_index              float64 `json:"heat_index"`
		Humidity                float64 `json:"humidity"`
		Motion                  int     `json:"motion"`
		Noise_level             float64 `json:"noise_level"`
		Pm_1_0_0                float64 `json:"pm_1_0_0"`
		Pm_2_5                  float64 `json:"pm_2_5"`
		Pm_4_0                  float64 `json:"pm_4_0"`
		Tamper                  int     `json:"tamper"`
		Temperature             float64 `json:"temperature"`
		Time                    int     `json:"time"`
		Tvoc                    int     `json:"tvoc"`
		Usa_air_quality_index   int     `json:"usa_air_quality_index"`
		Vape_index              int     `json:"vape_index"`
		Vape_index_experimental int     `json:"vape_index_experimental"`
	} `json:"data"`
	Device_id       string `json:"device_id"`
	Device_name     string `json:"device_name"`
	Device_serial   string `json:"device_serial"`
	Interval        string `json:"interval"`
	Next_page_token string `json:"next_page_token"`
}
