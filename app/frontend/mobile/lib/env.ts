import { z } from "zod";

export const withDevDefault = <T extends z.ZodTypeAny>(
    schema: T,
    val: z.infer<T>,
) => (process.env["NODE_ENV"] !== "production" ? schema.default(val) : schema);

const envSchema = z.object({
    IPV4: z.string(),
    NGROK: z.string().url(),
    AUTH_PORT: z.string(),
    USER_PORT: z.string(),
    ACTION_PORT: z.string(),
});

function getEnv() {
    const { success, data, error } = envSchema.safeParse({
        IPV4: process.env['IPV4'],
        NGROK: process.env['NGROK'],
        AUTH_PORT: process.env['AUTH_PORT'],
        USER_PORT: process.env['USER_PORT'],
        ACTION_PORT: process.env['ACTION_PORT'],
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

