"use server"

import { TriggerWorkspace } from './types'
import { env } from '@/lib/env';
import { cookies } from 'next/headers';
import { triggerSchema } from "@/app/home/lib/new-trigger"

export async function send_workspace(triggerWorkspace: TriggerWorkspace) {
    const access_token = cookies().get("Authorization")?.value;
    const res = await fetch(`${env.NEXT_PUBLIC_AUTH_SERVICE_URL}/api/workspace/id/${triggerWorkspace.id}`, {
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${access_token}`,
        },
        body: JSON.stringify({
            nodes: Object.keys(triggerWorkspace.nodes).map((k) => ({
                node_id: k,
                fields: triggerWorkspace.nodes[k].fields,
                parents: triggerWorkspace.nodes[k].parent_ids,
                children: triggerWorkspace.nodes[k].child_ids,
                x_pos: triggerWorkspace.nodes[k].x_pos,
                y_pos: triggerWorkspace.nodes[k].y_pos,
            }))
        }),
    });
    if (!res.ok) {
        throw new Error(`Failed to send workspace: ${res.status}`);
    }
    const { data, error } = triggerSchema.safeParse(await res.json());
    if (error) {
      console.error(error);
      throw new Error("could not parse api response");
    }
    return data
}
