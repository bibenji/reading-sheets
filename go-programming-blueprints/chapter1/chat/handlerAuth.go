package main

import (
	"encoding/json"
	"fmt"
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
	log.Println("currentUserInformationsJSON", currentUserInformationsJSON, err0)

	var currentUserInformations map[string]string

	json.Unmarshal([]byte(currentUserInformationsJSON), &currentUserInformations)

	log.Println("currentUserInformations.first_name", currentUserInformations["first_name"])

	// firstName, err1 := gothic.GetFromSession("first_name", r)
	// log.Println("firstName", firstName, err1)

	// lastName, err2 := gothic.GetFromSession("last_name", r)
	// nickname, err3 := gothic.GetFromSession("nickname", r)
	// AvatarURL, err4 := gothic.GetFromSession("avatar_url", r)

	// session, _ := gothic.Store.Get(r, "123456789")
	// log.Println(session.Values)

	// log.Println("lastName", lastName, err2)
	// log.Println("nickname", nickname, err3)
	// log.Println("AvatarURL", AvatarURL, err4)

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
		"FirstName": currentUserInformations["first_name"],
		"LastName":  currentUserInformations["last_name"],
		"Nickname":  currentUserInformations["nickname"],
		"AvatarURL": currentUserInformations["avatar_url"],
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

				log.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAA")
				log.Println("gothUser.FirstName", gothUser.FirstName)
				log.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAA")

				currentUserInformations := map[string]string{
					"first_name": gothUser.FirstName,
					"last_name":  gothUser.LastName,
					"nickname":   gothUser.NickName,
					"avatar_url": gothUser.AvatarURL}

				currentUserInformationsJSON, _ := json.Marshal(currentUserInformations)

				gothic.StoreInSession("current_user_informations", string(currentUserInformationsJSON), req, resW)

				// session, _ := gothic.Store.Get(req, "123456789")

				// session.Values["first_name"] = gothUser.FirstName

				// // Save it before we write to the response/return from the handler.
				// err := session.Save(req, resW)
				// if err != nil {
				// 	http.Error(resW, err.Error(), http.StatusInternalServerError)
				// 	return
				// }

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

			log.Println("IIIIIIIIIIIIIIIIIIIIIIIIIIIIIIi")
			log.Println(gothUser.FirstName)
			log.Println("IIIIIIIIIIIIIIIIIIIIIIIIIIIIIIi")

			currentUserInformations := map[string]string{
				"first_name": gothUser.FirstName,
				"last_name":  gothUser.LastName,
				"nickname":   gothUser.NickName,
				"avatar_url": gothUser.AvatarURL}

			currentUserInformationsJSON, _ := json.Marshal(currentUserInformations)

			gothic.StoreInSession("current_user_informations", string(currentUserInformationsJSON), req, resW)

			// session, _ := gothic.Store.Get(req, "123456789")

			// session.Values["first_name"] = gothUser.FirstName

			// // Save it before we write to the response/return from the handler.
			// err := session.Save(req, resW)
			// if err != nil {
			// 	http.Error(resW, err.Error(), http.StatusInternalServerError)
			// 	return
			// }

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
