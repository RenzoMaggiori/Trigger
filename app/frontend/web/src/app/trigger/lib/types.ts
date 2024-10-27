import { type Node } from "@xyflow/react";
import { ConfigMenuType } from "@/app/trigger/components/config-menu";

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
  type: string;
  fields: Record<string, unknown>;
  parent_ids: Array<string>;
  child_ids: Array<string>;
  x_pos: number;
  y_pos: number;
};
