const express = require('express');
const bodyParser = require('body-parser');
const app = express();

const { createLogger, format, transports } = require('winston');
const { json, combine, timestamp, label, prettyPrint } = format;

const logger = createLogger({
    level: 'debug',
    format: combine(
        timestamp(),
        prettyPrint(),
    ),
    transports: new transports.Console(),
});

app.use(bodyParser.json());

app.get('/', (req, res) => {
    res.json({ ok: 1 });
});

app.get('/health', (req, res) => {
    res.json({ ok: 1 });
});

app.get('/meeting', (req, res) => {
    const start = Date.now();
    logger.debug('starting a meeting')
    setTimeout(() => {
        const now = Date.now();
        logger.debug('finished a meeting')
        res.json({ duration: now - start });
    }, 5);
});


const port = process.env.PORT || 3000;
app.listen(port, () => logger.info(`Running on port: ${port}`));
