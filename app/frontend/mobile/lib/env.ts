import { z } from "zod";

export const withDevDefault = <T extends z.ZodTypeAny>(
    schema: T,
    val: z.infer<T>,
) => (process.env["NODE_ENV"] !== "production" ? schema.default(val) : schema);

const envSchema = z.object({
    NGROK: z.string().url(),
});

function getEnv() {
    const { success, data, error } = envSchema.safeParse({
        NGROK: process.env['NGROK'],
    });

    if (!success) {
        throw new Error(
            "❌ Invalid environment variables:" +
            JSON.stringify(error.format(), null, 4),
        );
    }
    return data;
}

export const Env = getEnv();

