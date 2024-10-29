"use server"

import { env } from "@/lib/env";
import { workspaces } from "@/app/home/lib/types";
import { z } from "zod";
import { cookies } from 'next/headers';

export async function getWorkspaces(): Promise<z.infer<typeof workspaces>> {
  const accessToken = cookies().get("Authorization")?.value;

  if (!accessToken) {
    throw new Error("could not get access token");
  }

  const res = await fetch(
    `${env.NEXT_PUBLIC_ACTION_SERVICE_URL}/api/workspace/me`,
    {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${accessToken}`,
      },
    },
  );
  if (!res.ok) {
    console.error("Invalid response status, received: ", res.status);
    throw new Error(`invalid status code: ${res.status}`);
  }

  const body = await res.json();
  const { data, error } = workspaces.safeParse(body);
  if (error) {
    console.error(error);
    throw new Error("could not parse api response");
  }
  return data;
}
