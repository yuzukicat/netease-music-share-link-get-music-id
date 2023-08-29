package main

import (
    "fmt"
    "log"
     "os"
     "regexp"
)

func main() {

    // Read raw netease music share links from txt file
    neteasemusicraw, err := os.ReadFile("neteasemusicraw.txt")
    if err != nil {
        log.Fatal(err)
    }

    // Convert to string
    neteasemusicstring := string(neteasemusicraw)
    // fmt.Println(string(neteasemusicraw))

    // Find the pattern we want which is "id=<song id>&userid"
    // (?i): case insensitive, .: matches any single character, *:matches the preceding element zero or more times(in case the shared link do not contain a song id), +: matches the preceding element zero or more times.
    re := regexp.MustCompile("(?i)id=.*(&userid)+")

    // -1 means search for all possible matches.
    found := re.FindAllString(neteasemusicstring, -1)
    // fmt.Printf("%q\n", found)

    if found == nil {
        fmt.Printf("no match found\n")
        os.Exit(1)
    }

    for _, content := range found {
        // replace "id=" as "/music " and "/lyric"
        re1 := regexp.MustCompile("id=")
        replacetomusic := re1.ReplaceAllString(content, "/music ")
        replacetolyric := re1.ReplaceAllString(content, "/lyric ")
        // replace "&userid" as ""
        re2 := regexp.MustCompile("&userid")
        replacedasmusic := re2.ReplaceAllString(replacetomusic, "")
        replacedaslyric := re2.ReplaceAllString(replacetolyric, "")
        // output
        fmt.Printf("%s\n", replacedasmusic)
        fmt.Printf("%s\n", replacedaslyric)
    }
}
