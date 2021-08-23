package main

import (
	"errors"
	"fmt"

	"github.com/digitalhouse-dev/dh-kit/logger"
)

func main() {
	log := logger.New(logger.SentryOption{Dsn: "a"}, logger.LogOption{})
	log.CatchMessage("hola")
	fmt.Println()
	log.CatchError(errors.New("Hola"))
	fmt.Println()

	fmt.Println("With mocks")
	logger.StartMock()
	log.CatchMessage("My text message")
	fmt.Println()
	log.CatchError(errors.New("My error message"))
	mock := logger.GetMock()
	fmt.Println(mock)
	logger.StopMock()
	fmt.Println()

	fmt.Println("With Debug")
	log = logger.New(logger.SentryOption{Dsn: "a"}, logger.LogOption{Debug: true})
	log.CatchMessage("mi mensaje")
	log.DebugMessage("mi debug del mensaje")
	fmt.Println()
	log.CatchError(errors.New("mi error"))
	log.DebugError(errors.New("Mi debug del error"))
	fmt.Println()

}
