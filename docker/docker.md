# DOCKER
Installazione: Vedi docker_install.sh
Comandi base:
```sh
docker info
docker version
docker login
```

# Come eseguire un container

docker run --publish LISTENPORT (es. 80:80) --name NOMECONTAINER IMMAGINE

sudo docker run -it nginx /bin/bash # Esegue container bash

# Avvio server nginx
sudo docker run -d -p 8080:80 --name webserver nginx
sudo docker ps #listo tutti i container in esecuzione
sudo docker images #listo tutte le imagini installate

sudo docker container exec -it webserver bash #per fare l'attach al container
(lo vedo come root@container-id)

Il nome con cui mi interfaccerò con il container è "webserver"
Può essere stoppato (stop) e rimosso dalla memoria (rm)
Per uscire dal terminare allatcher: CTRL+P CTRL+Q

# Avvio server node
1. inizializza progetto node con ```npm init -y``` e ```npm install```
2. da vscode CTRL+SHIFT+P -> docker add -> docker config files: node
3. fare la build dell'immagine container ```sudo docker build -t node-app-test .```
	> L'immagine è stata creata con successo (listabile con sudo docker images)
4. eseguire il contanier ```sudo docker run -p 3000:3000 node-app-test```
5. avvia il container in modalitò "interattiva" ```sudo docker run -it --rm node-app-test sh```

In Docker:
    Immagine: è una sorta di "modello" o "template" immutabile che contiene tutto il necessario per eseguire un'applicazione: codice, librerie, dipendenze, file di configurazione e variabili d'ambiente. Un'immagine viene utilizzata per creare container.
    Container: è un'istanza in esecuzione di un'immagine. Si tratta di un ambiente isolato e autonomo che esegue il software definito nell'immagine. I container sono leggeri e condividono il kernel del sistema operativo host, ma sono isolati dal resto del sistema e tra di loro.
In sintesi, un'immagine è come una "fotografia" di un'applicazione, mentre un container è l'esecuzione pratica di quella fotografia.
