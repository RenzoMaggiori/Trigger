import { getEnv } from "@/lib/env";


export async function login(email: string, password: string) {
  const env = getEnv();
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
    },
  );

  if (!res.ok) {
    throw new Error(`invalid status code: ${res.status}`);
  }
  console.log(res);
}

export async function register(email: string, password: string) {
  const env = getEnv();
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
    },
  );

  if (!res.ok) {
    throw new Error(`invalid status code: ${res.status}`);
  }
  console.log(res);
}

