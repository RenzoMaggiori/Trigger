import express from "express";
import bodyParser from "body-parser";
import dotenv from "dotenv";
import { handleConnection } from './db/connection.js';

dotenv.config();
const port = process.env.PORT;

const app = express();

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

app.listen(port, () => {
  handleConnection();
  console.log(`Server is running on port ${port}`);
});
