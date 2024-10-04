import React from "react";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { CustomNode, NodesArrayItem } from "../[triggerID]/page";
import { Label } from "@/components/ui/label";
import { Combox, Status } from "@/components/ui/combox";
import { SiGooglegemini } from "react-icons/si";
import { MenuProvider } from "./MenuProvider";

export function ConfigMenu({ menu, parentNodes, node }: { menu: React.JSX.Element, parentNodes: CustomNode[], node: CustomNode | null }) {
    const [selectedStatus, setSelectedStatus] = React.useState<Status | null>(
        null
    )

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

                {selectedStatus?.value === "Personalized" && (
                    <div className="p-4 border border-gray-300 rounded-md">
                        <h4 className="text-lg font-bold mb-2">Personalized Settings</h4>
                        <MenuProvider initialFields={ []}> {/*Add the fields if they are set*/}
                            {menu}
                        </MenuProvider>
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