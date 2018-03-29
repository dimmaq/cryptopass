package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

type Bookmark struct {
	Url      string `json:"url"`
	Username string `json:"username"`
	Length   int    `json:"length"`
}

func getBookmarksData() []byte {
	dat, err := ioutil.ReadFile("CryptopassBookmarks.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return dat
}

func getBookmarksArray(dat []byte) []Bookmark {
	var bookmarks = make([]Bookmark, 1)
	err := json.Unmarshal(dat, &bookmarks)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return bookmarks
}

func getPassword(bm Bookmark, pass string) string {
	iterCount := 5000
	salt := []byte(bm.Username + "@" + bm.Url)
	secret := []byte(pass)

	dk := pbkdf2.Key(secret, salt, iterCount, 32, sha256.New)
	b64 := base64.StdEncoding.EncodeToString(dk)
	return b64[:bm.Length]
}

func main() {
	dat := getBookmarksData()
	if dat == nil {
		return
	}
	bookmarks := getBookmarksArray(dat)
	if dat == nil {
		return
	}
	fmt.Print(bookmarks)

	fmt.Println("***Bookmarks:***")
	for idx, bm := range bookmarks {
		fmt.Println(idx, bm)
	}

	rc := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("----------------")
		fmt.Print("enter num: ")
		z, _ := rc.ReadString('\n')
		//fmt.Println(z)
		num, _ := strconv.Atoi(strings.TrimSpace(z))
		//fmt.Println(num,)
		if (num < 0) || (num >= len(bookmarks)) {
			fmt.Println("***wrong num***")
			continue
		}
		fmt.Print("enter secret: ")
		pw, _ := rc.ReadString('\n')
		pw = strings.TrimSpace(pw)
		//fmt.Print(pw)

		bm := bookmarks[num]
		fmt.Println("***", bm, ":", getPassword(bookmarks[num], pw))

	}

}
