import { z } from "zod";

const envClientSchema = z.object({
  NEXT_PUBLIC_AUTH_SERVICE_URL: z.string().url(),
});

export function getEnv() {
  if (typeof window === undefined) {
    throw new Error("Not in the client");
  }

  const { data, error } = envClientSchema.safeParse(process.env);
  if (error) {
    console.error(error);
    throw new Error(error.toString());
  }
  return data;
}
