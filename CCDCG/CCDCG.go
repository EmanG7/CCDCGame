package main

import (
	//"fmt"
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
	timestart int
	daycheck int
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
	/*
	for {
		_, err := fmt.Fprint(w, " ")
		if err != nil {
			log.Fatal("client is gone, shutting down")
			return
		}
		flusher := w.(http.Flusher)
		flusher.Flush()
		time.Sleep(time.Second)
	}
	*/
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
	/*
	for {
		_, err := fmt.Fprint(w, " ")
		if err != nil {
			log.Fatal("client is gone, shutting down")
			return
		}
		flusher := w.(http.Flusher)
		flusher.Flush()
		time.Sleep(time.Second)
	}
	*/
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
	/*
	for {
		_, err := fmt.Fprint(w, " ")
		if err != nil {
			log.Fatal("client is gone, shutting down")
			return
		}
		flusher := w.(http.Flusher)
		flusher.Flush()
		time.Sleep(time.Second)
	}
	*/
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
	/*
	for {
		_, err := fmt.Fprint(w, " ")
		if err != nil {
			log.Fatal("client is gone, shutting down")
			return
		}
		flusher := w.(http.Flusher)
		flusher.Flush()
		time.Sleep(time.Second)
	}
	*/
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
	/*
	for {
		_, err := fmt.Fprint(w, " ")
		if err != nil {
			log.Fatal("client is gone, shutting down")
			return
		}
		flusher := w.(http.Flusher)
		flusher.Flush()
		time.Sleep(time.Second)
	}
	*/
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
	/*
	for {
		_, err := fmt.Fprint(w, " ")
		if err != nil {
			log.Fatal("client is gone, shutting down")
			return
		}
		flusher := w.(http.Flusher)
		flusher.Flush()
		time.Sleep(time.Second)
	}
	*/
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
	/*
	for {
		_, err := fmt.Fprint(w, " ")
		if err != nil {
			log.Fatal("client is gone, shutting down")
			return
		}
		flusher := w.(http.Flusher)
		flusher.Flush()
		time.Sleep(time.Second)
	}
	*/
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

	log.Println("running at port 127.0.0.1:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
//End of main system

//Timer
func checkTime() (int,int){
	t := time.Now()
	h := t.Hour()
	m := t.Minute()
	s := t.Second()
	return (h * 360) + (m * 60) + s, t.Day()
}

func totalTime(startT int, stopT int, today int, nextday int) int {
	if today < nextday {
		stopT = stopT + (24*360)
	}
	return stopT - startT
}

func questionGen(index int) string {
	network := [4]string{"network1","network2","network3","Game Over, press the done button"}
	liaison := [4]string{"What format do you submit an inject as?","Which specialist would you task with Restriction of Facebook for all users?","Which distro is typically used as the centralized logging hub?","Game Over, press the done button"}
	windows := [4]string{"win1","win2","win3","Game Over, press the done button"}
	linux := [4]string{"linux1","linux2","linux3","Game Over, press the done button"}
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
	network := [3]string{"networkanswer1","networkanswer2","networkanswer3"}
	liaison := [3]string{"PDF","Network Specialist","Splunk"}
	windows := [3]string{"win1","win2","win3"}
	linux := [3]string{"linux1","linux2","linux3"}
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
