package pages

import (
	"Panorama/database"
	"Panorama/model"
	"encoding/json"
	"log"
)

func PictureData() map[string]interface{} {
	response := map[string]interface{}{"Title": "Pictures Data"}

	// get picture data from database and put into interface
	db := database.OpenDB()
	row, err := db.Query("select id,picture,year,place from picture")
	if err != nil {
		log.Println(err)
	}

	var id, year int
	var picture, place string
	pictures := []map[string]interface{}{}
	for row.Next() {
		row.Scan(&id, &picture, &year, &place)
		pictures = append(pictures, map[string]interface{}{"id": id, "picture": picture, "year": year, "place": place})
	}

	response["data"] = pictures

	return response
}

func InsertPicture(param map[string]interface{}) map[string]interface{} {
	response := map[string]interface{}{"Status": "Insert Failed, please send value as singular JSON"}

	mapping, err := json.Marshal(param)
	if err != nil {
		// do error check
		log.Println(err)
		return response
	}

	picData := model.InsertPicture{}

	if err := json.Unmarshal(mapping, &picData); err != nil {
		// do error check
		log.Println(err)
		return response
	}
	db := database.OpenDB()

	check1, err := db.Query("select count(*) from year where year = ?", picData.Year)
	check2, err := db.Query("select count(*) from place where place = ?", picData.Place)
	defer check1.Close()
	defer check2.Close()
	if err != nil {
		log.Println(err)
	}

	if database.CheckCount(check1) >= 1 {
		if database.CheckCount(check2) >= 1 {
			_, err := db.Exec("insert into picture (year,place,picture) values (?,?,?)", picData.Year, picData.Place, picData.Picture)
			if err != nil {
				log.Println(err)
				defer db.Close()
				return response
			}

			response["Status"] = "Insert Picture Successful"
			return response
		}

		response["Status"] = "This place doesn't exist in table"
		return response
	}
	response["Status"] = "This year doesn't exist in table"
	return response

}

func DeletePicture(param interface{}) map[string]interface{} {
	response := map[string]interface{}{"Status": "Delete Failed, please send value as singular JSON"}

	db := database.OpenDB()
	_, err := db.Exec("delete from picture where picture = ?", param)
	if err != nil {
		log.Println(err)
	}
	response["Status"] = "Data deleted successfully"

	return response
}

func UpdatePicture(param map[string]interface{}) map[string]interface{} {
	response := map[string]interface{}{"Status": "Update Failed, please send value as singular JSON"}

	picture, err := json.Marshal(param)
	if err != nil {
		// do error check
		log.Println(err)
		return response
	}

	structPic := model.UpdatePicture{}

	if err := json.Unmarshal(picture, &structPic); err != nil {
		// do error check
		log.Println(err)
		return response
	}

	db := database.OpenDB()

	check1, err := db.Query("select count(*) from year where year = ?", structPic.Year)
	check2, err := db.Query("select count(*) from place where place = ?", structPic.Place)
	defer check1.Close()
	defer check2.Close()
	if err != nil {
		log.Println(err)
	}

	if database.CheckCount(check1) >= 1 {
		if database.CheckCount(check2) >= 1 {

			_, err := db.Exec("update picture SET year=?,place=?,picture=? where picture=?", structPic.Year, structPic.Place, structPic.Newpicture, structPic.Oldpicture)

			if err != nil {
				log.Println(err)
			}
			response["Status"] = "Data updated successfully"
			return response
		}
		response["Status"] = "Place doesn't exist in table"
		return response
	}
	response["Status"] = "Year doesn't exist in table"
	return response
}
