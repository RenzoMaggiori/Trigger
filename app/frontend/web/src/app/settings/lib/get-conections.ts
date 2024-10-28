"use server";

import { cookies } from "next/headers";
import { env } from "@/lib/env";

export async function getConnections() {
  const accessToken = cookies().get("Authorization")?.value;
  if (!accessToken) {
    throw new Error("could not get access token");
  }

  const res = await fetch(
    `${env.NEXT_PUBLIC_SETTINGS_SERVICE_URL}/api/settings`,
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

//   const { data, error } = triggerSchema.safeParse(await res.json()); // TODO: settings schema
//   if (error) {
//     console.error(error);
//     throw new Error("could not parse api response");
//   }
//   return data;
}
