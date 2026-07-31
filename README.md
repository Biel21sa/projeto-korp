# Projeto Korp - Desafio Técnico

## Descrição

Este projeto foi desenvolvido como parte do processo seletivo da Korp.

A solução contempla:

* Serviço HTTP em Golang
* Containerização com Docker
* Orquestração com Docker Compose
* Proxy reverso com NGINX
* Monitoramento com Prometheus
* Visualização de métricas com Grafana
* Automação de infraestrutura utilizando Ansible
* Configuração por variáveis de ambiente (.env)
* Provisionamento automático do Grafana

---

## Arquitetura

```text
Cliente
   |
   v
NGINX (porta 80)
   |
   v
http-server-projeto-korp (porta 8080)
   |
   +--> /metrics
           |
           v
      Prometheus
           |
           v
        Grafana
```

---

## Estrutura do Projeto

```text
projeto-korp/

├── app/
│   ├── main.go
│   ├── go.mod
│   └── Dockerfile
│
├── nginx/
│   └── http-server-projeto-korp.conf
│
├── prometheus/
│   └── prometheus.yml
│
├── grafana/
│   ├── dashboards/
│   │   └── http-server-projeto-korp-dashboard.json
│   │
│   └── provisioning/
│       ├── dashboards/
│       │   └── dashboard.yml
│       │
│       └── datasources/
│           └── datasource.yml
│
├── ansible/
│   ├── inventory.ini
│   └── playbook.yml
│
├── docker-compose.yml
├── .env
├── .env.example
└── README.md
```

---

## Funcionalidades

* API REST desenvolvida em Golang
* Containerização utilizando Docker
* Proxy reverso com NGINX
* Coleta de métricas via Prometheus
* Dashboard de monitoramento com Grafana
* Provisionamento automático de datasource e dashboard
* Configuração via variáveis de ambiente
* Deploy automatizado com Ansible

---

## Requisitos

* Docker
* Docker Compose
* Ansible
* Git

---

## Configuração

Copie o arquivo de exemplo:

```bash
cp .env.example .env
```

---

## Variáveis de Ambiente

| Variável         | Descrição           | Valor padrão |
| ---------------- | ------------------- | ------------ |
| GRAFANA_USER     | Usuário do Grafana  | admin        |
| GRAFANA_PASSWORD | Senha do Grafana    | admin        |
| APP_PORT         | Porta da aplicação  | 8080         |
| NGINX_PORT       | Porta do NGINX      | 80           |
| PROMETHEUS_PORT  | Porta do Prometheus | 9090         |
| GRAFANA_PORT     | Porta do Grafana    | 3000         |

---

## Executando com Docker Compose

Na raiz do projeto:

```bash
docker compose up -d --build
```

Verificar containers:

```bash
docker ps
```

Parar o ambiente:

```bash
docker compose down
```

---

## Testando o Serviço

Executar:

```bash
curl http://localhost/projeto-korp
```

Resposta esperada:

```json
{
  "nome": "Projeto Korp",
  "horario": "2026-07-31T12:00:00Z"
}
```

---

## Monitoramento

### Endpoint de métricas

```text
http://localhost/metrics
```

### Prometheus

```text
http://localhost:9090
```

Consultas úteis:

```promql
http_requests_total
```

```promql
up
```

```promql
rate(http_requests_total[1m])
```

### Grafana

```text
http://localhost:3000
```

Credenciais padrão:

```text
Usuário: admin
Senha: admin
```

As credenciais podem ser alteradas através do arquivo `.env`.

---

## Provisionamento Automático do Grafana

O Grafana é configurado automaticamente durante a inicialização do ambiente.

Os seguintes recursos são provisionados sem intervenção manual:

* Datasource Prometheus
* Dashboard do Projeto Korp
* Integração automática com Prometheus

Arquivos responsáveis pelo provisionamento:

```text
grafana/provisioning/datasources/datasource.yml
grafana/provisioning/dashboards/dashboard.yml
grafana/dashboards/http-server-projeto-korp-dashboard.json
```

Após executar o ambiente, o dashboard estará disponível automaticamente no Grafana.

---

## Dashboard

O dashboard exibe:

* Disponibilidade do serviço
* Total de requisições recebidas
* Taxa de requisições por segundo
* Consumo de memória da aplicação Go
* Quantidade de goroutines em execução

Métricas monitoradas:

```promql
up
```

```promql
http_requests_total
```

```promql
rate(http_requests_total[1m])
```

```promql
go_memstats_alloc_bytes
```

```promql
go_goroutines
```

---

## Automação com Ansible

Executar:

```bash
cd ansible

ansible-playbook -i inventory.ini playbook.yml
```

O playbook realiza:

* Instalação do Docker
* Inicialização do serviço Docker
* Criação da rede Docker
* Build da aplicação
* Subida dos containers
* Configuração do NGINX
* Configuração do Prometheus
* Configuração do Grafana
* Provisionamento automático do datasource
* Provisionamento automático do dashboard
* Validação do endpoint da API
* Verificação dos serviços Prometheus e Grafana

---

## Validação do Ambiente

Verificar API:

```bash
curl http://localhost/projeto-korp
```

Verificar métricas:

```bash
curl http://localhost/metrics
```

Verificar containers:

```bash
docker ps
```

Verificar targets do Prometheus:

```text
http://localhost:9090/targets
```

---

## Bônus Implementados

* Provisionamento automático do Grafana
* Datasource configurado via código
* Dashboard configurado via código
* Configuração por variáveis de ambiente (.env)
* Deploy automatizado com Ansible
* Build Docker multi-stage
* Infraestrutura reproduzível através de Docker Compose

---

## Tecnologias Utilizadas

* Golang
* Docker
* Docker Compose
* NGINX
* Prometheus
* Grafana
* Ansible

---

## Autor

Gabriel Silva Alves
