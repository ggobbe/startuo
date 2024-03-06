package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/jlaffaye/ftp"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// TODO: quick test, refactor this
func (a *App) DownloadManifest() string {

	c, err := ftp.Dial("maj.uoresistance.com:2121", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login("uoresistance", "")
	if err != nil {
		log.Fatal(err)
	}

	// Do something with the FTP conn
	// Manifest.txt
	r, err := c.Retr("Manifest.txt")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	buf, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	println(string(buf))

	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s", string(buf))
}
