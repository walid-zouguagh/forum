package utils

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func ExecuteTemplate(w http.ResponseWriter, pages []string, data any) {
	base_pages := []string{
		"views/components/navbar.html",
		"views/components/footer.html",
		"views/base.html",
	}
	pages = append(pages, base_pages...)
	tmpl, err := template.ParseFiles(pages...)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), "sorry but there are Error in server try next time", err)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), "sorry but there are Error in server try next time", err)
		fmt.Printf("html page template error: %v", err.Error())
		return
	}
}

func ErrorHandler(w http.ResponseWriter, statusCode int, msg1, msg2 string, err error) {
	// print errors in case of intenal server error
	if err != nil && statusCode == 500 {
		log.Println(err)
	}

	Error := ErrorData{
		Msg1:       msg1,
		Msg2:       msg2,
		StatusCode: statusCode,
	}

	tmpl, err := template.ParseFiles("./views/pages/error.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, msg1, statusCode)
		return
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, Error); err != nil {
		fmt.Println(err)
		http.Error(w, msg1, statusCode)
		return
	}
	w.WriteHeader(statusCode)
	// If successful, write the buffer content to the ResponseWriter
	buf.WriteTo(w)
}
