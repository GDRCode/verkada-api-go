package client

type GetAlertsResponse struct {
	Next_page_token string `json:"next_page_token"`
	Notifications   []struct {
		Camera_id         string   `json:"camera_id"`
		Created           int      `json:"created"`
		Crowd_threshold   int      `json:"crowd_threshold"`
		Image_url         string   `json:"image_url"`
		Notification_type string   `json:"notification_type"`
		Objects           []string `json:"objects"`
		Person_label      string   `json:"person_label"`
		Video_url         string   `json:"video_url"`
	} `json:"notifications"`
}

type GetDashboardOTDataResponse struct {
	Dashboard_id   string  `json:"dashboard_id"`
	Dashboard_name string  `json:"dashboard_name"`
	End_time       int     `json:"end_time"`
	Occupancy      [][]int `json:"occupancy"`
	Org_id         string  `json:"org_id"`
	Org_name       string  `json:"org_name"`
	Start_time     int     `json:"start_time"`
	Trend_in       [][]int `json:"trend_in"`
	Trend_out      [][]int `json:"trend_out"`
}

type GetMaxCountsResponse struct {
	People_count  int `json:"people_count"`
	Vehicle_count int `json:"vehicle_count"`
}

type GetObjectCountsResponse struct {
	Next_page_token string `json:"next_page_token"`
	Object_counts   []struct {
		Detected_time int `json:"detected_time"`
		People_count  int `json:"people_count"`
		Vehicle_count int `json:"vehicle_count"`
	} `json:"object_counts"`
}

type SetMQTTConfigResponse struct {
	Broker_cert      string `json:"broker_cert"`
	Broker_host_port string `json:"broker_host_port"`
	Camera_id        string `json:"camera_id"`
	Client_password  string `json:"client_password"`
	Client_username  string `json:"client_username"`
}

type GetOTDataResponse struct {
	Camera_address string  `json:"camera_address"`
	Camera_id      string  `json:"camera_id"`
	Camera_name    string  `json:"camera_name"`
	Camera_site    string  `json:"camera_site"`
	End_time       int     `json:"end_time"`
	Org_id         string  `json:"org_id"`
	Org_name       string  `json:"org_name"`
	Preset_id      string  `json:"preset_id"`
	Start_time     int     `json:"start_time"`
	Trend_in       [][]int `json:"trend_in"`
	Trend_out      [][]int `json:"trend_out"`
	Type           string  `json:"type"`
}

type GetSeenPlatesResponse struct {
	Camera_id  string `json:"camera_id"`
	Detections []struct {
		Image_url         string `json:"image_url"`
		License_plate     string `json:"license_plate"`
		Timestamp         int    `json:"timestamp"`
		Vehicle_image_url string `json:"vehicle_image_url"`
	} `json:"detections"`
	Next_page_token int `json:"next_page_token"`
}

type DeleteLPOIResponse struct {
	Creation_time int    `json:"creation_time"`
	Description   string `json:"description"`
	License_plate string `json:"license_plate"`
}

type GetAllLPOIResponse struct {
	License_plate_of_interest []struct {
		Creation_time int    `json:"creation_time"`
		Description   string `json:"description"`
		License_plate string `json:"license_plate"`
	} `json:"license_plate_of_interest"`
	Next_page_token string `json:"next_page_token"`
}

type UpdateLPOIResponse struct {
	Creation_time int    `json:"creation_time"`
	Description   string `json:"description"`
	License_plate string `json:"license_plate"`
}

type CreateLPOIResponse struct {
	Creation_time int    `json:"creation_time"`
	Description   string `json:"description"`
	License_plate string `json:"license_plate"`
}

type DeleteLPOIByCSVResponse struct {
	AddedMs int `json:"addedMs"`
}

type CreateLPOIByCSVResponse struct {
	AddedMs int `json:"addedMs"`
}

type GetLicensePlateTSResponse struct {
	Camera_id       string `json:"camera_id"`
	Detections      []int  `json:"detections"`
	License_plate   string `json:"license_plate"`
	Next_page_token int    `json:"next_page_token"`
}

