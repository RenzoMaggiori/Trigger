import React from "react";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Label } from "@/components/ui/label";
import { Combox, Status } from "@/components/ui/combox";
import { SiGooglegemini } from "react-icons/si";
import { useMenu } from "@/app/trigger/components/MenuProvider";
import { CustomNode } from "@/app/trigger/lib/types";
import {
  DiscordSettings,
  EmailSettings,
  GithubSettings,
} from "@/app/trigger/components/service-settings";

export type ConfigMenuType = {
  menu: keyof typeof settingsComponentMap;
  parentNodes: CustomNode[];
  node: CustomNode | null;
};

const settingsComponentMap = {
  email: EmailSettings,
  discord: DiscordSettings,
  github: GithubSettings,
};

export function ConfigMenu({ menu, parentNodes, node }: ConfigMenuType) {
  const { triggerWorkspace, setFields } = useMenu();

  if (!node) return <div>custom node doesn't exist</div>;

  const nodeItem = triggerWorkspace?.nodes[node.id];
  if (!nodeItem) return <div>could not find node</div>;

  const nodeStatus = nodeItem.fields.selectedStatus || "None";

  const handleStatusChange = (status: Status | null) => {
    setFields(node.id, { selectedStatus: status?.value || "None" });
  };

  const combinedStatuses: Status[] = [
    {
      label: (
        <div className="flex flex-row items-center text-md font-bold">
          <SiGooglegemini className="mr-2" /> None
        </div>
      ),
      value: "None",
    },
    {
      label: (
        <div className="flex flex-row items-center text-md font-bold">
          <SiGooglegemini className="mr-2 fill-purple-500" /> Personalized
        </div>
      ),
      value: "Personalized",
    },
    ...parentNodes.map((parentNode) => ({
      label: parentNode.data.label as string,
      value: parentNode.id,
    })),
  ];

  const SettingsComponent = settingsComponentMap[menu];

  return (
    <Card className="h-full w-[500px]">
      <CardHeader>
        <CardTitle className="flex items-center text-xl font-bold">
          {node?.data?.label} Settings
        </CardTitle>
        <CardDescription className="ml-2 text-md">
          ID: {node?.id}
        </CardDescription>
      </CardHeader>
      <CardContent>
        <div className="mb-4">
          <Label
            htmlFor="parent-node-dropdown"
            className="block text-sm font-medium text-gray-700"
          >
            Information to Send
          </Label>
          <Combox
            statuses={combinedStatuses}
            setSelectedStatus={handleStatusChange}
            selectedStatus={combinedStatuses.find((status) => status.value === nodeStatus) || null}
            label="info"
            icon={<SiGooglegemini className="mr-2" />}
          />
        </div>

        {nodeStatus === "Personalized" && (
          <div className="p-4 border border-gray-300 rounded-md">
            <h4 className="text-lg font-bold mb-2">Personalized Settings</h4>
            <SettingsComponent node={nodeItem} />
          </div>
        )}

        {nodeStatus !== "Personalized" && nodeStatus !== "None" && (
          <div className="mt-4">
            <h4 className="font-bold">Selected Parent Node ID:</h4>
            <p>{nodeStatus}</p>
            <h4 className="font-bold">Parent Node Label:</h4>
            <p>
              {
                parentNodes.find((node) => node.id === nodeStatus)?.data
                  .label
              }
            </p>
          </div>
        )}
      </CardContent>
    </Card>
  );
}

