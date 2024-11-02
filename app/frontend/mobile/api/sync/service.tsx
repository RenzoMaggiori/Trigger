import { Env } from '@/lib/env';
import * as WebBrowser from 'expo-web-browser';
import { makeRedirectUri } from 'expo-auth-session';
import AsyncStorage from '@react-native-async-storage/async-storage';
import { SessionService } from '@/api/session/service';
import { UserService } from '@/api/user/service';

export class SyncService {
    static async handleSync(provider: string) {

        const redirectUrl = makeRedirectUri();
        const url = `${Env.NGROK}/api/sync/sync-with?provider=${provider}&redirect=${redirectUrl}&token=${await AsyncStorage.getItem('token')}`;

        console.log('URL:', url);
        try {
            const result = await WebBrowser.openAuthSessionAsync(
                url,
                redirectUrl,
            );
            if (result.type === 'success') {
                await AsyncStorage.setItem(provider, 'true');
                return true;
            } else if (result.type === 'cancel') {
                throw new Error('Browser Canceled');
            } else if (result.type === 'dismiss') {
                throw new Error('Browser Dismissed');
            }
            return false;
        } catch (error) {
            console.error('Failed to open URL:', error);
            throw error;
        }
        return false;
    }
}