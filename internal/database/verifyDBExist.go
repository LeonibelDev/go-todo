package database

func VerifyDBexist() bool {
	db := DbConnection()
	defer db.Close()

	createTable := `
		CREATE TABLE IF NOT EXISTS Tasks(
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			description TEXT,
			status TEXT DEFAULT 'pending',
			datecreated DEFAULT CURRENT_TIMESTAMP
		);
	`
	_, err := db.Exec(createTable)
	return err == nil
}
