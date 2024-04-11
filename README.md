# Github clean package versions

> Por enquanto funciona somente com organizações

Este projeto visa facilitar a gestão de pacotes hospedados no GitHub Registry. Através deste utilitário, os usuários podem listar os pacotes disponíveis no registry e remover pacotes antigos de forma eficiente.

## Parametros

| Nome | Descrição | Obrigatório | Valor default |
| --- | --- | --- | --- |
| repo_token | Token de autorização do seu usuário do github. Lembrando que tem que ter permissão de escrita e leitura no registry | Sim | --- |
| organization | Nome da organização | Sim | --- |
| list_package_monitored | Lista dos pacotes que deseja que versões sejam removidas. Para multiplos pacotes separe usando virgula | Sim | --- |
| package_type | Tipo dos pacotes informados. Ex: container, npm, ... | Não | container |
| number_versions | Número de versões do pacote que se deseja manter | Não | 1 |

## Example

```
- name: Github clear package registry
  uses: matheusrosmaninho/github-clean-packages-versions@v1
  with:
    repo_token: "${{ secrets.GITHUB_TOKEN }}"
    organization: "terminalbaka"
    list_package_monitored: "traefik-production"

```

## Contribuição
Contribuições são bem-vindas! Sinta-se à vontade para abrir issues relatando problemas ou propor novas funcionalidades. Se desejar contribuir diretamente, por favor, siga os padrões estabelecidos para solicitações de pull.

## Licença
Este projeto está licenciado sob a Licença MIT.
