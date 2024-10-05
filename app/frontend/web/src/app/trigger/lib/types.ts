import { ReactElement } from "react";
import { type Node } from '@xyflow/react';

export interface CustomNode extends Node {
    data: {
        label: React.ReactNode;
        settings?: Service["settings"];
    };
}

export type NodesArrayItem = {
    id: string
    type: string
    fields: Record<string, any>
    parent_ids: Array<string>
    child_ids: Array<string>
    x_pos: number
    y_pos: number
}

export interface Service {
    name: string;
    icon: React.ReactNode;
    settings: ({node}: {node: NodesArrayItem}) => React.JSX.Element
}

export type TriggerWorkspace = {
    id: number
    nodes: Array<NodesArrayItem>
}