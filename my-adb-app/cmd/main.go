package main

import (
	"github.com/serhiq/go_fyne_my_adb/internal/app"
	"github.com/serhiq/go_fyne_my_adb/internal/config"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {

	cfg, err := config.New()

	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app, err := app.New(*cfg)
	if err != nil {
		log.Fatalf("App init error: %s", err)
	}

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	go func() {
		sig := <-exit
		log.Printf("%s   received.", sig.String())
		log.Printf("\n Goroutines: %d", runtime.NumGoroutine())

		app.Stop()
		log.Println("\n Shutdown app")
	}()

	err = app.Start()
	if err != nil {
		log.Fatalf("Server start error: %s", err)
	}
}
