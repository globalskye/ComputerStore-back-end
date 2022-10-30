package cmd

import (
	"course_work"
	"log"
)

func main() {
	srv := new(course_work.Server)
	if err := srv.Run("8888"); err != nil {
		log.Fatalln("error to running http server : %s", err)
	}

}
