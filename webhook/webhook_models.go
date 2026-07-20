package webhook

import "encoding/json"

type WebhookDetails struct {
	Webhook_type string `json:"webhook_type"`
	Webhook_id   string `json:"webhook_id"`
	Org_id       string `json:"org_id"`
	Created_at   int    `json:"created_at"`
}

type CameraEventsWebhookObject struct {
	WebhookDetails
	Data struct {
		Camera_id         string   `json:"camera_id"`
		Created           int      `json:"created"`
		Crowd_threshold   int      `json:"crowd_threshold"`
		Image_url         string   `json:"image_url"`
		Notification_type string   `json:"notification_type"`
		Objects           []string `json:"objects"`
		Person_label      string   `json:"person_label"`
		Video_url         string   `json:"video_url"`
	} `json:"data"`
}

type LPRWebhookObject struct {
	WebhookDetails
	Data struct {
		Camera_id            string    `json:"camera_id"`
		Created              int       `json:"created"`
		Detected             int       `json:"detected"`
		License_plate_number string    `json:"license_plate_number"`
		Confidence           float64   `json:"confidence"`
		Crop                 []float64 `json:"crop"`
		Image_url            string    `json:"image_url"`
		License_plate_state  string    `json:"license_plate_state"`
	} `json:"data"`
}

type AccessEventWebhookObject struct {
	WebhookDetails
	Data struct {
		Notification_type string `json:"notification_type"`
		Device_id         string `json:"device_id"`
		Created           int    `json:"created"`
		Input_value       string `json:"input_value"`
		Door_id           string `json:"door_id"`
		Event_id          string `json:"event_id"`
		Device_type       string `json:"device_type"`
		Direction         string `json:"direction"`
		Aux_info          struct {
			Aux_input_id   string `json:"aux_input_id"`
			Aux_input_name string `json:"aux_input_name"`
			Site           struct {
				Name    string `json:"name"`
				Site_id string `json:"site_id"`
			} `json:"site"`
		} `json:"aux_info"`
		User_info struct {
			User_id    string `json:"user_id"`
			Name       string `json:"name"`
			Email      string `json:"email"`
			First_name string `json:"first_name"`
			Last_name  string `json:"last_name"`
			Phone      string `json:"phone"`
		} `json:"user_info"`
		Lockdown_info struct {
			Lockdown_id   string `json:"lockdown_id"`
			Lockdown_name string `json:"lockdown_name"`
			Message       string `json:"message"`
			Action        string `json:"action"`
			Enabled       string `json:"enabled"`
			Site          struct {
				Name    string `json:"name"`
				Site_id string `json:"site_id"`
			} `json:"site"`
		} `json:"lockdown_info"`
		Door_info struct {
			Acu_name string `json:"acu_name"`
			Acu_id   string `json:"acu_id"`
			Name     string `json:"name"`
			Door_id  string `json:"door_id"`
			Site     struct {
				Name    string `json:"name"`
				Site_id string `json:"site_id"`
			} `json:"site"`
			Api_control_enabled any `json:"api_control_enabled"`
			Camera_info         any `json:"camera_info"`
		} `json:"door_info"`
	} `json:"data"`
}

type SensorEventWebhookObject struct {
	WebhookDetails
	Data struct {
		Alert_event_id     string  `json:"alert_event_id"`
		Start_time         float64 `json:"start_time"`
		End_time           float64 `json:"end_time"`
		Most_extreme_value float64 `json:"most_extreme_value"`
		Is_above_max_event bool    `json:"is_above_max_event"`
		Threshold          float64 `json:"threshold"`
		Reading            string  `json:"reading"`
		Device_id          string  `json:"device_id"`
		Device_name        string  `json:"device_name"`
		Device_serial      string  `json:"device_serial"`
	} `json:"data"`
}

