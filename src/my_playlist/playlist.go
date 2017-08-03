package main

import (
    "os"
    "strings"
    "fmt"
    "path/filepath"
    "io/ioutil"
    "log"
    "strconv"
)

type Song struct {
    Title       string
    Filename    string
    Seconds     int
}

func main() {
    if len(os.Args) == 1 || !(strings.HasSuffix(os.Args[1], ".m3u") || strings.HasSuffix(os.Args[1], ".pls")) {
        fmt.Printf("usage: %s <file.[m3u|pls]>\n", filepath.Base(os.Args[0]))
        os.Exit(1)
    }
    
    if strings.HasSuffix(os.Args[1], ".m3u") {
        if rawBytes, err := ioutil.ReadFile(os.Args[1]); err != nil {
            log.Fatal(err)
        } else {
            songs := readM3uPlaylist(string(rawBytes))
            writePlsPlaylist(songs)
        }
    }
    
    if strings.HasSuffix(os.Args[1], ".pls") {
        if rawBytes, err := ioutil.ReadFile(os.Args[1]); err != nil {
            log.Fatal(err)
        } else {
            songs := readPlsPlaylist(string(rawBytes))
            writeM3uPlaylist(songs)
        }
    }
}

func readPlsPlaylist(data string) (songs []Song) {
    var song Song
    var arr []string
    var err error
    
    for _, line := range strings.Split(data, "\n") {
        arr = append(arr, line)
    }
    
    for i := 1; i < len(arr) - 2; i+=3 {
        song.Filename = strings.Map(mapPlatformDirSeparator, arr[i][strings.Index(arr[i], "=") + 1:])
        song.Title = arr[i+1][strings.Index(arr[i+1], "=") + 1:]
        if song.Seconds, err = strconv.Atoi(arr[i+2][strings.Index(arr[i+2], "=") + 1:]); err != nil {
            log.Printf("failed to read the duration for '%s': %v\n", song.Title, err)
                song.Seconds = -1
        }
        songs = append(songs, song)
    }
    return songs
}

func writeM3uPlaylist(songs []Song) {
    fmt.Println("#EXTM3U")
    for i, song := range songs {
        i++
        fmt.Printf("#EXTINF:%d,%s\n", song.Seconds, song.Title)
        fmt.Printf("%s\n", song.Filename)
    }
}

func readM3uPlaylist(data string) (songs []Song) {
    var song Song
    for _, line := range strings.Split(data, "\n") {
        line = strings.TrimSpace(line)
        
        if line == "" || strings.HasPrefix(line, "#EXTM3U") {
            continue
        }
        
        if strings.HasPrefix(line, "#EXTINF:") {
            song.Title, song.Seconds = parseExtinfLine(line)
        } else {
            song.Filename = strings.Map(mapPlatformDirSeparator, line)
        }
        
        if song.Filename != "" && song.Title != "" && song.Seconds != 0 {
            songs = append(songs, song)
            song = Song{}
        }
    }
    return songs
}

func parseExtinfLine(line string) (title string, seconds int) {
    if i := strings.IndexAny(line, "-0123456789"); i > -1 {
        const separator  = ","
        line = line[i:]
        if j := strings.Index(line, separator); j > -1 {
            title = line[j+len(separator):]
            var err error
            
            if seconds, err = strconv.Atoi(line[:j]); err != nil {
                log.Printf("failed to read the duration for '%s': %v\n", title, err)
                seconds = -1
            }
        }
    }
    return title, seconds
}

func mapPlatformDirSeparator(char rune) rune {
    if char == '/' || char == '\\' {
        return filepath.Separator
    }
    return char
}

func writePlsPlaylist(songs []Song) {
    fmt.Println("[playlist]")
    for i, song := range songs {
        i++
        fmt.Printf("File%d=%s\n", i, song.Filename)
        fmt.Printf("Title%d=%s\n", i, song.Title)
        fmt.Printf("Length%d=%d\n", i, song.Seconds)
    }
    fmt.Printf("NumnerOfEntries=%d\nVersion=2\n", len(songs))
}
