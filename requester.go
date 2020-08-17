package requester

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func patchJSON(url string, params url.Values, result interface{}) error {
	request, err := http.NewRequest(
		http.MethodPatch,
		url,
		strings.NewReader(params.Encode()),
	)
	if err != nil {
		return fmt.Errorf("Response error: %w", err)
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Bad status: %d", response.StatusCode)
	}
	if result == nil {
		return nil
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("Can't read response %w", err)
	}
	err = json.Unmarshal(body, result)
	if err != nil {
		return fmt.Errorf("Can't parse response %w", err)
	}
	return nil
}
