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

## O que é o Kafka Connect?

O Kafka Connect é um framework para conectar fontes de dados e sinks de dados com o Kafka. O Kafka Connect é uma ferramenta open-source que permite a integração de fontes de dados e sinks de dados com o Kafka.

### Dinâmica de funcionamento

O Apache Kafka conversa com o Kafka Connect, que por sua vez conversa com as fontes de dados e sinks de dados através de diversos tipos de conectores.

Temos conectores focados em fontes de dados, como por exemplo: JDBC, MongoDB, MySQL, PostgreSQL, etc. E temos conectores focados em sinks de dados, como por exemplo: JDBC, MongoDB, MySQL, PostgreSQL, etc.

E também temos conectores focados em fontes e sinks de dados, como por exemplo: Elastic Search, S3, AWS Lambda, etc.

## Conectores

Sink: é um conector que permite que os dados fluam do Kafka para um destino externo, como um banco de dados ou um sistema de armazenamento em nuvem. É usado quando você quer exportar dados do Kafka para outra plataforma ou serviço.

Source: é um conector que permite que os dados fluam de um sistema externo para o Kafka. É usado quando você deseja importar dados de outra plataforma ou serviço para o Kafka.

Replicator: é um tipo especial de conector que permite replicar dados de um cluster Kafka para outro. É usado para criar cópias de um cluster Kafka para fins de backup, recuperação de desastres, análise de dados ou outras finalidades.

Em resumo, o sink é usado para enviar dados do Kafka para outro sistema, o source é usado para receber dados de outro sistema e o replicator é usado para replicar dados de um cluster Kafka para outro.

## Standlone Workers

O Kafka Connect é composto por um ou mais workers. Cada worker é um processo que executa um ou mais conectores. Cada worker é composto por um ou mais tasks. Cada task é um processo que executa um ou mais conectores.

## Distributed Workers

O Kafka Connect também pode ser executado em modo distribuído. Nesse caso, cada worker é um processo que executa um ou mais tasks. Cada task é um processo que executa um ou mais conectores.

## Converters

Os converters são responsáveis por converter os dados de entrada e saída do Kafka Connect. Os converters são configurados no arquivo de configuração do worker.

## Dead Letter Queue

A dead letter queue é uma fila de mensagens que armazena mensagens que não foram processadas com sucesso. A dead letter queue é configurada no arquivo de configuração do worker.

No Kafka Connect, a dead letter queue é chamada de error queue. A dead letter queue é configurada no arquivo de configuração do worker.

Quando há um registro inválido, independente da razão, o erro pode ser tratado nas configurações do conector através do parâmetro errors.tolerance. Esse tipo de configuração só pode ser realizada para conectores do tipo Sink.

- none: Faz a tarefa parar imediatamente
- all: Erros s'ao ignorados e o processo continua normalmente
- errors.deadletterqueue.topic.name: Erros são enviados para a dead letter queue

## MSK: Apache Kafka Gerenciado pela AWS - Reduzindo a complexidade de instalar e administrar um cluster de Kafka na AWS

### Qual problema o Kafka resolve?

Mover e transformar um grande volume de dados em tempo real entre diferentes sistemas.

### O que é o MSK?

O Amazon MSK é um serviço totalmente gerenciado que facilita a implantação e a operação do Apache Kafka em escala na AWS. O Amazon MSK é um serviço totalmente gerenciado que facilita a implantação e a operação do Apache Kafka em escala na AWS.

- Zookeeper Incluído
- Altamente Disponível
- Altamente Integrado
- Altamente Seguro
- Configurável
- Escalável

### Vantagens do MSK

- Instalação simplificada através do console
- Automatização simplificada com Terraform
- Excelente desempenho e segurança
- Última versão do Kafka

### Desvantagens do MSK

- Expansão do cluster ainda não funciona via Terraform, apenas via aws cli
- Ausência de métricas importantes como o LAG
- Poucas opções de tipos de instâncias
- Não oferece instância reservada

### Dicas

- Para n consumidores lerem de um mesmo tópico simultaneamente é necessário que este tópico tenha n partições.
- Cuidado com o tamanho das partições. Não deixe muito pequeno, pois isso pode gerar muitas partições e isso pode gerar um overhead de processamento. Não deixe muito grande, pois isso pode gerar um overhead de armazenamento.
