package client

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/GDRCode/verkada-api-go/pkg/client/auth"
)

func (c *CameraClient) GetAlerts(options *GetAlertsOptions) (*GetAlertsResponse, error) {
	if options == nil {
		options = &GetAlertsOptions{}
	}
	// Notification type must be one of the following:
	notification_type_validation := map[string]bool{
		"person_of_interest":        true,
		"license_plate_of_interest": true,
		"tamper":                    true,
		"crowd":                     true,
		"motion":                    true,
		"camera_offline":            true,
		"camera_online":             true,
		"line_crossing":             true,
		"loitering":                 true,
	}
	for _, param := range options.Notification_type {
		if ok := notification_type_validation[param]; !ok {
			return nil, fmt.Errorf("could not validate parameter in notification_type: %s", param)
		}
	}
	// page_size must be between 1 and 200
	if options.Page_size != nil && (*options.Page_size < 1 || *options.Page_size > 200) {
		return nil, fmt.Errorf("parameter page_size (%d) is not between 1 and 200", *options.Page_size)
	}
	var ret GetAlertsResponse
	url := c.client.baseURL + "/cameras/v1/alerts"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	if err != nil {
		return nil, err
	}
	if c.client.AutoPaginate {
		for ret.Next_page_token != "" {
			var nextRet GetAlertsResponse
			options.Page_token = ret.Next_page_token
			err = c.client.MakeVerkadaRequest("GET", url, *options, nil, &nextRet, 0)
			if err != nil {
				return &ret, err
			}
			ret.Notifications = append(ret.Notifications, nextRet.Notifications...)
			ret.Next_page_token = nextRet.Next_page_token
		}
	}
	return &ret, err
}

func (c *CameraClient) GetDashboardOTData(dashboard_id string, options *GetDashboardOTDataOptions) (*GetDashboardOTDataResponse, error) {
	if options == nil {
		options = &GetDashboardOTDataOptions{}
	}
	options.dashboard_id = dashboard_id
	// only certain strings allowed for interval
	intervalValidation := map[string]bool{
		"":           true,
		"15_minutes": true,
		"1_hour":     true,
		"6_hours":    true,
		"12_hours":   true,
		"1_day":      true,
		"30_days":    true,
	}
	if !intervalValidation[options.Interval] {
		return nil, fmt.Errorf("could not validate interval parameter: %s", options.Interval)
	}
	var ret GetDashboardOTDataResponse
	url := c.client.baseURL + "/cameras/v1/analytics/dashboard_occupancy_trends"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	return &ret, err
}

func (c *CameraClient) GetMaxCounts(camera_id string, options *GetMaxCountsOptions) (*GetMaxCountsResponse, error) {
	if options == nil {
		options = &GetMaxCountsOptions{}
	}
	options.camera_id = camera_id
	// translating search zones to valid query
	var b strings.Builder
	for _, pair := range options.Search_zones {
		if len(pair) != 2 {
			return nil, fmt.Errorf("failer to parse GetMaxCountsOptions search_zones: inner arrays must have length 2")
		}
		fmt.Fprintf(&b, "%d.%d.", pair[0], pair[0])
	}
	options.search_zones = strings.TrimSuffix(b.String(), ".")
	options.Search_zones = nil
	var ret GetMaxCountsResponse
	url := c.client.baseURL + "/cameras/v1/analytics/max_object_counts"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	return &ret, err
}

func (c *CameraClient) GetObjectCounts(camera_id string, options *GetObjectCountsOptions) (*GetObjectCountsResponse, error) {
	if options == nil {
		options = &GetObjectCountsOptions{}
	}
	options.camera_id = camera_id
	// page_size must be between 1 and 200
	if options.Page_size != nil && (*options.Page_size < 1 || *options.Page_size > 200) {
		return nil, fmt.Errorf("parameter page_size (%d) is not between 1 and 200", *options.Page_size)
	}
	var ret GetObjectCountsResponse
	url := c.client.baseURL + "/cameras/v1/analytics/object_counts"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	if err != nil {
		return nil, err
	}
	if c.client.AutoPaginate {
		for ret.Next_page_token != "" {
			var nextRet GetObjectCountsResponse
			options.Page_token = ret.Next_page_token
			err = c.client.MakeVerkadaRequest("GET", url, *options, nil, &nextRet, 0)
			if err != nil {
				return &ret, err
			}
			ret.Object_counts = append(ret.Object_counts, nextRet.Object_counts...)
			ret.Next_page_token = nextRet.Next_page_token
		}
	}
	return &ret, err
}

