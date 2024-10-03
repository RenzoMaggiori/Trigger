import React from "react";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { CustomNode } from "../[triggerID]/page";
import { Edge } from "@xyflow/react";
import { Label } from "@/components/ui/label";

export function ConfigMenu({ menu, parentNodes, node }: { menu: React.JSX.Element, parentNodes: CustomNode[], node: CustomNode | null }) {
    const [selectedParent, setSelectedParent] = React.useState<string | null>(null);

    const handleParentSelection = (e: React.ChangeEvent<HTMLSelectElement>) => {
        const selectedId = e.target.value;
        setSelectedParent(selectedId);
    };

    return (
        <Card className="h-full w-[500px]">
            <CardHeader>
                <CardTitle className="flex items-center text-xl font-bold">{node?.data?.label} Settings</CardTitle>
                <CardDescription className="ml-2 text-md">ID: {node?.id}</CardDescription>
            </CardHeader>
            <CardContent>
                <div className="mb-4">
                    <Label htmlFor="parent-node-dropdown" className="block text-sm font-medium text-gray-700">
                        Information to Send
                    </Label>
                    <select
                        id="parent-node-dropdown"
                        value={selectedParent ?? ''}
                        onChange={handleParentSelection}
                        className="mt-1 block w-full p-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                    >
                        <option value="">None</option>
                        <option value="Personalized">Personalized</option>
                        {parentNodes.map((parentNode) => (
                            <option key={parentNode.id} value={parentNode.id}>
                                {parentNode.data.label}
                            </option>
                        ))}
                    </select>
                </div>

                {selectedParent === "Personalized" && (
                    <div className="p-4 border border-gray-300 rounded-md">
                        <h4 className="text-lg font-bold mb-2">Personalized Settings</h4>
                        {menu}
                    </div>
                )}

                {selectedParent && (
                    <div className="mt-4">
                        <h4 className="font-bold">Selected Parent Node ID:</h4>
                        <p>{selectedParent}</p>
                        <h4 className="font-bold">Parent Node Label:</h4>
                        <p>{parentNodes.find((node) => node.id === selectedParent)?.data.label}</p>
                    </div>
                )}
            </CardContent>
        </Card>
    );
}