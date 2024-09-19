import mysql from 'mysql2';
import dotenv from 'dotenv';

dotenv.config();

export const dbConnection = mysql.createConnection({
    host: process.env.DB_HOST,
    user: process.env.DB_USER,
    password: process.env.DB_PASSWORD,
    database: process.env.DB_DATABASE,
    port: process.env.DB_PORT,
});

export const handleConnection = () => {
    dbConnection.connect((error) => {
        if (error) {
            console.error('Error connecting to database:', error);
            setTimeout(handleConnection, 2000);
        } else {
            console.log('Connected to the database');
        }
    });
};
