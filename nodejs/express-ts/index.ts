import express, { Request, Response } from "express";
import middleware from "./middleware";
import bodyParser from "body-parser";

const app = express();
const jsonParser = bodyParser.json();
const port = 4002;

const webhookPluginAPI = (req: Request, res: Response) => {
  console.log("200 Ok");
  res.status(200).json({ msg: "200 Ok" });
};

app.post(
  "/nodejs/express-ts/api-key-signature-protected",
  jsonParser, // json body parser
  middleware.logger, // request logger
  middleware.apiKeyAuthValidator,
  middleware.timestampValidator,
  middleware.signatureValidator,
  webhookPluginAPI
);

app.listen(port, () => {
  console.log(`[server]: Server is running at http://localhost:${port}`);
});
