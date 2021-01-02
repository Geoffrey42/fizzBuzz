CRI=docker
COMPOSE=$(CRI)-compose
NAME=fizzbuzz

all: $(NAME)

$(NAME):
	$(COMPOSE) -f $(COMPOSE).yml -f $(COMPOSE).prod.yml up -d --build

dev:
	$(COMPOSE) up -d --build

hooks:
	@git config --local core.hooksPath "./.hooks"

commit:
	@gommit check range HEAD~2 HEAD

redis-cli:
	$(CRI) exec -it fb-redis /usr/local/bin/redis-cli

ps:
	$(COMPOSE) ps

re: fclean dev

fclean: clean
	$(CRI) image rm fb-api-img

clean:
	$(COMPOSE) down

.PHONY: clean fclean re all redis-cli dev hooks