func (c *CameraClient) SetMQTTConfig(broker_cert string, broker_host_port string, camera_id string, body *SetMQTTConfigBody) (*SetMQTTConfigResponse, error) {
	if body == nil {
		body = &SetMQTTConfigBody{}
	}
	fullBody := struct {
		Broker_cert      string `json:"broker_cert"`
		Broker_host_port string `json:"broker_host_port"`
		Camera_id        string `json:"camera_id"`
		Client_password  string `json:"client_password,omitempty"`
		Client_username  string `json:"client_username,omitempty"`
	}{
		Broker_cert:      broker_cert,
		Broker_host_port: broker_host_port,
		Camera_id:        camera_id,
		Client_password:  body.Client_password,
		Client_username:  body.Client_username,
	}
	var ret SetMQTTConfigResponse
	url := c.client.baseURL + "/cameras/v1/analytics/object_position_mqtt"
	err := c.client.MakeVerkadaRequest("POST", url, nil, fullBody, &ret, 0)
	return &ret, err
}

// Get Occupancy Trends data
//
// Requires: camera_id
//
// Optional: start_time, end_time, interval, type, preset_id
func (c *CameraClient) GetOTData(camera_id string, preset_id string, options *GetOTDataOptions) (*GetOTDataResponse, error) {
	if options == nil {
		options = &GetOTDataOptions{}
	}
	options.camera_id, options.preset_id = camera_id, preset_id
	// only certain strings allowed for interval
	intervalValidation := map[string]bool{
		"":           true,
		"15_minutes": true,
		"1_hour":     true,
		"6_hours":    true,
		"12_hours":   true,
		"1_day":      true,
		"30_days":    true,
	}
	if options.Interval != "" && !intervalValidation[options.Interval] {
		return nil, fmt.Errorf("could not validate interval parameter: %s", options.Interval)
	}
	var ret GetOTDataResponse
	url := c.client.baseURL + "/cameras/v1/analytics/occupancy_trends"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	return &ret, err
}

func (c *CameraClient) GetSeenPlates(camera_id string, options *GetSeenPlatesOptions) (*GetSeenPlatesResponse, error) {
	if options == nil {
		options = &GetSeenPlatesOptions{}
	}
	options.camera_id = camera_id
	// page_size must be between 1 and 200
	if options.Page_size != nil && (*options.Page_size < 1 || *options.Page_size > 200) {
		return nil, fmt.Errorf("parameter page_size (%d) is not between 1 and 200", *options.Page_size)
	}
	var ret GetSeenPlatesResponse
	url := c.client.baseURL + "/cameras/v1/analytics/lpr/images"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	if err != nil {
		return nil, err
	}
	if c.client.AutoPaginate {
		for ret.Next_page_token != 0 {
			var nextRet GetSeenPlatesResponse
			options.Page_token = Int64(int64(ret.Next_page_token))
			err = c.client.MakeVerkadaRequest("GET", url, *options, nil, &nextRet, 0)
			if err != nil {
				return &ret, err
			}
			ret.Detections = append(ret.Detections, nextRet.Detections...)
			ret.Next_page_token = nextRet.Next_page_token
		}
	}
	return &ret, err
}

func (c *CameraClient) DeleteLPOI(license_plate string) (*DeleteLPOIResponse, error) {
	options := &DeleteLPOIOptions{license_plate: license_plate}
	var ret DeleteLPOIResponse
	url := c.client.baseURL + "/cameras/v1/analytics/lpr/license_plate_of_interest"
	err := c.client.MakeVerkadaRequest("DELETE", url, *options, nil, &ret, 0)
	return &ret, err
}

