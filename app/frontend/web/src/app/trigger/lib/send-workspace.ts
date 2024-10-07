import React from 'react'
import { TriggerWorkspace } from './types'
import { env } from '@/lib/env';

export async function send_workspace(triggerWorkspace: TriggerWorkspace) {
    const res = await fetch(`${env.NEXT_PUBLIC_AUTH_SERVICE_URL}/api/workspaces`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(triggerWorkspace),
    });
    if (!res.ok) {
        throw new Error(`Failed to send workspace: ${res.status}`);
    }
}
