package Core

import (
	"encoding/json"
	"net/http"
)

func ServerSync() ([]DbMangaEntry, []DbChapterEntry) {
	r, err := http.Get("http://localhost:8080/MangaList")
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	var MangaList []DbMangaEntry
	if err = json.NewDecoder(r.Body).Decode(&MangaList); err != nil {
		panic(err)
	}
	r2, err := http.Get("http://localhost:8080/ChapterList")
	if err != nil {
		panic(err)

	}
	defer r2.Body.Close()
	var ChapterList []DbChapterEntry
	if err = json.NewDecoder(r2.Body).Decode(&ChapterList); err != nil {
		panic(err)

	}
	return MangaList, ChapterList
}
