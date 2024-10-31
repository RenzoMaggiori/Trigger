import { Env } from '@/lib/env';
import * as WebBrowser from 'expo-web-browser';

const BASE_URL = `http://${Env.IPV4}:${Env.AUTH_PORT}`;

export class ProvidersService {
    //? OAUTH
    static async handleOAuth(provider: string) {
        const url = `${Env.NGROK}/api/oauth2/login?provider=${provider}&redirect=${Env.NGROK}/api/oauth2/callback`;

        try {
            const result = await WebBrowser.openAuthSessionAsync(
                url,
                `${BASE_URL}/api/oauth2/callback`
            );
            if (result.type === 'cancel') {
                throw new Error('Browser Canceled');
            } else if (result.type === 'dismiss') {
                throw new Error('Browser Dismissed');
            } else {
                console.log('Browser opened successfully:', result);
            }
            if (result.type === 'success') {
                console.log('result: ', result?.url);
                return;
            } else {
                throw new Error('Browser Failed');
            }
        } catch (error) {
            console.error('Failed to open URL:', error);
            throw error;
        }
    }
}
