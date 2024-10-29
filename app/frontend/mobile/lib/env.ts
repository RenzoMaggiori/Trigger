import { z } from "zod";

export const withDevDefault = <T extends z.ZodTypeAny>(
    schema: T,
    val: z.infer<T>,
) => (process.env["NODE_ENV"] !== "production" ? schema.default(val) : schema);

const envSchema = z.object({
    IPv4: z.string().url(),
    ngrok: z.string().url(),
    AUTH_PORT: z.string().url(),
    USER_PORT: z.string().url(),
});

function getEnv() {
    const { success, data, error } = envSchema.safeParse({
        IPV4: process.env['IPv4'],
        NGROK: process.env['ngrok'],
        AUTH_PORT: process.env['AUTH_PORT'],
        USER_PORT: process.env['USER_PORT'],
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

