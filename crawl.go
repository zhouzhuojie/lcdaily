package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const (
	htmlTemplateName  = "email_template.html"
	leetcodeRandomURL = "https://leetcode.com/problems/random-one-question/algorithms"
	leetcodeBaseURL   = "https://leetcode.com"
)

// LeetcodePage is a struct to hold the information in leetcode
type LeetcodePage struct {
	// Title is the tile of the problem
	Title string

	// Description is the description of the project
	Description string

	// URL is the url of the project
	URL string

	// SolutionURL is the url linked to the solutions
	SolutionURL string

	// QuestionInfo is the question info
	QuestionInfo string

	// HostPort is the host port that we can use to hit our own server
	// in an adhoc url
	HostPort string
}

// NewRandomLeetcodePage creates a new random page
func NewRandomLeetcodePage() (*LeetcodePage, error) {
	var url string
	r, _ := http.NewRequest("GET", leetcodeRandomURL, nil)
	cl := http.Client{}
	cl.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		url = req.URL.RequestURI()
		return nil
	}
	cl.Do(r)

	url = leetcodeBaseURL + url
	return NewLeetcodePage(url)
}

// NewLeetcodePage creates a new leetcode page from url
func NewLeetcodePage(url string) (*LeetcodePage, error) {

	doc, err := goquery.NewDocument(url)

	if err != nil {
		return nil, err
	}

	descriptionDiv := doc.Find("meta").Eq(2)
	description, _ := descriptionDiv.Attr("content")

	title := doc.Find("h3").First().Text()
	questionInfo := doc.Find(".question-info").First().Text()

	l := &LeetcodePage{
		Title:        title,
		Description:  description,
		URL:          url,
		SolutionURL:  url + "#/solutions",
		QuestionInfo: questionInfo,
	}

	return l, nil
}

// ToHTML returns HTML representation of the page
func (l *LeetcodePage) ToHTML(host string, port string) []byte {

	l.HostPort = fmt.Sprintf("%s:%s", host, port)

	var b bytes.Buffer
	t := template.Must(template.New(htmlTemplateName).ParseFiles(htmlTemplateName))
	err := t.Execute(&b, l)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return b.Bytes()
}
