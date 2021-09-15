package main

import (
	"os"
	"sync"
	"regexp"
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/grokify/html-strip-tags-go"
	//"strings"
	"text/template"
	"crypto/md5"
	"sort"
	//"errors"
)

var HH_SEARCH_TEXT = os.Getenv("INPUT_HH_SEARCH_TEXT")
const VACANCIES_MAX_TOTAL = 1999 // HH.RU API allows up to 2000 vacancies only
const VACANCIES_PER_PAGE = 99    // ...again, this is upper limit posed by HH.RU
//const PAGES_COUNT = 2            
var PAGES_COUNT = VACANCIES_MAX_TOTAL / VACANCIES_PER_PAGE

const numDownloaders = 100 //20

var VACANCIES_AGE = os.Getenv("INPUT_VACANCY_AGE")
var VACANCIES_PAGE_BASE_URL = fmt.Sprintf(
	"https://api.hh.ru/vacancies?text=%s" +
	"&per_page=%d" +
	"&period=%s" +
	"&vacancy_search_order=publication_date&page=",
	HH_SEARCH_TEXT,
	VACANCIES_PER_PAGE,
	VACANCIES_AGE)

const API_VACANCIES_BASE_URL = "https://api.hh.ru/vacancies/"
const VACANCY_BASE_URL = "https://hh.ru/vacancy/"

