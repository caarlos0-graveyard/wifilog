package wifi

type Wifi interface {
	SSID() (string, error)
}

func New() Wifi {
	return OSX{}
}
