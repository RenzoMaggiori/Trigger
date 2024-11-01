"use client";

import React from "react";
import { IoLogoGithub } from "react-icons/io";
import { FaDiscord, FaTwitch, FaSpotify } from "react-icons/fa";
import { type Edge } from "@xyflow/react";
import "@xyflow/react/dist/style.css";
import { ConfigMenu } from "@/app/trigger/components/config-menu";
import { BiLogoGmail } from "react-icons/bi";
import { CustomNode, NodeItem, Service } from "@/app/trigger/lib/types";
import { useMenu } from "@/app/trigger/components/MenuProvider";
import { transformCustomNodes } from "@/app/trigger/lib/transform-custom-nodes";
import { useQuery } from "@tanstack/react-query";
import { ServicesComponent } from "@/app/trigger/components/service-page";
import { ReactFlowComponent } from "@/app/trigger/components/react-flow";
import { getWorkspace } from "@/app/trigger/lib/get-workspace";
import { getActions } from "../lib/get-actions";

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
    icon: <FaSpotify className="w-5 h-5 mr-2 text-green-500" />,
    name: "Spotify",
    settings: "spotify",
  },
  {
    icon: <FaTwitch className="w-5 h-5 mr-2 text-violet-500" />,
    name: "Twitch",
    settings: "twitch",
  }
];

export default function Page({ params }: { params: { triggerID: string } }) {
  const [customNodes, setCustomNodes] = React.useState<CustomNode[]>([]);
  const [edges, setEdges] = React.useState<Edge[]>([]);
  const [settings, setSettings] = React.useState<Service["settings"]>();
  const [parentNodes, setParentNodes] = React.useState<CustomNode[]>([]);
  const [selectedNode, setSelectedNode] = React.useState<CustomNode | null>(null);

  const { triggerWorkspace, setTriggerWorkspace, setNodes } = useMenu();

  const transformedNodes = React.useMemo(() => {
    return transformCustomNodes(customNodes, edges, triggerWorkspace?.nodes || {});
  }, [customNodes, edges]);

  React.useEffect(() => {
    if (customNodes.length > 0 || edges.length > 0) {
      setNodes(transformedNodes);
    }
  }, [customNodes, edges, transformedNodes]);


  const { data, isPending, error } = useQuery({
    queryKey: ["workspaceAndActions", params.triggerID],
    queryFn: async () => {
      const [workspace, actions] = await Promise.all([
        getWorkspace({ id: params.triggerID }),
        getActions(),
      ]);
      return { workspace, actions };
    },
  });

  React.useEffect(() => {
    if (!data?.workspace) return;

    setTriggerWorkspace({
      id: data.workspace.id,
      nodes: data.workspace.nodes.reduce((acc, n) => {
        acc[n.node_id] = {
          id: n.node_id,
          action_id: n.action_id || "",
          fields: n.input || {},
          parent_ids: n.parents || [],
          child_ids: n.children || [],
          x_pos: n.x_pos || 0,
          y_pos: n.y_pos || 0,
        };
        return acc;
      }, {} as { [key: string]: NodeItem }),
    });
  }, [data, setTriggerWorkspace]);

  const findService = (nodeId: string) => {
    const serviceName = nodeId.split("-")[0];
    return services.find((service) => service.name.toLowerCase() === serviceName.toLowerCase());
  };

  React.useEffect(() => {
    if (!data) return;

    const nodes = data.workspace.nodes.map((n) => {
      const service = findService(n.node_id);
      return {
        id: n.node_id,
        position: { x: n.x_pos || 0, y: n.y_pos || 0 },
        data: {
          label: (
            <div className="p-2 flex items-center gap-2">
              {service?.icon}
              <p className="font-bold">{service?.name || n.node_id}</p>
            </div>
          ),
          settings: service?.settings,
        },
        style: { border: "1px solid #ccc", padding: 10 },
        parents: n.parents || [],
        children: n.children || [],
      };
    });

    const newEdges = data.workspace.nodes.flatMap((n) =>
      (n.children || []).map((childId) => ({
        id: `${n.node_id}`,
        source: n.node_id,
        target: childId,
        style: { stroke: "#ddd", strokeWidth: 2 },
      }))
    );

    setCustomNodes(nodes);
    setEdges(newEdges);
  }, [data, setCustomNodes, setEdges]);

  if (error) return <div>failed to get workspace.</div>
  if (isPending) return <div>loading...</div>

  const updateParentNodes = (nodeId: string) => {
    const parentNodes = findParentNodes(nodeId, edges, customNodes);
    setParentNodes(parentNodes);
  };

  const handleNodeClick = (event: React.MouseEvent, node: CustomNode) => {
    if (node.data?.settings) {
      setSettings(node.data.settings);
      updateParentNodes(node.id);
      setSelectedNode(node);
    }
  };

  const handleDragStart = (e: React.DragEvent<HTMLDivElement>, service: Service): void => {
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

  return (
    <div className="flex h-screen w-full overflow-hidden">
      <div className="hidden md:flex h-full">

        <ServicesComponent
          services={services}
          handleDragStart={handleDragStart}
        />
      </div>
      <ReactFlowComponent
        customNodes={customNodes}
        setCustomNodes={setCustomNodes}
        edges={edges}
        setEdges={setEdges}
        handleNodeClick={handleNodeClick}
        handleDrop={handleDrop}
        handleDragOver={handleDragOver}
        updateParentNodes={updateParentNodes}
      />
      <div className="p-5 hidden md:flex h-full">
        {settings && (
          <ConfigMenu
            menu={settings}
            parentNodes={parentNodes}
            node={selectedNode}
            actions={data.actions}
          />
        )}
      </div>
    </div>
  );
}

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
