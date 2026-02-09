# go:embed and go:generate

## GO:EMBED - Inserindo arquivos no binario

Uma diretiva de pré-compilação introduzida no Go `1.16` que permite incorporar arquivos ou diretórios inteiros diretamente no binário compilado.

O conteúdo dos arquivos é lido durante a compilação e transformado em variáveis no seu programa Go.

Vantagens:

- Single Binary  
- Sem Erros de Caminho  
- Atomic Deployment  
- Tipos de Variáveis  

Pontos de Atenção:

- Cuidado com o Tamanho  
- Conteúdo Dinâmico  
- Segurança  

## GO:GENERATE – Automatizando geração de código

Uma diretiva especial que instrui a ferramenta `go generate` a executar comandos arbitrários durante o desenvolvimento, geralmente utilizada para gerar código, arquivos auxiliares ou artefatos a partir de fontes externas.  
Os comandos são executados manualmente pelo desenvolvedor e não fazem parte do processo de compilação padrão.

Pontos de Atenção:

- Documentação  
- Idempotência  
- Performance  
- Versionamento do Código Gerado  
