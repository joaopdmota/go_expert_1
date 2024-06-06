package queries

func ExecCreateTableQuery() string {
	return `CREATE TABLE IF NOT EXISTS exchange (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"createdAt" DATETIME DEFAULT CURRENT_TIMESTAMP,
		"value" VARCHAR(255)
	  );`
}

func ExecInserIntoQuery(bid string) string {
	return "INSERT INTO exchange (value) VALUES(" + bid + ")"
}
