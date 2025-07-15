dev:
	docker-compose -f docker-compose.yml up -d postgres redis

ps:
	docker-compose -f docker-compose.yml  ps -a

down:
	docker-compose -f docker-compose.yml down

rebuild:
	docker-compose -f docker-compose.yml up --build --force-recreate
air:
	cd app && air -c .air.toml
