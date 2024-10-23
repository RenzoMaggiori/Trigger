import { z } from "zod";

export const triggerSchema = z.object({
  id: z.string(),
  nodes: z.array(
    z.object({
      node_id: z.string(),
      action_id: z.string(),
      fields: z.record(z.string(), z.unknown()),
      parents: z.array(z.string()),
      children: z.array(z.string()),
      status: z.string(),
      x_pos: z.number(),
      y_pos: z.number(),
    }),
  ),
});
