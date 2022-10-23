package handler

import (
	"goweb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

// handler on routing dengan function
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Ini kalau di home"))
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Print(err)
		http.Error(w, "Error yah, yaudah sabar", http.StatusInternalServerError)
		return
	}

	data := map[string]string{
		"title":   "Judulnya belum ada",
		"content": "Yutup pesulap merah merona",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err)
		http.Error(w, "Error yah, yaudah sabar", http.StatusInternalServerError)
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello pal, lagi belajar golang yak"))
}

func MasHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello pal, ngapain di mas route path"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	// catch query id
	id := r.URL.Query().Get("id")
	// converte string to number
	idNumb, err := strconv.Atoi(id)

	// validation
	if err != nil && idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	// fmt.Fprintf(w, "Halaman produk %d", idNumb)

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Print(err)
		http.Error(w, "Error di product kan ya", http.StatusInternalServerError)
		return
	}

	// data parsing dari map
	// data := map[string]interface{}{
	// 	"id": idNumb,
	// }

	// data parsing dari struct
	data := entity.Product{ID: 3, Name: "Super73-S2", Price: 4000, Stock: 5}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err)
		http.Error(w, "Error tuh kan", http.StatusInternalServerError)
		return
	}
}