type GetCameraAudioStatusResponse struct {
	Camera_id string `json:"camera_id"`
	Enabled   bool   `json:"enabled"`
}

type UpdateCameraAudioResponse struct {
}

type GetCBSettingsResponse struct {
	Camera_id               string `json:"camera_id"`
	Days_to_preserve        string `json:"days_to_preserve"`
	Enabled                 int    `json:"enabled"`
	Last_updated_segment_hq string `json:"last_updated_segment_hq"`
	Last_updated_segment_sq string `json:"last_updated_segment_sq"`
	Time_to_preserve        string `json:"time_to_preserve"`
	Upload_timeslot         string `json:"upload_timeslot"`
	Video_quality           string `json:"video_quality"`
	Video_to_upload         string `json:"video_to_upload"`
}

type UpdateCBSettingsResponse struct {
}

type GetCameraDevicesResponse struct {
	Cameras []struct {
		Camera_id                string  `json:"camera_id"`
		Cloud_retention          int     `json:"cloud_retention"`
		Date_added               int     `json:"date_added"`
		Device_retention         any     `json:"device_retention"`
		Firmware                 string  `json:"firmware"`
		Firmware_update_schedule string  `json:"firmware_update_schedule"`
		Last_online              int     `json:"last_online"`
		Local_ip                 string  `json:"local_ip"`
		Location                 string  `json:"location"`
		Location_angle           float64 `json:"location_angle"`
		Location_lat             float64 `json:"location_lat"`
		Location_lon             float64 `json:"location_lon"`
		Mac                      string  `json:"mac"`
		Model                    string  `json:"model"`
		Name                     string  `json:"name"`
		People_history_enabled   bool    `json:"people_history_enabled"`
		Serial                   string  `json:"serial"`
		Site                     string  `json:"site"`
		Site_id                  string  `json:"site_id"`
		Status                   string  `json:"status"`
		Timezone                 string  `json:"timezone"`
		Vehicle_history_enabled  bool    `json:"vehicle_history_enabled"`
	} `json:"cameras"`
	Next_page_token string `json:"next_page_token"`
}

type GetOTCamerasResponse struct {
	Cameras []struct {
		Camera_id  string   `json:"camera_id"`
		Preset_ids []string `json:"preset_ids"`
		Presets    []struct {
			Object_class string `json:"object_class"`
			Preset_id    string `json:"preset_id"`
		} `json:"presets"`
	} `json:"cameras"`
}

type GetLinkToFootageResponse struct {
	Url string `json:"url"`
}

type GetThumbnailLinkResponse struct {
	Expiry    int    `json:"expiry"`
	Timestamp int    `json:"timestamp"`
	Url       string `json:"url"`
}

type GetStreamingTokenResponse struct {
	AccessibleCameras []string `json:"accessibleCameras"`
	AccessibleSites   []string `json:"accessibleSites"`
	Expiration        int      `json:"expiration"`
	ExpiresAt         int      `json:"expiresAt"`
	Jwt               string   `json:"jwt"`
	Permission        []string `json:"permission"`
}

type DeletePOIResponse struct {
	Created   int    `json:"created"`
	Label     string `json:"label"`
	Last_seen int    `json:"last_seen"`
	Person_id string `json:"person_id"`
}

type GetAllPOIResponse struct {
	Next_token          string `json:"next_token"`
	Persons_of_interest []struct {
		Created   int    `json:"created"`
		Label     string `json:"label"`
		Last_seen int    `json:"last_seen"`
		Person_id string `json:"person_id"`
	} `json:"persons_of_interest"`
}

type UpdatePOIResponse struct {
	Created   int    `json:"created"`
	Label     string `json:"label"`
	Last_een  int    `json:"last_seen"`
	Person_id string `json:"person_id"`
}

type CreatePOIResponse struct {
	Created   int    `json:"created"`
	Label     string `json:"label"`
	Last_een  int    `json:"last_seen"`
	Person_id string `json:"person_id"`
}
