package middleware

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// SessionStore menyimpan data session di memory
// Untuk production, bisa diganti dengan Redis atau database
type SessionStore struct {
	sessions map[string]*Session // Map token session ke data session
	mu       sync.RWMutex        // Mutex untuk thread-safety
}

// Session menyimpan informasi session pengguna
type Session struct {
	Username  string    // Username yang login
	CreatedAt time.Time // Waktu session dibuat
	ExpiresAt time.Time // Waktu session kadaluarsa
}

// sessionStore adalah instance global session store
var sessionStore = &SessionStore{
	sessions: make(map[string]*Session),
}

// Durasi session: 24 jam
const sessionDuration = 24 * time.Hour

// Nama cookie untuk menyimpan token session
const sessionCookieName = "admin_session"

// CreateSession membuat session baru untuk user yang berhasil login
// Mengembalikan token session yang disimpan di cookie
func CreateSession(c *gin.Context, username string) string {
	// Generate token random yang aman secara kriptografi
	token := generateToken()

	sessionStore.mu.Lock()
	defer sessionStore.mu.Unlock()

	// Simpan session di memory
	sessionStore.sessions[token] = &Session{
		Username:  username,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(sessionDuration),
	}

	// Set cookie di browser pengguna
	c.SetCookie(
		sessionCookieName,
		token,
		int(sessionDuration.Seconds()),
		"/",
		"",    // Domain kosong = domain saat ini
		false, // Secure = false untuk development (set true di production)
		true,  // HttpOnly = true untuk mencegah JS mengakses cookie
	)

	return token
}

// DestroySession menghapus session (logout)
func DestroySession(c *gin.Context) {
	token, err := c.Cookie(sessionCookieName)
	if err != nil {
		return
	}

	sessionStore.mu.Lock()
	defer sessionStore.mu.Unlock()

	// Hapus session dari memory
	delete(sessionStore.sessions, token)

	// Hapus cookie dari browser
	c.SetCookie(sessionCookieName, "", -1, "/", "", false, true)
}

// AuthRequired adalah middleware yang memastikan request berasal dari admin yang sudah login
// Jika belum login, redirect ke halaman login
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(sessionCookieName)
		if err != nil {
			// Tidak ada cookie session — redirect ke login
			c.Redirect(http.StatusFound, "/admin/login")
			c.Abort()
			return
		}

		sessionStore.mu.RLock()
		session, exists := sessionStore.sessions[token]
		sessionStore.mu.RUnlock()

		if !exists {
			// Token tidak valid — redirect ke login
			c.Redirect(http.StatusFound, "/admin/login")
			c.Abort()
			return
		}

		// Cek apakah session sudah kadaluarsa
		if time.Now().After(session.ExpiresAt) {
			// Session kadaluarsa — hapus dan redirect ke login
			DestroySession(c)
			c.Redirect(http.StatusFound, "/admin/login")
			c.Abort()
			return
		}

		// Set username di context agar bisa diakses oleh handler
		c.Set("admin_username", session.Username)
		c.Next()
	}
}

// generateToken membuat token random 32 byte (64 karakter hex)
// menggunakan crypto/rand yang aman secara kriptografi
func generateToken() string {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		// Fallback jika crypto/rand gagal (sangat jarang terjadi)
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}
	return hex.EncodeToString(bytes)
}
