package main

import (
	"gopro/internal/config"
	v1 "gopro/internal/http/v1"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	//fmt.Printf("this program uses postfix to send and receive mail. \n")
	//from := "nirajanchapagain@localhost"
	//to := []string{"nirajanchapagain@localhost"}
	//
	//msg := []byte("To: nirajanchapagain@localhost  \r\n+ " +
	//	"Subject: Hello from macos localhost" +
	//	"\r\n" +
	//	"This is email body")
	//
	//err := smtp.SendMail("localhost:25", nil, from, to, msg)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Printf("Email sent successfully")

	cfg := config.Load()
	server := v1.NewServer(cfg)
	server.Start()

}
