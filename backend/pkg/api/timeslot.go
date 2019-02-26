package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
)

// timeslotAPI holds all of the api functions related to Timeslots and all of the variables needed to access the backend
type timeslotAPI struct {
	timeslotReader  db.TimeslotReader
	timeslotWriter  db.TimeslotWriter
	timeslotUpdater db.TimeslotUpdater
	timeslotDeleter db.TimeslotDeleter
}

// CreateTimeslotRoutes makes all of the routes for room related calls
func CreateTimeslotRoutes(timeslotDBFacade db.TimeslotReaderWriterUpdaterDeleter) []Route {
	tAPI := timeslotAPI{
		timeslotReader:  timeslotDBFacade,
		timeslotWriter:  timeslotDBFacade,
		timeslotUpdater: timeslotDBFacade,
		timeslotDeleter: timeslotDBFacade,
	}

	routes := []Route{
		NewRoute("/api/v1/timeslot", tAPI.writeATimeslot, "POST"),
		NewRoute("/api/v1/timeslot", tAPI.updateATimeslot, "PUT"),
		NewRoute("/api/v1/timeslot", tAPI.deleteATimeslot, "DELETE"),
	}

	return routes
}

// writeATimeslotRequest request for writeATimeslot
type writeATimeslotRequest struct {
	StartTime string `example:"2019-02-18 21:00:00"`
	EndTime   string `example:"2019-10-01 23:00:00"`
}

// writeATimeslot Add a timeslot to the db
// @Summary Add a timeslot
// @Description Add a timeslot to the db
// @accept json
// @produce json
// @param timeslot body api.writeATimeslotRequest true "the timeslot to add"
// @Success 200 {} int "the id of the timeslot added"
// @Failure 400 {} string "the request was bad"
// @Failure 503 {} string "failed to access the db"
// @Router /api/v1/timeslot [post]
func (t timeslotAPI) writeATimeslot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	j, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := "unable to read body"
		log.Printf("%s: %v", msg, err)
		w.WriteHeader(http.StatusBadRequest)
		response, _ := json.Marshal(msg)
		w.Write(response)
		return
	}

	timeslotRequest := writeATimeslotRequest{}
	json.Unmarshal(j, &timeslotRequest)

	log.Println(timeslotRequest.StartTime)
	startTime, err := time.Parse(db.TimeFormat, timeslotRequest.StartTime)
	if err != nil {
		msg := fmt.Sprintf("start time invalid. please use the format %s", db.TimeFormat)
		log.Printf("%s: %v", msg, err)
		w.WriteHeader(http.StatusBadRequest)
		response, _ := json.Marshal(msg)
		w.Write(response)
		return
	}

	endTime, err := time.Parse(db.TimeFormat, timeslotRequest.EndTime)
	if err != nil {
		msg := fmt.Sprintf("end time invalid. please use the format %s", db.TimeFormat)
		log.Printf("%s: %v", msg, err)
		w.WriteHeader(http.StatusBadRequest)
		response, _ := json.Marshal(msg)
		w.Write(response)
		return
	}

	timeslot := db.Timeslot{
		StartTime: &startTime,
		EndTime:   &endTime,
	}

	id, err := t.timeslotWriter.WriteATimeslot(timeslot)
	if err != nil {
		msg := "failed to access the db"
		log.Printf("%s: %v", msg, err)
		w.WriteHeader(http.StatusServiceUnavailable)
		response, _ := json.Marshal(msg)
		w.Write(response)
		return
	}

	response := map[string]int64{"id": id}
	responseJSON, _ := json.Marshal(response)

	w.Write(responseJSON)
}

// writeATimeslotRequest request for updateATimeslot
type updateATimeslotRequest struct {
	ID        int64  `example:"1"`
	StartTime string `example:"2019-02-18 21:00:00"`
	EndTime   string `example:"2019-10-01 23:00:00"`
}

// updateATimeslot Update an existing timeslot in the db
// @Summary Update an existing timeslot in the db
// @Description Update an existing timeslot in the db
// @accept json
// @produce json
// @param timeslot body api.updateATimeslotRequest true "the timeslot to update with the new values"
// @Success 200 "Updated properly"
// @Failure 400 {} string "the request was bad"
// @Failure 503 {} string "failed to access the db"
// @Router /api/v1/timeslot [put]
func (t timeslotAPI) updateATimeslot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	j, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := "unable to read body"
		log.Printf("%s: %v", msg, err)
		w.WriteHeader(http.StatusBadRequest)
		response, _ := json.Marshal(msg)
		w.Write(response)
		return
	}

	timeslotRequest := updateATimeslotRequest{}
	json.Unmarshal(j, &timeslotRequest)

	startTime, err := time.Parse(db.TimeFormat, timeslotRequest.StartTime)
	if err != nil {
		msg := fmt.Sprintf("start time invalid. please use the format %s", db.TimeFormat)
		log.Printf("%s: %v", msg, err)
		w.WriteHeader(http.StatusBadRequest)
		response, _ := json.Marshal(msg)
		w.Write(response)
		return
	}

	endTime, err := time.Parse(db.TimeFormat, timeslotRequest.EndTime)
	if err != nil {
		msg := fmt.Sprintf("end time invalid. please use the format %s", db.TimeFormat)
		log.Printf("%s: %v", msg, err)
		w.WriteHeader(http.StatusBadRequest)
		response, _ := json.Marshal(msg)
		w.Write(response)
		return
	}

	timeslot := db.Timeslot{
		ID:        &timeslotRequest.ID,
		StartTime: &startTime,
		EndTime:   &endTime,
	}

	err = t.timeslotUpdater.UpdateATimeslot(timeslot)
	if err != nil {
		msg := ""
		switch err {
		case db.ErrNothingChanged:
			msg = "nothing in the db was changed. id probably does not exist"
			w.WriteHeader(http.StatusBadRequest)
		default:
			msg = "failed to access the db"
			w.WriteHeader(http.StatusServiceUnavailable)
		}
		log.Printf("%s: %v", msg, err)
		response, _ := json.Marshal(msg)
		w.Write(response)
		return
	}

	w.Write(nil)
}

// writeATimeslotRequest request for deleteATimeslot
type deleteATimeslotRequest struct {
	ID int64 `example:"1"`
}

// deleteATimeslot Delete an existing timeslot in the db
// @Summary Delete an existing timeslot in the db
// @Description Delete an existing timeslot in the db
// @accept json
// @produce json
// @param timeslot body api.deleteATimeslotRequest true "the timeslot to delete"
// @Success 200 "Deleted properly"
// @Failure 400 {} string "the request was bad"
// @Failure 503 {} string "failed to access the db"
// @Router /api/v1/timeslot [delete]
func (t timeslotAPI) deleteATimeslot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	j, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := "unable to read body"
		log.Printf("%s: %v", msg, err)
		w.WriteHeader(http.StatusBadRequest)
		response, _ := json.Marshal(msg)
		w.Write(response)
		return
	}

	timeslotRequest := deleteATimeslotRequest{}
	json.Unmarshal(j, &timeslotRequest)

	err = t.timeslotDeleter.DeleteATimeslot(timeslotRequest.ID)
	if err != nil {
		msg := ""
		switch err {
		case db.ErrNothingChanged:
			msg = "nothing in the db was changed. id probably does not exist"
			w.WriteHeader(http.StatusBadRequest)
		default:
			msg = "failed to access the db"
			w.WriteHeader(http.StatusServiceUnavailable)
		}
		log.Printf("%s: %v", msg, err)
		response, _ := json.Marshal(msg)
		w.Write(response)
		return
	}

	w.Write(nil)
}
