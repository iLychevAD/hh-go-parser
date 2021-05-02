package main
// The struct is generated from HH's response using https://mholt.github.io/json-to-go/
type Vacancy struct {
	Details                *VacancyDetails // custom field, not from HH respond
	ID                     string      `json:"id"`
	Premium                bool        `json:"premium"`
	Name                   string      `json:"name"`
	Department             interface{} `json:"department"`
	HasTest                bool        `json:"has_test"`
	ResponseLetterRequired bool        `json:"response_letter_required"`
	Area                   struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"area"`
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
	ApplyAlternateURL string        `json:"apply_alternate_url"`
	InsiderInterview  interface{}   `json:"insider_interview"`
	URL               string        `json:"url"`
	AlternateURL      string        `json:"alternate_url"`
	Relations         []interface{} `json:"relations"`
	Employer          struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		URL          string `json:"url"`
		AlternateURL string `json:"alternate_url"`
		LogoUrls     struct {
			Num90    string `json:"90"`
			Num240   string `json:"240"`
			Original string `json:"original"`
		} `json:"logo_urls"`
		VacanciesURL string `json:"vacancies_url"`
		Trusted      bool   `json:"trusted"`
	} `json:"employer"`
	Snippet struct {
		Requirement    string `json:"requirement"`
		Responsibility string `json:"responsibility"`
	} `json:"snippet"`
	Contacts interface{} `json:"contacts"`
	Schedule struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"schedule"`
	WorkingDays          []interface{} `json:"working_days"`
	WorkingTimeIntervals []interface{} `json:"working_time_intervals"`
	WorkingTimeModes     []interface{} `json:"working_time_modes"`
	AcceptTemporary      bool          `json:"accept_temporary"`
} 

type Vacancies struct {
	Items []Vacancy `json:"items"`
	// The following fields are in the HH's response but we dont use them
	// Found        int         `json:"found"`
	// Pages        int         `json:"pages"`
	// PerPage      int         `json:"per_page"`
	// Page         int         `json:"page"`
	// Clusters     interface{} `json:"clusters"`
	// Arguments    interface{} `json:"arguments"`
	// AlternateURL string      `json:"alternate_url"`
}

type VacancyDetails struct {
	ID          string `json:"id"`
	Premium     bool   `json:"premium"`
	BillingType struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"billing_type"`
	Relations              []interface{} `json:"relations"`
	Name                   string        `json:"name"`
	InsiderInterview       interface{}   `json:"insider_interview"`
	ResponseLetterRequired bool          `json:"response_letter_required"`
	Area                   struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"area"`
	Salary struct {
		From     int    `json:"from"`
		To       int    `json:"to"`
		Currency string `json:"currency"`
		Gross    bool   `json:"gross"`
	} `json:"salary"`
	Type struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"type"`
	Address       interface{} `json:"address"`
	AllowMessages bool        `json:"allow_messages"`
	Site          struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"site"`
	Experience struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"experience"`
	Schedule struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"schedule"`
	Employment struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"employment"`
	Department                 interface{} `json:"department"`
	Contacts                   interface{} `json:"contacts"`
	Description                string      `json:"description"`
	BrandedDescription         interface{} `json:"branded_description"`
	VacancyConstructorTemplate interface{} `json:"vacancy_constructor_template"`
	KeySkills                  []struct {
		Name string `json:"name"`
	} `json:"key_skills"`
	AcceptHandicapped bool        `json:"accept_handicapped"`
	AcceptKids        bool        `json:"accept_kids"`
	Archived          bool        `json:"archived"`
	ResponseURL       interface{} `json:"response_url"`
	Specializations   []struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		ProfareaID   string `json:"profarea_id"`
		ProfareaName string `json:"profarea_name"`
	} `json:"specializations"`
	Code                    interface{}   `json:"code"`
	Hidden                  bool          `json:"hidden"`
	QuickResponsesAllowed   bool          `json:"quick_responses_allowed"`
	DriverLicenseTypes      []interface{} `json:"driver_license_types"`
	AcceptIncompleteResumes bool          `json:"accept_incomplete_resumes"`
	Employer                struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		URL          string `json:"url"`
		AlternateURL string `json:"alternate_url"`
		LogoUrls     struct {
			Num90    string `json:"90"`
			Num240   string `json:"240"`
			Original string `json:"original"`
		} `json:"logo_urls"`
		VacanciesURL string `json:"vacancies_url"`
		Trusted      bool   `json:"trusted"`
	} `json:"employer"`
	PublishedAt          string        `json:"published_at"`
	CreatedAt            string        `json:"created_at"`
	NegotiationsURL      interface{}   `json:"negotiations_url"`
	SuitableResumesURL   interface{}   `json:"suitable_resumes_url"`
	ApplyAlternateURL    string        `json:"apply_alternate_url"`
	HasTest              bool          `json:"has_test"`
	Test                 interface{}   `json:"test"`
	AlternateURL         string        `json:"alternate_url"`
	WorkingDays          []interface{} `json:"working_days"`
	WorkingTimeIntervals []interface{} `json:"working_time_intervals"`
	WorkingTimeModes     []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"working_time_modes"`
	AcceptTemporary bool `json:"accept_temporary"`
}

/////////////////
