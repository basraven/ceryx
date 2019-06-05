# Ceryx
A message proxy written in Go with polymorphic integrations such as HTTP, Telegram and Facebook Messenger (future)

## How to run

* In Go
    ```shell
    go get -u github.com/basraven/ceryx
    go run github.com/basraven/ceryx
    ```
* In Docker
    ```shell
    cd docker
    docker-compose up
    ```
* In Kubernetes
    ```shell
    kubectl apply -f docker/k8s/ceryx.yaml
    ```
### Env variables
Ceryx requires the following environmental variables to be set to work:
```
TELEGRAM_API_KEY=
CHAT_ID=
```