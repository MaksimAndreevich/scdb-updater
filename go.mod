module gitlab.com/scdb/updater

go 1.24.3

require (
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.10.9
	github.com/vinay03/chalk v1.1.4
	gitlab.com/scdb/core v0.1.2
	gitlab.com/scdb/database v0.1.2
)

replace gitlab.com/scdb/core => ../core

replace gitlab.com/scdb/database => ../database
