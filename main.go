package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
)

func fetchLatestVersion() (string, error) {
	resp, err := http.Get("https://github.com/neovim/neovim/releases/tag/nightly")
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func getCurrentVersion() (string, error) {
	cmd := exec.Command("nvim", "--version")
	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(stdout), nil
}

func getVersion(content string) string {
	pattern := regexp.MustCompile("NVIM v.*-[a-z](?P<Commit>\\w{9})")

	matches := pattern.FindStringSubmatch(string(content))
	index := pattern.SubexpIndex("Commit")

	return matches[index]
}

func downloadLatestVersion(dest string) error {
	resp, err := http.Get("https://github.com/neovim/neovim/releases/download/nightly/nvim.appimage")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write body to file
	_, err = io.Copy(out, resp.Body)

	return err
}

func main() {
	// Init destination flag
	dest := flag.String("dest", "/usr/bin/nvim", "executable directory destination")
	check := flag.Bool("check", false, "check only if a new version is available")

	flag.Parse()

	// Latest version from GitHub
	log.Println("Fetching latest version of neovim")
	content, err := fetchLatestVersion()
	if err != nil {
		log.Fatalf("Fatal error when fetching neovim latest version: %v", err)
	}
	latest := getVersion(content)
	log.Printf("Latest neovim nightly version is %s", latest)

	// Current version
	content, err = getCurrentVersion()
	if err != nil {
		log.Fatalf("Error when getting nvim current version: %v", err)
	}
	current := getVersion(content)
	log.Printf("Current neovim version is %s", current)

	// Exit if already on latest version
	if latest == current {
		log.Printf("Already at the latest version: latest=%s current=%s", latest, current)
		os.Exit(0)
	} else if *check {
		log.Printf("A new version is available: latest=%s current=%s", latest, current)
		os.Exit(0)
	}

	// Download latest nightly release
	if err := downloadLatestVersion(*dest); err != nil {
		log.Fatalf("Fatal error when downloading latest nightly release: %v", err)
	}
	log.Printf("Successfully upgraded neovim to latest version: version=%s dest=%s", latest, *dest)
}
