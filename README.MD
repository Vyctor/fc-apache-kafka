# Apache Kafka

## Apresentação

Apache Kafka é uma plataforma de streaming de eventos distribuída, de código aberto, que permite a publicação e a subscrição de fluxos de registros. É uma das plataformas de streaming mais utilizadas no mundo, sendo utilizada por empresas como Netflix, Uber, LinkedIn, Twitter, etc.

## O mundo dos eventos

Cada dia precisamos realizar processamento de eventos de diversos tipos de plataformas, como por exemplo:

- Aplicações web
- Aplicações mobile
- Aplicações desktop
- Aplicações IoT
- Aplicações de Big Data
- Aplicações de Machine Learning

E cada uma dessas aplicações gera eventos que precisam ser processados. Por exemplo, uma aplicação web pode gerar eventos de login, logout, cadastro, etc. Uma aplicação mobile pode gerar eventos de login, logout, cadastro, etc. Uma aplicação desktop pode gerar eventos de login, logout, cadastro, etc. Uma aplicação IoT pode gerar eventos de temperatura, umidade, etc. Uma aplicação de Big Data pode gerar eventos de processamento de dados, etc. Uma aplicação de Machine Learning pode gerar eventos de treinamento de modelos, etc.

Devido a isso devemos nos fazer as seguintes perguntas:

- Onde salvar esses eventos?
- Como recuperar de forma rápida e simples, de forma que o feedbacke entre um processo e outro possa acontecer de forma fluida e em tempo real?
- Como escalar a solução?
- Como ter resiliência e alta disponibilidade?

## Os super poderes do Apache Kafka

- Altissímo throughput
- Latência extremamente baixa (2ms)
- Escalável
- Armazenamento
- Se conecta com quase tudo
- Bibliotecas prontas para as mais diversas tecnologias
- Ferramenta open-source

## Conceitos e dinâmica básica de funcionamento

### Tópicos

Tópicos são como canais de comunicação. Um tópico é um fluxo de eventos ordenados por tempo. Um tópico é composto por uma ou mais partições. Cada mensagem enviada para um tópico é anexada a uma única partição.

Cada partição é um fluxo ordenado de mensagens que é persistido em um log. Os logs são segmentados em segmentos, que são os arquivos físicos que são armazenados no disco. Os logs são compactados para remover mensagens antigas. Os logs são replicados para garantir a tolerância a falhas.

Cada mensagem enviada para um tópico é armazenada, de forma persistente, em um log. Cada log é composto por segmentos, que são os arquivos físicos que são armazenados no disco. Os logs são compactados para remover mensagens antigas. Os logs são replicados para garantir a tolerância a falhas.

### Anatomia de um registro

Um registro é composto por:

- Header
- Key
- Value
- Timestamp

### Partições

Cada tópico pode ter uma ou mais partições para garantir a distribuição e resiliência de seus dados.

### Garantindo a ordem de entrega

O Kafka garante a ordem de entrega de mensagens para cada partição. Isso significa que as mensagens são entregues em ordem para cada partição. Porém, não há garantia de ordem de entrega entre as partições.

### Partições distribuídas

Cada partição é distribuída em um ou mais brokers e cada broker é um servidor que armazena mensagens de uma ou mais partições.

#### Replicator Factor

- Número de cópias de cada partição que são armazenadas em diferentes brokers.
- É necessário no mínimo 2 cópias de cada partição para garantir a tolerância a falhas.
- O número de cópias de cada partição é definido no momento da criação do tópico.

### Partition Leadership

- Cada partição tem um líder. O líder é o único que pode receber e responder a requisições de leitura e escrita para a partição.
- As demais partições são chamadas de followers. Os followers são apenas leitores. Eles não podem receber requisições de escrita.
- O líder é eleito automaticamente pelo Kafka.

### Garantia de entrega de mensagens

Como garantir que uma mensagem foi entregue com sucesso?

Quando o producer envia a mensagem para o broker ele seta um parâmetro chamado ack.
Os tipos de ack são:

- acks=0: O producer não espera nenhuma confirmação do broker. O producer assume que a mensagem foi enviada com sucesso. Mais conhecido como fire and forget.
- acks=1: O producer espera a confirmação do líder da partição. O producer assume que a mensagem foi enviada com sucesso.
- acks=-1 ou all: O producer espera a confirmação do líder da partição e dos followers. O producer assume que a mensagem foi enviada com sucesso.

### Producer - Garantia de entrega de mensagens

#### At most once

O producer envia a mensagem e não espera nenhuma confirmação. O producer assume que a mensagem foi enviada com sucesso. Mais conhecido como fire and forget. Poder perder mensagens.

#### At least once

O producer envia a mensagem e espera a confirmação do líder da partição. O producer assume que a mensagem foi enviada com sucesso. Poderá duplicar mensagens.

#### Exactly once

O producer envia a mensagem e espera a confirmação do líder da partição e dos followers. O producer assume que a mensagem foi enviada com sucesso. Não poderá duplicar mensagens.

### Produtor Idempotente

O producer garante que uma mensagem será enviada uma única vez, mesmo que o producer envie a mesma mensagem várias vezes. Para isso, o producer precisa setar o parâmetro enable.idempotence=true.

### Consumer e Consumer Groups

O consumer é um processo que lê mensagens de um ou mais tópicos. O consumer pode ser um processo de batch ou um processo de streaming.

O consumer group é um conjunto de consumers que compartilham mensagens de um tópico. Cada mensagem é entregue para um único consumer dentro de um consumer group.