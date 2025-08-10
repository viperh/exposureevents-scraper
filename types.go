package main

type SearchReq struct {
	Gender          string `json:"Gender"`
	InviteType      string `json:"InviteType"`
	SportType       string `json:"sportType"`
	Page            int    `json:"Page"`
	SearchToken     string `json:"SearchToken"`
	StartDateString string `json:"StartDateString"`
	EndDateString   string `json:"EndDateString"`
}

type Results struct {
	Results []Event `json:"Results"`
}

type GetResult struct {
	Total int `json:"Total"`
}

type Event struct {
	ID                          int    `json:"Id"`
	Name                        string `json:"Name"`
	OrgId                       string `json:"OrganisationId"`
	OrgWebsite                  string `json:"OrganisationWebsite"`
	OrgName                     string `json:"OrganisationName"`
	Gender                      int
	StartDate                   string  `json:"StartDate"`
	EndDate                     string  `json:"EndDate"`
	Location                    string  `json:"Location"`
	StreetAddress               string  `json:"StreetAddress"`
	ExtendedAddress             string  `json:"ExtendedAddress"`
	City                        string  `json:"City"`
	StateRegion                 string  `json:"StateRegion"`
	StateRegionAbr              string  `json:"StateRegionAbbr"`
	PostalCode                  string  `json:"PostalCode"`
	Longitude                   float64 `json:"Longitude"`
	Latitude                    float64 `json:"Latitude"`
	LogoId                      int     `json:"LogoId"`
	OrgLogoId                   int     `json:"OrganisationLogoId"`
	DateCreated                 string  `json:"DateCreated"`
	MarketingState              int     `json:"MarketingState"`
	EventType                   int     `json:"EventType"`
	Type                        string  `json:"Type"`
	Featured                    int     `json:"Featured"`
	ExposureCertified           bool    `json:"ExposureCertified"`
	ContactName                 string  `json:"ContactName"`
	ContactEmail                string  `json:"ContactEmail"`
	ContactPhone                string  `json:"ContactPhone"`
	SportType                   int     `json:"SportType"`
	SportHost                   string  `json:"SportHost"`
	SportName                   string  `json:"SportName"`
	LogoLink                    string  `json:"Logo"`
	EnableRegistration          bool    `json:"EnableRegistration"`
	GamesGuaranteed             any     `json:"GamesGuaranteed"`
	MinCost                     any     `json:"MinCost"`
	MaxCost                     any     `json:"MaxCost"`
	Ability                     any     `json:"Ability"`
	Certifications              []any   `json:"Certifications"`
	YouthAgeGradesBoth          string  `json:"YouthAgeGradesBoth"`
	DateFormatted               string  `json:"DateFormatted"`
	PriceFormatted              any     `json:"PriceFormatted"`
	RegistrationEnded           bool    `json:"RegistrationEnded"`
	StateRegionLink             string  `json:"StateRegionLink"`
	CalendarLink                string  `json:"CalendarLink"`
	Slug                        string  `json:"Slug"`
	Link                        string  `json:"Link"`
	CityState                   string  `json:"CityState"`
	Website                     string  `json:"Website"`
	RegistrationLink            string  `json:"RegistrationLink"`
	ExternalRegistrationWebsite string  `json:"ExternalRegistrationWebsite"`
	ExternalScheduleWebsite     string  `json:"ExternalScheduleWebsite"`
	ShowSchedule                bool    `json:"ShowSchedule"`
	ScheduleLink                string  `json:"ScheduleLink"`
	IsSiteUrl                   bool    `json:"IsSiteUrl"`
	CollegeCoachLink            any     `json:"CollegeCoachLink"`
	GateLink                    any     `json:"GateLink"`
	PaymentsLink                any     `json:"PaymentsLink"`
}
