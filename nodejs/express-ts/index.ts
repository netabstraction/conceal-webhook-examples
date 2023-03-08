import express, { Express, NextFunction, Request, response, Response } from 'express';
import dotenv from 'dotenv';
import CryptoJS from 'crypto-js'

dotenv.config();

const app = express();
const port = 4002;

const signatureKeyConst = "signature-key"
const apiKeyKeyConst = "x-api-key"
const apiKeyValueConst = "sample-key"

app.post('/nodejs/express-ts/api-key-signature-protected', (req: Request, res: Response) => {
  res.send('Express + TypeScript Server');
});

app.listen(port, () => {
  console.log(`[server]: Server is running at http://localhost:${port}`);
});

const webhookPluginAPI = (req: Request, res: Response) => {
  return res.status(200); 
}

const apiKeyAuthValidator = (req: Request, res: Response, next: NextFunction) => {
  if(req.headers[apiKeyKeyConst] !== apiKeyValueConst) {
    return res.status(401).json({ msg: 'API Key missing' }); 
  }
  next();
}

const timestampValidator = (req: Request, res: Response, next: NextFunction) => {
  const requestTimestamp  = req.headers["conceal_timestamp"] as unknown as number
  const currentTimestamp = Math.floor(Date.now() / 1000)

  if (requestTimestamp-currentTimestamp < -60000 || requestTimestamp-currentTimestamp > 120000) {
    return res.status(400).json({ msg: 'Invalid Timestamp. Timestamp not in range' }); 
  }
  
  next();
}

const signatureValidator = (req: Request, res: Response, next: NextFunction) => {
  const requestTimestamp  = req.headers["conceal_timestamp"] as unknown as string
  const requestSignature  = req.headers["conceal_timestamp"] as unknown as string

  const message = `${requestTimestamp}|${req.url}`
  console.log(message)

  const verifierHashed = CryptoJS.SHA256(signatureKeyConst);
  const expectedSignature = verifierHashed
    .toString(CryptoJS.enc.Base64)
    .replace(/\+/g, '-')
    .replace(/\//g, '_')
    .replace(/=/g, '');

  if(requestSignature !== expectedSignature) {
    return res.status(401).json({ msg: 'Invalid Signature' }); 
  }

  next();
}