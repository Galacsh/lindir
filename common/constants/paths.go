package constants

import "os"

const (
	APP_DIR         = "." + APP_UNDERSCORE
	TRACKER         = APP_DIR + string(os.PathSeparator) + "tracker"
	CONNECTOR       = APP_DIR + string(os.PathSeparator) + "connector"
	IGNORE_PATTERNS = "." + APP_UNDERSCORE + "ignore"
)
