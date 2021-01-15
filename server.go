package main

import (
	toolbox "github.com/jfabdo/speak-easy-toobox"
	"net/http"
	"os"
)

//'evilrus.c/speak-easy/toolbox" (backup just in case-q)
//globals

var conn = toolbox.GetPool() //get pool of redis connections, that pulls from database(redis) to https servcer

func main() {
	// start w. 2 player focus and later add for up to 4 players.
	//building n player (2/4p) game(J)

	toolbox.GetServer(
		os.Getenv("ERU_SE_REDIS_IP"), "80", serverIndex)
}

//api w(single microservice (Q)
func serverIndex(w http.ResponseWriter, r *http.Request) {
	//(switch is if combined statement)
	switch r.URL.String() {

	case "Location":
		serveLocation() //when they move, the user will control the console but we only need when, not how(Q)

	case "scores":
		serveScores() //() indicate start of function

	case "Effects":
		serveEffects()
	default:
		serveDefault() //serve everythinng back but serve front page/end(J) Redirect to the lobby (j)/serve everythinng back but serve front page/end(J) Redirect to the lobby (j)
	}
}

func serveLocation() {
	//will be filled in when mechanices are flushed out
}

func serveScores() {
	//will be filled in when mechanics are flushed out
}

func serveEffects() {
	//will be filled in when mechanices are flushed out
}

func serveDefault() {
	//will be filled in when mechanices are flushed out
}
