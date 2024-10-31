import { type Node } from "@xyflow/react";
import { ConfigMenuType } from "@/app/trigger/components/config-menu";
import { z } from "zod";

export interface Service {
  name: string;
  icon: React.ReactNode;
  settings: ConfigMenuType["menu"];
}

export interface CustomNode extends Node {
  data: {
    label: React.ReactNode;
    settings?: Service["settings"];
  };
}

export type TriggerWorkspace = {
  id: string;
  nodes: Record<string, NodeItem>;
};

export type NodeItem = {
  id: string;
  action_id: string;
  fields: Record<string, unknown>;
  parent_ids: Array<string>;
  child_ids: Array<string>;
  x_pos: number;
  y_pos: number;
};


export const actionsSchema = z.array(z.object({
  id: z.string(),
  input: z.array(z.string()),
  output: z.array(z.string()),
  provider: z.string(),
  type: z.string(),
  action: z.string(),
}));

export type ActionType = z.infer<typeof actionsSchema>