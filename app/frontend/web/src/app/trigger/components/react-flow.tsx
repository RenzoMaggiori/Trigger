import React from "react";
import { ReactFlow, type Edge, type OnNodesChange, type OnEdgesChange, Background, Connection, addEdge, applyNodeChanges, applyEdgeChanges } from "@xyflow/react";
import { CustomNode } from "@/app/trigger/lib/types";
import { Card, CardContent } from "@/components/ui/card";

interface ReactFlowProps {
  customNodes: CustomNode[];
  setCustomNodes: React.Dispatch<React.SetStateAction<CustomNode[]>>;
  edges: Edge[];
  setEdges: React.Dispatch<React.SetStateAction<Edge[]>>;
  handleNodeClick: (event: React.MouseEvent, node: CustomNode) => void;
  handleDrop: (e: React.DragEvent<HTMLDivElement>) => void;
  handleDragOver: (e: React.DragEvent<HTMLDivElement>) => void;
  updateParentNodes: (nodeId: string) => void;
}

export const ReactFlowComponent: React.FC<ReactFlowProps> = ({
  customNodes,
  setCustomNodes,
  edges,
  setEdges,
  handleNodeClick,
  handleDrop,
  handleDragOver,
  updateParentNodes,
}) => {
  const onNodesChange: OnNodesChange = React.useCallback(
    (changes) => {
      setCustomNodes((nds) => applyNodeChanges(changes, nds) as CustomNode[]);
    },
    [setCustomNodes]
  );

  const onEdgesChange: OnEdgesChange = React.useCallback(
    (changes) => {
      setEdges((eds) => applyEdgeChanges(changes, eds));
    },
    [setEdges]
  );

  const onConnect = React.useCallback(
    (params: Connection | Edge) => {
      setEdges((eds) => addEdge(params, eds));
      if (params.target) {
        updateParentNodes(params.target);
      }
    },
    [setEdges, updateParentNodes]
  );

  return (
    <div className="w-full p-5">
      <Card className="w-full h-full">
        <CardContent
          className="flex flex-row flex-wrap py-5 gap-x-4 h-full"
          onDragOver={handleDragOver}
          onDrop={handleDrop}
        >
          <ReactFlow
            nodes={customNodes}
            edges={edges}
            onNodesChange={onNodesChange}
            onEdgesChange={onEdgesChange}
            onConnect={onConnect}
            onNodeClick={handleNodeClick}
          >
            <Background />
          </ReactFlow>
        </CardContent>
      </Card>
    </div>
  );
};
