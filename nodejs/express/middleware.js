const crypto = require('crypto');

const signatureKeyConst = "signature-key";
const apiKeyKeyConst = "x-api-key";
const apiKeyValueConst = "sample-key";
const webhookUrl =
  "http://127.0.0.1:4000/nodejs/express/api-key-signature-protected";

// API Key validator  
module.exports.apiKeyAuthValidator = (req, res, next) => {
  if (req.headers[apiKeyKeyConst] !== apiKeyValueConst) {
    console.log("API Key missing/API Key doesnot match");
    return res.status(401).json({ msg: "API Key missing/API Key doesnot match" });
  }
  next();
};

// Timestamp validator request timestamp is in the range of [current_timestamp-60sec, current_timestamp_120sec]
module.exports.timestampValidator = (req, res, next) => {
  const requestTimestamp = req.headers["conceal-timestamp"];
  const currentTimestamp = Math.floor(Date.now() / 1000);

  if (
    requestTimestamp - currentTimestamp < -60000 ||
    requestTimestamp - currentTimestamp > 120000
  ) {
    console.log("Invalid Timestamp. Timestamp not in range");
    return res
      .status(400)
      .json({ msg: "Invalid Timestamp. Timestamp not in range" });
  }

  next();
};

// Signature validator
module.exports.signatureValidator = (req, res, next) => {
  const requestTimestamp = req.headers["conceal-timestamp"];
  const requestSignature = req.headers["conceal-signature"];

  const message = `${requestTimestamp}|${webhookUrl}`;
  console.log(`Computed Signature Message: ${message}`);

  const expectedSignature = crypto.createHmac('sha256', signatureKeyConst).update(message).digest('hex');

  console.log(`Computed Signature: ${expectedSignature}`);
  console.log(`Request Signature: ${requestSignature}`);

  if (requestSignature !== expectedSignature) {
    console.log("Invalid Signature");
    return res.status(401).json({ msg: "Invalid Signature" });
  }

  next();
};
