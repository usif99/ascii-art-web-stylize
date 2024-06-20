package main

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type PageData struct {
	Text []string
}

func main() {
	fmt.Println("CTRL + CLICK TO VIEW THE PROJECT --> http://localhost:8080/")
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil) //listen starts HTTP server
}

func Handler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("thetext")
	fileName := r.FormValue("chose")
	_, error := os.Stat(fileName + ".txt")
	indexTemplate, _ := template.ParseFiles("template/index.html")

	if (!CheckLetter(text) || text == "") && r.Method == "POST" {
		w.WriteHeader(http.StatusBadRequest)
		http.ServeFile(w, r, "./template/400.html")
		return
	} else if (os.IsNotExist(error) || len(r.FormValue("thetext")) > 2000) && r.Method != "GET" {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "template/500.html")
		return
	}

	textInASCII := serveIndex(text, fileName)
	pageData := PageData{
		Text: textInASCII,
	}

	if r.URL.Path == "/style.css" {
		http.ServeFile(w, r, "./template/style.css")
		return
	} else if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "./template/404.html")
		return
	} else if r.Method == "GET" || r.Method == "POST" {
		err := indexTemplate.Execute(w, pageData)
		if err != nil {
			fmt.Print(err)
		}
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		http.ServeFile(w, r, "./template/400.html") // should be 400
		return
	}
}

func serveIndex(text, filename string) []string {
	var Text []string
	WordsInArr := strings.Split(text, "\r\n")
	var Words []string
	for l := 0; l < len(WordsInArr); l++ {
		var Words [][]string
		Text1 := strings.ReplaceAll(WordsInArr[l], "\\t", "   ")
		if Text1 != "" {
			for j := 0; j < len(Text1); j++ {
				Words = append(Words, ReadLetter(Text1[j], filename))
			}
			for x := 0; x < 8; x++ {
				Lines := ""
				for n := 0; n < len(Words); n++ {
					Lines += Words[n][x]
				}
				Text = append(Text, Lines)
			}
		} else {
			Text = append(Text, "\n")
		}
	}
	return append(Words, strings.Join(Text, "\n"))
}

func ReadLetter(Text1 byte, fileName string) []string {
	var Letter []string
	ReadFile, _ := os.Open(fileName + ".txt")
	FileScanner := bufio.NewScanner(ReadFile)
	stop := 1
	i := 0
	letterLength := (int(Text1)-32)*9 + 2
	for FileScanner.Scan() {
		i++
		if i >= letterLength {
			stop++
			Letter = append(Letter, FileScanner.Text())
			if stop > 8 {
				break
			}
		}
	}
	ReadFile.Close()
	return Letter
}

func CheckLetter(s string) bool {
	WordsInArr := strings.Split(s, "\r\n")
	for l := 0; l < len(WordsInArr); l++ {
		for g := 0; g < len(WordsInArr[l]); g++ {
			if WordsInArr[l][g] > 126 || WordsInArr[l][g] < 32 {
				return false
			}
		}
	}
	return true
}