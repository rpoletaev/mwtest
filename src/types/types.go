package types

type Response struct {
	Name          string `json:"name"`
	Default       bool   `json:"default"`
	Url           string `json:"url"`
	GmailLogo     string `json:"gmailLogo"`
	HangoutslLogo string `json:"hangoutsLogo"`
	DriveLogo     string `json:"driveLogo"`
}
