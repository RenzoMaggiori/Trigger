import { env } from "@/lib/env";
import { cookies } from "next/headers";

export async function sync(provider: string) {
  const accessToken = cookies().get("Authorization")?.value;
  if (!accessToken) {
    throw new Error("could not get access token");
  }

  const res = await fetch(
    `${env.NEXT_PUBLIC_SYNC_SERVICE_URL}/api/sync/sync-with?provider=${provider}&redirect=${env.NEXT_PUBLIC_WEB_URL}/settings`,
    {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${accessToken}`,
      },
      credentials: "include",
    },
  );

  if (!res.ok) {
    throw new Error(`invalid status code: ${res.status}`);
  }
}
