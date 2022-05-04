package main

import (
    "manga-g/app"
)

// Entrypoint for the program.
func main() {
    MangaG := new(app.App)

    //MangaUrl := "https://mangadex.org/chapter/ef0a1ebe-c57d-49f2-bae8-75a64cfbd2a9/1"
    //fmt.Println("Starting MangaG...")
    //	fmt.Println("Please Enter a URL for a Manga's first page to download:")
    //	MangaUrl := MangaG.GetInput()
    //	fmt.Println("trying to grab manga from:", MangaUrl)

    //time.Sleep(time.Second * 8)
    //MangaG.SaveHtml(MangaUrl, "manga.html")
    //fmt.Println("Saved HTML")

    //fmt.Println("Attempting to detect Manga From Site...")
    //html := MangaG.LoadHtml("manga.html")
    //fmt.Println("Attempting to load HTML from a file...")

    //html := MangaG.StringifyHtml(MangaUrl)
    //fmt.Println(html)
    //	fmt.Println("Html was loaded into memory")

    //	fmt.Println("Attempting to retrieve all manga pages from the site.")
    MangaG.NewDir("images")

    //pages, _ := MangaG.GetMangaDex(html)
    //fmt.Println("Pages:", pages)

    //fmt.Println("Manga pages were retrieved.", len(pages))

    //fmt.Println("found[1]:", found[1])
    //	found := MangaG.FindImageUrls(html)
    //	MangaG.DownloadAllWebp(found)
    //MangaG.DownloadWebp(found[1], "images/001.webp")

    //MangaG.DeleteFile("manga.html")
    //fmt.Println("Deleted HTML no longer needed")
    //MangaG.DeleteFile("images/")
    //fmt.Println("Deleted images no longer needed")
}
