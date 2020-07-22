package handler

import (
	"Panorama/pages"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/bitly/go-simplejson"
)

// #region handlers
func Pictures(w http.ResponseWriter, r *http.Request) {
	response := pages.PictureData()
	jsonResponse(w, response)
}

func HandleInsertPicture(w http.ResponseWriter, r *http.Request) { // this function no upload file yet, for now only input name only
	jsonData := bodyToJson(r)

	year := assertInt(getParameter(jsonData, "year"))
	place := assertString(getParameter(jsonData, "place"))
	picture := assertString(getParameter(jsonData, "picture"))

	//still lacking upload image into projects folder
	param := map[string]interface{}{"year": year, "place": place, "picture": picture}
	response := pages.InsertPicture(param)
	jsonResponse(w, response)

}

func HandleDeletePicture(w http.ResponseWriter, r *http.Request) {
	jsonData := bodyToJson(r)

	picture := assertString(getParameter(jsonData, "picture"))
	// still lacking deleting file in folder, for now just delete the filename that puts into the database
	response := pages.DeletePicture(picture)
	jsonResponse(w, response)
}

func HandleUpdatePicture(w http.ResponseWriter, r *http.Request) {
	jsonData := bodyToJson(r)

	year := assertInt(getParameter(jsonData, "year"))
	place := assertString(getParameter(jsonData, "place"))
	oldPicture := assertString(getParameter(jsonData, "oldpicture"))
	newPicture := assertString(getParameter(jsonData, "newpicture"))

	//still lacking update image into projects folder and deleting the older file
	param := map[string]interface{}{"oldpicture": oldPicture, "newpicture": newPicture, "year": year, "place": place}
	response := pages.UpdatePicture(param)
	jsonResponse(w, response)

}

func Tags(w http.ResponseWriter, r *http.Request) {
	response := pages.TagData()
	jsonResponse(w, response)
}
func HandleDeleteTag(w http.ResponseWriter, r *http.Request) {
	jsonData := bodyToJson(r)

	id := assertInt(getParameter(jsonData, "id"))

	response := pages.DeleteTag(id)
	jsonResponse(w, response)
}

func HandleInsertTag(w http.ResponseWriter, r *http.Request) {
	jsonData := bodyToJson(r)

	picture := assertString(getParameter(jsonData, "picture"))
	tag := assertString(getParameter(jsonData, "tag"))
	date := assertString(getParameter(jsonData, "date"))
	xval := assertFloat64(getParameter(jsonData, "xval"))
	yval := assertFloat64(getParameter(jsonData, "yval"))
	theta := assertFloat64(getParameter(jsonData, "theta"))
	phi := assertFloat64(getParameter(jsonData, "phi"))
	param := map[string]interface{}{"picture": picture, "tag": tag, "date": date, "xval": xval, "yval": yval, "theta": theta, "phi": phi}

	response := pages.InsertTag(param)
	jsonResponse(w, response)
}

func HandleUpdateTag(w http.ResponseWriter, r *http.Request) {
	jsonData := bodyToJson(r)

	id := assertInt(getParameter(jsonData, "id"))
	tag := assertInt(getParameter(jsonData, "tag"))
	picture := assertString(getParameter(jsonData, "picture"))
	date := assertString(getParameter(jsonData, "date"))
	xval := assertFloat64(getParameter(jsonData, "xval"))
	yval := assertFloat64(getParameter(jsonData, "yval"))
	theta := assertFloat64(getParameter(jsonData, "theta"))
	phi := assertFloat64(getParameter(jsonData, "phi"))
	param := map[string]interface{}{"id": id, "picture": picture, "tag": tag, "date": date, "xval": xval, "yval": yval, "theta": theta, "phi": phi}

	response := pages.UpdateTag(param)
	jsonResponse(w, response)
}

func Places(w http.ResponseWriter, r *http.Request) {
	response := pages.PlaceData()
	jsonResponse(w, response)
}

func HandleInsertPlace(w http.ResponseWriter, r *http.Request) {
	jsonData := bodyToJson(r)

	place := assertString(getParameter(jsonData, "place"))

	response := pages.InsertPlace(place)
	jsonResponse(w, response)
}

func HandleDeletePlace(w http.ResponseWriter, r *http.Request) {
	jsonData := bodyToJson(r)

	place := assertString(getParameter(jsonData, "place"))

	response := pages.DeletePlace(place)
	jsonResponse(w, response)
}

