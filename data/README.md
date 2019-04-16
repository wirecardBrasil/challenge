# Wirecard Data Challenge


## Case 1: - Modelagem dimensional

    Atualmente um dos nossos desafios é definir o modelo dimensional para a tarifação
das nossas contas. Temos que efetuar consultas SQL em uma série de tabelas e
considerando diversas regras para conseguirmos obter a informação desejada. Diante
disto, seu primeiro desafio será definir um modelo dimensional que otimize o acesso a este
tipo de informação.

    Além das informações sobre as taxas, através do modelo dimensional proposto
devemos ser capazes de acessar algumas informações detalhadas das contas (login,
nome, etc.).

    Como resultado da modelagem dimensional, você deve enviar o modelo proposto.
O modelo poder ser desenvolvido através da ferramenta open source MySQLWorkbench,
ou através de alguma outra ferramenta. Caso outra ferramenta seja utilizada, Você
precisará converter o modelo final em um dos seguintes formatos: xls, pdf, png ou jpeg.
Os parágrafos posteriores irão te contextualizar quanto as regras de negócio e os
arquivos que devem ser utilizados durante o desenvolvimento.


#### Regras de Negócio
    Obrigatoriamente, devemos aplicar duas taxas em todo e qualquer pagamento: uma
taxa fixa (definida em reais); e uma taxa percentual, que deve ser aplicada com base no
valor total do pagamento. Os valores dessas taxas podem variar de acordo com o modelo
de precificação.

    Além dessas duas taxas, caso o cliente queira antecipar o valor que ele tem para
receber, devemos aplicar uma terceira taxa, a taxa de antecipação. Para esta taxa, as
contas Moip são tarifadas de acordo com o modelo de precificação definido em contrato.

    Atualmente temos três modelos de precificação:

- Modelo 1:
    Taxa percentual: aplicada de acordo com o tipo de pagamento.
    Taxa de antecipação: aplicar uma taxa fixa de 1,99% a.m. (juros compostos)

- Modelo 2:
    Taxa percentual: aplicada de acordo com o tipo de pagamento e agrupamento de quantidade de parcelas.
    Taxa de antecipação: aplicar uma taxa percentual de acordo com o meio de
pagamento. Portanto, cada meio de pagamento terá sua taxa específica. Caso um meio de
pagamento não tenha uma taxa atrelada, devemos considerar a mesma taxa de
antecipação utilizada pelo meio de pagamento Cartão de crédito.

- Modelo 3:

    Taxa percentual: aplicada de acordo com o tipo de pagamento e quantidade
de parcelas.
    
    Taxa de antecipação: aplicar uma taxa fixa de 2,89% a.m. (juros simples)
sobre a quantidade de dias que serão antecipados.
    
    Além dos 3 modelos descritos acima, embora cada uma das contas Moip possua
suas respectivas taxas, caso ela transacione (receba pagamentos) através de um channel,
as taxas que serão aplicadas àquela conta serão as taxas do channel, desde que a conta
não possua uma taxa negociada.
    
    Portanto, a precedência das taxas ocorre da seguinte forma (maior para menor
prioridade): taxa da conta se houver negociação, taxa do channel e taxa da conta.
    
    Um channel também possui uma conta Moip. Portanto, as taxas do channel são
definidas da mesma forma que as taxas de uma conta Moip comum. Ou seja, seguindo as
mesmas regras descritas acima, com exceção de que um channel não pode transacionar
por outro channel.

#### Arquivos:

    Os arquivos que você deve utilizar durante a modelagem são os listados abaixo e
que estão no zip no repositório:

- account: contém a lista de contas que deverão ser utilizadas durante os cases
(52 contas)

- fixed_table_fee: contém as taxas default das contas para os modelos 1 e 3 de
precificação

- tax_applied_to_account: contém as taxas customizadas das contas que
possuem o modelo 1 de precificação

- account_payment_fee: contém as taxas das contas que possuem o modelo 2
de precificação

- account_fixed_table_fee: contém as taxas customizadas das contas que
possuem o modelo 3 de precificação

- accounts_channels: contém as contas que transacionam por channels

- channel: contém as informações dos channels

- channel_fixed_table_fee: contém as taxas default dos channels (caso as taxas
não sejam informadas em outra tabela)

- member: contém informações mais detalhadas sobre as contas

- payment_form: contém as possíveis formas de pagamento

#### Observações:

- O modelo de precificação da conta pode ser obtido pelo campo account.fee_type
- É possível verificar se a conta possui uma taxa negociada através do campo
account.negotiated_tax

## Case 2: - ETL

    Agora que definimos o modelo dimensional, podemos seguir para a etapa de
implementação.

    Neste case seu desafio será implementar o modelo definido no case anterior. Você
pode utilizar tecnologias open source para isso (HDFS, sqoop, kafka, flume, hive, spark,
etc.). O foco desse case será nas etapas de extração e transformação dos dados.

    Como resultado da implementação do modelo, você deve enviar todos os scripts
desenvolvidos. É importante que você envie tudo que foi utilizado durante este etapa, pois
estes scripts serão validados e testados. Portanto, caso algum script esteja faltando não
será possível efetuar a validação.

    Use e abuse da sua criatividade. Lembre-se que aqui você participará de todo o
desenvolvimento do processo ETL, desde a definição do modelo até a disponibilização dos
dados.


## Regras Gerais

1. Você precisa utilizar todos os arquivos listados no case #01.

2. Caso você não tenha familiaridade com tecnologias Big Data, pode utilizar o SGBD
MySQL (open source) e alguma ferramenta de ETL (pentaho, scripts SQL, python,
etc.).

3. Não esqueça de comentar seu código, pois isto ajuda muito no processo de
manutenção. Lembre-se que outra pessoa irá avaliar seu código. Quanto mais bem
documentado, mais fácil e menos passível de ambiguidades será a análise.

4. O resultado dos cases deve ser enviado em até 3 dias a partir do envio do e-mail.

5. Você deve enviar todos os arquivos utilizados durante a resolução dos cases.
