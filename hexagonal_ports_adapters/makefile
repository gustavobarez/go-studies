DB_DSN := postgresql://postgres:example@localhost:5432/database?sslmode=disable
JET_OUT := ./adapter/output
JET_TMP := ./jetgen

jet:
	@echo "→ Limpando geração anterior..."
	rm -rf $(JET_TMP)

	@echo "→ Gerando código Jet..."
	jet -dsn="$(DB_DSN)" -path=$(JET_TMP)

	@echo "→ Criando pasta final..."
	mkdir -p $(JET_OUT)

	@echo "→ Movendo apenas conteúdo dentro de public/..."
	mv $(JET_TMP)/*/public/* $(JET_OUT)/

	@echo "→ Removendo pasta temporária..."
	rm -rf $(JET_TMP)

	@echo "✔ Jet gerado com sucesso!"
