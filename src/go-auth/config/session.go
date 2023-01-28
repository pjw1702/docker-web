// go get github.com/gorilla/sessions
package config

import "github.com/gorilla/sessions"

const SESSOION_ID = "go_auth_sess"

// Store Session where Web browsers as cookie
var Store = sessions.NewCookieStore([]byte("afsafdsa092743qwreqwr"))
