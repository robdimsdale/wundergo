package oauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/pivotal-golang/lager"
	"github.com/robdimsdale/wundergo"
)

// TaskPositionsForListID returns the positions of all Tasks in the List
// associated with the provided listID.
// The returned TaskPosition.Values might be empty if the Tasks have never been reordered.
func (c oauthClient) TaskPositionsForListID(listID uint) ([]wundergo.Position, error) {
	if listID == 0 {
		return nil, errors.New("listID must be > 0")
	}

	url := fmt.Sprintf(
		"%s/task_positions?list_id=%d",
		c.apiURL,
		listID,
	)

	req, err := c.newGetRequest(url)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected response code %d - expected %d", resp.StatusCode, http.StatusOK)
	}

	taskPositions := []wundergo.Position{}
	err = json.NewDecoder(resp.Body).Decode(&taskPositions)
	if err != nil {
		c.logger.Debug("", lager.Data{"response": resp})
		return nil, err
	}
	return taskPositions, nil
}

// TaskPosition returns the TaskPosition associated with the provided taskPositionID.
func (c oauthClient) TaskPosition(taskPositionID uint) (wundergo.Position, error) {
	if taskPositionID == 0 {
		return wundergo.Position{}, errors.New("taskPositionID must be > 0")
	}

	url := fmt.Sprintf(
		"%s/task_positions/%d",
		c.apiURL,
		taskPositionID,
	)

	req, err := c.newGetRequest(url)
	if err != nil {
		return wundergo.Position{}, err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return wundergo.Position{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return wundergo.Position{}, fmt.Errorf("Unexpected response code %d - expected %d", resp.StatusCode, http.StatusOK)
	}

	taskPosition := wundergo.Position{}
	err = json.NewDecoder(resp.Body).Decode(&taskPosition)
	if err != nil {
		c.logger.Debug("", lager.Data{"response": resp})
		return wundergo.Position{}, err
	}
	return taskPosition, nil
}

// UpdateTaskPosition updates the provided TaskPosition.
// This will reorder the Tasks.
func (c oauthClient) UpdateTaskPosition(taskPosition wundergo.Position) (wundergo.Position, error) {
	body, err := json.Marshal(taskPosition)
	if err != nil {
		return wundergo.Position{}, err
	}

	url := fmt.Sprintf(
		"%s/task_positions/%d",
		c.apiURL,
		taskPosition.ID,
	)

	req, err := c.newPatchRequest(url, body)
	if err != nil {
		return wundergo.Position{}, err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return wundergo.Position{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return wundergo.Position{}, fmt.Errorf("Unexpected response code %d - expected %d", resp.StatusCode, http.StatusOK)
	}

	returnedTaskPosition := wundergo.Position{}
	err = json.NewDecoder(resp.Body).Decode(&returnedTaskPosition)
	if err != nil {
		c.logger.Debug("", lager.Data{"response": resp})
		return wundergo.Position{}, err
	}
	return returnedTaskPosition, nil
}