func (c *CameraClient) GetAllLPOI(options *GetAllLPOIOptions) (*GetAllLPOIResponse, error) {
	if options == nil {
		options = &GetAllLPOIOptions{}
	}
	// page_size must be between 1 and 10,000
	if options.Page_size != nil && (*options.Page_size < 1 || *options.Page_size > 10000) {
		return nil, fmt.Errorf("parameter page_size (%d) is not between 1 and 10,000", *options.Page_size)
	}
	var ret GetAllLPOIResponse
	url := c.client.baseURL + "/cameras/v1/analytics/lpr/license_plate_of_interest"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	if err != nil {
		return nil, err
	}
	if c.client.AutoPaginate {
		for ret.Next_page_token != "" {
			var nextRet GetAllLPOIResponse
			options.Page_token = ret.Next_page_token
			err = c.client.MakeVerkadaRequest("GET", url, *options, nil, &nextRet, 0)
			if err != nil {
				return &ret, err
			}
			ret.License_plate_of_interest = append(ret.License_plate_of_interest, nextRet.License_plate_of_interest...)
			ret.Next_page_token = nextRet.Next_page_token
		}
	}
	return &ret, err
}

func (c *CameraClient) UpdateLPOI(license_plate string, description string) (*UpdateLPOIResponse, error) {
	options := &UpdateLPOIOptions{license_plate: license_plate}
	body := struct {
		Description string `json:"description"`
	}{
		Description: description,
	}
	var ret UpdateLPOIResponse
	url := c.client.baseURL + "/cameras/v1/analytics/lpr/license_plate_of_interest"
	err := c.client.MakeVerkadaRequest("PATCH", url, *options, body, &ret, 0)
	return &ret, err
}

func (c *CameraClient) CreateLPOI(license_plate string, description string) (*CreateLPOIResponse, error) {
	body := struct {
		License_plate string `json:"license_plate"`
		Description   string `json:"description"`
	}{
		License_plate: license_plate,
		Description:   description,
	}
	var ret CreateLPOIResponse
	url := c.client.baseURL + "/cameras/v1/analytics/lpr/license_plate_of_interest"
	err := c.client.MakeVerkadaRequest("POST", url, nil, body, &ret, 0)
	return &ret, err
}

func (c *CameraClient) DeleteLPOIByCSV(filename string) (*DeleteLPOIByCSVResponse, error) {
	var ret DeleteLPOIByCSVResponse
	url := c.client.baseURL + "/cameras/v1/analytics/lpr/license_plate_of_interest/batch"
	err := c.client.MakeVerkadaRequestWithFile("DELETE", url, nil, filename, "text/csv", &ret, 0)
	return &ret, err
}

func (c *CameraClient) CreateLPOIByCSV(filename string) (*CreateLPOIByCSVResponse, error) {
	var ret CreateLPOIByCSVResponse
	url := c.client.baseURL + "/cameras/v1/analytics/lpr/license_plate_of_interest/batch"
	err := c.client.MakeVerkadaRequestWithFile("POST", url, nil, filename, "text/csv", &ret, 0)
	return &ret, err
}

func (c *CameraClient) GetLicensePlateTS(camera_id string, license_plate string, options *GetLicensePlateTSOptions) (*GetLicensePlateTSResponse, error) {
	if options == nil {
		options = &GetLicensePlateTSOptions{}
	}
	options.camera_id, options.license_plate = camera_id, license_plate
	var ret GetLicensePlateTSResponse
	url := c.client.baseURL + "/cameras/v1/analytics/lpr/timestamps"
	err := c.client.MakeVerkadaRequest("GET", url, options, nil, &ret, 0)
	if err != nil {
		return nil, err
	}
	if c.client.AutoPaginate {
		for ret.Next_page_token != 0 {
			var nextRet GetLicensePlateTSResponse
			options.Page_token = Int64(int64(ret.Next_page_token))
			err = c.client.MakeVerkadaRequest("GET", url, *options, nil, &nextRet, 0)
			if err != nil {
				return &ret, err
			}
			ret.Detections = append(ret.Detections, nextRet.Detections...)
			ret.Next_page_token = nextRet.Next_page_token
		}
	}
	return &ret, err
}

func (c *CameraClient) GetCameraAudioStatus(camera_id string) (*GetCameraAudioStatusResponse, error) {
	options := &GetCameraAudioStatusOptions{camera_id: camera_id}
	var ret GetCameraAudioStatusResponse
	url := c.client.baseURL + "/cameras/v1/audio/status"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	return &ret, err
}

func (c *CameraClient) UpdateCameraAudio(camera_id string, enabled bool) (*UpdateCameraAudioResponse, error) {
	body := struct {
		Camera_id string `json:"camera_id"`
		Enabled   bool   `json:"enabled"`
	}{
		Camera_id: camera_id,
		Enabled:   enabled,
	}
	var ret UpdateCameraAudioResponse
	url := c.client.baseURL + "/cameras/v1/audio/status"
	err := c.client.MakeVerkadaRequest("POST", url, nil, body, &ret, 0)
	return &ret, err
}

