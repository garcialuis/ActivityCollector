package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	client_models "github.com/garcialuis/ActivityCollector/client/models"
)

// GetActivity fetches and returns an activity with a specified originID from the database
func (activity *ActivityCollector) GetActivity(originID uint64, activityID uint64) (Activity client_models.Activity) {

	reqURL := fmt.Sprint(activity.config.hostURL, "activity/")
	base, err := url.Parse(reqURL)
	if err != nil {
		log.Println("ActivityClient: Unable to parse reqURL: ", err.Error())
		return Activity
	}

	base.Path += fmt.Sprintf("%d/%d", originID, activityID)

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

func (activity *ActivityCollector) GetActivities(originID uint64) (Activities client_models.Activities) {

	reqURL := fmt.Sprint(activity.config.hostURL, "activities/")
	base, err := url.Parse(reqURL)
	if err != nil {
		log.Println("ActivityClient: Unable to parse reqURL: ", err.Error())
		return Activities
	}

	base.Path += fmt.Sprintf("%d", originID)

	resp, err := http.Get(base.String())
	if err != nil {
		log.Println("ActivityClient: Unable to complete request due to: ", err.Error())
		return Activities
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &Activities)
	if err != nil {
		log.Println("ActivityClient: Unable to unmarshall Activities data retrieved: ", err.Error())
		return Activities
	}

	return Activities
}

func (activity *ActivityCollector) GetActivitiesInRange(originID uint64, starttime uint64, endtime uint64) (Activities client_models.Activities) {

	reqURL := fmt.Sprint(activity.config.hostURL, "activities/")
	base, err := url.Parse(reqURL)
	if err != nil {
		log.Println("ActivityClient: Unable to parse reqURL: ", err.Error())
		return Activities
	}

	base.Path += fmt.Sprintf("%d/%d/%d", originID, starttime, endtime)

	resp, err := http.Get(base.String())
	if err != nil {
		log.Println("ActivityClient: Unable to complete request due to: ", err.Error())
		return Activities
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &Activities)
	if err != nil {
		log.Println("ActivityClient: Unable to unmarshall Activities data retrieved: ", err.Error())
		return Activities
	}

	return Activities
}
