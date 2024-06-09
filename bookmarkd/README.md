# generate new table
cd internal
go run -mod=mod entgo.io/ent/cmd/ent new Account

# generate the ent files
go generate ./internal/ent
