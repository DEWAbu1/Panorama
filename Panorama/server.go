package main

import (
	"Panorama/handler"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// #region handlers
func main() {

	// #region places
	http.HandleFunc("/places", handler.Places)
	http.HandleFunc("/places/insert", handler.HandleInsertPlace)
	http.HandleFunc("/places/delete", handler.HandleDeletePlace)
	http.HandleFunc("/places/update", handler.HandleUpdatePlace)
	// #endregion

	// #region year
	http.HandleFunc("/years", handler.Years)
	http.HandleFunc("/years/insert", handler.HandleInsertYear)
	http.HandleFunc("/years/delete", handler.HandleDeleteYear)
	http.HandleFunc("/years/update", handler.HandleUpdateYear)
	// #endregion

	// #region picture
	http.HandleFunc("/pictures", handler.Pictures)
	http.HandleFunc("/pictures/insert", handler.HandleInsertPicture)
	http.HandleFunc("/pictures/delete", handler.HandleDeletePicture)
	http.HandleFunc("/pictures/update", handler.HandleUpdatePicture)
	// #endregion

	// #region tag
	http.HandleFunc("/tags", handler.Tags)
	http.HandleFunc("/tags/insert", handler.HandleInsertTag)
	http.HandleFunc("/tags/delete", handler.HandleDeleteTag)
	http.HandleFunc("/tags/update", handler.HandleUpdateTag)
	// #endregion

	log.Println("Server is listening on port :28080")
	http.ListenAndServe(":28080", loggingMiddleware(http.DefaultServeMux))
}

// #endregion

// #region log config
const logFileName = "Panorama.log"

var logFile os.File

func init() {
	logFile, err := os.OpenFile("./"+logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("cannot open " + logFileName + ":" + err.Error())
	}
	log.SetOutput(io.MultiWriter(logFile, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

// #endregion

// #region middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		// setupCORS(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

// #endregion

// #region CORS
// func setupCORS(w *http.ResponseWriter, req *http.Request) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:9090")
// 	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
// 	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT")
// 	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Set-Cookie")
// }

// #endregion
