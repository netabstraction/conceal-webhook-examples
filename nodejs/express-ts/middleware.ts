import { NextFunction, Request, Response } from "express";
import CryptoJS from "crypto-js";

const signatureKeyConst = "signature-key";
const apiKeyKeyConst = "x-api-key";
const apiKeyValueConst = "sample-key";
const webhookUrl = "http://127.0.0.1:4002/nodejs/express-ts/api-key-signature-protected"

const apiKeyAuthValidator = (
  req: Request,
  res: Response,
  next: NextFunction
) => {
  if (req.headers[apiKeyKeyConst] !== apiKeyValueConst) {
    return res.status(401).json({ msg: "API Key missing" });
  }
  next();
};

const timestampValidator = (
  req: Request,
  res: Response,
  next: NextFunction
) => {
  const requestTimestamp = req.headers[
    "conceal_timestamp"
  ] as unknown as number;
  const currentTimestamp = Math.floor(Date.now() / 1000);

  if (
    requestTimestamp - currentTimestamp < -60000 ||
    requestTimestamp - currentTimestamp > 120000
  ) {
    return res
      .status(400)
      .json({ msg: "Invalid Timestamp. Timestamp not in range" });
  }

  next();
};

const signatureValidator = (
  req: Request,
  res: Response,
  next: NextFunction
) => {
  const requestTimestamp = req.headers[
    "conceal_timestamp"
  ] as unknown as string;
  const requestSignature = req.headers[
    "conceal_timestamp"
  ] as unknown as string;

  const message = `${requestTimestamp}|${webhookUrl}`;
  console.log(message);

  const expectedSignature = CryptoJS.HmacSHA256(
    message,
    signatureKeyConst
  ) as unknown as string;

  console.log(expectedSignature)
  if (requestSignature !== expectedSignature) {
    return res.status(401).json({ msg: "Invalid Signature" });
  }

  next();
};

const validator = {
  apiKeyAuthValidator,
  timestampValidator,
  signatureValidator,
};

export default validator;