func (c *CameraClient) GetCBSettings(camera_id string) (*GetCBSettingsResponse, error) {
	options := &GetCBSettingsOptions{camera_id: camera_id}
	var ret GetCBSettingsResponse
	url := c.client.baseURL + "/cameras/v1/cloud_backup/settings"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	return &ret, err
}

func (c *CameraClient) UpdateCBSettings(camera_id string, days_to_preserve string, enabled int, time_to_preserve string, upload_timeslot string, video_quality string, video_to_upload string) (*UpdateCBSettingsResponse, error) {
	// check formatting on days_to_preserve (7 characters 0/1, 6 delimiters ",")
	if len(days_to_preserve) != 13 {
		return nil, fmt.Errorf("parameter days_to_preserve is not the correct length (13) - %s length %d", days_to_preserve, len(days_to_preserve))
	}
	validLetters := map[string]bool{
		"0": true,
		"1": true,
		",": true,
	}
	for _, l := range days_to_preserve {
		if !validLetters[string(l)] {
			return nil, fmt.Errorf("parameter days_to_preserve includes invalid letters - should only be 0 or 1 with \",\" as a delimiter - received %s", days_to_preserve)
		}
	}
	// enabled is int but can only be 0 or 1
	if !(enabled == 0 || enabled == 1) {
		return nil, fmt.Errorf("parameter enabled can only be 0 or 1 - received %d", enabled)
	}
	// check if time_to_preserve is delimited correctly (does not check valid values)
	// valid values should be 0 <= start_time < end_time < 86399
	splitTimeToPreserve := strings.Split(time_to_preserve, ",")
	switch len(splitTimeToPreserve) {
	case 1:
		return nil, fmt.Errorf("did not delimit time_to_preserve with \",\" - received %s", time_to_preserve)
	case 2:
	default:
		return nil, fmt.Errorf("too many delimiters detected in time_to_preserve - received %s", time_to_preserve)
	}
	// check if upload_timeslot is delimited correctly (does not check valid values)
	// valid values should be 0 <= start_time < end_time < 86399
	splitUploadTimeslot := strings.Split(upload_timeslot, ",")
	switch len(splitUploadTimeslot) {
	case 1:
		return nil, fmt.Errorf("did not delimit upload_timeslot with \",\" - received %s", upload_timeslot)
	case 2:
	default:
		return nil, fmt.Errorf("too many delimiters detected in upload_timeslot - received %s", upload_timeslot)
	}
	// video_quality can only be STANDARD_QUALITY or HIGH_QUALITY
	if !(video_quality == "STANDARD_QUALITY" || video_quality == "HIGH_QUALITY") {
		return nil, fmt.Errorf("parameter video_quality can only be \"STANDARD_QUALITY\" or \"HIGH_QUALITY\" - received %s", video_quality)
	}
	// video_to_upload can only be MOTION or ALL
	if !(video_to_upload == "MOTION" || video_to_upload == "ALL") {
		return nil, fmt.Errorf("parameter video_to_upload can only be \"MOTIONS\" or \"ALL\" - received %s", video_to_upload)
	}
	fullBody := struct {
		Camera_id        string `json:"camera_id"`
		Days_to_preserve string `json:"days_to_preserve"`
		Enabled          int    `json:"enabled"`
		Time_to_preserve string `json:"time_to_preserve"`
		Upload_timeslot  string `json:"upload_timeslot"`
		Video_quality    string `json:"video_quality"`
		Video_to_upload  string `json:"video_to_upload"`
	}{
		Camera_id:        camera_id,
		Days_to_preserve: days_to_preserve,
		Enabled:          enabled,
		Time_to_preserve: time_to_preserve,
		Upload_timeslot:  upload_timeslot,
		Video_quality:    video_quality,
		Video_to_upload:  video_to_upload,
	}
	var ret UpdateCBSettingsResponse
	url := c.client.baseURL + "/cameras/v1/cloud_backup/settings"
	err := c.client.MakeVerkadaRequest("POST", url, nil, fullBody, &ret, 0)
	return &ret, err
}

