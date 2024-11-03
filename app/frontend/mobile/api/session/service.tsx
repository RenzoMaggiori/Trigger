import { Env } from "@/lib/env";
import AsyncStorage from "@react-native-async-storage/async-storage";

export class SessionService {
    static async getBaseUrl() {
        return `${Env.NGROK}/api/session`;
    }

    static async getSessionByAccessToken(accessToken: string) {
        try {
            const baseUrl = await this.getBaseUrl();
            const response = await fetch(`${baseUrl}/access/${accessToken}`, {
                method: 'GET',
                headers: {
                    'Authorization': `Bearer ${accessToken}`,
                    'Content-Type': 'application/json'
                }
            });

            if (response.status !== 200) {
                console.log('get session failed', response.status);
                throw new Error('Something went wrong.');
            }
            const data = await response.json();
            return data[0];
        } catch (error) {
            console.error("Catched Get Session By Access Token Error:", error);
            throw error;
        }
    }

    static async getSessionByUserId(userId: string) {
        try {
            const baseUrl = await this.getBaseUrl();
            const response = await fetch(`${baseUrl}/user/${userId}`, {
                method: 'GET',
                headers: {
                    'Authorization': `Bearer ${await AsyncStorage.getItem('token')}`,
                    'Content-Type': 'application/json'
                }
            });

            if (response.status !== 200) {
                console.log('get session by user id failed', response.status);
                throw new Error('Something went wrong.');
            }
            const data = await response.json();
            console.log('-----data:', data);
            return data[0];
        } catch (error) {
            console.error("Catched Get Session By User Id Error:", error);
            throw error;
        }
    }
}