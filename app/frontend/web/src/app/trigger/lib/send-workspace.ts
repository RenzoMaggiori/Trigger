"use server"

import { TriggerWorkspace } from './types'
import { env } from '@/lib/env';
import { cookies } from 'next/headers';

export async function send_workspace(triggerWorkspace: TriggerWorkspace) {
    const access_token = cookies().get("access_token")?.value;
    const res = await fetch(`${env.NEXT_PUBLIC_AUTH_SERVICE_URL}/api/workspaces`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${access_token}`,
        },
        body: JSON.stringify(triggerWorkspace),
    });
    if (!res.ok) {
        throw new Error(`Failed to send workspace: ${res.status}`);
    }
}
