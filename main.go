package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func fatal(args ...interface{}) {
	fmt.Println(args...)
	os.Exit(1)
}

func main() {
	dirname := "/usr/local/Cellar/"

	app := cli.NewApp()

	app.Name = "brewservice"
	app.Usage = "Easily start and stop services installed via homebrew"
	app.Version = "0.1.0"

	app.Commands = []cli.Command{
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "Parse homebrew folders for plist files and cache them for later",
			Action: func(c *cli.Context) {
				cache, err := update(dirname)
				if err != nil {
					fatal(err)
				}

				err = cache.Save()
				if err != nil {
					fatal(err)
				}

				fmt.Printf("Loaded %d service(s) from %s\n", len(cache), dirname)
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "List available services",
			Action: func(c *cli.Context) {

				cache := loadCacheOrFail()
				for service, _ := range cache {
					fmt.Println("* " + service)
				}
			},
		},
		{
			Name:  "start",
			Usage: "Start a service",
			Action: func(c *cli.Context) {
				cache := loadCacheOrFail()
				serviceName := c.Args().First()
				if service, ok := cache[serviceName]; ok {
					err := service.Start()
					if err != nil {
						fatal(err)
					}
					fmt.Println(serviceName + " started")
				} else {
					fatal("Service not found")
				}
			},
		},
		{
			Name:  "stop",
			Usage: "Stop a service",
			Action: func(c *cli.Context) {
				cache := loadCacheOrFail()
				serviceName := c.Args().First()
				if service, ok := cache[serviceName]; ok {
					err := service.Stop()
					if err != nil {
						fatal(err)
					}
					fmt.Println(serviceName + " stopped")
				} else {
					fatal("Service not found")
				}
			},
		},
		{
			Name:  "restart",
			Usage: "Restart a service",
			Action: func(c *cli.Context) {
				cache := loadCacheOrFail()
				serviceName := c.Args().First()
				if service, ok := cache[serviceName]; ok {
					err := service.Restart()
					if err != nil {
						fatal(err)
					}
					fmt.Println(serviceName + " restarted")
				} else {
					fatal("Service not found")
				}
			},
		},
	}

	app.Run(os.Args)

}
