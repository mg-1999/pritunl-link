package constants

import (
	"path"
	"time"
)

const (
	Version                   = "1.0.548.17"
	VarDir                    = "/var/lib/pritunl"
	LogPath                   = "/var/log/pritunl.log"
	ConfPath                  = "/etc/pritunl-link.json"
	PublicIpServer            = "https://app.pritunl.com/ip"
	PublicIp6Server           = "https://app6.pritunl.com/ip"
	DefaultDiconnectedTimeout = 45 * time.Second
)

var (
	RoutesPath    = path.Join(VarDir, "routes")
	CurRoutesPath = path.Join(VarDir, "cur_routes")
)
