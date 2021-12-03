# Project
DH-kit is a project with utilities packages to microservices 

# Setup
```sh
go get github.com/digitalhouse-dev/dh-kit
```

# Logger

```go

// catch messege and errors
log := logger.New(logger.SentryOption{Dsn: "a"}, logger.LogOption{})
log.CatchMessage("hola")
log.CatchError(errors.New("Hola"))


// start with mock
logger.StartMock()
log.CatchMessage("My text message")
log.CatchError(errors.New("My error message"))
mock := logger.GetMock() // get mocks
fmt.Println(mock) 
logger.StopMock()

// with debus
log = logger.New(logger.SentryOption{Dsn: "a"}, logger.LogOption{Debug: true})
log.CatchMessage("mi mensaje")
log.DebugMessage("mi debug del mensaje") // only in logger.LogOption
log.CatchError(errors.New("mi error"))
log.DebugError(errors.New("Mi debug del error")) // only in logger.LogOption

```