func (c *CameraClient) GetCameraDevices(options *GetCameraDevicesOptions) (*GetCameraDevicesResponse, error) {
	if options == nil {
		options = &GetCameraDevicesOptions{}
	}
	// page_size must be between 1 and 10,000
	if options.Page_size != nil && (*options.Page_size < 1 || *options.Page_size > 10000) {
		return nil, fmt.Errorf("parameter page_size (%d) is not between 1 and 10,000", *options.Page_size)
	}
	var ret GetCameraDevicesResponse
	url := c.client.baseURL + "/cameras/v1/devices"
	err := c.client.MakeVerkadaRequest("GET", url, options, nil, &ret, 0)
	if err != nil {
		return nil, err
	}
	if c.client.AutoPaginate {
		for ret.Next_page_token != "" {
			var nextRet GetCameraDevicesResponse
			options.Page_token = ret.Next_page_token
			err = c.client.MakeVerkadaRequest("GET", url, *options, nil, &nextRet, 0)
			if err != nil {
				return &ret, err
			}
			ret.Cameras = append(ret.Cameras, nextRet.Cameras...)
			ret.Next_page_token = nextRet.Next_page_token
		}
	}
	return &ret, err
}

func (c *CameraClient) GetOTCameras() (*GetOTCamerasResponse, error) {
	var ret GetOTCamerasResponse
	url := c.client.baseURL + "/cameras/v1/occupancy_trend_enabled"
	err := c.client.MakeVerkadaRequest("GET", url, nil, nil, &ret, 0)
	return &ret, err
}

func (c *CameraClient) GetLinkToFootage(camera_id string, options *GetLinkToFootageOptions) (*GetLinkToFootageResponse, error) {
	if options == nil {
		options = &GetLinkToFootageOptions{}
	}
	options.camera_id = camera_id
	var ret GetLinkToFootageResponse
	url := c.client.baseURL + "/cameras/v1/footage/link"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	return &ret, err
}

func (c *CameraClient) GetThumbnailImage(camera_id string, options *GetThumbnailImageOptions, filename string) error {
	if options == nil {
		options = &GetThumbnailImageOptions{}
	}
	options.camera_id = camera_id
	// resolution can only be low_res or hi_res
	resolutionValidation := map[string]bool{
		"":        true,
		"low-res": true,
		"hi-res":  true,
	}
	if !resolutionValidation[options.Resolution] {
		return fmt.Errorf("could not validate resolution parameter: %s", options.Resolution)
	}
	// filename validation and replacement if left blank
	if filename == "" {
		filename = "thumbnail.jpg"
	}
	if !strings.HasSuffix(filename, ".jpg") {
		return fmt.Errorf("included filename is not blank but does not end with \".jpg\" - received %s", filename)
	}
	url := c.client.baseURL + "/cameras/v1/footage/thumbnails"
	err := c.client.MakeVerkadaRequestForFile("GET", url, *options, filename, 0)
	return err
}

func (c *CameraClient) GetLatestThumbnailImage(camera_id string, options *GetLatestThumbnailImageOptions, filename string) error {
	if options == nil {
		options = &GetLatestThumbnailImageOptions{}
	}
	options.camera_id = camera_id
	// resolution can only be low_res or hi_res
	resolutionValidation := map[string]bool{
		"":        true,
		"low-res": true,
		"hi-res":  true,
	}
	if !resolutionValidation[options.Resolution] {
		return fmt.Errorf("could not validate resolution parameter: %s", options.Resolution)
	}
	// filename validation and replacement if left blank
	if filename == "" {
		filename = "thumbnail.jpg"
	}
	if !strings.HasSuffix(filename, ".jpg") {
		return fmt.Errorf("included filename is not blank but does not end with \".jpg\" - received %s", filename)
	}
	url := c.client.baseURL + "/cameras/v1/footage/thumbnails/latest"
	err := c.client.MakeVerkadaRequestForFile("GET", url, *options, filename, 0)
	return err
}

func (c *CameraClient) GetThumbnailLink(camera_id string, options *GetThumbnailLinkOptions) (*GetThumbnailLinkResponse, error) {
	if options == nil {
		options = &GetThumbnailLinkOptions{}
	}
	options.camera_id = camera_id
	var ret GetThumbnailLinkResponse
	url := c.client.baseURL + "/cameras/v1/footage/thumbnails/link"
	err := c.client.MakeVerkadaRequest("GET", url, options, nil, &ret, 0)
	return &ret, err
}

