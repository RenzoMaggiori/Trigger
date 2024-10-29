import { z } from "zod";

export const triggerSchema = z.object({
  id: z.string(),
  nodes: z.array(
    z.object({
      node_id: z.string(),
      action_id: z.string(),
      input: z.record(z.string(), z.string()),
      output: z.record(z.string(), z.string()),
      parents: z.array(z.string()),
      children: z.array(z.string()),
      status: z.string(),
      x_pos: z.number(),
      y_pos: z.number(),
    }),
  ),
});

export const workspaces = z.array(triggerSchema.pick({
  id: true
}))

export type TriggerSchemaType = z.infer<typeof triggerSchema>;
export type Workspaces = z.infer<typeof workspaces>
