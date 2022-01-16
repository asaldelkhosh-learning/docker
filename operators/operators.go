package operators

var (
	Database   = "database transaction"
	Preprocess = "user process"
	Syslog     = "system log"
)

func List() []string {
	return []string{Database, Preprocess, Syslog}
}
