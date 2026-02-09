# Error Wrapping and Error Chaining in Go

## Error Wrapping (`%w`) — Encadeando erros

O Go permite encapsular (wrap) um erro dentro de outro utilizando `fmt.Errorf` com o verbo `%w`.

Isso cria uma **cadeia de erros**, onde o erro original é preservado dentro do novo erro.

Exemplo:
```go
return fmt.Errorf("buscar usuário: %w", ErrNotFound)
```
O que acontece:

- Um novo erro é criado
- O erro original (`ErrNotFound`) é mantido internamente
- O contexto é adicionado ("buscar usuário")
- Permite verificar o erro original posteriormente

Vantagens:

- Preserva contexto do erro
- Mantém rastreabilidade entre camadas (repo → service → controller)
- Permite usar `errors.Is` e `errors.As`
- Facilita debugging

---

## Error Formatting (`%v`) — Apenas mensagem

Usar `%v` dentro do `fmt.Errorf` **não encapsula o erro**.

Exemplo:
```go
return fmt.Errorf("buscar usuário: %v", ErrNotFound)
```
O que acontece:

- Apenas a string do erro é utilizada
- O erro original NÃO é preservado
- A cadeia de erros é perdida

Consequências:

- `errors.Is` não funciona
- Não é possível detectar o erro original

---

## Verificando erros com `errors.Is`

A função `errors.Is` percorre a cadeia de erros procurando um erro específico.
```go
errors.Is(err, ErrNotFound)
```
Se o erro foi criado usando `%w`:

- retorna `true`

Se foi criado usando `%v`:

- retorna `false`

---

## Cadeia de erros (Error Chain)

Quando usamos `%w`, criamos uma estrutura semelhante a:
```
buscar usuário
└── not found (erro original)
```
Isso permite adicionar contexto em múltiplas camadas sem perder a causa raiz.

---

## Boas práticas

- Use `%w` ao **propagar erros** entre funções
- Use `%v` apenas para formatação de texto ou logs
- Sempre prefira adicionar contexto ao retornar erros
- Evite comparar mensagens de erro diretamente (`err.Error()`)

---

## Observações importantes

- `%w` só funciona dentro de `fmt.Errorf`
- Apenas um `%w` é permitido por chamada
- `fmt.Println(err)` já chama `err.Error()` internamente
- Erros em Go são valores, não exceções
