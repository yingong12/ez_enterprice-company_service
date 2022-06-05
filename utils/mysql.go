package utils

import "github.com/go-sql-driver/mysql"

func IsMysqlDupKeyErr(err error) bool {
	// make sure err is a mysql.MySQLError.
	if errMySQL, ok := err.(*mysql.MySQLError); ok {
		switch errMySQL.Number {
		case 1062:
			// TODO handle Error 1062: Duplicate entry '%s' for key %d
			return true
		}
	}
	return false
}
