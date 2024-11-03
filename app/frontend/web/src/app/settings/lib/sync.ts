"use server"

import { env } from "@/lib/env";
import { cookies } from "next/headers";

export async function sync(provider: string) {
  const redirect = `${env.NEXT_PUBLIC_WEB_URL}/settings`;
  const accessToken = cookies().get("Authorization")?.value;
  if (!accessToken) {
    throw new Error("could not get access token");
  }

  return `${env.NEXT_PUBLIC_SERVER_URL}/api/sync/sync-with?provider=${provider}&redirect=${redirect}&token=${accessToken}`
}
