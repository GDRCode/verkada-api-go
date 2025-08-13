package client

func (c *VXClient) GetVXDevices() (*GetVXDevicesResponse, error) {
	var ret GetVXDevicesResponse
	url := c.client.baseURL + "/viewing_station/v1/devices"
	err := c.client.MakeVerkadaRequest("GET", url, nil, nil, &ret, 0)
	return &ret, err
}
