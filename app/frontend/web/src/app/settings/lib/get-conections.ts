"use server";

import { cookies } from "next/headers";
import { env } from "@/lib/env";
import { settingsSchema } from "@/app/settings/lib/types";

export async function getConnections() {
  const accessToken = cookies().get("Authorization")?.value;
  if (!accessToken) {
    throw new Error("could not get access token");
  }

  const res = await fetch(
    `${env.NEXT_PUBLIC_SETTINGS_SERVICE_URL}/api/settings/me`,
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

  const { data, error } = settingsSchema.safeParse(await res.json());
  if (error) {
    console.error(error);
    throw new Error("could not parse api response");
  }
  return data;
}
