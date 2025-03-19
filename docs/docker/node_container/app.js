const express = require('express');
const fs = require('fs');
const path = require('path');
const bodyParser = require('body-parser');

const app = express();
const port = 3000;	

/* ---------------------------------------------- */
/*												  */	
/*  Progetto test: semplici operazioni CRUD       */
/*  Su file di testo						      */
/*												  */
/* ---------------------------------------------- */

const dataDir = path.join(__dirname, 'data');		 	// Percorso cartella
const dataFilePath = path.join(dataDir, 'data.txt'); 	// Percorso file di testo
app.use(bodyParser.json()); 							// Middleware

/* CREAZIONE ----------------------------------- */

fs.mkdirSync(dataDir, { recursive: true });

if (!fs.existsSync(dataFilePath)) {
    fs.writeFileSync(dataFilePath, '');
}

/* LETTURA ------------------------------------- */ 

app.get('/read', (req, res) => {
    const content = fs.readFileSync(dataFilePath, 'utf8');
    res.send(content);
});

/* (SOVRA) SCRITTURA --------------------------- */ 

app.post('/write', (req, res) => {
    const { text } = req.body;
    fs.writeFileSync(dataFilePath, text);
    res.send('File aggiornato con successo');
});

/* SCRITTURA IN APPEND ------------------------- */ 

app.post('/append', (req, res) => {
    const { text } = req.body;
    fs.appendFileSync(dataFilePath, `\n${text}`);
    res.send('Testo aggiunto con successo');
});

/* PULIZIA CONTENUTO --------------------------- */ 

app.delete('/delete', (req, res) => {
    fs.writeFileSync(dataFilePath, '');
    res.send('File svuotato con successo');
});

/* ------------ Server in ascolto -------------- */

app.listen(port, () => {
    console.log(`Server running at http://localhost:${port}`);
});
