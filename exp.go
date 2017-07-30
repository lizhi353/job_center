package main

import (
	"runtime"
	"time"
	"math/rand"

	"log"
	"net/http"
	"github.com/job-schedule/trigger/rest"
	//"util"
	"fmt"
	"github.com/job-schedule/schedule"
	"errors"
)

func exp() error {
	fmt.Println("hello")
	return errors.New("hello error")
}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UTC().UnixNano())

	myjob := new(schedule.Job)
	myjob.JobFunc = exp
	go func() {

	}()
	log.Fatal(http.ListenAndServe(":8080", func() http.Handler {
		s, _ := rest.CollectRouters()
		return s
	}()))
}