func HandleUpdatePlace(w http.ResponseWriter, r *http.Request) {
	jsonData := bodyToJson(r)

	oldPlace := assertString(getParameter(jsonData, "oldplace"))
	newPlace := assertString(getParameter(jsonData, "newplace"))
	param := map[string]interface{}{"oldplace": oldPlace, "newplace": newPlace}

	response := pages.UpdatePlace(param)
	jsonResponse(w, response)
}

func Years(w http.ResponseWriter, r *http.Request) {
	response := pages.YearData()
	jsonResponse(w, response)
}

func HandleInsertYear(w http.ResponseWriter, r *http.Request) {
	jsonData := bodyToJson(r)

	year := assertInt(getParameter(jsonData, "year"))

	response := pages.InsertYear(year)
	jsonResponse(w, response)
}

func HandleDeleteYear(w http.ResponseWriter, r *http.Request) {
	jsonData := bodyToJson(r)

	year := assertInt(getParameter(jsonData, "year"))

	response := pages.DeleteYear(year)
	jsonResponse(w, response)
}

func HandleUpdateYear(w http.ResponseWriter, r *http.Request) {
	jsonData := bodyToJson(r)

	oldYear := assertInt(getParameter(jsonData, "oldyear"))
	newYear := assertInt(getParameter(jsonData, "newyear"))
	param := map[string]interface{}{"oldyear": oldYear, "newyear": newYear}

	response := pages.UpdateYear(param)
	jsonResponse(w, response)
}

// #endregion

// #region utilities
func jsonResponse(w http.ResponseWriter, data map[string]interface{}) { // write response interface as json
	response, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func getParameter(jsonData *simplejson.Json, key string) (parameter interface{}) { // getting parameter from jsonData
	//possible improvement with goroutines?
	str, err := jsonData.Get(key).String() // get string param
	if err == nil {
		parameter = str
		return parameter
	}

	strArry, err := jsonData.Get(key).StringArray() // get array of string param
	if err == nil {
		parameter = strArry
		return parameter
	}

	i, err := jsonData.Get(key).Int() // get int param
	if err == nil {
		parameter = i
		return parameter
	}

	float, err := jsonData.Get(key).Float64() // get float64 param
	if err == nil {
		parameter = float
		return parameter
	}

	arrayIntf, err := jsonData.Get(key).Array() // get any array
	if err == nil {
		_, err = arrayIntf[0].(json.Number).Int64() // check if array type is int64
		if err == nil {
			var intParam []int64
			for _, v := range arrayIntf {
				value, _ := v.(json.Number).Int64()
				intParam = append(intParam, value)
			}
			parameter = intParam
			return parameter
		}
		_, err = arrayIntf[0].(json.Number).Float64() // check if array type is float64
		if err == nil {
			var floatParam []float64
			for _, v := range arrayIntf {
				value, _ := v.(json.Number).Float64()
				floatParam = append(floatParam, value)
			}
			parameter = floatParam
			return parameter
		}
	}

	intf := jsonData.Get(key).Interface() // get interface param
	if parameter, ok := intf.([]string); ok {
		parameter = intf.([]string)
		return parameter
	}

	parameter = "" // default
	return parameter
}

func bodyToJson(r *http.Request) *simplejson.Json { // read all request body and turn it into json

	buf, _ := ioutil.ReadAll(r.Body)
	rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
	rdr3 := ioutil.NopCloser(bytes.NewBuffer(buf))

	json, err := simplejson.NewFromReader(rdr1)
	if err != nil {
		//
	} else if json != nil {
		r.Body.Close()
		r.Body = rdr2
		return json
	}
	r.Body.Close()
	r.Body = rdr2

	json = simplejson.New()

	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}

	for key, values := range r.Form {
		json.Set(key, values)
	}

	r.Body.Close()
	r.Body = rdr3

	return json
}

func assertString(param interface{}) (asserted string) {
	switch param.(type) {
	case string: // handling json input
		asserted = param.(string)
		return asserted
	case []string: // handling form-data input
		asserted = param.([]string)[0]
		return asserted
	default:
		return asserted
	}
}

func assertInt(param interface{}) (asserted int) {
	switch param.(type) {
	case int: // handling json input
		asserted = param.(int)
	case []string: // handling form-data input
		asserted, _ = strconv.Atoi(param.([]string)[0])
	default:
		return asserted
	}
	return asserted
}

func assertFloat64(param interface{}) (asserted float64) {
	switch param.(type) {
	case float64: // handling json input
		asserted = param.(float64)
		return asserted
	case []string: // handling form-data input
		asserted, _ = strconv.ParseFloat(param.([]string)[0], 64)
		return asserted
	default:
		return asserted
	}
}

// #endregion
