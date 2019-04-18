/**
 * @author Jose Nidhin
 */
package system

import (
	"context"
	"database/sql"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

func NewPGDB(appConfig *Config) *sql.DB {
	db, err := sql.Open("postgres", pgConnStr(appConfig.Database.PG.Default))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		panic(err)
	}

	return db
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

func pgConnStrPart(key, val string) string {
	if val != "" {
		return key + "=" + val + " "
	}
	return ""
}
