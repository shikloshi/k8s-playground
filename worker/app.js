const express = require('express');
const bodyParser = require('body-parser');
//const superagent = require('superagnet');
const fetch = require('node-fetch');
const app = express();

const numOfMeeting = 2;
app.use(bodyParser.json());

app.get('/', (req, res) => {
    res.json({ status: "okay"})
});

app.get('/work', async (req, res) => {
    const start = Date.now();
    for (let i = 0; i < numOfMeeting; i++) {
        console.log('[DEBUG] going into meeting' + i);
        const res = await goToMeeting();
        console.log('[DEBUG] meeting-v1 service respond with: ' + res.status);
    }
    const end = Date.now();
    res.json({ duration: end - start})
});

async function goToMeeting() {
    try {
        return await fetch('http://localhost:3000/meeting');
        //return await fetch('http://meeting-v1:3000/meeting');
        //return res;
    } catch (e) {
        console.error(e);
    }
}


const port = process.env.PORT || 4000;
app.listen(port, () => console.log(`Running on port: ${port}`));
