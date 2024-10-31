"use server";

import { cookies } from "next/headers";
import { env } from "@/lib/env";
import { actionsSchema } from "@/app/trigger/lib/types";

export async function getActions() {
  const accessToken = cookies().get("Authorization")?.value;
  if (!accessToken) {
    throw new Error("could not get access token");
  }

  const res = await fetch(
    `${env.NEXT_PUBLIC_ACTION_SERVICE_URL}/api/action`,
    {
      method: "GET",
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    },
  );
  if (!res.ok) {
    console.error("invalid status code: ", res.status);
    throw new Error(`invalid status code: ${res.status}`);
  }

  const body = await res.json();
  const { data, error } = actionsSchema.safeParse(body);
  if (error) {
    console.error(error);
    throw new Error("could not parse api response");
  }
  return data;
}
