# Práctica #2
* Sistemas Operativos 1
* Erick José André Villatoro Revolorio
* 201900907


## Manual Técnico

### Deploy de Kafka
Para el levantamiento de kafka se utilizó `Strimzi`. En primer lugar se debe utilizar un namespace llamado `kafka`.
```
kubectl create namespace kafka
```

Además se instalaron los archivos de `Strimzi`.
```
kubectl create -f 'https://strimzi.io/install/latest?namespace=kafka' -n kafka
```

Posteriormente se levanta kafka y zookeeper. 

```
kubectl apply -f kafka-deployment.yml -n kafka 
```

## Creacion de Proyecto de gRPC

go mod init github.com/Villa01/practica2/gRPC-Client