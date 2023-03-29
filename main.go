package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/lesismal/nbio/nbhttp"
)

func main() {
	app := iris.New()
	//app.Favicon("./resources/favicon.ico")
	tmpl := iris.HTML("./", ".html")
	// Set custom delimeters.
	tmpl.Delims("{{", "}}")
	// Enable re-build on local template files changes.
	tmpl.Reload(true)

	app.RegisterView(tmpl)
	app.HandleDir("/js", iris.Dir("./js"))
	app.HandleDir("/css", iris.Dir("./css"))
	app.HandleDir("/icon", iris.Dir("./icon"))
	app.HandleDir("/pages", iris.Dir("./pages"))

	app.Get("/", func(ctx iris.Context) {

		if err := ctx.View("index.html"); err != nil {
			ctx.HTML("<h3>%s</h3>", err.Error())
			return
		}
	})

	err := app.Build()
	if err != nil {
		panic(err)
	}

	svr := nbhttp.NewServer(nbhttp.Config{
		Network: "tcp",
		Addrs:   []string{"localhost:8080"},
		Handler: app,
	})

	err = svr.Start()
	if err != nil {
		log.Fatalf("nbio.Start failed: %v\n", err)
	}

	log.Println("serving [go-chi/chi] on [nbio]")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	svr.Shutdown(ctx)
}
