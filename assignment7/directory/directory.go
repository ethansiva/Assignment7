package directory

import (
	"log"
	"net/http"
	"fmt"
)

func DirectoryFunc() {
	dir := http.Dir("/Users/ethansiva/Documents/")
	http.Handle("/", http.directory(dir))

	fmt.Println("Serving on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}