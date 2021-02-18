package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/garcialuis/ActivityCollector/api/responses"
	"github.com/garcialuis/ActivityCollector/api/services"
	"github.com/gorilla/mux"
)

func (server *Server) GetActivitySample(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	originID, err := strconv.ParseUint(vars["origin_id"], 10, 64)

	if err != nil {
		err = fmt.Errorf("OriginID must be a number")
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	activity, _ := services.GetActivitySample(originID)
	fmt.Println(activity)

	responses.JSON(w, http.StatusOK, activity)
}

func (server *Server) GetActivity(w http.ResponseWriter, r *http.Request) {

	var err error

	vars := mux.Vars(r)
	originIDParam := vars["origin_id"]
	activityIDParam := vars["activity_id"]

	if len(originIDParam) < 1 {
		err = fmt.Errorf("Missing parameter: origin_id")
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	} else if len(activityIDParam) < 1 {
		err = fmt.Errorf("Missing parameter: activity_id")
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	originID, err := strconv.ParseUint(originIDParam, 10, 64)

	if err != nil {
		err = fmt.Errorf("OriginID must be a number")
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	activityID, err := strconv.ParseUint(activityIDParam, 10, 64)

	if err != nil {
		err = fmt.Errorf("ActivityID must be a number")
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	activity, err := services.GetActivity(server.DB, originID, activityID)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, activity)
}

func (server *Server) GetActivities(w http.ResponseWriter, r *http.Request) {

	var err error

	vars := mux.Vars(r)
	originIDParam := vars["origin_id"]

	if len(originIDParam) < 1 {
		err = fmt.Errorf("Missing parameter: origin_id")
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	originID, err := strconv.ParseUint(originIDParam, 10, 64)

	if err != nil {
		err = fmt.Errorf("OriginID must be a number")
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	activities, err := services.GetActivities(server.DB, originID)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, activities)
}

func (server *Server) GetActivitiesInRange(w http.ResponseWriter, r *http.Request) {

	var err error

	vars := mux.Vars(r)
	originIDParam := vars["origin_id"]
	timeStartParam, timeEndParam := vars["starttime"], vars["endtime"]

	if len(originIDParam) < 1 || len(timeStartParam) < 1 || len(timeEndParam) < 1 {
		err = fmt.Errorf("Missing one or more parameter: origin_id, starttime, endtime")
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	originID, err := strconv.ParseUint(originIDParam, 10, 64)

	if err != nil {
		err = fmt.Errorf("OriginID must be a number")
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	startUnixTime, err := strconv.ParseInt(timeStartParam, 10, 64)

	if err != nil {
		err = fmt.Errorf("starttime must be a number : unixtime")
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	endUnixTime, err := strconv.ParseInt(timeEndParam, 10, 64)

	if err != nil {
		err = fmt.Errorf("endtime must be a number : unixtime")
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	activities, err := services.GetActivitiesInRange(server.DB, originID, startUnixTime, endUnixTime)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, activities)
}
