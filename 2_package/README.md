# Módulo
Significa um conujunto de pacotes - criados / externos(do github)/ do próprio Go

```
    go mod init modulo OR github.com.br/thallesrangel/nome-repo
```

Quando se tem um módulo, pode-se rodar com o comando - toda modificação deve se rodar novamente
```
    go build
```

Para rodar um código com o build - toda modificação deve se rodar novamente
./modulo.exe


### Visibilidade

Visilidade do Go
Pacotes do mesmo nível conseguem visualizar os respectivos métodos, mesmo que público ou privado
Pacotes diferentes com letra minúsculas não se pode acessar o método
