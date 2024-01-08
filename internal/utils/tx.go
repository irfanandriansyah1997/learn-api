//nolint:typecheck // disable lint since PanicIfError still on same package
package utils

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()

	if err != nil {
		errorRollback := tx.Rollback()
		PanicIfError(errorRollback)

		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicIfError(errorCommit)
	}
}
