import { Env } from '@/lib/env';
import { NetworkInfo } from 'react-native-network-info';

const BASE_URL = `http://${Env.IPv4}:${Env.AUTH_PORT}`;

export class CredentialsService {
    static async getBaseUrl() {
        // const ip = await NetworkInfo.getIPV4Address();
        // return `http://${ip}:8000/api/auth`;
        return `${BASE_URL}/api/auth`;
    }

    //? REGISTER
    static async register(email: string, password: string) {
        try {
            const baseUrl = await this.getBaseUrl();
            const response = await fetch(`${baseUrl}/register`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    "user": {
                        email,
                        password,
                    }
                }),
            });
            if (response.status !== 200) {
                console.log('register failed', response.status);
                throw new Error('Something went wrong.');
            }
            console.log('successful register');
            return;
        } catch (error) {
            console.error("Catched Register Error:", error);
            throw error;
        }
    }

    //? LOGIN
    static async login(email: string, password: string) {
        try {
            const baseUrl = await this.getBaseUrl();
            const response = await fetch(`${baseUrl}/login`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    email,
                    password
                }),
            });

            if (response.status !== 200) {
                console.log('login failed', response.status);
                throw new Error('Incorrect username or password.');
            }
            console.log('successful login');
            return;
        } catch (error) {
            console.error("Catched Login Error:", error);
            throw error;
        }
    }
}
