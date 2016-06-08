package wifi

import (
	"os/exec"
	"regexp"
)

const airport = "/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport"

type OSX struct {
}

func (osx OSX) SSID() (string, error) {
	output, err := exec.Command(airport, "-I").Output()
	if err != nil {
		return "", err
	}
	regexp := regexp.MustCompile(`.* SSID: (.*)`)
	return regexp.FindStringSubmatch(string(output[:]))[1], nil
}
