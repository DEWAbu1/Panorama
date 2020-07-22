package pages

import (
	"Panorama/database"
	"Panorama/model"
	"encoding/json"
	"log"
)

func TagData() map[string]interface{} {
	response := map[string]interface{}{"Title": "Tag Data"}

	// get tag data from database and put into interface
	db := database.OpenDB()
	row, err := db.Query("select id,picture,tag,date,xval,yval,theta,phi from tag")
	if err != nil {
		log.Println(err)
	}

	var id int
	var picture, date, tag string
	var xval, yval, theta, phi float64

	tags := []map[string]interface{}{}
	for row.Next() {
		row.Scan(&id, &picture, &tag, &date, &xval, &yval, &theta, &phi)
		tags = append(tags, map[string]interface{}{"id": id, "picture": picture, "tag": tag, "date": date, "xval": xval, "yval": yval, "theta": theta, "phi": phi})
	}

	response["data"] = tags

	return response
}

func InsertTag(param map[string]interface{}) map[string]interface{} {
	response := map[string]interface{}{"Status": "Insert Failed, please send value as singular"}

	mapping, err := json.Marshal(param)
	if err != nil {
		// do error check
		log.Println(err)
		return response
	}

	tag := model.InsertTag{}

	if err := json.Unmarshal(mapping, &tag); err != nil {
		// do error check
		log.Println(err)
		return response
	}

	db := database.OpenDB()
	check3, err := db.Query("select count(*) from picture where picture = ?", tag.Picture)
	defer check3.Close()
	if err != nil {
		log.Println(err)
		return response
	}

	if database.CheckCount(check3) >= 1 { //check if picture exist in table
		_, err = db.Exec("insert into tag (picture,tag,date,xval,yval,theta,phi) values (?,?,?,?,?,?,?)",
			tag.Picture, tag.Tag, tag.Date, tag.Xval, tag.YVal, tag.Theta, tag.Phi)
		if err != nil {
			log.Println(err)
			return response
		}

		response["Status"] = "Insert tag Successful"
		return response
	}
	response["Status"] = "Picture doesn't exist"

	return response
}

func DeleteTag(param interface{}) map[string]interface{} {
	response := map[string]interface{}{"Status": "Delete Failed, please send value as singular"}

	db := database.OpenDB()
	_, err := db.Exec("delete from tag where id = ?", param)
	if err != nil {
		log.Println(err)
	}
	response["Status"] = "Data deleted successfully"

	return response
}

func UpdateTag(param map[string]interface{}) map[string]interface{} {
	response := map[string]interface{}{"Status": "Update Failed, please send value as singular"}

	mapping, err := json.Marshal(param)
	if err != nil {
		// do error check
		log.Println(err)
		return response
	}

	updateTag := model.UpdateTag{}

	if err := json.Unmarshal(mapping, &updateTag); err != nil {
		// do error check
		log.Println(err)
		return response
	}

	db := database.OpenDB()
	check3, err := db.Query("select count(*) from picture where picture = ?", updateTag.Picture)
	defer check3.Close()
	if err != nil {
		log.Println(err)
		return response
	}

	if database.CheckCount(check3) >= 1 { //check if picture exist in table

		_, err = db.Exec("update tag SET picture=?,tag=?,date=?,xval=?,yval=?,theta=?,phi=? where id=?",
			updateTag.Picture, updateTag.Tag, updateTag.Date, updateTag.Xval, updateTag.YVal, updateTag.Theta, updateTag.Phi, updateTag.Id)

		if err != nil {
			log.Println(err)
		}
		response["Status"] = "Data updated successfully"
		return response
	}
	response["Status"] = "Picture doesn't exist"
	return response

}
