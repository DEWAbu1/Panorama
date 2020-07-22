package pages

import (
	"Panorama/database"
	"Panorama/model"
	"encoding/json"
	"log"
)

func PlaceData() map[string]interface{} {
	response := map[string]interface{}{"Title": "Places Data"}

	// get places datafrom database and put into interface
	db := database.OpenDB()
	row, err := db.Query("select id,place from place")
	if err != nil {
		log.Println(err)
	}

	var id int
	var place string
	places := []map[string]interface{}{}
	for row.Next() {
		row.Scan(&id, &place)
		places = append(places, map[string]interface{}{"id": id, "place": place})
	}

	response["data"] = places

	return response
}

func InsertPlace(param interface{}) map[string]interface{} {
	response := map[string]interface{}{"Status": "Insert Failed, please send value as singular"}

	db := database.OpenDB()

	row, err := db.Query("select count(*) from place where place = ?", param)
	defer row.Close()
	if err != nil {
		log.Println(err)
	}

	if database.CheckCount(row) < 1 {

		_, err := db.Exec("insert into place (place) values (?)", param)
		if err != nil {
			log.Println(err)
			defer db.Close()
			return response
		}

		response["Status"] = "Insert place Successful"
		return response
	}
	response["Status"] = "This place already exist"
	return response

}

func DeletePlace(param interface{}) map[string]interface{} {
	response := map[string]interface{}{"Status": "Delete Failed, please send value as singular"}

	db := database.OpenDB()
	_, err := db.Exec("delete from place where place = ?", param)
	if err != nil {
		log.Println(err)
	}
	response["Status"] = "Data deleted successfully"

	return response
}

func UpdatePlace(param map[string]interface{}) map[string]interface{} {
	response := map[string]interface{}{"Status": "Update Failed, please send value as singular"}

	place, err := json.Marshal(param)
	if err != nil {
		// do error check
		log.Println(err)
		return response
	}

	structPlace := model.UpdatePlace{}

	if err := json.Unmarshal(place, &structPlace); err != nil {
		// do error check
		log.Println(err)
		return response
	}

	db := database.OpenDB()

	row, err := db.Query("select count(*) from place where place = ?", structPlace.Oldplace)
	defer row.Close()
	if err != nil {
		log.Println(err)
	}

	if database.CheckCount(row) < 1 {

		_, err := db.Exec("update  place SET place=? where place=?", structPlace.Newplace, structPlace.Oldplace)

		if err != nil {
			log.Println(err)
		}
		response["Status"] = "Data updated successfully"
	}
	response["Status"] = "Place doesn't exist"
	return response
}
