package migrations

import "database/sql"

type CreateAccountsTable struct{}

func (m *CreateAccountsTable) GetName() string {
	return "create_accounts_table"
}

func (m *CreateAccountsTable) Up(db *sql.DB) error {
	stmt := "CREATE TABLE accounts (" +
		"id VARCHAR(36) NOT NULL," +
		"nickname VARCHAR(12) NOT NULL," +
		"slug VARCHAR(12) NOT NULL," +
		"dt_created DATETIME," +
		"PRIMARY KEY (id)," +
		"UNIQUE KEY unique_nickname (nickname)," +
		"UNIQUE KEY unique_slug (slug)" +
		");"
	_, err := db.Exec(stmt)
	return err
}

func (m *CreateAccountsTable) Down(db *sql.DB) error {
	stmt := "DROP TABLE accounts;"
	_, err := db.Exec(stmt)
	return err
}
