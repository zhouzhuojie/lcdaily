package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"text/template"

	leetcodegraphql "github.com/WindomZ/leetcode-graphql"
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
	titleSlug := strings.Split(url, "/")[4]

	q := new(leetcodegraphql.BaseQuestion)
	q.Do(titleSlug)

	l := &LeetcodePage{
		Title:        q.QuestionTitle,
		Description:  q.Content,
		URL:          leetcodeBaseURL + q.QuestionDetailURL,
		SolutionURL:  leetcodeBaseURL + q.QuestionDetailURL + "discuss",
		QuestionInfo: fmt.Sprintf("ID: %s, Difficulty: %s", q.QuestionID, q.Difficulty),
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
