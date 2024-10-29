import { z } from "zod";

export const withDevDefault = <T extends z.ZodTypeAny>(
  schema: T,
  val: z.infer<T>,
) => (process.env["NODE_ENV"] !== "production" ? schema.default(val) : schema);

const envSchema = z.object({
  NEXT_PUBLIC_AUTH_SERVICE_URL: z.string().url(),
  NEXT_PUBLIC_ACTION_SERVICE_URL: z.string().url(),
  NEXT_PUBLIC_SYNC_SERVICE_URL: z.string().url(),
  NEXT_PUBLIC_SETTINGS_SERVICE_URL: z.string().url(),
  NEXT_PUBLIC_WEB_URL: z.string().url(),
});

function getEnv() {
  const { success, data, error } = envSchema.safeParse({
    NEXT_PUBLIC_AUTH_SERVICE_URL: process.env.NEXT_PUBLIC_AUTH_SERVICE_URL,
    NEXT_PUBLIC_ACTION_SERVICE_URL: process.env.NEXT_PUBLIC_ACTION_SERVICE_URL,
    NEXT_PUBLIC_SYNC_SERVICE_URL: process.env.NEXT_PUBLIC_SYNC_SERVICE_URL,
    NEXT_PUBLIC_SETTINGS_SERVICE_URL: process.env.NEXT_PUBLIC_SETTINGS_SERVICE_URL,
    NEXT_PUBLIC_WEB_URL: process.env.NEXT_PUBLIC_WEB_URL,
  });

  if (!success) {
    throw new Error(
      "❌ Invalid environment variables:" +
        JSON.stringify(error.format(), null, 4),
    );
  }
  return data;
}

export const env = getEnv();
