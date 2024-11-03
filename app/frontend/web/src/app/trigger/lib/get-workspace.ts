"use server";

import { cookies } from "next/headers";
import { env } from "@/lib/env";
import { triggerSchema } from "@/app/home/lib/types";

export async function getWorkspace({id}: {id: string}) {
  const accessToken = cookies().get("Authorization")?.value;
  if (!accessToken) {
    throw new Error("could not get access token");
  }

  const res = await fetch(
    `${env.NEXT_PUBLIC_SERVER_URL}/api/workspace/id/${id}`,
    {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${accessToken}`,
      },
    },
  );
  if (!res.ok)
    throw new Error(`invalid status code: ${res.status}`);

  const body = await res.json();
  const { data, error } = triggerSchema.safeParse(body);
  if (error) {
    console.error(error);
    throw new Error("could not parse api response");
  }
  return data;
}
