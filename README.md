# Urubu do Pix Back-end

O back-end do Urubu do Pix é uma API feita para estudos. A API é totalmente fictícia e não utiliza dados reais.

![example](images/image.png)

## Tecnologias:

- Go v1.23.4
- Fiber
- MongoDB
- JWT

## TO-DO
- [ ] **Saque**
  - [X] Criar endpoint para saque;
  - [X] Validar saldo do usuário;
  - [X] Atualizar saldo após saque;
  - [ ] Verificar minimo de 30 dias.
- [X] **Depósito**
  - [X] Criar endpoint para depósito;
  - [X] Validar valor do depósito;
  - [X] Atualizar saldo após depósito.
- [X] **Transferência**
  - [X] Criar endpoint para transferência;
  - [X] Validar autenticacao;
  - [X] Validar saldo e destinatário;
  - [X] Atualizar saldos após transferência.
- [ ] **Rendimento**
   - [X] Cronjob para rendimento diario;
   - [X] Rendimento variado entre 4% e 8%;
   - [ ] Adicionar o rendimento nas transacoes.


## Contribuição:

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues e pull requests.