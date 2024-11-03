import { type Edge } from "@xyflow/react";
import { CustomNode, NodeItem, TriggerWorkspace } from "@/app/trigger/lib/types";

export function transformCustomNodes(customNodes: CustomNode[], edges: Edge[], existingNodes: TriggerWorkspace["nodes"]): Record<string, NodeItem> {
  const nodes: Record<string, NodeItem> = {};

  for (const node of customNodes) {
    nodes[node.id] = {
      id: node.id,
      status: (existingNodes && existingNodes[node.id]) ? existingNodes[node.id]?.action_id : "inactive",
      action_id:  (existingNodes && existingNodes[node.id]) ? existingNodes[node.id]?.action_id : "default",
      fields: {},
      parent_ids: edges
        .filter((edge) => edge.target === node.id)
        .map((edge) => edge.source),
      child_ids: edges
        .filter((edge) => edge.source === node.id)
        .map((edge) => edge.target),
      x_pos: node.position.x,
      y_pos: node.position.y,
    };
  }
  return nodes;
};

