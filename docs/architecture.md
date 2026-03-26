# Zenvy architecture

Updated at 25.03.2026

# Scheme

// Under development

# Microservices
*missing technologies, dbs and etc. now*

### Logs Service

The logs service processes all logs from pipelines

### Pipelines Service

The pipelines service processes pipelines creating, updating, deletion, aggregation and validating pipelines' spec files

### Runners Service

The runners service asynchronously(via Kafka) executes the pipelines on free workers

Workers can be both hosted near the runners service and provided by user

### Gateway

The gateway serves as a bridge between the client and the entire Zenvy application
