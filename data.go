package legiondata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const repoURL = "https://raw.githubusercontent.com/zacharyp/sw-legion-data"

// GetData fetches the latest out/legion-data.json file, parses it and returns a Data struct
func GetData() (Data, error) {
	dataURL := repoURL + "/master/lib/legion-data.json"
	var data Data
	client := http.Client{}
	resp, err := client.Get(dataURL)
	if err != nil {
		return data, err
	}

	if resp.StatusCode != 200 {
		return data, fmt.Errorf("non-Ok response code %s (%d)", resp.Status, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}

	_ = json.Unmarshal(body, &data)

	return data, nil
}
