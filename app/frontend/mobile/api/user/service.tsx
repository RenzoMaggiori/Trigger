import { Env } from "@/lib/env";
import AsyncStorage from "@react-native-async-storage/async-storage";

const BASE_URL = `http://${Env.IPV4}:${Env.USER_PORT}`;

export class UserService {
    static async getBaseUrl() {
        return `${BASE_URL}/api/user`;
    }

    static async getUser(email: string) {
        try {
            const baseUrl = await this.getBaseUrl();
            const response = await fetch (`${baseUrl}/email/${email}`, {
                method: 'GET',
                headers: {
                    'Authorization': `Bearer ${await AsyncStorage.getItem('token')}`,
                    'Content-Type': 'application/json'
                }
            });

            if (response.status !== 200) {
                console.log('get user failed', response.status);
                throw new Error('Something went wrong.');
            }
            const data = await response.json();
            return data;
        } catch (error) {
            console.error("Catched Get User Error:", error);
            throw error;
        }
    }

    static async getUserById(id: string) {
        try {
            const baseUrl = await this.getBaseUrl();
            const response = await fetch (`${baseUrl}/id/${id}`, {
                method: 'GET',
                headers: {
                    'Authorization': `Bearer ${await AsyncStorage.getItem('token')}`,
                    'Content-Type': 'application/json'
                }
            });

            if (response.status !== 200) {
                console.log('get user by id failed', response.status);
                throw new Error('Something went wrong.');
            }
            const data = await response.json();
            return data;
        } catch (error) {
            console.error("Catched Get User By Id Error:", error);
            throw error;
        }
    }
}
