package lvlup

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Service represents single service from ListServices func result.
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

// ListServices represents result of ListServices func.
type ListServicesResult struct {
	Services []Service
}

// ListServices allows to list all services like VPS or domains.
// It returns request result or any errors encountered.
func (lc LvlClient) ListServices() (*ListServicesResult, error) {
	response, err := lc.get(
		"/services",
		withHeaders(map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", lc.ApiKey),
		}),
	)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		message, err := io.ReadAll(response.Body)

		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("status: %v, message: %s", response.Status, message)
	}

	var result ListServicesResult
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DDoSAttack represents single DDoS Attack.
type DDoSAttack struct {
	Id        int    `json:"id"`
	Ip        string `json:"ip"`
	StartedAt int    `json:"startedAt"`
	EndedAt   int    `json:"endedAt"`
}

// ListDDoSAttacksResult represents result of ListDDoSAttacks func.
type ListDDoSAttacksResult struct {
	Count int          `json:"count"`
	Items []DDoSAttack `json:"items"`
}

// ListDDoSAttacks allows to access list of DDoS attacks for specific VPS.
func (lc LvlClient) ListDDoSAttacks(vpsId string) (*ListDDoSAttacksResult, error) {
	response, err := lc.get(
		fmt.Sprintf("/services/vps/%v/attacks", vpsId),
		withHeaders(map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", lc.ApiKey),
		}),
	)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		message, err := io.ReadAll(response.Body)

		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("status: %v, message: %s", response.Status, message)
	}

	var result ListDDoSAttacksResult
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetUDPFilterResult represents result of GetUDPFilter func.
type GetUDPFilterResult struct {
	FilteringEnabled bool   `json:"filteringEnabled"`
	State            string `json:"state"`
}

// GetUDPFilter allows to check UDP filtering status for specified VPS.
func (lc LvlClient) GetUDPFilter(vpsId string) (*GetUDPFilterResult, error) {
	response, err := lc.get(
		fmt.Sprintf("/services/vps/%v/filtering", vpsId),
		withHeaders(map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", lc.ApiKey),
		}),
	)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		message, err := io.ReadAll(response.Body)

		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("status: %v, message: %s", response.Status, message)
	}

	var result GetUDPFilterResult
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// SetUDPFilteringOptions represents available options for SetUDPFiltering func.
type SetUDPFilteringOptions struct {
	FilteringEnabled bool `json:"filteringEnabled"`
}

// SetUDPFilteringResult represents result of SetUDPFiltering func.
type SetUDPFilteringResult struct {
	FilteringEnabled bool   `json:"filteringEnabled"`
	State            string `json:"state"`
}

// SetUDPFiltering allows to switch UDP filtering status for specified VPS on and off.
func (lc LvlClient) SetUDPFiltering(vpsId string, filteringEnabled bool) (*SetUDPFilteringResult, error) {
	options := SetUDPFilteringOptions{
		FilteringEnabled: filteringEnabled,
	}

	payload, err := json.Marshal(options)

	if err != nil {
		return nil, err
	}

	response, err := lc.put(
		fmt.Sprintf("/services/vps/%v/filtering", vpsId),
		withBody(payload),
		withHeaders(map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", lc.ApiKey),
		}),
	)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		message, err := io.ReadAll(response.Body)

		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("status: %v, message: %s", response.Status, message)
	}

	var result SetUDPFilteringResult
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UDPFilterExceptionPorts represents options for UDP filter exception ports.
type UDPFilterExceptionPorts struct {
	From int `json:"form"`
	To   int `json:"to"`
}

// UDPFilterWhitelistException represents single exception.
type UDPFilterException struct {
	Id       int                       `json:"id"`
	Ports    []UDPFilterExceptionPorts `json:"ports"`
	Protocol string                    `json:"protocol"`
	State    string                    `json:"state"`
}

