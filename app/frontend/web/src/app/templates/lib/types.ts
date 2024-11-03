import { z } from "zod";

export const templatesSchema = z.object({
    name: z.string(),
    nodes: z.array(
        z.object({
          node_id: z.string(),
          action_id: z.string(),
          input: z.record(z.string(), z.string()).nullable(),
          parents: z.array(z.string()),
          children: z.array(z.string()),
          x_pos: z.number(),
          y_pos: z.number(),
        }),
      ),
});

export const templateArray = z.array(templatesSchema);

export type TemplatesType = z.infer<typeof templatesSchema>
export type TemplatesArrayType = z.infer<typeof templateArray>