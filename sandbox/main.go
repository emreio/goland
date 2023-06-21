package main

import (
	"crypto/rand"
	"encoding/base64"
	"html/template"
	"net/http"
	"sync"
	"time"
)

// Define the credentials
const (
	username = "emre"
	password = "123"
)

// Define the session struct
type Session struct {
	ID        string
	Username  string
	ExpiresAt time.Time
}

// Define the session manager
type SessionManager struct {
	mux      sync.Mutex
	sessions map[string]*Session
}

// Generate a random session ID
func generateSessionID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {

		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// Create a new session for the given username
func (sm *SessionManager) CreateSession(username string) (*Session, error) {
	sm.mux.Lock()
	defer sm.mux.Unlock()

	// Generate a new session ID
	sessionID, err := generateSessionID()
	if err != nil {
		return nil, err
	}

	// Create a new session with the given username and expiration time
	session := &Session{
		ID:        sessionID,
		Username:  username,
		ExpiresAt: time.Now().Add(time.Hour),
	}

	// Add the session to the session manager
	sm.sessions[sessionID] = session

	return session, nil
}

// Get the session with the given session ID
func (sm *SessionManager) GetSession(sessionID string) *Session {
	sm.mux.Lock()
	defer sm.mux.Unlock()

	// Get the session with the given session ID
	session, ok := sm.sessions[sessionID]
	if !ok {
		return nil
	}

	// Delete the session if it has expired
	if session.ExpiresAt.Before(time.Now()) {
		delete(sm.sessions, sessionID)
		return nil
	}

	return session
}

// Define the session manager instance
var sessionManager = &SessionManager{
	sessions: make(map[string]*Session),
}

func main() {
	// Define the HTTP route for the login page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Check if the user is already logged in
		sessionID, err := r.Cookie("session_id")
		if err == nil {
			session := sessionManager.GetSession(sessionID.Value)
			if session != nil {
				// If the user is logged in, redirect them to the home page
				http.Redirect(w, r, "/home", http.StatusSeeOther)
				return
			}
		}

		// If the user is not logged in, display the login page
		if r.Method != http.MethodPost {
			// Render the login page template
			tmpl, err := template.ParseFiles("login.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			tmpl.Execute(w, nil)
			return
		}

		// // If the user submitted the login form, check the credentials
		// u := r.FormValue("username")
		// p := r.FormValue("password")
		// if u == username && p == password
	})
}
