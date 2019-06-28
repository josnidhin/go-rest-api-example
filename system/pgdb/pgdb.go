/**
 * @author Jose Nidhin
 */
package pgdb

import (
	"context"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//
type PGConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Port     string `json:"port"`
	DBName   string `json:"dbname"`
	SSLMode  string `json:"sslMode"`
}

//
func New(pgConfig *PGConfig) (db *sqlx.DB) {
	connStr := pgConnStr(pgConfig)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	db, err := sqlx.ConnectContext(ctx, "postgres", connStr)
	if err != nil {
		panic(err)
	}

	return
}

//
func pgConnStr(pgConfig *PGConfig) string {
	var connStr strings.Builder

	connStr.WriteString(pgConnStrPart("host", pgConfig.Host))
	connStr.WriteString(pgConnStrPart("user", pgConfig.User))
	connStr.WriteString(pgConnStrPart("password", pgConfig.Password))
	connStr.WriteString(pgConnStrPart("port", pgConfig.Port))
	connStr.WriteString(pgConnStrPart("dbname", pgConfig.DBName))
	connStr.WriteString(pgConnStrPart("sslmode", pgConfig.SSLMode))

	return connStr.String()
}

//
func pgConnStrPart(key, val string) string {
	if val != "" {
		return key + "=" + val + " "
	}
	return ""
}
