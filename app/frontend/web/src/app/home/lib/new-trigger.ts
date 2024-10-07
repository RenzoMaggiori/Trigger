"use server";

import { cookies } from "next/headers";
import { z } from "zod";

import { env } from "@/lib/env";

export const triggerSchema = z.object({
  id: z.string(),
  userId: z.string(),
  nodes: z.array(
    z.object({
      node_id: z.string(),
      action_id: z.string(),
      fields: z.array(z.any()),
      parents: z.array(z.string()),
      children: z.array(z.string()),
      status: z.string(),
      x_pos: z.number(),
      y_pos: z.number(),
    }),
  ),
});

export async function newTrigger(): Promise<z.infer<typeof triggerSchema>> {
  const accessToken = cookies().get("Authorization")?.value;
  if (!accessToken) {
    throw new Error("could not get access token");
  }

  const res = await fetch(
    `${env.NEXT_PUBLIC_ACTION_SERVICE_URL}/api/workspace/add`,
    {
      method: "POST",
      body: JSON.stringify({ nodes: [] }),
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${accessToken}`,
      },
    },
  );
  if (!res.ok) {
    throw new Error(`invalid status code: ${res.status}`);
  }

  const { data, error } = triggerSchema.safeParse(await res.json());
  if (error) {
    console.error(error);
    throw new Error("could not parse api response");
  }
  return data;
}
