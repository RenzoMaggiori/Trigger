import { Env } from "@/lib/env";
import AsyncStorage from "@react-native-async-storage/async-storage";

export class TriggersService {
    static async getBaseUrl() {
        return `http://${Env.IPV4}:${Env.ACTION_PORT}/api`;
    }

    static async getActions() {
        try {
            const baseUrl = await this.getBaseUrl();
            const response = await  fetch (`${baseUrl}/action`, {
                method: 'GET',
                headers: {
                    'Authorization': `Bearer ${await AsyncStorage.getItem('token')}`,
                    'Content-Type': 'application/json'
                }
            });

            if (response.status !== 200) {
                console.log('get triggers failed', response.status);
                throw new Error('Something went wrong.');
            }
            const data = await response.json();
            // console.log('[get actions] success: ', data);
            return data;
        } catch (error) {
            console.error("Catched Get Triggers Error:", error);
            throw error;
        }
    }

    static async getTriggersByProvider(provider: string) {
        if (!provider) {
            throw new Error('Provider is required');
        }
        try {
            const actionList = await this.getActions();
            const actionsProvider = actionList.filter((action: any) => action.provider === provider);
            const triggers = actionsProvider.filter((action: any) => action.type === 'trigger');
            // console.log('[get actions by provider] success: ', actions);
            return triggers;
        } catch (error) {
            console.error("Catched Get Actions By Provider Error:", error);
            throw error;
        }
    }

    static async getReactionsByProvider(provider: string) {
        if (!provider) {
            throw new Error('Provider is required');
        }
        try {
            const actionList = await this.getActions();
            const actionsProvider = actionList.filter((action: any) => action.provider === provider);
            const reactions = actionsProvider.filter((action: any) => action.type === 'reaction');
            // console.log('[get actions by provider] success: ', actions);
            return reactions;
        } catch (error) {
            console.error("Catched Get Actions By Provider Error:", error);
            throw error;
        }
    }

    static async addTrigger(trigger: any) {
        try {
            // const baseUrl = this.getBaseUrl();
            const token = await AsyncStorage.getItem('token');
            const response = await fetch(`https://ample-social-elk.ngrok-free.app/api/workspace/add`, {
                method: 'POST',
                headers: {
                    'Authorization': `Bearer ${token}`,
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(trigger),
            });
            if (response.status !== 201 && response.status !== 200) {
                console.log('add trigger failed', response.status);
                throw new Error(`Failed to add trigger. Status: ${response.status}`);
            }
            const data = await response.json();
            console.log('[add trigger] success:', data);
            return data;
        } catch (error) {
            console.error("Catched Add Trigger Error:", error);
            throw error;
        }
    }

    static async getTriggers() {
        try {
            const baseUrl = await this.getBaseUrl();
            const response = await fetch(`${baseUrl}/workspace/me`, {
                method: 'GET',
                headers: {
                    'Authorization': `Bearer ${await AsyncStorage.getItem('token')}`,
                    'Content-Type': 'application/json'
                }
            });
            if (response.status !== 200) {
                console.log('get triggers failed', response.status);
                throw new Error('Something went wrong.');
            }
            const data = await response.json();
            console.log('[get triggers] success: ', data);
            return data;
        } catch (error) {
            console.error("Catched Get Triggers Error:", error);
            throw error;
        }
    }
}
