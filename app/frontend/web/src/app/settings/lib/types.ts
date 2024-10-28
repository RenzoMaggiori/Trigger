import { z } from "zod";

export const settingsSchema = z.array(z.object({
    id: z.string().optional(),
    userId: z.string().optional(),
    providerName: z.string(),
    active: z.boolean(),
}));

export type SettingsType = z.infer<typeof settingsSchema>