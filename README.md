##goserve

Simple Server in Go

It is an alternative to `python -m SimpleHTTPServer".

Just install it by typing `go get github.com/hariharan-uno/goserve` and cd into any directory and run `goserve`.

You can also specify the port in which you want to serve. For example:
~~~sh
goserve -port=:6080
~~~

It is very useful if you are working with JS frameworks such as ReactJS/AngularJS, which require http for many things.