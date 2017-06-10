deps:
	@printf "Installing all dependencies\n\n"
	@go get -u -v github.com/githubnemo/CompileDaemon
	@printf "Installing glide dependencies\n\n"
	@glide install

install:
	@printf "Installing CryptoPunks\n\n"
	go install -v github.com/disiqueira/CryptoPunks/cmd/cps

local:
	@printf "Installing Pre-commit\n\n"
	@brew install pre-commit

	@printf "Installing Pre-commit hooks\n\n"
	@pre-commit install
