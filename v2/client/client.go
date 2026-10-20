package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/GDRCode/verkada-api-go/v2/auth"
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
	PassSDK        *PassSDKClient
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

type PassSDKClient struct {
	client *Client
}

// Potential options for initiating a new Client.
// Made into a type struct to allow for future non-breaking option additions.
//
// It is recommended to use a .env file for auth.GetEnvFromFile() or os.Setenv() for the API key.
// The field here is provided in case environment variables are not used.
type ClientOptions struct {
	Region       string
	AutoPaginate bool
	KeyName      string
	APIKey       string
}

// New returns a Client and any errors relating to configuration options.
// Region (and therefore base URL for requests) is set at Client creation and cannot be changed.
// Auto-pagination can be enabled so that paginated responses are combined into one response.
func New(options *ClientOptions) (*Client, error) {
	auth.GetEnvFromFile()
	if options.KeyName == "" {
		options.KeyName = "API_KEY"
	}
	envKey := os.Getenv(options.KeyName)
	if envKey == "" {
		if options.APIKey != "" {
			envKey = options.APIKey
		} else {
			return nil, fmt.Errorf("error: no environment variable \"%s\", either set in a .env file for auth.GetEnvFromFile() or set directly using os.SetEnv()", options.KeyName)
		}
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
	c.PassSDK = &PassSDKClient{client: c}
	switch options.Region {
	case "prod1":
		c.baseURL = "https://api.verkada.com"
	case "prod2":
		c.baseURL = "https://api.eu.verkada.com"
	case "au":
		c.baseURL = "https://api.au.verkada.com"
	case "ohio":
		c.baseURL = "https://api.oh.verkada.com"
	case "gov":
		c.baseURL = "https://api.verkadagov.com"
	default:
		return nil, fmt.Errorf("error: invalid region/shard, must be \"prod1\", \"prod2\", \"au\", \"ohio\", or \"gov\" - received %s", options.Region)
	}
	tokenResponse, err := auth.GetAuthToken(envKey, c.baseURL)
	if err != nil {
		return c, err
	} else {
		c.TokenContainer = tokenResponse
	}
	return c, nil
}

type APIError struct {
	Id         string          `json:"id"`
	Error_code string          `json:"error_code"`
	Message    string          `json:"message"`
	Data       json.RawMessage `json:"data"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error %s: %s", e.Error_code, e.Message)
}

type PaginatedFields struct {
	Limit  *int   `name:"limit"`
	Cursor string `name:"cursor"`
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

	req.URL.RawQuery, err = encodeQuery(params)
	if err != nil {
		return err
	}

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
	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response body - %s", err.Error())
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		var apiErr APIError
		if err := json.Unmarshal(buf, &apiErr); err != nil {
			return fmt.Errorf("request unsuccessful - %s, response: %s", res.Status, string(buf))
		}
		return &apiErr
	}

	decode := json.NewDecoder(bytes.NewReader(buf))
	decode.DisallowUnknownFields()

	if err := decode.Decode(target); err != nil {
		return fmt.Errorf("%s, status: %s, response: %s", err.Error(), res.Status, string(buf))
	}
	return nil
}

// Used by all methods that require file upload (typically csv or pictures).
// Handles auth token refresh automatically based on the Client's API key.
//
// Exported so custom requests can be made and can also be used in case new endpoints are not reflected in the package.
func (c *Client) MakeVerkadaRequestWithFile(method string, url string, params any, bodyParams any, filename string, target any, retry int) error {
	boundary := "WebKitFormBoundaryPublicAPIClient"
	bodyBuf := &bytes.Buffer{}
	writer := multipart.NewWriter(bodyBuf)
	writer.SetBoundary(boundary)

	// Add body parameters if necessary
	if bodyParams != nil {
		jsonBytes, err := json.Marshal(bodyParams)
		if err != nil {
			return fmt.Errorf("failed to marshal body object: %w", err)
		}
		var bodyMap map[string]any
		if err := json.Unmarshal(jsonBytes, &bodyMap); err != nil {
			return fmt.Errorf("failed to unmarshal body to map: %w", err)
		}
		for key, value := range bodyMap {
			// Handle nested structs or arrays by converting them back to strings
			var stringVal string
			switch v := value.(type) {
			case string:
				stringVal = v
			default:
				stringVal = fmt.Sprintf("%v", v)
			}

			if err := writer.WriteField(key, stringVal); err != nil {
				return fmt.Errorf("failed to write field %s: %w", key, err)
			}
		}
	}

	// Add the file with the field name "file"
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	fileWriter, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return fmt.Errorf("failed to create form file for %s: %w", filename, err)
	}
	if _, err := io.Copy(fileWriter, file); err != nil {
		return fmt.Errorf("failed to copy file content for %s: %w", filename, err)
	}

	if err := writer.Close(); err != nil {
		return err
	}

	req, _ := http.NewRequest(method, url, bodyBuf)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", writer.FormDataContentType())
	if time.Now().After(c.TokenContainer.Expires) {
		tokenResponse, err := auth.GetAuthToken(c.Key, c.baseURL)
		if err != nil {
			return err
		} else {
			c.TokenContainer = tokenResponse
		}
	}
	req.Header.Add("x-verkada-auth", c.TokenContainer.Token)

	req.URL.RawQuery, err = encodeQuery(params)
	if err != nil {
		return err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode == 429 {
		retryPeriod := 50 * math.Pow(2, float64(retry))
		time.Sleep(time.Millisecond * 50 * time.Duration(retryPeriod))
		c.MakeVerkadaRequestWithFile(method, url, params, bodyParams, filename, target, retry+1)
	}

	defer res.Body.Close()
	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response body - %s", err.Error())
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		var apiErr APIError
		if err := json.Unmarshal(buf, &apiErr); err != nil {
			return fmt.Errorf("request unsuccessful - %s, response: %s", res.Status, string(buf))
		}
		return &apiErr
	}

	decode := json.NewDecoder(bytes.NewReader(buf))
	decode.DisallowUnknownFields()

	if err := decode.Decode(target); err != nil {
		return fmt.Errorf("%s, status: %s, response: %s", err.Error(), res.Status, string(buf))
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

	var err error
	req.URL.RawQuery, err = encodeQuery(params)
	if err != nil {
		return err
	}

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
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		var apiErr APIError
		buf, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("error reading response body - %s", err.Error())
		}
		if err := json.Unmarshal(buf, &apiErr); err != nil {
			return fmt.Errorf("request unsuccessful - %s, response: %s", res.Status, string(buf))
		}
		return &apiErr
	}

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
				case reflect.Int, reflect.Int64:
					fmt.Fprintf(&b, "%s=%d&", fld.Tag.Get("name"), val.Field(i).Elem().Int())
				case reflect.Float64:
					fmt.Fprintf(&b, "%s=%f&", fld.Tag.Get("name"), val.Field(i).Elem().Float())
				case reflect.Pointer:
					fmt.Fprintf(&b, "%s=%v&", fld.Tag.Get("name"), val.Field(i).Elem().Interface())
				}
			}
		}
	}
	return strings.TrimSuffix(b.String(), "&")
}

func encodeQuery(src any) (string, error) {
	v := reflect.ValueOf(src)

	// if !v.IsValid() {
	// 	return "", fmt.Errorf("expected struct or pointer to struct, got nil")
	// }

	// Allow a struct or a single pointer to a struct.
	if v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return "", nil
		}
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return "", fmt.Errorf("expected struct, got %s", v.Kind())
	}

	values := make(url.Values)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		name := field.Tag.Get("name")

		if name == "" || name == "-" || !v.Field(i).CanInterface() {
			continue
		}

		fv := v.Field(i)

		// Dereference field pointers; omit nil fields.
		for fv.Kind() == reflect.Pointer {
			if fv.IsNil() {
				fv = reflect.Value{}
				break
			}
			fv = fv.Elem()
		}

		if !fv.IsValid() {
			continue
		}

		// Collect slice/array elements into a comma-separated value.
		if fv.Kind() == reflect.Slice || fv.Kind() == reflect.Array {
			items := make([]string, 0, fv.Len())

			for j := 0; j < fv.Len(); j++ {
				elem := fv.Index(j)

				for elem.Kind() == reflect.Pointer {
					if elem.IsNil() {
						fmt.Printf("warning: skipping nil element at %s[%d]", name, j)
						elem = reflect.Value{}
						break
					}
					elem = elem.Elem()
				}

				if !elem.IsValid() {
					continue
				}

				var s string

				switch elem.Kind() {
				case reflect.String:
					s = elem.String()
				case reflect.Bool:
					s = strconv.FormatBool(elem.Bool())
				case reflect.Int64:
					s = strconv.FormatInt(elem.Int(), 10)
				case reflect.Float64:
					s = strconv.FormatFloat(elem.Float(), 'f', -1, 64)
				default:
					return "", fmt.Errorf("%s[%d]: unsupported type %s", name, j, elem.Type())
				}

				items = append(items, s)
			}

			if len(items) > 0 {
				values.Set(name, strings.Join(items, ","))
			}
			continue
		}

		var s string

		switch fv.Kind() {
		case reflect.String:
			s = fv.String()
		case reflect.Bool:
			s = strconv.FormatBool(fv.Bool())
		case reflect.Int64, reflect.Int:
			s = strconv.FormatInt(fv.Int(), 10)
		case reflect.Float64:
			s = strconv.FormatFloat(fv.Float(), 'g', -1, 64)
		default:
			return "", fmt.Errorf("%s: unsupported type %s", name, fv.Type())
		}

		values.Set(name, s)
	}

	return values.Encode(), nil
}
