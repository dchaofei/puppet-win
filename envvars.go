package puppet_win

import (
	"fmt"
	"os"
)

func envWinAPIServer(server string) (string, error) {
	if server != "" {
		return server, nil
	}

	server = os.Getenv("WECHATY_PUPPET_WIN_API_SERVER")
	if server != "" {
		return server, nil
	}

	return "", fmt.Errorf("puppet-win: WECHATY_PUPPET_WIN_API_SERVER is empty")
}

func envWebsocketHost(host string) (string, error) {
	if host != "" {
		return host, nil
	}

	host = os.Getenv("WECHATY_PUPPET_WIN_WEBSOCKET_HOST")
	if host != "" {
		return host, nil
	}

	return "", fmt.Errorf("puppet-win: WECHATY_PUPPET_WIN_WEBSOCKET_HOST is empty")
}

func envWebsocketPort(port string) (string, error) {
	if port != "" {
		return port, nil
	}

	port = os.Getenv("WECHATY_PUPPET_WIN_WEBSOCKET_PORT")
	if port != "" {
		return port, nil
	}

	return "", fmt.Errorf("puppet-win: WECHATY_PUPPET_WIN_WEBSOCKET_PORT is empty")
}
