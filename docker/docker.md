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
Per uscire dal terminare allatcher: CTRL+P

# Avvio server node
1. inizializza progetto node con ```npm init -y```
2. da vscode CTRL+SHIFT+P -> docker run -> docker config files: node



