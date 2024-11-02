"use server";

import { cookies } from "next/headers";
import { env } from "@/lib/env";
import { triggerSchema } from "@/app/home/lib/types";

export async function deployWorkspace({id}: {id: string}) {
  const accessToken = cookies().get("Authorization")?.value;
  if (!accessToken) {
    throw new Error("could not get access token");
  }

  const res = await fetch(
    `${env.NEXT_PUBLIC_SERVER_URL}/api/workspace/start/id/${id}`,
    {
      method: "PATCH",
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    },
  );
  if (!res.ok)
    throw new Error(`invalid status code: ${res.status}`);

  const { data, error } = triggerSchema.safeParse(await res.json());
  if (error) {
    console.error(error);
    throw new Error("could not parse api response");
  }
  return data;
}
