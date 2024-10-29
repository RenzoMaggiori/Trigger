import { Env } from '@/lib/env';
import AsyncStorage from '@react-native-async-storage/async-storage';
import { NetworkInfo } from 'react-native-network-info';

const BASE_URL = `http://${Env.IPV4}:${Env.AUTH_PORT}`;

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
            const cookieHeader = response.headers.get('set-cookie');
            if (cookieHeader) {
                const tokenMatch = cookieHeader.match(/Authorization=([^;]+)/);
                if (tokenMatch) {
                    const token = tokenMatch[1];
                    await AsyncStorage.setItem('token', token);
                } else {
                    console.error('No token found in cookie header:', cookieHeader);
                }
            } else {
                console.error('No cookie header found in response:', response.headers);
            }
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

            const cookieHeader = response.headers.get('set-cookie');
            if (cookieHeader) {
                const tokenMatch = cookieHeader.match(/Authorization=([^;]+)/);
                if (tokenMatch) {
                    const token = tokenMatch[1];
                    await AsyncStorage.setItem('token', token);
                } else {
                    console.error('No token found in cookie header:', cookieHeader);
                }
            } else {
                console.error('No cookie header found in response:', response.headers);
            }
            return;
        } catch (error) {
            console.error("Catched Login Error:", error);
            throw error;
        }
    }
}
