package client

type GetAlertsOptions struct {
	Start_time        *int     `name:"start_time"`
	End_time          *int     `name:"end_time"`
	Include_image_url *bool    `name:"include_image_url"`
	Page_token        string   `name:"page_token"`
	Page_size         *int     `name:"page_size"`
	Notification_type []string `name:"notification_type"`
}

type GetDashboardOTDataOptions struct {
	dashboard_id string `name:"dashboard_id"`
	Start_time   *int   `name:"start_time"`
	End_time     *int   `name:"end_time"`
	Interval     string `name:"interval"`
}

type GetMaxCountsOptions struct {
	camera_id    string `name:"camera_id"`
	Start_time   *int   `name:"start_time"`
	End_time     *int   `name:"end_time"`
	Search_zones [][]int64
	search_zones string `name:"search#zones"`
}

type GetObjectCountsOptions struct {
	camera_id  string `name:"camera_id"`
	Start_time *int   `name:"start_time"`
	End_time   *int   `name:"end_time"`
	Page_token string `name:"page_token"`
	Page_size  *int   `name:"page_size"`
}

type SetMQTTConfigBody struct {
	Client_password string
	Client_username string
}

type GetOTDataOptions struct {
	camera_id  string `name:"camera_id"`
	Start_time *int   `name:"start_time"`
	End_time   *int   `name:"end_time"`
	Interval   string `name:"interval"`
	preset_id  string `name:"preset_id"`
}

type GetDashboardWidgetTrendDataOptions struct {
	End_time     string   `json:"end_time,omitempty"`
	Interval     string   `json:"interval,omitempty"`
	Site_ids     []string `json:"site_ids,omitempty"`
	Start_time   string   `json:"start_time,omitempty"`
	Widget_ids   []string `json:"widget_ids,omitempty"`
	Widget_types []string `json:"widget_types,omitempty"`
}

type GetSeenPlatesOptions struct {
	camera_id     string `name:"camera_id"`
	License_plate string `name:"license_plate"`
	Start_time    *int   `name:"start_time"`
	End_time      *int   `name:"end_time"`
	Page_token    *int   `name:"page_token"`
	Page_size     *int   `name:"page_size"`
}

type DeleteLPOIOptions struct {
	license_plate string `name:"license_plate"`
}

type GetAllLPOIOptions struct {
	Page_size  *int   `name:"page_size"`
	Page_token string `name:"page_token"`
}

type UpdateLPOIOptions struct {
	license_plate string `name:"license_plate"`
}

type GetLicensePlateTSOptions struct {
	camera_id     string `name:"camera_id"`
	license_plate string `name:"license_plate"`
	Start_time    *int   `name:"start_time"`
	End_time      *int   `name:"end_time"`
	Page_token    *int   `name:"page_token"`
	Page_size     *int   `name:"page_size"`
}

type GetCameraAudioStatusOptions struct {
	camera_id string `name:"camera_id"`
}

type GetCBSettingsOptions struct {
	camera_id string `name:"camera_id"`
}

type GetCameraDevicesOptions struct {
	Page_token string `name:"page_token"`
	Page_size  *int   `name:"page_size"`
}

type GetLinkToFootageOptions struct {
	camera_id string `name:"camera_id"`
	Timestamp string `name:"timestamp"`
}

type GetThumbnailImageOptions struct {
	camera_id  string `name:"camera_id"`
	Timestamp  string `name:"timestamp"`
	Resolution string `name:"resolution"`
}

type GetLatestThumbnailImageOptions struct {
	camera_id  string `name:"camera_id"`
	Resolution string `name:"resolution"`
}

type GetThumbnailLinkOptions struct {
	camera_id string `name:"camera_id"`
	Timestamp *int   `name:"timestamp"`
	Expiry    *int   `name:"expiry"`
}

type GetFootageOptions struct {
	org_id     string `name:"org_id"`
	camera_id  string `name:"camera_id"`
	Start_time *int   `name:"start_time"`
	End_time   *int   `name:"end_time"`
	Resolution string `name:"resolution"`
	jwt        string `name:"jwt"`
}

type DeletePOIOptions struct {
	person_id string `name:"person_id"`
}

type GetAllPOIOptions struct {
	Page_size  *int   `name:"page_size"`
	Page_token string `name:"page_token"`
}

type UpdatePOIOptions struct {
	person_id string `name:"person_id"`
}
