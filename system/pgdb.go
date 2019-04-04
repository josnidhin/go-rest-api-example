/**
 * @author Jose Nidhin
 */
package system

import (
	"database/sql"
	"strings"

	_ "github.com/lib/pq"
)

var db *sql.DB

func NewPGDB(appConfig *Config) {

}

func pgConnStr(pgConfig *PGConfig) string {
	var connStr strings.Builder

	connStr.WriteString(pgConnStrPart("host", pgConfig.Host))
	connStr.WriteString(pgConnStrPart("user", pgConfig.User))
	connStr.WriteString(pgConnStrPart("password", pgConfig.Password))
	connStr.WriteString(pgConnStrPart("port", pgConfig.Port))
	connStr.WriteString(pgConnStrPart("dbname", pgConfig.DBName))

	return connStr.String()
}

func pgConnStrPart(key, val string) {
	if val != "" {
		return key + "=" val + " "
	}
	return ""
}
