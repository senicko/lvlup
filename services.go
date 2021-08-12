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

// ListServices allows to list all services like VPS or domains.
// It returns request result or any errors encountered.
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

// DDoSAttack represents single DDoS Attack.
type DDoSAttack struct {
	Id        int    `json:"id"`
	Ip        string `json:"ip"`
	StartedAt int    `json:"startedAt"`
	EndedAt   int    `json:"endedAt"`
}

// ListVpsDDoSResult represents result of GET /services/vps/{id}/attacks request.
type ListVpsDDoSResult struct {
	Count int          `json:"count"`
	Items []DDoSAttack `json:"items"`
}

// ListVpsDDoS allows to access list of DDoS attacks for specific VPS.
func (lc LvlClient) ListVpsDDoS(vpsId string) (*ListVpsDDoSResult, error) {
	response, err := lc.get(
		"/services/vps/"+vpsId+"/attacks",
		withHeaders(map[string]string{
			"Authorization": "Bearer " + lc.ApiKey,
		}),
	)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %v", response.Status)
	}

	var body ListVpsDDoSResult
	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
}

// GetUDPFilterResponse represents result of GET /services/vps/{id}/filtering.
type GetUDPFilterResponse struct {
	FilteringEnabled bool   `json:"filteringEnabled"`
	State            string `json:"state"`
}

// GetUDPFilter allows to check UDP filtering status for specified VPS.
func (lc LvlClient) GetUDPFilter(vpsId string) (*GetUDPFilterResponse, error) {
	response, err := lc.get(
		"/services/vps/"+vpsId+"/filtering",
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

	var body GetUDPFilterResponse
	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
}
