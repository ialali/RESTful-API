// package handlers

// import (
// 	"html/template"
// 	"net/http"
// )

// func HandleHomePage(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		http.Error(w, "404 Page Not Found", http.StatusNotFound)
// 		return
// 	}

// 	data := map[string]interface{}{
// 		"Artists": locations,
// 	}
// 	// Serve the HTML page with the filtered artists
// 	tmpl := template.Must(template.ParseFiles("templates/template.html"))
// 	err := tmpl.Execute(w, data)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }
