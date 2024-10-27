import { NextRequest, NextResponse } from "next/server";
import { cookies } from "next/headers";
import { env } from "@/lib/env";

export async function GET(request: NextRequest) {
  const redirect = `${env.NEXT_PUBLIC_WEB_URL}/settings`;
  const accessToken = cookies().get("Authorization")?.value;
  const { searchParams } = new URL(request.url);
  const provider = searchParams.get("provider");

  if (!accessToken || !provider) {
    return NextResponse.redirect(redirect);
  }

  const res = await fetch(
    `${env.NEXT_PUBLIC_SYNC_SERVICE_URL}/api/sync/sync-with?provider=${provider}&redirect=${redirect}`,
    {
      method: "GET",
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
      credentials: "include",
    },
  );
  return res;
}
