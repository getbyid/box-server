package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type siteHandler struct {
	site *Site
}

func (h siteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path)

	// подставляем индексный файл здесь, т.к. нужен для Content-Type
	p := r.URL.Path
	if strings.HasSuffix(p, "/") {
		p += h.site.index
	}

	rc, err := h.site.ContentByPath(p)
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}

	defer rc.Close()

	contentType := textContentType(p)

	buffer := make([]byte, 4096)
	var size int
	for {
		n, err := rc.Read(buffer)
		if err == io.EOF && n == 0 {
			break
		}

		if err != io.EOF && err != nil {
			http.Error(w, "Failed to read data", http.StatusInternalServerError)
			return
		}

		if size == 0 {
			if contentType == "" {
				contentType = http.DetectContentType(buffer)
			}
			w.Header().Set("Content-Type", contentType)

			// кеширование на стороне клиента позволяет избежать лишних запросов
			w.Header().Set("Cache-Control", "max-age=3600, public")
		}

		_, err = w.Write(buffer[:n])
		if err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}

		size += n
	}
	log.Printf("--> %s %d bytes\n", contentType, size)
}

func main() {
	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "Usage: %s website.zip\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}

	port := flag.Int("port", 8080, "listen port")
	index := flag.String("index", "index.html", "file for root path")

	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		return
	}

	site := NewSite(*index)
	site.LoadFromZip(flag.Arg(0))

	address := fmt.Sprintf("localhost:%d", *port)
	server := &http.Server{Addr: address, Handler: siteHandler{site}}

	log.Printf("Server is starting on http://%s\n", address)
	log.Fatal(server.ListenAndServe())
}
