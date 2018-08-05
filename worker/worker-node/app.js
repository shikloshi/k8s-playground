const express = require('express');
const bodyParser = require('body-parser');
const { createLogger, format, transports } = require('winston');
const { json, combine, timestamp, label, prettyPrint } = format;

const fetch = require('node-fetch');
const app = express();

const numOfMeeting = 3;

const meetingHost = process.env.MEETING_HOST || 'localhost';

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
    res.json({ status: "okay"})
});

app.get('/health', (req, res) => {
    res.json({ status: "okay"})
});

app.get('/work', async (req, res) => {
    let resBody = {};
    const start = Date.now();
    logger.debug('request hit the /work endpoint');
    if (req.headers['x-skip-meeting']) {
        logger.debug('request hit but skipped going to meeting');
        //return res.json({ message: 'meeting skipped'});
    }
    else {
        // TODO: add header for easy return here
        for (let i = 0; i < numOfMeeting; i++) {
            logger.debug('going into meeting' + i);
            try {
                response = await goToMeeting();
                logger.debug('meeting-v1 service respond with: ' + response.status);
                resBody['meeting_' + i] = response.status;
            } catch (e) {
                logger.error('error from meeting service:' + e.message);
                resBody = { error: e.message }
            }
        }
    }
    const end = Date.now();
    const duration = end - start;
    logger.info('duration of request:' +  duration);
    resBody = Object.assign({}, resBody, { duration });
    console.log('Going to respond with: ', resBody);
    return res.json(resBody)
});

async function goToMeeting() {
    try {
        return await fetch('http://' + meetingHost + ':3000/meeting');
        //return await fetch('http://meeting-v1:3000/meeting');
        //return res;
    } catch (e) {
        logger.error(e);
        throw new Error('Could not finish call to meeting service');
    }
}


const port = process.env.PORT || 4000;
app.listen(port, () => { 
    logger.info(`Running on port: ${port}`);
    logger.info(`Calling meeting service on: ${meetingHost}`);
});
