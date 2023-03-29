const express = require('express');
const bodyParser = require('body-parser');
const crypto = require('crypto');

const signatureKeyConst = "signature-key";
const apiKeyKeyConst = "x-api-key";
const apiKeyValueConst = "sample-key";
const webhookUrl = "http://127.0.0.1:8080/webhook";

const app = express();
const jsonParser = bodyParser.json();
const port = 8080;

const handleWebhook = (req, res, next) => {
    const requestTimestamp = req.headers["conceal-timestamp"];


    const requestSignature = req.headers["conceal-signature"];

    const requestApiKey = req.headers[apiKeyKeyConst];

    // API Key Validation
    if (requestApiKey !== apiKeyValueConst) {
        console.log("Invalid API Key");
        return res.status(401).json({msg: "Invalid API Key"});
    }

    // Timestamp Validation
    if (!isValidTimestamp(requestTimestamp)) {
        console.log("Invalid Timestamp");
        return res.status(400).json({msg: "Invalid Timestamp"});
    }

    // Signature Validation
    if (!isValidSignature(requestTimestamp, requestSignature)) {
        console.log("Invalid Signature");
        return res.status(401).json({msg: "Invalid Signature"});
    }

    // Process the webhook payload
    // ..
    logRequest(req);
    // ..

    // Return a success response
    console.log("200 Ok");
    res.status(200).json({msg: "200 Ok" });
    next()
};


// Validate timestamp. Timestamp is in the range of [current_timestamp-60sec, current_timestamp_120sec]
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

// Validate Signature
const isValidSignature = (requestTimestamp, requestSignature) => {
    const message = `${requestTimestamp}|${webhookUrl}`;
    console.log(`Computed Signature Message: ${message}`);

    const expectedSignature = crypto.createHmac('sha256', signatureKeyConst).update(message).digest('hex');

    return requestSignature == expectedSignature;
};

// Log Request
const logRequest = (req) => {
      console.log(`req [${req.method}] Url : ${req.url}`);
      console.log(`headers :`);
      console.log(req.headers);
      console.log(`Body :`);
      console.log(req.body);
};


app.use(bodyParser.json());

app.post('/webhook', handleWebhook);

app.listen(port, '127.0.0.1');
console.log(`[server]: Server is running at http://localhost:${port}`);
