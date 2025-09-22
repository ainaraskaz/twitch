package auth

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/sessions"
)

type tHandler struct {
	twitch  *TwitchAuth
	store   *sessions.CookieStore
	session string
}

func NewHandler(twitch *TwitchAuth, store *sessions.CookieStore) *tHandler {
	return &tHandler{
		twitch:  twitch,
		store:   store,
		session: "twitch-auth-session",
	}
}

// GET /login
func (h *tHandler) Login(w http.ResponseWriter, r *http.Request) {
	state := "random-state" // in real app: generate securely and store in session
	sess, _ := h.store.Get(r, h.session)
	sess.Values["state"] = state
	sess.Save(r, w)

	http.Redirect(w, r, h.twitch.GetAuthURL(state), http.StatusFound)
}

// GET /callback
func (h *tHandler) Callback(w http.ResponseWriter, r *http.Request) {
	sess, _ := h.store.Get(r, h.session)
	stateFromSession, _ := sess.Values["state"].(string)

	if r.URL.Query().Get("state") != stateFromSession {
		http.Error(w, "Invalid state", http.StatusBadRequest)
		return
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Missing code", http.StatusBadRequest)
		return
	}

	token, err := h.twitch.ExchangeCode(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange code: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Save token in session
	sess.Values["token"] = token.AccessToken
	sess.Save(r, w)

	// Fetch user info
	req, _ := http.NewRequest("GET", "https://api.twitch.tv/helix/users", nil)
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	req.Header.Set("Client-Id", h.twitch.OAuthConfig.ClientID)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Failed to get user info: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	// Debug print or return as JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

// GET /profile (protected route)
func (h *tHandler) Profile(w http.ResponseWriter, r *http.Request) {
	sess, _ := h.store.Get(r, h.session)
	token, ok := sess.Values["token"].(string)
	if !ok || token == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	fmt.Fprintf(w, "You are logged in with token: %s", token)
}
