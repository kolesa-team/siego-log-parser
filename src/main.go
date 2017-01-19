package main

import (
	"io/ioutil"
	"os"

	"./siego-log-parser/schema"
	"./siego-log-parser/statsd"

	"github.com/codegangsta/cli"
	"fmt"
)

func main() {
	app := cli.NewApp()

	app.Name = "siego-log-parser"
	app.Usage = "Log parser for siego. Bridge from XML files to statsd."
	app.Version = "0.0.1"
	app.Author = "Igor Borodikhin"
	app.Email = "iborodikhin@gmail.com"
	app.Action = actionRun
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Usage: "FILE, siego log in XML format",
		},
		cli.StringFlag{
			Name:  "address, a",
			Usage: "Statsd address in form address:port",
		},
		cli.StringFlag{
			Name:  "prefix, p",
			Usage: "Prefix for statsd",
		},
	}

	app.Run(os.Args)
}

func actionRun(c *cli.Context) (err error) {
	client, err := statsd.NewStatsd(c.String("address"), c.String("prefix"))
	if err != nil {
		return
	}

	data, err := ioutil.ReadFile(c.String("file"))
	if err != nil {
		return
	}

	result, err := schema.ParseStatistics(data)
	if err != nil {
		return
	}

	err = statsd.Save(client, result)
	if err == nil {
		fmt.Printf("All done! Bye-bye!\r\n")
	}

	return
}
