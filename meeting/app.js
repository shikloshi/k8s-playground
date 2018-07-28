const express = require('express');
const bodyParser = require('body-parser');
const app = express();

app.use(bodyParser.json());

app.get('/', (req, res) => {
    res.json({ ok: 1 });
});

app.get('/health', (req, res) => {
    res.json({ ok: 1 });
});

app.get('/meeting', (req, res) => {
    const start = Date.now();
    console.log('[DEBUG] starting a meeting')
    setTimeout(() => {
        const now = Date.now();
        console.log('[DEBUG] finished a meeting')
        res.json({ duration: now - start });
    }, 250);
});


const port = process.env.PORT || 3000;
app.listen(port, () => console.log(`Running on port: ${port}`));
