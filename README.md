# authenticator
Authenticator middleware for Negroni

Authenticator is a simple authentication library that does not use basic authentication, password files or any other local authentication. Authenticator checks every url for "admin" being in the path, if there is such a string in the path, it checks for the cookie "user" and if the value of the cookie is correctly encrypted. If the cookie is in place, and is correctly encrypted, it will allow you to continue. If the cookie is not in place or if the cookie has in any way been tampered with,it will call http code 401 and redirect the user to a login page.

What the contents of the cookie should be is up to the programmer. Popular method is to keep an encoded json object with the user information, that way it is easy to get the cookie, decrypt it and check the user settings, is_admin, has rights and so forth.

This module is based on the authenticator from Python Tornado

This does not slow the server down if the page does not contain admin in the url.

To install
~~~~
go get -u github.com/codegangsta/negroni
go get -u github.com/traustitj/authenticator
go get -u github.com/gorilla/securecookie
~~~~
To Use in a webserver

~~~ go
package main

import (
  "github.com/codegangsta/negroni"
  "github.com/traustitj/authenticator"
  "net/http"
  "fmt"
  )
  
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the login page!")
	})
	
	mux.HandleFunc("/admin", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the super secret admin page!")
	})

	n := negroni.Classic()
	n.Use(authenticator.NewAuthenticator())
	n.UseHandler(mux)
	n.Run(":8080")
}
~~~~

What remains to be done.

- Add example code
- Add new creator like NewAuthenticator but with custom values
