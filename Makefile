dev:
	docker compose up --build --remove-orphans --detach --force-recreate;
	code --folder-uri vscode-remote://attached-container+$$(printf "tester_dummy_example-golang-1" | xxd -p)/var/www/back

exec:
	docker compose exec golang bash

exec-root:
	docker compose exec --user=root golang bash

stop:
	docker compose stop

test:
	docker composer run golang go test ./...


# inside container

test:
	go test ./...
