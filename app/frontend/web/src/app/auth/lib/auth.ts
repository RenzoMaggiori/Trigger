"use server";

import { redirect } from "next/navigation";

import { env } from "@/lib/env";

export async function login(email: string, password: string) {
  const res = await fetch(
    `${env.NEXT_PUBLIC_AUTH_SERVICE_URL}/api/auth/login`,
    {
      method: "POST",
      body: JSON.stringify({
        email,
        password,
      }),
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
    },
  );

  if (!res.ok) {
    throw new Error(`invalid status code: ${res.status}`);
  }
  console.log("user logged in");
  redirect("/home");
}

export async function register(email: string, password: string) {
  const res = await fetch(
    `${env.NEXT_PUBLIC_AUTH_SERVICE_URL}/api/auth/register`,
    {
      method: "POST",
      body: JSON.stringify({
        user: {
          email,
          password,
        },
      }),
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
    },
  );

  if (!res.ok) {
    throw new Error(`invalid status code: ${res.status}`);
  }
  console.log("user registered");
  redirect("/home");
}
