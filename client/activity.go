package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	client_models "github.com/garcialuis/ActivityCollector/client/models"
)

// GetActivity fetches and returns an activity with a specified originID from the database
func (activity *ActivityCollector) GetActivity(originID uint64) (Activity client_models.Activity) {

	reqURL := fmt.Sprint(activity.config.hostURL, "activity/")
	base, err := url.Parse(reqURL)
	if err != nil {
		log.Println("ActivityClient: Unable to parse reqURL: ", err.Error())
		return Activity
	}

	base.Path += strconv.FormatUint(originID, 64)

	resp, err := http.Get(base.String())
	if err != nil {
		log.Println("ActivityClient: Unable to complete request due to: ", err.Error())
		return Activity
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &Activity)
	if err != nil {
		log.Println("ActivityClient: Unable to unmarshall data retrieved: ", err.Error())
		return Activity
	}

	return Activity
}
