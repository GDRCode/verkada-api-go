package client

// Gets information about the all the Viewing Stations in an organization.
//
// [Verkada API Docs - Get Viewing Station Devices]
//
// [Verkada API Docs - Get Viewing Station Devices]: https://apidocs.verkada.com/reference/getviewingstationdevicesviewv1
func (c *VXClient) GetVXDevices() (*GetVXDevicesResponse, error) {
	var ret GetVXDevicesResponse
	url := c.client.baseURL + "/viewing_station/v1/devices"
	err := c.client.MakeVerkadaRequest("GET", url, nil, nil, &ret, 0)
	return &ret, err
}
