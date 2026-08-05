package web_page_utils

import (
	"fmt"
	"net/http"
)

func HandlePreloadLinks(w http.ResponseWriter, links []string) {
	for i := 0; i < len(links); i++ {
		link := fmt.Sprintf("<%s>; rel=preload; as=style", links[i])
		w.Header().Add("Link", link)
	}
	w.WriteHeader(103)
}
