const bodyParser = require('body-parser');
const crypto = require('crypto');

const signatureKeyConst = "signature-key";
const apiKeyKeyConst = "x-api-key";
const apiKeyValueConst = "sample-key";
const webhookUrl = "http://127.0.0.1:8080/webhook";
const jsonParser = bodyParser.json();

// Require the framework and instantiate it
const app = require('fastify')({
    logger: false
})

const handleWebhook = (req, res) => {
    const requestTimestamp = req.headers["conceal-timestamp"];
    const requestSignature = req.headers["conceal-signature"];
    const requestApiKey = req.headers[apiKeyKeyConst];

    // API Key Validation
    if (requestApiKey !== apiKeyValueConst) {
        console.log("Invalid API Key");
        res.code(401).send({ msg: 'Invalid API Key' });
        return;
    }

    // Timestamp Validation
    if (!isValidTimestamp(requestTimestamp)) {
        console.log("Invalid Timestamp");
        res.code(400).send({ msg: 'Invalid Timestamp' });
        return;
    };

    // Signature Validation
    if (!isValidSignature(requestTimestamp, requestSignature)) {
        console.log("Invalid Signature");
        res.code(401).send({ msg: 'Invalid Signature' });
        return;
    }

    // Process the webhook payload
    // ..
    logRequest(req);
    // ..

    // Return a success response
    console.log("200 Ok");
    res.code(200).send({ msg: '200 Ok' });
};

const isValidTimestamp = (requestTimestamp) => {
    const currentTimestamp = Math.floor(Date.now() / 1000);
    if (!requestTimestamp) {
        return false;
    }

    return (
        requestTimestamp - currentTimestamp > -60000 &&
        requestTimestamp - currentTimestamp < 120000
    );
};

const isValidSignature = (requestTimestamp, requestSignature) => {
    const message = `${requestTimestamp}|${webhookUrl}`;
    const expectedSignature = crypto.createHmac('sha256', signatureKeyConst).update(message).digest('hex');

    return requestSignature == expectedSignature;
};

const logRequest = (req) => {
    console.log(`req [${req.method}] Url : ${req.url}`);
    console.log(`headers : `);
    console.log(req.headers);
    console.log(`Body : `);
    console.log(req.body);
};

// Webhook route
app.post('/webhook', handleWebhook);

// Run
app.listen({port: 8080}, function (err, address) {
    if (err) {
        app.log.error(err)
        process.exit(1)
    }
})
