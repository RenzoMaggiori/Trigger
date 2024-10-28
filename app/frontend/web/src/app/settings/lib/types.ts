import { z } from "zod";

export const settingsSchema = z.array(z.object({
    id: z.string(),
    userId: z.string(),
    providerName: z.string(),
    active: z.boolean(),
}));

export type SettingsType = {
    id: string,
    userId: string,
    providerName: string,
    active: boolean
};