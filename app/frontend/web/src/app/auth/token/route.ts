import { NextRequest, NextResponse } from "next/server";

import { env } from "@/lib/env";
import { saveToken } from "@/app/auth/lib/save-token";

export async function GET(request: NextRequest) {
  const token = request.nextUrl.searchParams.get("token");
  saveToken(token ?? "");
  return NextResponse.redirect(`${env.NEXT_PUBLIC_WEB_URL}/home`);
}
