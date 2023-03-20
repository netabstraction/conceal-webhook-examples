const express = require('express');
const bodyParser = require('body-parser');
const crypto = require('crypto');
const middleware = require('./middleware');

const app = express();
const jsonParser = bodyParser.json();
const port = 4000;

const webhookPluginAPI = (req, res, next) => {
    console.log("200 Ok");
    res.json({msg: "200 Ok" });
    next()
};

app.use(bodyParser.json());
app.use(middleware.apiKeyAuthValidator);
app.use(middleware.timestampValidator);
app.use(middleware.signatureValidator);
app.use(webhookPluginAPI);

app.get('/nodejs/express/api-key-signature-protected',
    jsonParser,
    webhookPluginAPI
);

app.listen(port, () => {
    console.log(`[server]: Server is running at http://localhost:${port}`);
});
