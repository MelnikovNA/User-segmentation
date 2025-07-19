DSN=sqlite3://segments.db

all: 
	migrate -source file://migrations/sqlite \
	-database "${DSN}" up



migrate:	
	go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest