const express = require('express');
const fs = require('fs');
const path = require('path');

const app = express();
const port = 3000;

// Percorso della directory e del file di log
const logDir = path.join(__dirname, 'logs');
const logFilePath = path.join(logDir, 'logs.txt');

// Assicurati che la directory e il file di log esistano
fs.mkdirSync(logDir, { recursive: true });
if (!fs.existsSync(logFilePath)) {
  fs.writeFileSync(logFilePath, 'Log avviato\n'); // Crea il file con una riga iniziale
}

// Middleware per scrivere i log
app.use((req, res, next) => {
  const log = `${new Date().toISOString()} - ${req.method} ${req.url}\n`;
  fs.appendFileSync(logFilePath, log);
  next();
});

app.get('/', (req, res) => {
  res.send('Hello, World!');
});

app.listen(port, () => {
  console.log(`Server running at http://localhost:${port}`);
});
