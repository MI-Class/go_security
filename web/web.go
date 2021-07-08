package main

import "github.com/webview/webview"

func main() {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("https://www.baidu.com")
	go w.Run()
	select {}
}

func RunWeb() {

}