// The struct is generated from HH's response using https://mholt.github.io/json-to-go/
// Some fields are not used and thus removed
type Vacancy struct {
	Schedule struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"schedule"`
	Details                *VacancyDetails // custom field, not from HH respond
	ID                     string      `json:"id"`
	Name                   string      `json:"name"`
	Area                   struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"area"`
	AdditionalAreas        []string // if we find the vacancy in other regions, add them in here
	NormalizedSalary       string // convert any salary to a common value ($)
	Salary interface{} `json:"salary"`
	Type   struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"type"`
	Address           interface{}   `json:"address"`
	ResponseURL       interface{}   `json:"response_url"`
	SortPointDistance interface{}   `json:"sort_point_distance"`
	PublishedAt       string        `json:"published_at"`
	CreatedAt         string        `json:"created_at"`
	Archived          bool          `json:"archived"`
	Employer          struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		URL          string `json:"url"`
		Trusted      bool   `json:"trusted"`
	} `json:"employer"`
	Snippet struct {
		Requirement    string `json:"requirement"`
		Responsibility string `json:"responsibility"`
	} `json:"snippet"`
	Contacts interface{} `json:"contacts"`
	AcceptTemporary      bool          `json:"accept_temporary"`
} 

type Vacancies struct {
	Items []Vacancy `json:"items"`
}

type VacancyDetails struct {
	ID          string `json:"id"`
	Name                   string        `json:"name"`
	Description                string      `json:"description"`
	Excerpt                string // HTMLed excerpt of the description where keyword found
	Hash                   string // md5 of the description (with no HTML tags)
}

/////////////////

func pf(s string) {
	fmt.Printf("[Print] %s\n", s)
}

func getAllVacanciesList(done <-chan struct{}) (chan Vacancy) {
	pf(VACANCIES_PAGE_BASE_URL)
	allVacancies := make(chan Vacancy)
	go func() {
		defer close(allVacancies)
		for page := 1; page <= PAGES_COUNT; page++ {
			url := fmt.Sprintf("%s" + "%d", VACANCIES_PAGE_BASE_URL, page)
			resp, err := http.Get(url)
			if err != nil {
				pf(fmt.Sprintf("Error %s for %s\n", err, url))
				// log.Fatalln(err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)

			vacancies := &Vacancies{}
			err = json.Unmarshal(body, &vacancies)
			if err != nil {
				//log.Println(err)
				fmt.Println(err)
			}
		
			items := vacancies.Items
			for i := 0; i < len(items); i++ {
				//vac := &items[i]
				select {
					case allVacancies <- items[i]:
						// pf(
						// 	fmt.Sprintf(
						// 		"page %d - getVacancies(): № %d '%s' in '%s'\n", page, i + 1, vac.Name, vac.Employer.Name))
					case <-done:
						return //nil //errors.New("walk canceled")
				}
				
			}
		}
		return //nil
	}()
	return allVacancies
}

func getVacancyDescription(done <-chan struct{}, allVacancies chan Vacancy, withDescriptions chan<- Vacancy, myId int) {
	//pf(fmt.Sprintf("start getVacancyDescription() № %d\n", myId))
	for vac := range allVacancies {
		url := API_VACANCIES_BASE_URL + vac.ID
		//pf("Fetching " + url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error %s for %s\n", err, url)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		details := &VacancyDetails{}
		err = json.Unmarshal(body, &details)
		if err != nil {
			//log.Println(err)
			fmt.Println(err)
		}
		plain := strip.StripTags(details.Description)
		byted := []byte(plain)
		details.Hash = fmt.Sprintf("%x", md5.Sum(byted))

		//fmt.Printf("getVacancyDescription() № %d Vacancy: %s, descr: %s\n", myId, vac.Name, details.Description[:12])
		vac.Details = details
		select {
		case withDescriptions <- vac:
		case <-done:
			{
		    fmt.Printf("\t%d got done, EXIT\n", myId)
			return
		  }
		}
	}
	//fmt.Printf("\t%d EXITing as allVacancies closed\n", myId)
}
/* 
func normalizeSalary(vac Vacancy) string {
	salary := ""
	RurToDoll := 74
	var from int
	var to int
	if vac.Salary == nil {
		return salary
	}
	if vac.Salary == "USD" {
		salary += "$ "
	}
	if vac.Salary.from != nil {
		from = vac.Salary.from 
	}
	if vac.Salary.to != nil {
		to = vac.Salary.from 
	}
	if vac.Salary == "RUR" {
        to = to / RurToDoll
		from = from / RurToDoll
	}
    return salary
}
 */
/* 
func getExcerpt(keyword, txt string) (string) {
	txt = strip.StripTags(txt)
	contextSize := 50
	fullLen := len(txt)
	pos := strings.Index(txt, keyword)
	startPos := pos - contextSize
	endPos := pos + len(keyword) + contextSize
	if startPos < 0 { startPos = 0 }
	if endPos > fullLen - 1 { endPos = fullLen - 1}
	excerpt := fmt.Sprintf("...%s<b>%s</b>%s...",
		txt[startPos:pos],
		keyword,
		txt[pos + len(keyword):endPos])
	return excerpt
}
 */
 func getExcerpt(txt string) (string) {
	excerpt := ""
	pattern := "удален|remote"
	r, _ := regexp.Compile(pattern)
	txt = strip.StripTags(txt)
	pos := r.FindStringIndex(txt)
	if pos == nil {
		return excerpt
	}
	contextSize := 50
	fullLen := len(txt)

	startPos := pos[0] - contextSize
	endPos := pos[1] + contextSize
	if startPos < 0 { startPos = 0 }
	if endPos > fullLen - 1 { endPos = fullLen - 1}
	excerpt = fmt.Sprintf("...%s<b>%s</b>%s...",
		txt[startPos:pos[0]],
		txt[pos[0]:pos[1]],
		txt[pos[1]:endPos])
	return excerpt
}


func main() {
	done := make(chan struct{})
	defer close(done)

	allVacancies := getAllVacanciesList(done)
	
	withDescription := make(chan Vacancy)

	var downloadersWg sync.WaitGroup
	downloadersWg.Add(numDownloaders)
	downloaderId := 1
	for d := 0; d < numDownloaders; d++ {
		downloaderId++
		go func() {
            getVacancyDescription(done, allVacancies, withDescription, downloaderId)
			downloadersWg.Done()
		}()
	}
	go func() {
		downloadersWg.Wait()
		close(withDescription)
	}()
	i := 0
	notRemote := 0
	var remoteOnes []Vacancy 
	type RemoteVacancies struct{
		Title string
		Items []Vacancy
	}

	seenBefore := map[string]bool{}
	seenBeforeCount := 0

	for vac := range withDescription {
		i++
		//pos = 0
		   //pos := strings.Index(vac.Details.Description, "удален"); pos != 0 ||{
		vac.Details.Excerpt = getExcerpt(vac.Details.Description)
		if  vac.Schedule.ID == "remote" || len(vac.Details.Excerpt) > 0 {
			pf(fmt.Sprintf("%d (%s) [sal: %v] '%s' in '%s'\n\t%s", i, VACANCY_BASE_URL + vac.ID, vac.Salary,vac.Name, vac.Employer.Name, vac.Details.Excerpt))
			if vac.Salary == nil {
				vac.NormalizedSalary  = ""
			} else {
				vac.NormalizedSalary = fmt.Sprintf("%v", vac.Salary)
			}
			_, seen := seenBefore[vac.Details.Hash]
			if seen {
				pf("Seen before")
				seenBeforeCount++
			} else {
				seenBefore[vac.Details.Hash] = true
				remoteOnes = append(remoteOnes, vac)
			}
		} else {
			notRemote++
		}
	}
	sort.Slice(remoteOnes, func(i, j int) bool {
		return remoteOnes[i].PublishedAt > remoteOnes[j].PublishedAt
	})
	data := RemoteVacancies{Title: "vacancies", Items: remoteOnes}
	//data.Items = remoteOnes
	pf(fmt.Sprintf("Skipped not remote ones: %d", notRemote))
	pf(fmt.Sprintf("Duplicated: %d", seenBeforeCount))
	tmpl, _ := template.ParseFiles("templates/index.html")
	f, err := os.Create("index.html")
	if err != nil {
		// handle error
	}
	tmpl.Execute(f, data)
}

