package client

import "fmt"

// Deletes a license plate from License Plates of Interest using a specified license plate number.
//
// [Verkada API Docs - Delete a License Plate of Interest]
//
// [Verkada API Docs - Delete a License Plate of Interest]: https://apidocs.verkada.com/reference/deletelicenseplateofinterestviewv1
func (c *EventsClient) GetEvents(options *GetEventsOptions) (*GetEventsResponse, error) {
	if options == nil {
		options = &GetEventsOptions{}
	}
	// Product type (if specified) must be one of the following
	product_type_validation := map[string]bool{
		"cameras":        true,
		"sensor":         true,
		"access_control": true,
		"intercom":       true,
	}
	for _, param := range options.Product_types {
		if ok := product_type_validation[param]; !ok {
			return nil, fmt.Errorf("could not validate parameter in product_type: %s", param)
		}
	}
	// Only certain event types are supported by the products
	event_type_validation := map[string]bool{
		"camera_online":                    true,
		"camera_offline":                   true,
		"motion":                           true,
		"person_of_interest":               true,
		"license_plate_of_interest":        true,
		"tamper":                           true,
		"crowd":                            true,
		"line_crossing":                    true,
		"loitering":                        true,
		"inactivity":                       true,
		"door_opened":                      true,
		"door_granted":                     true,
		"door_rejected":                    true,
		"door_forced_open":                 true,
		"door_held_open":                   true,
		"door_tailgating":                  true,
		"door_crowd_detection":             true,
		"door_tamper":                      true,
		"door_offline":                     true,
		"door_poi_detection":               true,
		"door_initialized":                 true,
		"door_armed":                       true,
		"door_armed_button_pressed":        true,
		"door_aux_unlock":                  true,
		"door_locked":                      true,
		"door_unlocked":                    true,
		"door_unarmed_event":               true,
		"door_code_entered_event":          true,
		"door_button_press_entered_event":  true,
		"door_lock_state_changed":          true,
		"door_lockdown":                    true,
		"door_auxinput_change_state":       true,
		"door_auxinput_held":               true,
		"door_low_battery":                 true,
		"door_critical_battery":            true,
		"door_mobile_nfc_scan_accepted":    true,
		"door_mobile_nfc_scan_rejected":    true,
		"door_user_database_corrupt":       true,
		"door_keycard_entered_accepted":    true,
		"door_keycard_entered_rejected":    true,
		"door_code_entered_accepted":       true,
		"door_code_entered_rejected":       true,
		"door_remote_unlock_accepted":      true,
		"door_remote_unlock_rejected":      true,
		"door_press_to_exit_accepted":      true,
		"door_ble_unlock_attempt_accepted": true,
		"door_ble_unlock_attempt_rejected": true,
		"door_acu_offline":                 true,
		"door_fire_alarm_triggered":        true,
		"door_fire_alarm_released":         true,
		"door_acu_fire_alarm_triggered":    true,
		"door_acu_fire_alarm_released":     true,
		"door_schedule_toggle":             true,
		"door_acu_dpi_cut":                 true,
		"door_acu_dpi_short":               true,
		"door_acu_rex_cut":                 true,
		"door_acu_rex_short":               true,
		"door_acu_rex2_cut":                true,
		"door_acu_rex2_short":              true,
		"door_acu_auxinput_cut":            true,
		"door_acu_auxinput_short":          true,
		"door_lockdown_debounced":          true,
		"door_lp_presented_accepted":       true,
		"door_lp_presented_rejected":       true,
		"door_apb_double_entry":            true,
		"door_apb_double_exit":             true,
		"all_access_granted":               true,
		"all_access_rejected":              true,
		"door_auxoutput_activated":         true,
		"door_auxoutput_deactivated":       true,
		"door_schedule_override_applied":   true,
		"door_schedule_override_removed":   true,
		"temperature":                      true,
		"humidity":                         true,
		"tvoc":                             true,
		"co2":                              true,
		"pm25":                             true,
		"noise":                            true,
		"sensor_motion":                    true,
		"air_quality":                      true,
	}
	for _, param := range options.Event_types {
		if ok := event_type_validation[param]; !ok {
			return nil, fmt.Errorf("could not validate parameter in event_type: %s", param)
		}
	}
	// Limit must be between 1 and 200
	if options.Limit != nil && (*options.Limit < 1 || *options.Limit > 200) {
		return nil, fmt.Errorf("parameter page_size (%d) is not between 1 and 200", *options.Limit)
	}
	var ret GetEventsResponse
	url := c.client.baseURL + "/v2/events"
	err := c.client.MakeVerkadaRequest("GET", url, *options, nil, &ret, 0)
	return &ret, err
}
