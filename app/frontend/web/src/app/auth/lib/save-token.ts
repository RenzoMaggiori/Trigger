"use server";

import { cookies } from "next/headers";

export async function saveToken(token: string) {
  cookies().set("Authorization", token, {});
}
