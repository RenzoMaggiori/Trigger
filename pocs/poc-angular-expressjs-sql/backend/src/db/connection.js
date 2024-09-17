import mysql from 'mysql2';
import dotenv from 'dotenv';

dotenv.config();

const DATABASE = process.env.MYSQL_DATABASE;
const HOST = process.env.MYSQL_HOST;
const USER = process.env.MYSQL_USER;
const PASSWORD = process.env.MYSQL_ROOT_PASSWORD;

const connection = mysql.createConnection({
    database: DATABASE,
    host: HOST,
    user: USER,
    password: PASSWORD,
});

export const handleConnection = () => {
    connection.connect((error) => {
        if (error) {
            console.error('Error connecting to database:', error);
            setTimeout(handleConnection, 2000);
        } else {
            console.log('Connected to the database');
        }
    });
};
