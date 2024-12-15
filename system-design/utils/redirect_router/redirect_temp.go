package redirect_router

import (
	"errors"
	"fmt"
	"net/http"
)

type Redirect struct {
	ShortLink string `json:"shortlink"`
	FullLink  string `json:"fullLink"`
}

type RedirectList []Redirect

var routerList RedirectList

func (rl RedirectList) GetFullLink(key string) (string, error) {
	for _, value := range rl {
		if value.ShortLink == key {
			return value.FullLink, nil
		}
	}

	return "", errors.New("cannot find short link")
}

func redirectHandler(w http.ResponseWriter, r *http.Request, location string) {
	http.Redirect(w, r, location, http.StatusFound)
}

func shortLinkUrlHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	fullLink, err := routerList.GetFullLink(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	redirectHandler(w, r, fullLink)
}

func createRouterList() {
	routerList = RedirectList{
		Redirect{
			"/mail",
			"https://mail.google.com",
		},
		Redirect{
			"/ll",
			"https://www.linkedin.com/feed/",
		},
		Redirect{
			"/yt",
			"https://www.youtube.com",
		},
	}
}

func routerHandler() {
	for _, value := range routerList {
		http.HandleFunc(value.ShortLink, shortLinkUrlHandler)
	}
}

func ListenAndServe() {
	createRouterList()
	routerHandler()
	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		fmt.Errorf("failed to start server at port 9001")
		return
	}
}
