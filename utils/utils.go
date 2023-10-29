package utils

import (
	"gomx/config"
	"log"
	"net/http"
)

func ExecTemplate(w http.ResponseWriter, template string, data interface{}) {
	if err := config.Tpl.ExecuteTemplate(w, template, data); err != nil {
		log.Println("Failed !! Error While Parsing"+template+"Template: ", err)
	}
}

// func OpenInBrowser(browser string, url string, private bool) {

// 	var c *exec.Cmd
// 	doPrivate := ""

// 	if private {
// 		doPrivate = "--incognito"
// 	}

// 	switch runtime.GOOS {
// 	case "windows":
// 		c = exec.Command("cmd", "/C", "start", browser, url, doPrivate)
// 	default:
// 		c = exec.Command(browser, url, doPrivate)
// 	}

// 	if err := c.Run(); err != nil {
// 		log.Println("Failed !! Error While Opening Browser: ", err)
// 	}

// }
