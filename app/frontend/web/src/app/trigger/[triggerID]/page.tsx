"use client"
import { Card, CardContent } from '@/components/ui/card'
import React from 'react'
import { IoLogoGithub } from "react-icons/io";
import { FaDiscord } from "react-icons/fa";
import { Service, TriggerDraggable } from '../components/trigger-draggable';
import { Button } from '@/components/ui/button';
import { addEdge, ReactFlow, useEdgesState, useNodesState, type Node, type Edge, type OnNodesChange, type OnEdgesChange, applyNodeChanges, applyEdgeChanges, Background } from '@xyflow/react';
import '@xyflow/react/dist/style.css';
import { ConfigMenu } from '../components/config-menu';
const initialNodes = [
    { id: '1', position: { x: 0, y: 0 }, data: { label: '1' } },
    { id: '2', position: { x: 0, y: 100 }, data: { label: '2' } },
];
const initialEdges = [{ id: 'e1-2', source: '1', target: '2' }];

const services: Service[] = [
    {
        icon: <IoLogoGithub className='w-5 h-5 mr-2' />, name: "Github", settings: [
            { type: "input", label: "text" }
        ]
    },
    {
        icon: <FaDiscord className='w-5 h-5 mr-2 text-blue-600' />, name: "Discord", settings: [
            { type: "textarea", label: "writeHere", options: ["Message to Send"] }
        ]
    },
];

interface CustomNode extends Node {
    data: {
        label: React.ReactNode;
        settings?: Service["settings"];
    };
}

const Page = () => {
    const [nodes, setNodes] = React.useState<CustomNode[]>([]);
    const [edges, setEdges] = React.useState<Edge[]>([]);
    const [settings, setSettings] = React.useState<Service["settings"]>([]);
    const [draggedIndex, setDraggedIndex] = React.useState<number | null>(null);

    const onNodesChange: OnNodesChange = React.useCallback(
        (changes) => setNodes((nds) => applyNodeChanges(changes, nds) as CustomNode[]),
        [setNodes]
    );
    const onEdgesChange: OnEdgesChange = React.useCallback(
        (changes) => setEdges((eds) => applyEdgeChanges(changes, eds)),
        [setEdges],
    );

    const onConnect = React.useCallback((params: any) => setEdges((eds) => addEdge(params, eds)), [setEdges]);
    const handleDragStart = (e: React.DragEvent<HTMLDivElement>, service: Service, index?: number): void => {
        const serviceData = { name: service.name, index };
        e.dataTransfer.setData("service", JSON.stringify(serviceData));
        if (index !== undefined) setDraggedIndex(index);
    };

    const handleDragOver = (e: React.DragEvent<HTMLDivElement>): void => {
        e.preventDefault();
    };

    const handleDrop = (e: React.DragEvent<HTMLDivElement>): void => {
        e.preventDefault();
        const reactFlowBounds = e.currentTarget.getBoundingClientRect();
        const droppedService = JSON.parse(e.dataTransfer.getData("service")) as { name: string; index?: number };
        const position = {
            x: e.clientX - reactFlowBounds.left,
            y: e.clientY - reactFlowBounds.top,
        };

        const newService = services.find((service) => service.name === droppedService.name);
        if (newService) {
            const newNode: CustomNode = {
                id: `${droppedService.name}-${nodes.length}`,
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

            setNodes((nds) => [...nds, newNode]);
        }

        setDraggedIndex(null);
    };

    const handleNodeClick = (event: React.MouseEvent, node: CustomNode) => {
        if (node.data?.settings) {
            setSettings(node.data.settings);
        }
    };

    return (
        <div className='flex h-screen w-full overflow-hidden'>
            <div className='w-auto p-5'>
                <Card className='h-full'>
                    <p className='font-bold text-2xl p-3'>Services</p>
                    <CardContent className='flex flex-col items-center justify-start h-full py-5 gap-4'>
                        {services.map((item, key) => (
                            <div
                                key={key}
                                draggable
                                onDragStart={(e) => handleDragStart(e, item)}
                                className='cursor-move'
                            >
                                <TriggerDraggable service={item} />
                            </div>
                        ))}
                    </CardContent>
                </Card>
            </div>

            <div className='w-full p-5'>
                <Card className='w-full h-full'>
                    <CardContent
                        className='flex flex-row flex-wrap py-5 gap-x-4 h-full'
                        onDragOver={handleDragOver}
                        onDrop={handleDrop}
                    >
                        <ReactFlow
                            nodes={nodes}
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
            <div className='p-5'>
                {settings.length > 0 && (
                    <ConfigMenu menu={settings}/>
                )}
            </div>
        </div>
    );
};

export default Page;
