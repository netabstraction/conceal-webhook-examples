import express, { Request, Response } from "express";
import CryptoJS from "crypto-js";

const app = express();
const port = 8080;

const signatureKeyConst = "signature-key";
const apiKeyKeyConst = "x-api-key";
const apiKeyValueConst = "sample-key";
const webhookUrl = "http://127.0.0.1:8080/webhook";

const handleWebhook = (req: Request, res: Response) => {
  const requestTimestamp = req.headers[
    "conceal-timestamp"
  ] as unknown as number;

  const requestSignature = req.headers[
    "conceal-signature"
  ] as unknown as string;

  const requestApiKey = req.headers[apiKeyKeyConst] as unknown as string;

  // API Key validation
  if (requestApiKey !== apiKeyValueConst) {
    console.log("Invalid API Key");
    return res.status(401).json({ error: "Invalid API Key" });
  }

  // Timestamp validation
  if (!isValidTimestamp(requestTimestamp)) {
    console.log("Invalid Timestamp");
    return res.status(400).json({ error: "Invalid Timestamp" });
  }

  //Signature validation
  if (!isValidSignature(requestTimestamp, requestSignature)) {
    console.log("Invalid Signature");
    return res.status(401).json({ error: "Invalid Signature" });
  }

  // Process the webhook payload
  // ..
  logRequest(req);
  // ..

  // Return a success response
  console.log("Ok");
  res.status(200).json({ status: "OK" });
};

// Validate timestamp timestamp is in the range of [current_timestamp-60sec, current_timestamp_120sec]
const isValidTimestamp = (requestTimestamp: number) => {
  const currentTimestamp = Math.floor(Date.now() / 1000);

  if (!requestTimestamp) {
    return false;
  }
  
  return (
    requestTimestamp - currentTimestamp > -60000 &&
    requestTimestamp - currentTimestamp < 120000
  );
};

// Validate signature
const isValidSignature = (
  requestTimestamp: number,
  requestSignature: string
) => {
  const message = `${requestTimestamp}|${webhookUrl}`;
  console.log(`Computed Signature Message: ${message}`);

  const expectedSignature = CryptoJS.HmacSHA256(
    message,
    signatureKeyConst
  ).toString(CryptoJS.enc.Hex) as string;

  return requestSignature === expectedSignature;
};

// Log request
const logRequest = (req: Request) => {
  console.log(`req [${req.method}] Url : ${req.url}`);
  console.log(`headers :`);
  console.log(req.headers);
  console.log(`Body :`);
  console.log(req.body);
};

//The body value should be json
app.use(express.json());

app.post("/webhook", handleWebhook);

app.listen(port, () => {
  console.log(`[server]: Server is running at http://localhost:${port}`);
});
