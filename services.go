package lvlup

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Service represents single service from GET /services request response.
type Service struct {
	Id        int    `json:"id"`
	PlanName  string `json:"planName"`
	Active    bool   `json:"active"`
	CreatedAt string `json:"createdAt"`
	PayedTo   string `json:"payedTo"`
	Ip        string `json:"ip"`
	Name      string `json:"name"`
	NodeId    int    `json:"nodeId"`
	ServiceId int    `json:"serviceId"`
}

// ListServices represents result of GET /services request.
type ListServicesResult struct {
	Services []Service
}

// ListServices makes a request to GET /services.
// It returns request result and any errors occured
func (lc LvlClient) ListServices() (*ListServicesResult, error) {
	response, err := lc.get(
		"/services",
		withHeaders(map[string]string{
			"Authorization": "Bearer " + lc.ApiKey,
		}),
	)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %v", response.Status)
	}

	var body ListServicesResult
	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
}
