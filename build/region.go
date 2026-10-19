package main

// region.go — what differs between the ru and com deployments apart from env
// config: the zone's domain and the language of the app's shell. The values
// live in region-ru.go and region-com.go. Infrastructure (auth-center) stays
// in env.
//
// The shell is localized; the QR lab is not. The control rack (web/qr/schema.js,
// shapes.js), the readout, the scan diagnostics, the ai tab's JSON validation
// errors and the prompt it copies stay English in both zones. Technical text
// (logs, JSON error bodies) stays English too.

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"reflect"
)

type regionDef struct {
	Domain string // sh-development.<zone>: the login link and the menu redirect
	Lang   string // <html lang>

	Text pageText
}

// pageText holds the strings of web/index.html. Fields, not a map: a typo in
// the template fails at render instead of silently printing nothing.
type pageText struct {
	// login screen
	Login string // login button
	About string // text under the login card

	// header
	TabCodes  string
	TabAI     string
	TabManual string
	Apps      string // profile menu
	Readme    string
	Logout    string

	// qcode tab
	Create     string
	EmptyTitle string
	EmptyHint  template.HTML // trusted markup: holds <b>

	// editor savebar and export bar (manual + ai tabs)
	Name         string
	Unsaved      string // tooltip of the dirty marker
	Save         string
	DownloadPNG  string
	DownloadWebP string
	DownloadSVG  string
	CopySVG      string
	Randomize    string
	Reset        string
	RasterScale  string // aria-label of the scale picker

	// ai tab
	AIStep1          string
	AIStep2          string
	AIStep3          string
	CopyInstructions string
	Apply            string
	Images           string
	Logo             string
	Background       string
	Choose           string
	Clear            string

	JS jsText // strings set by script.js and web/qr/*.js
}

// jsText is rendered into the page as JSON (window.TEXT).
type jsText struct {
	Copied           string `json:"copied"`
	ClipboardBlocked string `json:"clipboardBlocked"`
	CopySVG          string `json:"copySVG"` // the button label restored after "Copied"
	Applied          string `json:"applied"`
	Uploaded         string `json:"uploaded"` // ai tab image row state
	None             string `json:"none"`
	Edit             string `json:"edit"` // card buttons
	Delete           string `json:"delete"`
	UnsavedMsg       string `json:"unsavedMsg"` // %s = code name
	Save             string `json:"save"`
	DontSave         string `json:"dontSave"`
	DeleteMsg        string `json:"deleteMsg"`     // %s = code name
	DeleteConfirm    string `json:"deleteConfirm"` // the modal's button
	Cancel           string `json:"cancel"`
	ExportNoSupport  string `json:"exportNoSupport"` // %s = PNG | WEBP
	ExportFailed     string `json:"exportFailed"`    // %s = PNG | WEBP
}

var regions = map[string]*regionDef{
	"ru":  &regionRU,
	"com": &regionCOM,
}

// Active region, set once by initRegion before anything serves.
var (
	regionName string
	region     *regionDef
	domain     string // region.Domain; auth-server.go builds the menu redirect from it
)

// initRegion picks the region by REGION (ru | com, set by the deploy matrix).
// Empty means local dev and falls back to ru; anything else is a config error
// and stops the process rather than serving the wrong zone — the health check
// then fails and deploy.sh rolls back. A string forgotten in either region is
// fatal too, whichever region this instance serves.
func initRegion() {
	for name, r := range regions {
		if err := checkFilled(reflect.ValueOf(*r), "region"); err != nil {
			log.Fatalf("region %s: %v", name, err)
		}
	}

	regionName = os.Getenv("REGION")
	if regionName == "" {
		regionName = "ru"
	}
	r, ok := regions[regionName]
	if !ok {
		log.Fatalf("unknown REGION %q (want ru or com)", regionName)
	}
	region = r
	domain = r.Domain
	log.Printf("region name=%s domain=%s lang=%s", regionName, domain, region.Lang)
}

// checkFilled walks structs and fails on the first empty string.
func checkFilled(v reflect.Value, path string) error {
	switch v.Kind() {
	case reflect.String:
		if v.Len() == 0 {
			return fmt.Errorf("%s is empty", path)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if err := checkFilled(v.Field(i), path+"."+v.Type().Field(i).Name); err != nil {
				return err
			}
		}
	}
	return nil
}
