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

	activity, err := services.GetActivity(server.DB, originID)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, activity)
}
