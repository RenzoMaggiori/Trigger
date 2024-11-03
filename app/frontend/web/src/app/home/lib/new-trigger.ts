"use server";

import { cookies } from "next/headers";
import { z } from "zod";

import { env } from "@/lib/env";
import { triggerSchema } from "@/app/home/lib/types";
import { TemplatesType } from "@/app/templates/lib/types";

export async function newTrigger(workspace: TemplatesType): Promise<z.infer<typeof triggerSchema>> {
  const accessToken = cookies().get("Authorization")?.value;
  if (!accessToken) {
    throw new Error("could not get access token");
  }

  const res = await fetch(
    `${env.NEXT_PUBLIC_SERVER_URL}/api/workspace/add`,
    {
      method: "POST",
      body: JSON.stringify({ name: workspace.name, nodes: workspace.nodes || [] }),
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${accessToken}`,
      },
    },
  );
  if (!res.ok) {
    console.error(`invalid status code: ${res.status}`);
    throw new Error(`invalid status code: ${res.status}`);
  }

  const { data, error } = triggerSchema.safeParse(await res.json());
  if (error) {
    console.error(error);
    throw new Error("could not parse api response");
  }
  return data;
}
