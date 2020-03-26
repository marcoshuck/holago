package main

func main() {
	var app IApplication
	app = &Application{}
	ctx := app.Prepare()
	go app.Start(ctx)
	for {
		if !app.HasFinished() {
			continue
		}
		return
	}
}