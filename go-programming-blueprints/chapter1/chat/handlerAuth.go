package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/markbates/goth/gothic"
)

type authHandler struct {
	next HandlerWithData
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	currentUserInformationsJSON, err0 := gothic.GetFromSession("current_user_informations", r)

	if err0 != nil {
		// some other error
		http.Error(w, err0.Error(), http.StatusInternalServerError)
		return
	}

	var currentUserInformations map[string]string
	json.Unmarshal([]byte(currentUserInformationsJSON), &currentUserInformations)

	// _, err := r.Cookie("auth")

	// if err == http.ErrNoCookie {
	// 	// not authenticated
	// 	w.Header().Set("Location", "/login")
	// 	w.WriteHeader(http.StatusTemporaryRedirect)
	// 	return
	// }

	// if err != nil {
	// 	// some other error
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// success - call the next handler

	i := map[string]interface{}{
		"UserID":    currentUserInformations["user_id"],
		"FirstName": currentUserInformations["first_name"],
		"LastName":  currentUserInformations["last_name"],
		"Nickname":  currentUserInformations["nickname"],
		"AvatarURL": currentUserInformations["avatar_url"],
		"Email":     currentUserInformations["email"],
	}
	h.next.SetData(i)
	h.next.ServeHTTP(w, r)
}

// MustAuth mark auth as required
func MustAuth(handler HandlerWithData) http.Handler {
	return &authHandler{next: handler}
}

// loginHandler handles the third-party login process.
// format: /auth/{action}/{provider}
func loginHandler(resW http.ResponseWriter, req *http.Request) {
	segs := strings.Split(req.URL.Path, "/")
	// provider := segs[2]
	action := segs[2]
	log.Println(action)
	switch action {
	case "login":
		{
			// log.Println("TODO handle login for", provider)
			// try to get the user without re-authenticating
			if gothUser, err := gothic.CompleteUserAuth(resW, req); err == nil {
				// t, _ := template.New("foo").ParseFiles(filepath.Join("templates", "user.template"))
				// t.Execute(resW, gothUser)

				m := md5.New()
				io.WriteString(m, strings.ToLower(gothUser.Email))
				userID := fmt.Sprintf("%x", m.Sum(nil))

				currentUserInformations := map[string]string{
					"user_id":    userID,
					"first_name": gothUser.FirstName,
					"last_name":  gothUser.LastName,
					"nickname":   gothUser.NickName,
					"email":      gothUser.Email,
					"avatar_url": gothUser.AvatarURL}
				currentUserInformationsJSON, _ := json.Marshal(currentUserInformations)
				gothic.StoreInSession("current_user_informations", string(currentUserInformationsJSON), req, resW)

				resW.Header().Set("Location", "/user")
				resW.WriteHeader(http.StatusTemporaryRedirect)
			} else {
				gothic.BeginAuthHandler(resW, req)
			}
			break
		}
	case "logout":
		{
			// log.Println("TODO handle logout for", provider)
			gothic.Logout(resW, req)
			resW.Header().Set("Location", "/login")
			resW.WriteHeader(http.StatusTemporaryRedirect)
			break
		}
	case "callback":
		{
			// log.Println("TODO handle callback for", provider)
			gothUser, err := gothic.CompleteUserAuth(resW, req)
			if err != nil {
				fmt.Fprintln(resW, err)
				return
			}

			m := md5.New()
			io.WriteString(m, strings.ToLower(gothUser.Email))
			userID := fmt.Sprintf("%x", m.Sum(nil))

			currentUserInformations := map[string]string{
				"user_id":    userID,
				"first_name": gothUser.FirstName,
				"last_name":  gothUser.LastName,
				"nickname":   gothUser.NickName,
				"email":      gothUser.Email,
				"avatar_url": gothUser.AvatarURL}
			currentUserInformationsJSON, _ := json.Marshal(currentUserInformations)
			gothic.StoreInSession("current_user_informations", string(currentUserInformationsJSON), req, resW)

			// t, _ := template.New("foo").ParseFiles(filepath.Join("templates", "user.template"))
			// t.Execute(resW, user)

			resW.Header().Set("Location", "/user")
			resW.WriteHeader(http.StatusTemporaryRedirect)
			break
		}
	default:
		resW.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(resW, "Auth action %s not supported", action)
	}
}
