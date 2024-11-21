package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

var (
	staticDir     = "/home/workspace/dashboard/frontend/build" // Directory for static files
	backendURLStr = "http://localhost:4000"                    // Backend server URL
)

func init(){
	staticDir = os.Getenv("STATIC_DIR")
	if staticDir == "" {
		staticDir = "/home/workspace/dashboard/frontend/build"
	}
	backendURLStr = os.Getenv("BACKEND_URL")
	if backendURLStr == "" {
		backendURLStr = "http://localhost:4000"
	}
}

func main() {
	// Parse the backend URL
	backendURL, err := url.Parse(backendURLStr)
	if err != nil {
		log.Fatalf("Failed to parse backend URL: %v", err)
	}

	// Create a reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(backendURL)

	// Static file server handler
	fs := http.FileServer(http.Dir(staticDir))

	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		// Add CORS headers if needed
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Strip "/api" prefix before proxying
		// r.URL.Path = r.URL.Path[4:]
		proxy.ServeHTTP(w, r)
	})

	// Main handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Sanitize path to prevent directory traversal
		cleanPath := filepath.Clean(r.URL.Path)
		staticFilePath := filepath.Join(staticDir, cleanPath)

		// If the path is "/" serve index.html
		if r.URL.Path == "/" {
			serveStaticFile(w, r, "index.html")
			return
		}

		// Ensure the path is still within the static directory
		if !strings.HasPrefix(staticFilePath, staticDir) {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		if fileExists(staticFilePath) {
			fs.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not found"))
			log.Printf("File not found: %s", staticFilePath)
		}
	})

	// Start the server
	log.Println("Starting server on :3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

// serveStaticFile serves a specific file from the static directory
func serveStaticFile(w http.ResponseWriter, r *http.Request, filePath string) {
	fullPath := filepath.Join(staticDir, filePath)
	http.ServeFile(w, r, fullPath)
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}
