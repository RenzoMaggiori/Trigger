"use server";

import { TriggerWorkspace } from "./types";
import { env } from "@/lib/env";
import { cookies } from "next/headers";
import { triggerSchema } from "@/app/home/lib/types";

export async function send_workspace(triggerWorkspace: TriggerWorkspace) {
  const access_token = cookies().get("Authorization")?.value;

  const res = await fetch(
    `${env.NEXT_PUBLIC_SERVER_URL}/api/workspace/id/${triggerWorkspace.id}`,
    {
      method: "PATCH",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${access_token}`,
      },
      body: JSON.stringify({
        nodes: Object.keys(triggerWorkspace.nodes).map((k) => ({
          node_id: k,
          action_id: triggerWorkspace.nodes[k].action_id,
          input: triggerWorkspace.nodes[k].fields || {},
          parents: triggerWorkspace.nodes[k].parent_ids,
          children: triggerWorkspace.nodes[k].child_ids,
          x_pos: triggerWorkspace.nodes[k].x_pos,
          y_pos: triggerWorkspace.nodes[k].y_pos,
        })),
      }),
    },
  );
  if (!res.ok) {
    throw new Error(`Failed to send workspace: ${res.status}`);
  }

  const body = await res.json()
  const { data, error } = triggerSchema.safeParse(body);
  if (error) {
    console.error(error);
    throw new Error("could not parse api response");
  }
  return data;
}
