package pages

import (
	"Panorama/database"
	"Panorama/model"
	"encoding/json"
	"log"
)

func YearData() map[string]interface{} {
	response := map[string]interface{}{"Title": "Year Data"}
	// get year data from database and put into interface

	db := database.OpenDB()

	row, err := db.Query("select id,year from year")
	if err != nil {
		log.Println(err)
	}
	defer row.Close()

	var id, year int
	years := []map[string]int{}
	for row.Next() {
		row.Scan(&id, &year)
		years = append(years, map[string]int{"id": id, "year": year})
	}

	response["data"] = years

	return response
}

func InsertYear(param interface{}) map[string]interface{} {
	response := map[string]interface{}{"Status": "Insert Failed, please send value as singular"}

	db := database.OpenDB()

	row, err := db.Query("select count(*) from year where year = ?", param)
	defer row.Close()
	if err != nil {
		log.Println(err)
	}

	if database.CheckCount(row) < 1 {

		_, err := db.Exec("insert into year (year) values (?)", param)
		if err != nil {
			log.Println(err)
			defer db.Close()
			return response
		}

		response["Status"] = "Insert Year Successful"
		return response
	}
	response["Status"] = "This year already exist"
	return response

}

func DeleteYear(param interface{}) map[string]interface{} {
	response := map[string]interface{}{"Status": "Delete Failed, please send value as singular"}

	db := database.OpenDB()
	_, err := db.Exec("delete from year where year = ?", param)
	if err != nil {
		log.Println(err)
	}
	response["Status"] = "Data deleted successfully"

	return response
}

func UpdateYear(param map[string]interface{}) map[string]interface{} {
	response := map[string]interface{}{"Status": "Update Failed, please send value as singular"}

	year, err := json.Marshal(param)
	if err != nil {
		// do error check
		log.Println(err)
		return response
	}

	structYear := model.UpdateYear{}

	if err := json.Unmarshal(year, &structYear); err != nil {
		// do error check
		log.Println(err)
		return response
	}

	db := database.OpenDB()

	row, err := db.Query("select count(*) from year where year = ?", structYear.Oldyear)
	defer row.Close()
	if err != nil {
		log.Println(err)
	}

	if database.CheckCount(row) < 1 {

		_, err := db.Exec("update  year SET year=? where year=?", structYear.Newyear, structYear.Oldyear)

		if err != nil {
			log.Println(err)
		}
		response["Status"] = "Data updated successfully"
	}
	response["Status"] = "Year doesn't exist"
	return response
}