type NewAlarmsWebhookObject struct {
	WebhookDetails
	Data struct {
		Event_type          string   `json:"event_type"`
		Event_time          string   `json:"event_time"`
		Alarm_id            string   `json:"alarm_id"`
		Site_id             string   `json:"site_id"`
		Site_name           string   `json:"site_name"`
		Partition_id        string   `json:"partition_id"`
		Partition_name      string   `json:"partition_name"`
		Response_id         string   `json:"response_id"`
		Response_level      string   `json:"response_level"`
		Is_silent           bool     `json:"is_silent"`
		Trigger_time        int      `json:"trigger_time"`
		Trigger_device_id   string   `json:"trigger_device_id"`
		Trigger_device_type string   `json:"trigger_device_type"`
		Trigger_device_name string   `json:"trigger_device_name"`
		Trigger_type        string   `json:"trigger_type"`
		Context_camera_ids  []string `json:"context_camera_ids"`
		Incident_link       string   `json:"incident_link"`
	} `json:"data"`
}

type ClassicAlarmsWebhookObject struct {
	WebhookDetails
	Data struct {
		Event_type          string `json:"event_type"`
		Site_id             string `json:"site_id"`
		Site_name           string `json:"site_name"`
		Site_state          string `json:"site_state"`
		Site_security_level string `json:"site_security_level"`
		Timestamp           int    `json:"timestamp"`
	} `json:"data"`
}

type AccessCredentialWebhookObject struct {
	WebhookDetails
	Data struct {
		Event_type          string `json:"eventType"`
		Timestamp           string `json:"timestamp"`
		Grantor_id          string `json:"grantorId"`
		Grantor_employee_id string `json:"grantorEmployeeId"`
		Grantor_external_id string `json:"grantorExternalId"`
		Event_id            string `json:"eventId"`
		Events              []struct {
			User_id     string           `json:"userId"`
			Employee_id string           `json:"employeeId"`
			External_id string           `json:"externalId"`
			New_record  CredentialRecord `json:"newRecord"`
			Old_record  CredentialRecord `json:"oldRecord"`
		} `json:"events"`
	} `json:"data"`
}

type CredentialRecord struct {
	// Credential_details field population depends on Credential_type
	Credential_type    string          `json:"credentialType"`
	Credential_id      string          `json:"credentialId"`
	Credential_details json.RawMessage `json:"credentialDetails"`
}

type CardDetails struct {
	Active      bool   `json:"active"`
	Card_status string `json:"cardStatus"`
	Card_type   string `json:"cardType"`
	Parameters  struct {
		Card_number     int    `json:"card_number"`
		Facility_code   int    `json:"facility_code"`
		Card_number_hex string `json:"card_number_hex"`
		Fascn           string `json:"fascn"`
		Guid            string `json:"guid"`
	} `json:"parameters"`
}

type EntryCodeDetails struct {
	Active      bool `json:"active"`
	Code_length int  `json:"codeLength"`
}

type LicensePlateDetails struct {
	Active bool `json:"active"`
}

type MFACodeDetails struct {
	Active      bool `json:"active"`
	Code_length int  `json:"codeLength"`
}

type QRCodeDetails struct {
	Active        bool   `json:"active"`
	Encoding_type string `json:"encodingType"`
}

type AccessMethodDetails struct {
	Active bool `json:"active"`
}

type FaceDetails struct {
	Active bool   `json:"active"`
	Source string `json:"source"`
}

type EventBasedWebhookObject struct {
	// Details object varies widely based on event and product type
	// See API docs (https://apidocs.verkada.com/reference/event-based-webhooks) for fields
	WebhookDetails
	Data struct {
		Alert_id     string          `json:"alert_id"`
		Alert_name   string          `json:"alert_name"`
		Alert_type   string          `json:"alert_type"`
		Device_id    string          `json:"device_id"`
		Site_id      string          `json:"site_id"`
		Product_type string          `json:"product_type"`
		Image_url    string          `json:"image_url"`
		Created      string          `json:"created"`
		Details      json.RawMessage `json:"details"`
	} `json:"data"`
}

type AIPoweredAlertDetails struct {
	Ai_query string `json:"ai_query"`
}

type ActivityDetectionDetails struct {
	Activity_type string `json:"activity_type"`
}

type SmartListDetails struct {
	Smart_list_name string `json:"smart_list_name"`
	Person_type     string `json:"person_type"`
}

type MotionDetectionDetails struct {
	Object_type string `json:"object_type"`
}

type LineCrossingDetectionDetails struct {
	Object_type string `json:"object_type"`
}

type CrowdDetectionDetails struct {
	Object_type string `json:"object_type"`
}

