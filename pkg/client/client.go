package client

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/GDRCode/verkada-api-go/pkg/client/auth"
)

// A Client contains the overarching information needed to make API calls.
// All API requests are made via an underlying http.Client.
// {Product}Client fields are used to organize which methods apply to which products.
// The API Key is used to obtain a TokenContainer for a short-lived auth token and its expiration.
type Client struct {
	httpClient     *http.Client
	Key            string
	TokenContainer auth.TokenContainer
	baseURL        string
	AutoPaginate   bool
	Helix          *HelixClient
	Camera         *CameraClient
	Core           *CoreClient
	Sensor         *SensorClient
	Guest          *GuestClient
	Access         *AccessClient
	ClassicAlarms  *ClassicAlarmsClient
	VX             *VXClient
}

type HelixClient struct {
	client *Client
}

type CameraClient struct {
	client *Client
}

type CoreClient struct {
	client *Client
}

type SensorClient struct {
	client *Client
}

type GuestClient struct {
	client *Client
}

type AccessClient struct {
	client *Client
}

type ClassicAlarmsClient struct {
	client *Client
}

type VXClient struct {
	client *Client
}

// Potential options for initiating a new Client.
// Made into a type struct to allow for future non-breaking option additions.
type ClientOptions struct {
	Region       string
	AutoPaginate bool
}

// New returns a Client and any errors relating to configuration options.
// Region (and therefore base URL for requests) is set at Client creation and cannot be changed.
// Auto-pagination can be enabled so that paginated responses are combined into one response.
func New(options *ClientOptions) (*Client, error) {
	auth.GetEnvFromFile()
	envKey := os.Getenv("API_KEY")
	if envKey == "" {
		return nil, fmt.Errorf("error: no environment variable \"API_KEY\", either set in a .env file for auth.GetEnvFromFile() or set directly using os.SetEnv()")
	}
	c := &Client{
		httpClient:   &http.Client{},
		Key:          envKey,
		AutoPaginate: options.AutoPaginate,
	}
	c.Helix = &HelixClient{client: c}
	c.Camera = &CameraClient{client: c}
	c.Core = &CoreClient{client: c}
	c.Sensor = &SensorClient{client: c}
	c.Guest = &GuestClient{client: c}
	c.Access = &AccessClient{client: c}
	c.ClassicAlarms = &ClassicAlarmsClient{client: c}
	c.VX = &VXClient{client: c}
	switch options.Region {
	case "prod1":
		c.baseURL = "https://api.verkada.com"
	case "prod2":
		c.baseURL = "https://api.eu.verkada.com"
	case "au":
		c.baseURL = "https://api.au.verkada.com"
	default:
		return nil, fmt.Errorf("error: invalid region/shard, must be \"prod1\", \"prod2\", or \"au\" - received %s", options.Region)
	}
	tokenResponse, err := auth.GetAuthToken(envKey, c.baseURL)
	if err != nil {
		return c, err
	} else {
		c.TokenContainer = tokenResponse
	}
	return c, nil
}

// Helper function to one-line a bool to *bool conversion.
// Required because a nullable boolean value is needed to identify disincluded boolean parameters in options structs.
func Bool(b bool) *bool {
	return &b
}

// Helper function to one-line a Int64 to *Int64 conversion.
// Required because a nullable number value is needed to identify disincluded boolean parameters in options structs.
// *Int64 is used for all whole-number values since the Go encoding/json package decodes JSON numbers as Int64.
func Int64(i int64) *int64 {
	return &i
}

// Helper function to one-line a Int64 to *Int64 conversion.
// Required because a nullable number value is needed to identify disincluded boolean parameters in options structs.
// *Int64 is used for all whole-number values since the Go encoding/json package decodes JSON numbers as Int64.
func Float64(f float64) *float64 {
	return &f
}

// Used by all methods that don't require file upload or download.
// Handles auth token refresh automatically based on the Client's API key.
//
// Exported so custom requests can be made and can also be used in case new endpoints are not reflected in the package.
func (c *Client) MakeVerkadaRequest(method string, url string, params any, body any, target any, retry int) error {
	b, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to parse body params via JSON marshal: %+v", body)
	}
	req, _ := http.NewRequest(method, url, bytes.NewReader(b))
	req.Header.Add("accept", "application/json")
	if body != nil {
		req.Header.Add("content-type", "application/json")
	}
	if time.Now().After(c.TokenContainer.Expires) {
		tokenResponse, err := auth.GetAuthToken(c.Key, c.baseURL)
		if err != nil {
			return err
		} else {
			c.TokenContainer = tokenResponse
		}
	}
	req.Header.Add("x-verkada-auth", c.TokenContainer.Token)

	req.URL.RawQuery = assembleQueryParams(params)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return errors.New("request error: " + err.Error() + ", request for: " + req.URL.String())
	}

	if res.StatusCode == 429 {
		retryPeriod := 50 * math.Pow(2, float64(retry))
		time.Sleep(time.Millisecond * 50 * time.Duration(retryPeriod))
		c.MakeVerkadaRequest(method, url, params, body, target, retry+1)
	}

	defer res.Body.Close()
	var buf bytes.Buffer
	tee := io.TeeReader(res.Body, &buf)
	decode := json.NewDecoder(tee)
	decode.DisallowUnknownFields()
	err = decode.Decode(target)
	if err != nil {
		return errors.New(err.Error() + ", status: " + res.Status + ", response: " + buf.String())
	}
	return nil
}

