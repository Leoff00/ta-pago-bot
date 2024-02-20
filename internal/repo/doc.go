/*
	repo orchestrate DB queries and persistences.

- Should be used by services to perform DB operations.

- Permits inject/plug any DB-repo type(eg. Sqlite, MYSql, PgSQL, MongoDB) in service without changing the business logic.

- Should be 'mockable' for testing.
*/
package repo
