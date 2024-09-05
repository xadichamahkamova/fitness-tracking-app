MIGRATE_FLAGS=-path migrations -database "postgres://postgres:1@localhost:5432/postgres?sslmode=disable"

mig-create:
	@migrate create -ext sql -dir migrations -seq create_users_table

mig-up:
	@migrate $(MIGRATE_FLAGS) up

mig-down:
	@migrate $(MIGRATE_FLAGS) down
	
mig-force:
	@migrate $(MIGRATE_FLAGS) force 1