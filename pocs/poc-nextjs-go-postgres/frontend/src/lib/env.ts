import { z } from "zod";

export const withDevDefault = <T extends z.ZodTypeAny>(schema: T, val: z.infer<T>) =>
    process.env["NODE_ENV"] !== "production" ? schema.default(val) : schema;

export const envSchema = z.object({
    NEXT_PUBLIC_API_URL: withDevDefault(z.string(), "http://localhost:8000"),
});

const { success, data, error } = envSchema.safeParse(process.env);

if (!success) {
    throw new Error(
        "❌ Invalid environment variables:" +
            JSON.stringify(error.format(), null, 4),
    );
}

export const env = data;

