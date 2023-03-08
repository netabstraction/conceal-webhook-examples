import express, {
  Express,
  NextFunction,
  Request,
  response,
  Response,
} from "express";
import dotenv from "dotenv";
import validator from "./middleware";

dotenv.config();

const app = express();
const port = 4002;

const webhookPluginAPI = (req: Request, res: Response) => {
  return res.status(200);
};

app.post(
  "/nodejs/express-ts/api-key-signature-protected",
  validator.apiKeyAuthValidator,
  validator.timestampValidator,
  validator.signatureValidator,
  webhookPluginAPI
);

app.listen(port, () => {
  console.log(`[server]: Server is running at http://localhost:${port}`);
});


