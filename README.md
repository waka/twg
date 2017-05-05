# twg

twg is twitter client for cli, written by golang.

## Install

```
$ go get -u https://github.com/waka/twg-go
```

## Start options

```
$ twg -h
Usage:
  twg [OPTIONS]

Application Options:
  -m, --mode=    you can select timeline(default) or mentions or list:slug
  -t, --tweet=   update your status, and finish
  -r, --reload   do auto reload by 2 minutes
  -v, --version

Help Options:
  -h, --help     Show this help message
```

## Key binds

| Key        | Action        |
|:-----------|:--------------|
| ctrl + q   | Quit app      |
| ctrl + r   | Reload tweets |
| up or k    | Up cursor     |
| down or j  | Up cursor     |
| :          | Command mode  |
| ctrl + c   | Normal mode   |

## Command mode

- `:tl` - show home timeline
- `:mentions` - show mentions tweets
- `:list {slug}` - show list timeline
- `:tweet {text}` - post tweet
- `:reply {text}` - reply to selected tweet
- `:fav` - fav to selected tweet
- `:rt` - retweet to selected tweet

## Reset access token

Remove configuration file.

```
$ rm ~/.twg
```

## Build in local machine

twg use [dep](https://github.com/golang/dep) to resolve dependencies.

```
$ make deps
$ make
$ ./bin/twg
```
