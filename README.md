brewservice - easily start and stop services installed by homebrew
==================================================================

If you have a whole bunch of services, it's probably not a good idea to have them all start automatically when you boot your machine. This tool parses your `/usr/local/Cellar/` directory to find all of the `homebrew.*.*.plist` files and lets you turn them on and off without typing out the tedious `launchctl` commands every time.

## Installation
```
$ go get github.com/jraede/brewservice
```

## Commands
### Help
```
$ brewservice --help
```

View available commands.

### Update
```
$ brewservice update
```

Parse the directory for `homebrew.*.*.plist` files and save them to a cached file for easy access. You should run this whenever you run `brew install` or `brew update`.

### List
```
$ brewservice list
```

List available services

### Status
```
$ brewservice status
```

Show on/off status of all available services

### Start
``` 
$ brewservice start <service>
```

Start a service

### Stop
```
$ brewservice stop <service>
```

Stop a service

### Restart
```
$ brewservice restart <service>
```

Restart a service
