# Servito
[![Build Status](https://travis-ci.org/GHawk1ns/Servito.svg?branch=master)](https://travis-ci.org/GHawk1ns/Servito)
[![GoDoc](https://godoc.org/github.com/ghawk1ns/Servito?status.svg)](https://godoc.org/github.com/ghawk1ns/Servito)

A small HTTP server in Go. A great start for small projects and prototyping.

# Install
~~~~
$ go get github.com/ghawk1ns/servito
~~~~


# Example
https://github.com/GHawk1ns/Servito/blob/master/example/main.go

# Config
You can set these values individually or pass servito a JSON string, file, or path to a json file
~~~
{
  "debug": true,
  "address": string,
  "port": string,
  "readTimeout": int,
  "writeTimeout": int,
  "csrfEnable", bool,
  "CSRFKey", string
}
~~~
