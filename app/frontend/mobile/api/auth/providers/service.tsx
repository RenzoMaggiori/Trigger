import { Env } from '@/lib/env';
import * as WebBrowser from 'expo-web-browser';
import { makeRedirectUri } from 'expo-auth-session';
import AsyncStorage from '@react-native-async-storage/async-storage';
import { SessionService } from '@/api/session/service';
import { UserService } from '@/api/user/service';

export class ProvidersService {
    //? OAUTH
    static async handleOAuth(provider: string) {
        const redirectUrl = makeRedirectUri();
        const url = `${Env.NGROK}/api/oauth2/login?provider=${provider}&redirect=${redirectUrl}`;

        console.log('URL:', url);
        try {
            const result = await WebBrowser.openAuthSessionAsync(
                url,
                redirectUrl,
            );
            if (result.type === 'success') {
                const { url } = result;
                const parsedUrl = new URL(url);
                const token = parsedUrl.searchParams.get('token');
                if (!token) {
                    throw new Error('No token found');
                }
                await AsyncStorage.setItem('token', token);
                const session = await SessionService.getSessionByAccessToken(token);
                const user = await UserService.getUserById(session.user_id);
                await AsyncStorage.setItem('user', JSON.stringify(user));

            } else if (result.type === 'cancel') {
                throw new Error('Browser Canceled');
            } else if (result.type === 'dismiss') {
                throw new Error('Browser Dismissed');
            }
        } catch (error) {
            console.error('Failed to open URL:', error);
            throw error;
        }
    }
}
