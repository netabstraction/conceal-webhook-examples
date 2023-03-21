import { NextFunction, Request, Response } from "express";
import CryptoJS from "crypto-js";

const signatureKeyConst = "signature-key";
const apiKeyKeyConst = "x-api-key";
const apiKeyValueConst = "sample-key";
const webhookUrl =
  "http://127.0.0.1:4002/nodejs/express-ts/api-key-signature-protected";

// API Key validator  
const apiKeyAuthValidator = (
  req: Request,
  res: Response,
  next: NextFunction
) => {
  if (req.headers[apiKeyKeyConst] !== apiKeyValueConst) {
    console.log("API Key missing/API Key doesnot match");
    return res.status(401).json({ msg: "API Key missing/API Key doesnot match" });
  }
  next();
};

// Timestamp validator request timestamp is in the range of [current_timestamp-60sec, current_timestamp_120sec]
const timestampValidator = (
  req: Request,
  res: Response,
  next: NextFunction
) => {
  const requestTimestamp = req.headers[
    "conceal-timestamp"
  ] as unknown as number;
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
const signatureValidator = (
  req: Request,
  res: Response,
  next: NextFunction
) => {
  const requestTimestamp = req.headers[
    "conceal-timestamp"
  ] as unknown as string;
  const requestSignature = req.headers[
    "conceal-signature"
  ] as unknown as string;

  const message = `${requestTimestamp}|${webhookUrl}`;
  console.log(`Computed Signature Message: ${message}`);

  const expectedSignature = CryptoJS.HmacSHA256(
    message,
    signatureKeyConst
  ).toString(CryptoJS.enc.Hex) as string;

  console.log(`Computed Signature: ${expectedSignature}`);
  console.log(`Request Signature: ${requestSignature}`);

  if (requestSignature !== expectedSignature) {
    console.log("Invalid Signature");
    return res.status(401).json({ msg: "Invalid Signature" });
  }

  next();
};

const logger = (req: Request, resp: Response, next: NextFunction) => {
  console.log("REQUEST");
  // console.log(req)
  console.log(`Url : ${req.url}`);
  console.log(`Query: `);
  console.log(req.query);
  console.log(`Method : ${req.method}`);
  console.log(`Header :`);
  console.log(req.headers);
  console.log(`Body :`);
  console.log(req.body);
  next();
};

const middleware = {
  apiKeyAuthValidator,
  timestampValidator,
  signatureValidator,
  logger,
};

export default middleware;
