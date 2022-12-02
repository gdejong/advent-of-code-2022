package gen

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func WebInput(year, day int) []byte {
	// Fetch from Advent of Code website
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	c := http.Client{Timeout: time.Duration(3) * time.Second}

	session := &http.Cookie{
		Name:   "session",
		Value:  os.Getenv("AOC_SESSION"),
		MaxAge: 0,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error with HTTP request: %s", err.Error())
	}
	req.AddCookie(session)

	resp, err := c.Do(req)
	if err != nil {
		logrus.Errorf("Error: %s", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		logrus.Error("Error: Failed to contact AOC server, status code: " + resp.Status)
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("Error: %s", err)
		return nil
	}
	return body
}
