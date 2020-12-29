hooks:
	@git config --local core.hooksPath "./.hooks"

commit:
	@gommit check range HEAD~2 HEAD