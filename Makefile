hooks:
	@git config --local core.hooksPath "./.hooks"

commit:
	@gommit check range HEAD~2 HEAD

redis-cli:
	docker exec -it fb-redis /usr/local/bin/redis-cli

ps:
	docker-compose ps