// ListUDPFilterExceptions allows to list all exceptions for UDP filter.
func (lc LvlClient) ListUDPFilterExceptions(vpsId string) ([]UDPFilterException, error) {
	response, err := lc.get(
		fmt.Sprintf("/services/vps/%v/filtering/whitelist", vpsId),
		withHeaders(map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", lc.ApiKey),
		}),
	)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		message, err := io.ReadAll(response.Body)

		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("status: %v, message: %s", response.Status, message)
	}

	var result []UDPFilterException
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

// AddUDPFilterException allows to add exception for UDP filter.
func (lc LvlClient) AddUDPFilterException(vpsId string, exception *UDPFilterException) error {
	payload, err := json.Marshal(exception)

	if err != nil {
		return err
	}

	response, err := lc.post(
		fmt.Sprintf("/services/vps/%v/filtering/whitelist", vpsId),
		withBody(payload),
		withHeaders(map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", lc.ApiKey),
		}),
	)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		message, err := io.ReadAll(response.Body)

		if err != nil {
			return err
		}

		return fmt.Errorf("status: %v, message: %s", response.Status, message)
	}

	return nil
}

// RemoveUDPFilterException allows to remove exception for UDP filter.
func (lc LvlClient) RemoveUDPFilterException(vpsId string, exceptionId string) error {
	response, err := lc.delete(
		fmt.Sprintf("/services/vps/%v/filtering/whitelist/%v", vpsId, exceptionId),
		withHeaders(map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", lc.ApiKey),
		}),
	)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		message, err := io.ReadAll(response.Body)

		if err != nil {
			return err
		}

		return fmt.Errorf("status: %v, message: %s", response.Status, message)
	}

	return nil
}

// ProxmoUser represents proxmo user.
type ProxmoUser struct {
	Password string `json:"password"`
	Url      string `json:"url"`
	Username string `json:"username"`
}

// GetProxmoUser allows to create new proxmo user, or reset password if already exists.
func (lc LvlClient) GetProxmoUser(vpsId string) (*ProxmoUser, error) {
	response, err := lc.post(
		fmt.Sprintf("/services/vps/%v/proxmo", vpsId),
		withHeaders(map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", lc.ApiKey),
		}),
	)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		message, err := io.ReadAll(response.Body)

		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("status: %v, message: %s", response.Status, message)
	}

	var proxmo ProxmoUser
	if err := json.NewDecoder(response.Body).Decode(&proxmo); err != nil {
		return nil, err
	}

	return &proxmo, nil
}

// StartVps allows to start specified VPS server.
func (lc LvlClient) StartVPS(vpsId string) error {
	response, err := lc.post(
		fmt.Sprintf("/services/vps/%v/start", vpsId),
		withHeaders(map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", lc.ApiKey),
		}),
	)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		message, err := io.ReadAll(response.Body)

		if err != nil {
			return err
		}

		return fmt.Errorf("status: %v, message: %s", response.Status, message)
	}

	return nil
}

// GetVPSStateResult represents result of GetVPSState.
type GetVPSStateResult struct {
	Status    string `json:"status"`
	VmUptimeS int    `json:"vmUptimeS"`
}

// GetVPSState allows to get specified VPS state.
func (lc LvlClient) GetVPSState(vpsId string) (*GetVPSStateResult, error) {
	response, err := lc.get(
		fmt.Sprintf("/services/vps/%v/state", vpsId),
		withHeaders(map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", lc.ApiKey),
		}),
	)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		message, err := io.ReadAll(response.Body)

		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("status: %v, message: %s", response.Status, message)
	}

	var result GetVPSStateResult
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// StopVPS allows to stop specified VPS.
func (lc LvlClient) StopVPS(vpsId string) error {
	response, err := lc.post(
		fmt.Sprintf("/services/vps/%v/stop", vpsId),
		withHeaders(map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", lc.ApiKey),
		}),
	)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		message, err := io.ReadAll(response.Body)

		if err != nil {
			return err
		}

		return fmt.Errorf("status: %v, message: %s", response.Status, message)
	}

	return nil
}
