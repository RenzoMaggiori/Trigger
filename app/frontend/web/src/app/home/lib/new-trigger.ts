"use server"

import { env } from '@/lib/env';
import { cookies } from 'next/headers';

export async function newtrigger() {
    const access_token = cookies().get("access_token")?.value;
    const res = await fetch(
        `${env.NEXT_PUBLIC_AUTH_SERVICE_URL}/api/`,
        {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${access_token}`,
            },
        },
    );

    if (!res.ok) {
        throw new Error(`invalid status code: ${res.status}`);
    }
    const data = await res.json();
    return data.trigger_id;
}
