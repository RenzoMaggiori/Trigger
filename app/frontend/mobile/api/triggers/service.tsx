import { Env } from "@/lib/env";
import AsyncStorage from "@react-native-async-storage/async-storage";

export class TriggersService {
    static async getBaseUrl() {
        return `http://${Env.IPV4}:${Env.ACTION_PORT}/api/action`;
    }

    static async getActions() {
        try {
            const baseUrl = await this.getBaseUrl();
            const response = await  fetch (`${baseUrl}/`, {
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
            const baseUrl = await this.getBaseUrl();
            const response = await fetch(`${baseUrl}/create`, {
                method: 'POST',
                headers: {
                    'Authorization': `Bearer ${await AsyncStorage.getItem('token')}`,
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(trigger),
            });

            if (response.status !== 201) {
                console.log('add trigger failed', response.status);
                throw new Error('Failed to add trigger');
            }
            const data = await response.json();
            console.log('[add trigger] success:', data);
            return data;
        } catch (error) {
            console.error("Catched Add Trigger Error:", error);
            throw error;
        }
    }
}
