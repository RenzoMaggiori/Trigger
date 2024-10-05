import { type Edge } from "@xyflow/react";
import { CustomNode, NodesArrayItem } from "./types";

export const transformCustomNodes = (customNodes: CustomNode[], edges: Edge[]): NodesArrayItem[] => {
    return customNodes.map((node) => ({
        id: node.id,
        type: node.type || 'default',
        fields: node.data.settings || {},
        parent_ids: edges.filter(edge => edge.target === node.id).map(edge => edge.source),
        child_ids: edges.filter(edge => edge.source === node.id).map(edge => edge.target),
        x_pos: node.position.x,
        y_pos: node.position.y,
    }));
};