type LPOIDetectionDetails struct {
	Plate_number    string `json:"plate_number"`
	Lpr_capture_url string `json:"lpr_capture_url"`
}

type CameraStatusDetails struct {
	Status string `json:"status"`
}

type AudioDetectionDetails struct {
	Audio_type string `json:"audio_type"`
}

type POIDetectionDetails struct {
}

type OcclusionDetectionDetails struct {
}

type TamperDetectionDetails struct {
}

type AccessEventWebhookDetails struct {
	Input_value string `json:"input_value"`
	Door_id     string `json:"door_id"`
	Direction   string `json:"direction"`
	Aux_info    struct {
		Aux_input_id   string `json:"aux_input_id"`
		Aux_input_name string `json:"aux_input_name"`
		Site           struct {
			Name    string `json:"name"`
			Site_id string `json:"site_id"`
		} `json:"site"`
	} `json:"aux_info"`
	User_info struct {
		User_id    string `json:"user_id"`
		Name       string `json:"name"`
		Email      string `json:"email"`
		First_name string `json:"first_name"`
		Last_name  string `json:"last_name"`
		Phone      string `json:"phone"`
	} `json:"user_info"`
	Lockdown_info struct {
		Lockdown_id   string `json:"lockdown_id"`
		Lockdown_name string `json:"lockdown_name"`
		Message       string `json:"message"`
		Action        string `json:"action"`
		Enabled       string `json:"enabled"`
		Site          struct {
			Name    string `json:"name"`
			Site_id string `json:"site_id"`
		} `json:"site"`
	} `json:"lockdown_info"`
	Door_info struct {
		Acu_name string `json:"acu_name"`
		Acu_id   string `json:"acu_id"`
		Name     string `json:"name"`
		Door_id  string `json:"door_id"`
		Site     struct {
			Name    string `json:"name"`
			Site_id string `json:"site_id"`
		} `json:"site"`
		Api_control_enabled any `json:"api_control_enabled"`
		Camera_info         any `json:"camera_info"`
	} `json:"door_info"`
}

type PanelStatusDetails struct {
	Status string `json:"status"`
}

type PanelPowerStatusDetails struct {
	Power_source string `json:"power_source"`
}

type PanelCellModuleStatusDetails struct {
	Connectivity_source string `json:"connectivity_source"`
}

type AlarmsDeviceStatusDetails struct {
	Status string `json:"status"`
}

type WiredSensorProblemDetails struct {
	Issue string `json:"issue"`
}

type PartitionStatusDetails struct {
	Partition_status string `json:"partition_status"`
	Partition_name   string `json:"partition_name"`
	Partition_id     string `json:"partition_id"`
}

type MisconfiguredResponseDetails struct {
	Issue string `json:"issue"`
}

type AlarmSensorEventDetails struct {
	Trigger_device_id   string   `json:"trigger_device_id"`
	Trigger_device_type string   `json:"trigger_device_type"`
	Trigger_type        string   `json:"trigger_type"`
	Trigger_device_name string   `json:"trigger_device_name"`
	Partition_id        string   `json:"partition_id"`
	Partition_name      string   `json:"partition_name"`
	Site_name           string   `json:"site_name"`
	Trigger_time        int      `json:"trigger_time"`
	Context_camera_ids  []string `json:"context_camera_ids"`
	Incident_link       string   `json:"incident_link"`
}

type SensorBatteryLowDetails struct {
}

type SensorTamperDetails struct {
}

type VLinkInterferenceDetails struct {
}

type GuestEventDetails struct {
	Visit_id        string `json:"visit_id"`
	Visit_type      string `json:"visit_type"`
	Checked_in_time int    `json:"checked_in_time"`
	Full_name       string `json:"full_name"`
	Email           string `json:"email"`
	Phone_number    string `json:"phone_number"`
	Hosts           []struct {
		Full_name      string `json:"full_name"`
		Host_id        string `json:"host_id"`
		Student_id     string `json:"student_id"`
		Student_number string `json:"student_number"`
	} `json:"hosts"`
	Device_name    string `json:"device_name"`
	Is_contactless bool   `json:"is_contactless"`
	Site_id        string `json:"site_id"`
}
