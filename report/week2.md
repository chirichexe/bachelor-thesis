# Settimana 2 11/03/2021 - 18/03/2021

## Installazione di k3s (su uno o più nodi), comprensione dell'architettura e delle API 

- Ho installato un nodo k3s su un'unica macchina, comprendendo come avrei potuto estenderlo ad altri nodi.

- Ho compreso il funzionamento dell'architettura e come si coordinano i vari agenti per offrire il servizio

## Creazione app di test nginx con Deployment, Service e IngressController
1. Deployment (nginx-deployment.yaml)

Mediante il Deployment ho definito il modo in cui il pod Nginx viene gestito in Kubernetes. Nel mio caso avvia 3 repliche (pods) dell'immagine nginx:latest. Ogni replica è una copia indipendente dell'applicazione. Anche se uno dei pod dovesse fallire, Kubernetes creerà un nuovo pod per mantenerne il numero di repliche definito.

2. Service (nginx-service.yaml)

Ho configurato il Service per fungere da bilanciatore di carico per i pod Nginx (il type è quindi "LoadBalancer"), distribuendo il traffico in ingresso verso i vari pod di Nginx. Essi saranno accessibili tramite un IP e una porta fissa.

3. Ingress (nginx-ingress.yaml)

L'ingress controller "traefik" fornisce un hostname (nel mio caso, nginx:local) per accedere al servizio.

> [Resoconto](kubernetes/kubernetes.md)

## Provisioning mediante Crossplane

Ho installato Crossplane.

> [Resoconto](crossplane/crossplane.md)