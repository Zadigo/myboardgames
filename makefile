runbasegame:
	cd services/basegame && go run .
rungenerateprotofiles:
	buf lint && buf generate && cd gen && go mod init myboardgames/gen && go mod tidy && cd .. && go work init ./gen ./services/minimalgo ./services/minimalgame
