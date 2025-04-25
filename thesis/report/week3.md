## Settimana 3 01/04/2021 - 08/04/2021

## Creo le basi per l'architettura del progetto

- Ho creato un Vagrantfile con due macchine virtuali e ho creato un cluster k3s con un master e un worker.

- Ho definito una CustomResourceDefinition per i dispositivi IoT, che include vari campi come IP, stato, data di ultima modifica dello stato, ecc.

- Ho definito un Deployment statico che crea due pod che eseguono un'immagine Docker (mock) per i dispositivi IoT.

# Creo un'applicazione di test in Go

- Ho creato un'applicazione in Go ( dockerizzata, caricando l'immagine su un repository remoto ) che funziona come un agente per i dispositivi, prende come variabili d'ambiente la porta ed espone un server HTTP.

> [Resoconto](../README.md)