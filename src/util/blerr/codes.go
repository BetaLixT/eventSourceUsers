package blerr

type ErrorCode struct {
	code    int
	message string
}

// - Startup errors
var InvalidPortCode = ErrorCode{code: 1000, message: "InvalidPort"}
var ConfigLoadFailureCode = ErrorCode{code: 1001, message: "ConfigLoadFailure"}
var LoggerCreateFailureCode = ErrorCode{
	code:    1002,
	message: "LoggerCreateFailure",
}

// - Common errors
var UnexpectedErrorCode = ErrorCode{
	code:    10000,
	message: "UnexpectedError",
}
var NotImplementedCode = ErrorCode{
	code:    10001,
	message: "NotImplemented",
}
var UnsetResponse = ErrorCode{
	code:    10002,
	message: "UnsetResponse",
}
var InvalidStatusCode = ErrorCode{
	code:    10003,
	message: "InvalidStatusCode",
}

// - Database errors
var DatabaseConnectionOpenFailure = ErrorCode{
	code:    2000,
	message: "DatabaseConnectionOpenFailure",
}
var DatabasePingFailure = ErrorCode{
	code:    2001,
	message: "DatabasePingFailure",
}
var DatabaseMigrationMigrationHistoryMismatchCode = ErrorCode{
	code:    2002,
	message: "DatabaseMigrationMigrationHistoryMismatch",
}
var DatabaseMigrationProcedureCheckFailedCode = ErrorCode{
	code:    2003,
	message: "DatabaseMigrationProcedureCheckFailed",
}
var DatabaseMigrationMigrationCheckFailedCode = ErrorCode{
	code:    2004,
	message: "DatabaseMigrationMigrationCheckFailed",
}
var DatabaseMigrationMigrationFetchFailedCode = ErrorCode{
	code:    2005,
	message: "DatabaseMigrationMigrationFetchFailed",
}

// - Auth errors
// - Auth errors
var TokenInvalidCode = ErrorCode{
	code:    4000,
	message: "TokenInvalid",
}
