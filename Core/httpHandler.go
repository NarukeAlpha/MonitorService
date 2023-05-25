package Core

import (
	"encoding/json"
	"net/http"
	"sync"
)

func MangaSync(c chan []DbMangaEntry, wg *sync.WaitGroup) {
	defer wg.Done()
	r, err := http.Get("http://localhost:8080/MangaList")
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	var MangaList []DbMangaEntry
	if err = json.NewDecoder(r.Body).Decode(&MangaList); err != nil {
		panic(err)
	}
	c <- MangaList
}

//func ChapterSync(c chan []DbChapterEntry, wg *sync.WaitGroup) {
//	defer wg.Done()
//	r, err := http.Get("http://localhost:8080/ChapterList")
//	if err != nil {
//		panic(err)
//	}
//	defer r.Body.Close()
//	var ChapterList []DbChapterEntry
//	if err = json.NewDecoder(r.Body).Decode(&ChapterList); err != nil {
//		panic(err)
//	}
//	c <- ChapterList
//}