// Used by all methods that require file upload (typically csv or pictures).
// Handles auth token refresh automatically based on the Client's API key.
//
// Exported so custom requests can be made and can also be used in case new endpoints are not reflected in the package.
func (c *Client) MakeVerkadaRequestWithFile(method string, url string, params any, filename string, filetype string, target any, retry int) error {
	boundary := "WebKitFormBoundaryPublicAPIClient"
	var b strings.Builder
	fmt.Fprintf(&b, "--%s\r\nContent-Disposition: form-data; name=\"file\"; filename=\"%s\"\r\nContent-Type: %s\r\n\r\n", boundary, filename, filetype)
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	if filetype == "text/csv" {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Fprintf(&b, "%s\r\n", line)
		}
	} else {
		buf, err := io.ReadAll(file)
		if err != nil {
			return fmt.Errorf("failure to read file %s into bytes buffer", filename)
		}
		fmt.Fprintf(&b, "%s", string(buf))
	}
	fmt.Fprintf(&b, "\r\n--%s--\r\n", boundary)
	body := strings.NewReader(b.String())
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "multipart/form-data; boundary="+boundary)
	if time.Now().After(c.TokenContainer.Expires) {
		tokenResponse, err := auth.GetAuthToken(c.Key, c.baseURL)
		if err != nil {
			return err
		} else {
			c.TokenContainer = tokenResponse
		}
	}
	req.Header.Add("x-verkada-auth", c.TokenContainer.Token)

	req.URL.RawQuery = assembleQueryParams(params)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode == 429 {
		retryPeriod := 50 * math.Pow(2, float64(retry))
		time.Sleep(time.Millisecond * 50 * time.Duration(retryPeriod))
		c.MakeVerkadaRequestWithFile(method, url, params, filename, filetype, target, retry+1)
	}

	defer res.Body.Close()
	var buf bytes.Buffer
	tee := io.TeeReader(res.Body, &buf)
	decode := json.NewDecoder(tee)
	decode.DisallowUnknownFields()
	err = decode.Decode(target)
	if err != nil {
		return errors.New(err.Error() + ", status: " + res.Status + ", response: " + buf.String())
	}
	return nil
}

// Used by all methods that require file download (typically csv or pictures).
// Handles auth token refresh automatically based on the Client's API key.
//
// Exported so custom requests can be made and can also be used in case new endpoints are not reflected in the package.
func (c *Client) MakeVerkadaRequestForFile(method string, url string, params any, filename string, retry int) error {
	req, _ := http.NewRequest(method, url, nil)
	if time.Now().After(c.TokenContainer.Expires) {
		tokenResponse, err := auth.GetAuthToken(c.Key, c.baseURL)
		if err != nil {
			return err
		} else {
			c.TokenContainer = tokenResponse
		}
	}
	req.Header.Add("x-verkada-auth", c.TokenContainer.Token)

	req.URL.RawQuery = assembleQueryParams(params)
	fmt.Println(req.URL.RawQuery)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return errors.New("request error: " + err.Error() + ", request for: " + req.URL.String())
	}

	if res.StatusCode == 429 {
		retryPeriod := 50 * math.Pow(2, float64(retry))
		time.Sleep(time.Millisecond * 50 * time.Duration(retryPeriod))
		c.MakeVerkadaRequestForFile(method, url, params, filename, retry+1)
	}

	defer res.Body.Close()
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	written, err1 := io.Copy(file, res.Body)
	if err1 != nil {
		return fmt.Errorf("error writing to file from request for %s, %d bytes written", req.URL.String(), written)
	}
	return nil
}

// Take any options struct and assemble query parameter string.
// Nil/zero values are treated differently per type.
//
// Strings: Included if not empty
//
// *Bool, *Int64, *Float64: Included if not nil (includes zero and non-zero)
//
// Slices and Arrays: Comma-delimited list of values if not empty
//
// Structs: assembleQueryParams is called recursively and added to the overall result
func assembleQueryParams(params any) string {
	if params == nil {
		return ""
	}
	typ := reflect.TypeOf(params)
	if typ.Kind() != reflect.Struct {
		fmt.Printf("Params (%+v type %s %v) is not a struct and cannot be parsed", params, typ.Kind().String(), params == nil)
		return ""
	}
	var b strings.Builder
	val := reflect.ValueOf(params)
	for i := 0; i < typ.NumField(); i++ {
		fld := typ.Field(i)
		if fld.Type.Kind() == reflect.String {
			if val.Field(i).String() != "" {
				fmt.Fprintf(&b, "%s=%s&", fld.Tag.Get("name"), val.Field(i).String())
			}
		} else if !val.Field(i).IsNil() {
			if val.Field(i).Kind() == reflect.Struct {
				fmt.Fprintf(&b, "%s&", assembleQueryParams(val.Field(i).Elem()))
			} else if val.Field(i).Kind() == reflect.Array || val.Field(i).Kind() == reflect.Slice {
				fmt.Fprintf(&b, "%s=", fld.Tag.Get("name"))
				for j := 0; j < val.Field(i).Len(); j++ {
					fmt.Fprintf(&b, "%v", val.Field(i).Index(j))
					if j != val.Field(i).Len()-1 {
						fmt.Fprintf(&b, ",")
					} else {
						fmt.Fprintf(&b, "&")
					}
				}
			} else {
				switch val.Field(i).Elem().Kind() {
				case reflect.Bool:
					fmt.Fprintf(&b, "%s=%t&", fld.Tag.Get("name"), val.Field(i).Elem().Bool())
				case reflect.Int64:
					fmt.Fprintf(&b, "%s=%d&", fld.Tag.Get("name"), val.Field(i).Elem().Int())
				case reflect.Float64:
					fmt.Fprintf(&b, "%s=%f&", fld.Tag.Get("name"), val.Field(i).Elem().Float())
				}
			}
		}
	}
	return strings.TrimSuffix(b.String(), "&")
}
