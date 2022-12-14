package main

import (
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"time"
)

//Custom variables
type player struct {
	username string
	question string
	questionindex int
	entry string
	score int
}

//Declaration of custom variables
var current player

//Start of main system.
func initH(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("html/init.html")
	if err != nil {
		log.Println(err)
		return
	}
	if err := tmpl.Execute(w,nil); err != nil {
		log.Println(err)
	}
}

func loginH(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	current.username = r.PostForm.Get("uname")
	current.questionindex = 99
	//log.Println(current.username)
	tmpl, err := template.ParseFiles("html/index.html")
	if err != nil {
		log.Println(err)
		return
	}
	if err := tmpl.Execute(w,current.username); err != nil {
		log.Println(err)
	}
}

func networkingStart(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	current.entry = r.PostForm.Get("entry")
	if current.entry != "" { scorer() }
	if (current.questionindex / 10) != 0 {
		current.questionindex = 0
	} else {
		current.questionindex += 1
	}
	current.question = questionGen(current.questionindex)
	tmpl, err := template.ParseFiles("html/network.html")
	if err != nil {
		log.Println(err)
		return
	}
	if err := tmpl.Execute(w,current.question); err != nil {
		log.Println(err)
	}
}
func liaisonStart(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	current.entry = r.PostForm.Get("entry")
	if current.entry != "" { scorer() }
	if (current.questionindex / 10) != 1 {
		current.questionindex = 10
	} else {
		current.questionindex += 1
	}
	current.question = questionGen(current.questionindex)
	tmpl, err := template.ParseFiles("html/liaison.html")
	if err != nil {
		log.Println(err)
		return
	}
	if err := tmpl.Execute(w,current.question); err != nil {
		log.Println(err)
	}
}
func windowsStart(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	current.entry = r.PostForm.Get("entry")
	if current.entry != "" { scorer() }
	if (current.questionindex / 10) != 2 {
		current.questionindex = 20
	} else {
		current.questionindex += 1
	}
	current.question = questionGen(current.questionindex)
	tmpl, err := template.ParseFiles("html/windows.html")
	if err != nil {
		log.Println(err)
		return
	}
	if err := tmpl.Execute(w,current.question); err != nil {
		log.Println(err)
	}
}

func linuxStart(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	current.entry = r.PostForm.Get("entry")
	if current.entry != "" { scorer() }
	if (current.questionindex / 10) != 3 {
		current.questionindex = 30
	} else {
		current.questionindex += 1
	}
	current.question = questionGen(current.questionindex)
	tmpl, err := template.ParseFiles("html/linux.html")
	if err != nil {
		log.Println(err)
		return
	}
	if err := tmpl.Execute(w,current.question); err != nil {
		log.Println(err)
	}
}

func scoreboard(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err := template.ParseFiles("html/scoreboard.html")
	if err != nil {
		log.Println(err)
		return
	}
	if err := tmpl.Execute(w, current.score); err != nil {
		log.Println(err)
	}
}

func main() {
	http.HandleFunc("/", initH)
	http.HandleFunc("/index", loginH)
	http.HandleFunc("/networking", networkingStart)
	http.HandleFunc("/liaison", liaisonStart)
	http.HandleFunc("/windows", windowsStart)
	http.HandleFunc("/linux", linuxStart)
	http.HandleFunc("/scoreboard", scoreboard)

	go func() {
		<-time.After(100 * time.Millisecond)
		err := exec.Command("cmd", "/C", "start", "http://127.0.0.1:8000").Run() //Windows Default Browser
		if err != nil {
			log.Println(err)
		}
	}()

	//log.Println("running at port 127.0.0.1:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
//End of main system

func questionGen(index int) string {
	network := [4]string{"What command do you use on Cisco network devices to set a minimum length of \"5\" for passwords ","What command do you use on Cisco network devices to show QOS Map?","What command do you use in the Palo Alto Network's CLI to set up the management interface with an ip address of \"172.16.1.100\", a subnet of \"255.255.255.0\", a default gateway of \"172.16.1.1\", and a primary DNS server of \"8.8.8.8\"?","Game Over, press the done button"}
	liaison := [4]string{"What format do you submit an inject as?","Which specialist would you task with Restriction of Facebook for all users?","Which distro is typically used as the centralized logging hub?","Game Over, press the done button"}
	windows := [4]string{"Use Powershell command to clear all exisiting Firewall Rules","Use Powershell commands to enable the firewall","Use Powershell commands to set the default firewall policy to block all traffic","Game Over, press the done button"}
	linux := [4]string{"Command to Delete all iptables","Command to Accept Outbound TCP traffic on ports 80 and 443","Command to stop SSH service","Game Over, press the done button"}
	var results string
	if (index / 10) == 0 {
		results = network[index%10]
	} else if (index / 10) == 1 {
		results = liaison[index%10]
	} else if (index / 10) == 2 {
		results = windows[index%10]
	} else if (index / 10) == 3 {
		results = linux[index%10]
	} else {
		log.Fatal("questionGen function failed to produce results based on the questionindex variable of the current player structure.")
	}
	return results
}

func answerGen(index int) string {
	network := [3]string{"security passwords min-length 5","show mls qos maps","set deviceconfig system ip-address 172.16.1.100 netmask 255.255.255.0 default-gateway 172.16.1.1 dns-setting servers primary 8.8.8.8"}
	liaison := [3]string{"PDF","Network Specialist","Splunk"}
	windows := [3]string{"netsh advfirewall firewall delete rule name= \"all\"","Netsh advfirewall set currentprofile state on","Netsh advfirewall set allprofiles firewallpolicy \"blockinbound,blockoutbound\""}
	linux := [3]string{"iptables -D","iptables -A OUTPUT -p tcp -m multiport -dports 80,443 -j ACCEPT","service sshd stop"}
	var results string
	if (index / 10) == 0 {
		results = network[index%10]
	} else if (index / 10) == 1 {
		results = liaison[index%10]
	} else if (index / 10) == 2 {
		results = windows[index%10]
	} else if (index / 10) == 3 {
		results = linux[index%10]
	} else {
		log.Fatal("answerScorer function failed to produce results based on the questionindex variable of the current player structure.")
	}
	return results
}

func scorer() {
	if current.entry == answerGen(current.questionindex) {
		current.score += 100
		//log.Println(current.score)
	} else {
		current.score -= 10
		//log.Println(current.score)
	}
}
