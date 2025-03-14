# TIROCINIO
**Argomento: Studio e implementazione di Provisioning e Deployment di Workload su Dispositivi IoT attraverso Kubernetes Operators**

0. Documentazione generale e comprensione delle API di docker e di kubernetes:
> Documenti: 
> - generale.md
> - docker/docker.md

## FASE 1: Installazione e configurazione di un cluster k3s su una o più macchine, creando un ambiente Kubernetes leggero per gestire applicazioni distribuite.

1. Deployment (nginx-deployment.yaml)

Il Deployment definisce il modo in cui il pod Nginx viene gestito in Kubernetes. In questo caso, hai creato un Deployment che avvia 3 repliche (pods) dell'immagine nginx:latest. Ogni replica è una copia indipendente dell'applicazione Nginx, garantendo alta disponibilità. Questo significa che anche se uno dei pod dovesse fallire, Kubernetes creerà un nuovo pod per mantenerne il numero di repliche definito.

In breve:

- 3 repliche di Nginx.
- Ogni replica è un pod che espone la porta 80, sulla quale il container di Nginx serve contenuti web.

2. Service (nginx-service.yaml)

Il Service è configurato per fungere da bilanciatore di carico per i pod Nginx, distribuendo il traffico in ingresso verso i vari pod di Nginx.

- **selector**: il Service seleziona i pod con l'etichetta app: nginx, che corrisponde ai pod creati dal Deployment di Nginx.
- **ports**: definisce che la porta 80 sul Service sarà esposta ai client, e la porta 80 dei pod sarà utilizzata per il traffico in entrata.
- **nodePort**: consente di accedere al servizio dall'esterno del cluster Kubernetes sulla porta 30080 (TCP).
    Tipo LoadBalancer: con il tipo LoadBalancer, Kubernetes tenterà di configurare automaticamente un bilanciatore di carico esterno, che in ambienti cloud può essere un servizio di bilanciamento del traffico (ad esempio, su AWS o Google Cloud). Se non sei in un ambiente cloud, il tipo LoadBalancer non avrà effetto pratico, ma puoi comunque usare il nodePort per accedere ai tuoi servizi tramite il nodo.

3. Ingress (nginx-ingress.yaml)

L'Ingress gestisce il traffico HTTP (e HTTPS) verso i servizi del tuo cluster Kubernetes.

- **Ingress Controller**: ho specificato l'uso di un IngressController chiamato traefik che si occupa di instradare il traffico verso i servizi interni in base alle regole di routing definite.
- **Regola del percorso**: la regola host: nginx.local significa che qualsiasi richiesta HTTP che arriva a nginx.local sarà instradata al Service di Nginx sulla porta 80.
- In pratica, la richiesta HTTP a http://nginx.local/ verrà indirizzata al tuo nginx-service e gestita dal pod Nginx.

4. Comandi eseguiti:

- ```kubectl apply```: questi comandi applicano la configurazione YAML per il Deployment, il Service e l'Ingress, facendo in modo che Kubernetes crei le risorse specificate nel cluster.
- ```kubectl get nodes```: mostra lo stato dei nodi del cluster.
- ```kubectl get pods```: mostri lo stato dei pod, che dovrebbero essere 3, poiché il deployment ha 3 repliche.
- ```kubectl get deployments```: mostra i dettagli del deployment di Nginx.
- ```kubectl get services```: mostri lo stato del servizio, che dovrebbe essere configurato come tipo LoadBalancer.

**Cosa comporta questo setup?**

- La configurazione di 3 repliche di Nginx e il bilanciamento del carico permettono di gestire più traffico, riducendo il rischio di downtime se un pod fallisce.
- Accesso tramite Ingress: il traffico esterno arriva tramite nginx.local, che viene instradato all'interno del cluster tramite l'Ingress Controller traefik. Questo ti permette di separare il traffico interno da quello esterno, migliorando la sicurezza e la gestione.
- Carico bilanciato: il Service distribuisce il traffico tra i pod di Nginx, mentre il tipo LoadBalancer aiuta a esporre il servizio a livello esterno, se configurato correttamente.

**Visualizzazione**

2. Installare e configurare strumenti di monitoraggio come Prometheus, integrandoli con Kubernetes e configurando gli operatori per una gestione automatica.

3. Creare e configurare un operatore Kubernetes per il deployment e la gestione automatica dei workload che eseguono su dispositivi IoT, verificando che l'operatore possa scalare e gestire i container in base alle esigenze di carico.

4. Simulare flussi di dati provenienti da dispositivi IoT per testare il sistema, creando applicazioni containerizzate che elaborano i flussi di dati e monitorano le loro performance.

5. Implementare meccanismi di scalabilità automatica dei container e configurare il bilanciamento del carico tra i nodi del cluster per garantire efficienza e ridondanza.

6. Simulare guasti nel sistema per testare la resilienza del cluster Kubernetes, implementando strategie di recovery e failover per garantire la disponibilità continua dei workload sui dispositivi IoT.

> Documenti
> - kubernetes/kubernetes.md

# Bibliografia

- docker: video 
- k3s: corso ufficiale rancher.academy
https://www.youtube.com/watch?v=QDwhbMvikGQ
