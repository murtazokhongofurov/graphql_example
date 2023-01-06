run:
	go run cmd/app/main.go

proto-gen:
	./script/gen-proto.sh
	
add-submodule:
	git submodule add git@gitlab.com:crm117/protos.git

update-submodule:
	git submodule update --remote --merge

migrate_groups:
	migrate create -ext sql -dir migrations -seq create_groups_table

migrate_up:
	migrate -path migrations/ -database postgres://developer:2002@localhost:5432/groups?sslmode=disable up

migrate_down:
	migrate -path migrations/ -database postgres://developer:2002@localhost:5432/groups?sslmode=disable down

migrate_force:
	migrate -path migrations/ -database postgres://developer:2002@localhost:5432/groups?sslmode=disable force 1

up-version:
	go get github.com/99designs/gqlgen@v0.17.22
	go get github.com/99designs/gqlgen/internal/imports@v0.17.22
	go get github.com/99designs/gqlgen/codegen/config@v0.17.22
	go get github.com/99designs/gqlgen/internal/imports@v0.17.22
	go run github.com/99designs/gqlgen generated