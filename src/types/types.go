package types

type Response struct {
	Name          string `json:"name" xml:"name"`
	Default       bool   `json:"default" xml:"default"`
	Url           string `json:"url" xml:"url"`
	GmailLogo     string `json:"gmailLogo" xml:"gmailLogo"`
	HangoutslLogo string `json:"hangoutsLogo" xml:"hangoutsLogo"`
	DriveLogo     string `json:"driveLogo" xml:"driveLogo"`
}
