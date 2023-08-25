package core

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func MainRoutine() error {
	log.Debug("Main routine is called")
	fmt.Println("Hello from golang-template main-routine")
	log.Debug("Main routine is complete")

	return nil
}
