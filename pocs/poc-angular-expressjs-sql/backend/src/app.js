import express from 'express';
import bodyParser from 'body-parser';
import dotenv from 'dotenv';
import cors from 'cors';
import { handleConnection } from './db/connection.js';
import todosRoutes from './routes/todos.routes.js';

dotenv.config();
const port = process.env.PORT;

const app = express();

app.use(cors());
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

app.use('/', todosRoutes);

app.listen(port, () => {
  handleConnection();
  console.log(`Server is running on port ${port}`);
});
