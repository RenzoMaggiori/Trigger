import { Env } from "@/lib/env";

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
}