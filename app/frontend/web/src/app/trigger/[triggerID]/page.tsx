"use client";

import { Card, CardContent } from "@/components/ui/card";
import React from "react";
import { IoLogoGithub } from "react-icons/io";
import { FaDiscord } from "react-icons/fa";
import { TriggerDraggable } from "@/app/trigger/components/trigger-draggable";
import {
  addEdge,
  ReactFlow,
  type Edge,
  type OnNodesChange,
  type OnEdgesChange,
  applyNodeChanges,
  applyEdgeChanges,
  Background,
  Connection,
} from "@xyflow/react";
import "@xyflow/react/dist/style.css";
import { ConfigMenu } from "@/app/trigger/components/config-menu";
import { BiLogoGmail } from "react-icons/bi";
import { PiMicrosoftOutlookLogo } from "react-icons/pi";
import { CustomNode, NodeItem, Service } from "@/app/trigger/lib/types";
import { useMenu } from "@/app/trigger/components/MenuProvider";
import { transformCustomNodes } from "@/app/trigger/lib/transform-custom-nodes";
import { Button } from "@/components/ui/button";
import { useMutation } from "@tanstack/react-query";
import { send_workspace } from "@/app/trigger/lib/send-workspace";
import { useRouter } from "next/router";

const services: Service[] = [
  {
    icon: <IoLogoGithub className="w-5 h-5 mr-2" />,
    name: "Github",
    settings: "github",
  },
  {
    icon: <FaDiscord className="w-5 h-5 mr-2 text-blue-600" />,
    name: "Discord",
    settings: "discord",
  },
  {
    icon: <BiLogoGmail className="w-5 h-5 mr-2 text-red-600" />,
    name: "Gmail",
    settings: "email",
  },
  {
    icon: <PiMicrosoftOutlookLogo className="w-5 h-5 mr-2 text-sky-500" />,
    name: "Outlook",
    settings: "email",
  },
];

export default function Page({ params }: { params: { triggerID: string } }) {
  const [customNodes, setCustomNodes] = React.useState<CustomNode[]>([]);
  const [edges, setEdges] = React.useState<Edge[]>([]);
  const [settings, setSettings] = React.useState<Service["settings"]>();
  const [parentNodes, setParentNodes] = React.useState<CustomNode[]>([]);
  const [selectedNode, setSelectedNode] = React.useState<CustomNode | null>(
    null,
  );
  const { triggerWorkspace, setTriggerWorkspace, setNodes } = useMenu();

  React.useEffect(() => {
    setTriggerWorkspace((prev) => prev || { id: params.triggerID, nodes: {} });
  }, []);

  React.useEffect(() => {
    if (customNodes.length > 0 || edges.length > 0) {
      const transformedNodes = transformCustomNodes(customNodes, edges);
      setNodes(transformedNodes);
    }
  }, [customNodes, edges]);

  const onNodesChange: OnNodesChange = React.useCallback(
    (changes) => {
      setCustomNodes((nds) => applyNodeChanges(changes, nds) as CustomNode[]);
    },
    [setCustomNodes],
  );
  const onEdgesChange: OnEdgesChange = React.useCallback(
    (changes) => {
      setEdges((eds) => applyEdgeChanges(changes, eds));
    },
    [setEdges],
  );

  const onConnect = React.useCallback(
    (params: Connection | Edge) => {
      setEdges((eds) => addEdge(params, eds));
      if (params.target) {
        updateParentNodes(params.target);
      }
    },
    [setEdges],
  );

  const handleDragStart = (
    e: React.DragEvent<HTMLDivElement>,
    service: Service,
  ): void => {
    const serviceData = { name: service.name };
    e.dataTransfer.setData("service", JSON.stringify(serviceData));
  };

  const handleDragOver = (e: React.DragEvent<HTMLDivElement>): void => {
    e.preventDefault();
  };

  const handleDrop = (e: React.DragEvent<HTMLDivElement>): void => {
    e.preventDefault();
    const reactFlowBounds = e.currentTarget.getBoundingClientRect();
    const droppedService = JSON.parse(e.dataTransfer.getData("service")) as {
      name: string;
    };
    const position = {
      x: e.clientX - reactFlowBounds.left,
      y: e.clientY - reactFlowBounds.top,
    };

    const newService = services.find(
      (service) => service.name === droppedService.name,
    );
    if (newService) {
      const newNode: CustomNode = {
        id: `${droppedService.name}-${customNodes.length}`,
        position,
        data: {
          label: (
            <div className="p-2 flex items-center gap-2">
              {newService.icon}
              <p className="font-bold">{newService.name}</p>
            </div>
          ),
          settings: newService.settings,
        },
        style: { border: "1px solid #ccc", padding: 10 },
      };

      setCustomNodes((nds) => [...nds, newNode]);
    }
  };

  const handleNodeClick = (event: React.MouseEvent, node: CustomNode) => {
    if (node.data?.settings) {
      setSettings(node.data.settings);
      updateParentNodes(node.id);
      setSelectedNode(node);
    }
  };

  const updateParentNodes = (nodeId: string) => {
    const parentNodes = findParentNodes(nodeId, edges, customNodes);
    setParentNodes(parentNodes);
  };

  const findParentNodes = (
    nodeId: string,
    edges: Edge[],
    nodes: CustomNode[],
    visited: Set<string> = new Set(),
  ): CustomNode[] => {
    if (visited.has(nodeId)) return [];
    visited.add(nodeId);

    const parentEdges = edges.filter((edge) => edge.target === nodeId);
    const parentNodes = parentEdges
      .map((edge) => nodes.find((node) => node.id === edge.source))
      .filter(Boolean) as CustomNode[];

    return [
      ...parentNodes,
      ...parentNodes.flatMap((parentNode) =>
        findParentNodes(parentNode.id, edges, nodes, visited),
      ),
    ];
  };

  const mutation = useMutation({
    mutationFn: send_workspace,
    onSuccess: (data) => {
      const nodes: Record<string, NodeItem> = {};
      for (const n of data.nodes) {
        nodes[n.node_id] = {
          id: n.node_id,
          type: n.action_id,
          fields: n.fields,
          parent_ids: n.parents,
          child_ids: n.children,
          x_pos: n.x_pos,
          y_pos: n.y_pos,
        };
      }
      setTriggerWorkspace({
        id: data.id,
        nodes: nodes,
      });
    },
  });

  const handleOnClick = () => {
    if (!triggerWorkspace) return;
    mutation.mutate(triggerWorkspace);
  };

  return (
    <div className="flex h-screen w-full overflow-hidden">
      <div className="w-auto p-5">
        <Card className="h-full">
          <p className="font-bold text-2xl p-3">Services</p>
          <CardContent className="flex flex-col items-center justify-start h-full py-5 gap-4">
            {services.map((item, key) => (
              <div
                key={key}
                draggable
                onDragStart={(e) => handleDragStart(e, item)}
                className="cursor-move"
              >
                <TriggerDraggable service={item} className="w-[200px]" />
              </div>
            ))}
            <Button
              className="w-full text-md rounded-full bg-gradient-to-r from-blue-500 via-violet-500 to-fuchsia-500 hover:bg-gradient-to-r hover:from-blue-600 hover:via-violet-600 hover:to-fuchsia-600 animate-gradient text-white"
              onClick={handleOnClick}
            >
              Deploy Trigger
            </Button>
          </CardContent>
        </Card>
      </div>
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
      <div className="p-5">
        {settings && (
          <ConfigMenu
            menu={settings}
            parentNodes={parentNodes}
            node={selectedNode}
          />
        )}
      </div>
    </div>
  );
}
