//
// GO TO: http://localhost:8000/start to use
//

package main

// Example that shows: 1) HTML with upload form 2) handling upload in BE 3) serving local image files
// Using fmt.Fprintf to send html to FE to make code shorter

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

// HTML for the /start page where a HTML FORM with input of type file will be presented.
// Note: the FORM element has the attribute enctype="multipart/form-data" - this makes the browser
// send the data to the Backend in a format that easiest to open with Go's http - is there a better way?
var start_html = `
<form enctype="multipart/form-data" action="/upload" method="POST">
  <label for="myfile">Select a file:</label>
  <input type="file" id="myfile" name="myfile"><br><br>
  <input type="submit" value="Upload">
</form>
`

// Template for the /end page where the link to the uploaded image will be placed in <img>
var end_html = `
<h1>Here's your image:</h1><br><img src="/uploads/%+v">
`

func main() {

	// HTTP handler that will serve (Send to browser) any file in the current directory (./)
	fs := http.FileServer(http.Dir("./"))
	http.HandleFunc("/start", handleStart)
	http.HandleFunc("/upload", handleUpload)
	http.HandleFunc("/end/", handleEnd)
	// This Handle will get all requests to /upload/<something> and return a file named <something>
	// if exists, using the http.FileServer handler we defined above
	http.Handle("/uploads/", http.StripPrefix("/uploads/", fs))
	http.ListenAndServe(":8000", nil)
}

func handleStart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, start_html)
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	var err error
	// FormFile method returns the file parts from POST request
	f, fh, err := r.FormFile("myfile")
	if err != nil {
		fmt.Println("Failed to extract uploaded file from Request")
		panic(err)
	}
	// File name can be extracted from the FileHandler object
	fo, err := os.Create(fh.Filename)
	if err != nil {
		fmt.Println("Failed to create local file.")
		panic(err)
	}
	// The actual file data is read from the File object,
	fileBytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("Failed to extract uploaded file from Request")
		panic(err)
	}
	_, err = fo.Write(fileBytes)
	if err != nil {
		fmt.Println("Failed to write data to local file")
		panic(err)
	}
	http.Redirect(w, r, "/end/"+fh.Filename, http.StatusSeeOther)
}

func handleEnd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	fmt.Fprintf(w, end_html, path.Base(r.RequestURI))
}
