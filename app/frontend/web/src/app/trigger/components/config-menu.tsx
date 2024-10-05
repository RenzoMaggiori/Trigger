import React from "react";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Label } from "@/components/ui/label";
import { Combox, Status } from "@/components/ui/combox";
import { SiGooglegemini } from "react-icons/si";
import { useMenu } from "./MenuProvider";
import { CustomNode, NodesArrayItem } from "@/app/trigger/lib/types";

type ConfigMenuType = {
    menu: ({ node }: { node: NodesArrayItem }) => React.JSX.Element
    parentNodes: CustomNode[]
    node: CustomNode | null
}

export function ConfigMenu({ menu, parentNodes, node }: ConfigMenuType) {
    const [selectedStatus, setSelectedStatus] = React.useState<Status | null>(null)
    const { triggerWorkspace } = useMenu()
    const nodeItem = triggerWorkspace?.nodes.filter((n) => n.id === node?.id)

    const combinedStatuses: Status[] = [
        { label: <div className="flex flex-row items-center text-md font-bold"><SiGooglegemini className="mr-2" /> None</div>, value: "None" },
        { label: <div className="flex flex-row items-center text-md font-bold"><SiGooglegemini className="mr-2 fill-purple-500" /> Personalized</div>, value: "Personalized" },
        ...parentNodes.map((parentNode) => ({
            label: parentNode.data.label as string,
            value: parentNode.id,
        })),
    ];

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
                    <Combox statuses={combinedStatuses} setSelectedStatus={setSelectedStatus} selectedStatus={selectedStatus} label="info" icon={<SiGooglegemini className="mr-2" />} />
                </div>

                {selectedStatus?.value === "Personalized" && nodeItem && nodeItem.length === 1 && (
                    <div className="p-4 border border-gray-300 rounded-md">
                        <h4 className="text-lg font-bold mb-2">Personalized Settings</h4>
                        {menu({ node: nodeItem[0] })}
                    </div>
                )}

                {selectedStatus && (selectedStatus.value != "Personalized" && selectedStatus.value != "None") && (
                    <div className="mt-4">
                        <h4 className="font-bold">Selected Parent Node ID:</h4>
                        <p>{selectedStatus.value}</p>
                        <h4 className="font-bold">Parent Node Label:</h4>
                        <p>{parentNodes.find((node) => node.id === selectedStatus.value)?.data.label}</p>
                    </div>
                )}
            </CardContent>
        </Card>
    );
}