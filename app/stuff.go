package app

import (
	"fmt"
)

func DoStuff(MangaG *MG, MangaUrl string) {
	nameLimit := 15
	// fmt.Println("trying to grab manga from:", MangaUrl)
	// time.Sleep(time.Second * 8)
	// MangaG.SaveHtml(MangaUrl, "manga.html")
	// fmt.Println("Saved HTML")

	// fmt.Println("Attempting to detect Manga From Site...")
	// html := MangaG.LoadHtml("manga.html")
	// fmt.Println("Attempting to load HTML from a file...")

	html := MangaG.StringifyHtml(MangaUrl)
	// fmt.Println(html)
	fmt.Println("Html was loaded into memory")

	// fmt.Println("Attempting to retrieve all manga pages from the site.")
	MangaG.NewDir("images")
	newFolderName := string(MangaG.TitleToDirName(MangaG.FindMangaTitle(html)))

	// if newFolderName is longer than nameLimit set newFolderName to the length of nameLimit
	if nameLimit < len(newFolderName) {
		newFolderName = newFolderName[:nameLimit]
	}

	MangaG.NewDir("images/" + newFolderName)

	//	found := MangaG.FindImageUrls(html)
	//	MangaG.DownloadAllWebp(found)
	//  MangaG.DownloadWebp(found[1], "images/001.webp")

	//MangaG.DeleteFile("manga.html")
	//fmt.Println("Deleted HTML no longer needed")

	//MangaG.DeleteFile("images/")
	//fmt.Println("Deleted images no longer needed")
}
