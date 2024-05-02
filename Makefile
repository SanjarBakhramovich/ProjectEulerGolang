.PHONY: commit push

commit:
	@read -p "Enter commit message: " message; \
	git add .; \
	git commit -m "$$message"; \

push: commit
	git push -u origin main
