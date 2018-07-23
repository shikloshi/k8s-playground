const express = require('express');
const bodyParser = require('body-parser');
const app = express();

app.use(bodyParser.json());


app.get('/meeting', (req, res) => {
    const start = Date.now();
    setTimeout(() => {
        const now = Date.now();
        res.json({ duration: now - start });
    }, 250);
});


const port = process.env.PORT || 3000;
app.listen(3000, () => console.log(`Running on port: ${port}`));
