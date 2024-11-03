import { Env } from '@/lib/env';
import AsyncStorage from '@react-native-async-storage/async-storage';

export class CredentialsService {
    static async getBaseUrl() {
        return `${Env.NGROK}/api/auth`;
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
            const token = response.headers.get('Authorization');
            if (token) {
                await AsyncStorage.setItem('token', token);
            } else {
                console.error('No token found in response:');
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
            const token = response.headers.get('Authorization');
            if (token) {
                await AsyncStorage.setItem('token', token);
            } else {
                console.error('No token found in response:');
            }
            return;
        } catch (error) {
            console.error("Catched Login Error:", error);
            throw error;
        }
    }

    //? LOGOUT
    static async logout() {
        try {
            const baseUrl = await this.getBaseUrl();
            const response = await fetch(`${baseUrl}/logout`, {
                method: 'POST',
                headers: {
                    'Authorization': `Bearer ${await AsyncStorage.getItem('token')}`,
                    'Content-Type': 'application/json'
                }
            });
            if (response.status !== 200) {
                console.log('logout failed', response.status);
                throw new Error('Something went wrong.');
            }
            console.log('successful logout');
            await AsyncStorage.removeItem('token');
            return;
        } catch (error) {
            console.error("Catched Logout Error:", error);
            throw error;
        }
    }
}