func (c *CameraClient) GetStreamingToken() (*GetStreamingTokenResponse, error) {
	var ret GetStreamingTokenResponse
	buf, jwt, err := auth.GetStreamingToken(c.client.Key, c.client.baseURL)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf.Bytes(), &ret)
	if err != nil {
		return nil, fmt.Errorf("%v - error unmarshing streaming token response - jwt maybe be found \"%s\" <-- empty string if not", err, jwt)
	}
	return &ret, nil
}

func (c *CameraClient) StreamFootage(org_id string, camera_id string, jwt string, options *GetFootageOptions, filename string) (string, error) {
	if options == nil {
		options = &GetFootageOptions{}
	}
	options.org_id, options.camera_id, options.jwt = org_id, camera_id, jwt
	// check for request duration validity
	if (options.Start_time != nil) != (options.End_time != nil) {
		return "", fmt.Errorf("start_time is provided without end_time (or vice versa) for streaming timestamps")
	} else if (options.Start_time != nil && options.End_time != nil) && *options.End_time-*options.Start_time > 3600 {
		return "", fmt.Errorf("difference between start_time and end_time is too large: %d - %d = %d", *options.End_time, *options.Start_time, (*options.End_time - *options.Start_time))
	}
	// check for resolution validity
	resolutionValidation := map[string]bool{
		"":         true,
		"low_res":  true,
		"high_res": true,
	}
	if !resolutionValidation[options.Resolution] {
		return "", fmt.Errorf("could not validate resolution parameter: %s", options.Resolution)
	}
	// filename validation and replacement if left blank
	if filename == "" {
		filename = "stream.m3u8"
	}
	if !strings.HasSuffix(filename, ".m3u8") {
		return "", fmt.Errorf("included filename is not blank but does not end with \".m3u8\" - received %s", filename)
	}
	url := c.client.baseURL + "/stream/cameras/v1/footage/stream/stream.m3u8"
	err := c.client.MakeVerkadaRequestForFile("GET", url, *options, filename, 0)
	return url + "?" + assembleQueryParams(*options), err
}

func (c *CameraClient) DeletePOI(person_id string, options *DeletePOIOptions) (*DeletePOIResponse, error) {
	if options == nil {
		options = &DeletePOIOptions{}
	}
	options.person_id = person_id
	var ret DeletePOIResponse
	url := c.client.baseURL + "/cameras/v1/people/person_of_interest"
	err := c.client.MakeVerkadaRequest("DELETE", url, *options, nil, &ret, 0)
	return &ret, err
}

func (c *CameraClient) GetAllPOI(options *GetAllPOIOptions) (*GetAllPOIResponse, error) {
	if options == nil {
		options = &GetAllPOIOptions{}
	}
	var ret GetAllPOIResponse
	url := c.client.baseURL + "/cameras/v1/people/person_of_interest"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	if err != nil {
		return nil, err
	}
	if c.client.AutoPaginate {
		for ret.Next_token != "" {
			var nextRet GetAllPOIResponse
			options.Page_token = ret.Next_token
			err = c.client.MakeVerkadaRequest("GET", url, *options, nil, &nextRet, 0)
			if err != nil {
				return &ret, err
			}
			ret.Persons_of_interest = append(ret.Persons_of_interest, nextRet.Persons_of_interest...)
			ret.Next_token = nextRet.Next_token
		}
	}
	return &ret, err
}

func (c *CameraClient) UpdatePOI(person_id string, label string) (*UpdatePOIResponse, error) {
	options := UpdatePOIOptions{person_id: person_id}
	body := struct {
		Label string `json:"label"`
	}{
		Label: label,
	}
	var ret UpdatePOIResponse
	url := c.client.baseURL + "/cameras/v1/people/person_of_interest"
	err := c.client.MakeVerkadaRequest("PATCH", url, options, body, &ret, 0)
	return &ret, err
}

func (c *CameraClient) CreatePOI(base64_image string, label string) (*CreatePOIResponse, error) {
	body := struct {
		Base64_image string `json:"base64_image"`
		Label        string `json:"label"`
	}{
		Base64_image: base64_image,
		Label:        label,
	}
	var ret CreatePOIResponse
	url := c.client.baseURL + "/cameras/v1/people/person_of_interest"
	err := c.client.MakeVerkadaRequest("POST", url, nil, body, &ret, 0)
	return &ret, err
}
