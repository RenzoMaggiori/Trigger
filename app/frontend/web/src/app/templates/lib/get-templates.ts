"use server";

import { env } from "@/lib/env";
import { z } from "zod";
import { cookies } from "next/headers";
import { templateArray } from "@/app/templates/lib/types";

export async function getTemplates(): Promise<z.infer<typeof templateArray>> {
  const accessToken = cookies().get("Authorization")?.value;

  if (!accessToken) {
    throw new Error("could not get access token");
  }

  const res = await fetch(`${env.NEXT_PUBLIC_SERVER_URL}/api/workspace/templates`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${accessToken}`,
    },
  });
  if (!res.ok) {
    console.error("Invalid response status, received: ", res.status);
    throw new Error(`invalid status code: ${res.status}`);
  }

  const body = await res.json();
  const { data, error } = templateArray.safeParse(body);
  if (error) {
    console.error(error);
    throw new Error("could not parse api response");
  }
  return data;
}
