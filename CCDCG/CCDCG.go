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
	current.question = "test"
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

func main() {
	http.HandleFunc("/", initH)
	http.HandleFunc("/index", loginH)
	http.HandleFunc("/networking", networkingStart)

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
	network := [3]string{"question1","question2","question3"}
	liaison := [3]string{"question1","question2","question3"}
	windows := [3]string{"question1","question2","question3"}
	linux := [3]string{"question1","question2","question3"}
	var results string
	if (index / 10) == 0 {
		results := network[index%10]
	} else if (index / 10) == 1 {

	} else if (index / 10) == 2 {

	} else if (index / 10) == 3 {

	} else {

	}
	return results
}