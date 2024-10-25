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

const configOptions = [
  {
    label: (
      <div className="flex flex-row items-center text-md font-bold">
        <SiGooglegemini className="mr-2 fill-blue-500" /> Trigger
      </div>
    ),
    value: "trigger",
  },
  {
    label: (
      <div className="flex flex-row items-center text-md font-bold">
        <SiGooglegemini className="mr-2 fill-green-500" /> Reaction
      </div>
    ),
    value: "reaction",
  },
];

export function ConfigMenu({ menu, parentNodes, node }: ConfigMenuType) {
  const { triggerWorkspace } = useMenu();
  const nodeItem = triggerWorkspace?.nodes[node?.id || ""];
  const [configType, setConfigType] = React.useState("trigger");
  const [configState, setConfigState] = React.useState<Record<string, unknown>>(
    () => ({
      [configType]: "Personalized",
    })
  );

  if (!node) return <div>{"custom node doesn't exist"}</div>;
  if (!nodeItem) return <div>could not find node</div>;

  const handleStatusChange = (
    status: Status | null,
    configType: "trigger" | "reaction",
  ) => {
    const newStatus = status?.value || "Personalized";
    setConfigState((prev) => ({
      ...prev,
      [configType]: newStatus,
    }));
  };

  const handleConfigTypeChange = (selectedConfigType: "trigger" | "reaction") => {
    setConfigType(selectedConfigType);
    setConfigState({
      [selectedConfigType]: configState[selectedConfigType] || "Personalized",
    });
  };

  const combinedStatuses: Status[] = [
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
  console.log(nodeItem);
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
        <div className="flex justify-between items-center">
          <div className="mb-4">
            <Label
              htmlFor="configType-dropdown"
              className="block text-sm font-medium text-gray-700"
            >
              Choose Configuration Type
            </Label>
            <Combox
              statuses={configOptions}
              setSelectedStatus={(selected) => handleConfigTypeChange(selected?.value === "trigger" ? "trigger" : "reaction")}
              selectedStatus={configOptions.find((option) => option.value === configType) || null}
              label="info"
              icon={<SiGooglegemini className="mr-2" />}
            />
          </div>

          <div className="mb-4">
            <Label
              htmlFor={`${configType}-dropdown`}
              className="block text-sm font-medium text-gray-700"
            >
              {configType === "trigger" ? "Trigger" : "Reaction"} Configuration
            </Label>
            <Combox
              statuses={combinedStatuses}
              setSelectedStatus={(status) =>
                handleStatusChange(status, configType as "trigger" | "reaction")
              }
              selectedStatus={
                combinedStatuses.find(
                  (status) => status.value === configState[configType],
                ) || combinedStatuses[0]
              }
              label="info"
              icon={<SiGooglegemini className="mr-2" />}
            />
          </div>
        </div>

        {configState[configType] === "Personalized" && (
          <div className="p-4 border border-gray-300 rounded-md">
            <h4 className="text-lg font-bold mb-2">
              Personalized {configType === "trigger" ? "Trigger" : "Reaction"} Settings
            </h4>
            <SettingsComponent key={configType} node={nodeItem} type={configType} />
          </div>
        )}

        {configState[configType] !== "Personalized" && (
          <div className="mt-4">
            <h4 className="font-bold">Selected Parent Node ID:</h4>
            <p>{configState[configType] as string}</p>
            <h4 className="font-bold">Parent Node Label:</h4>
            <p>
              {
                parentNodes.find(
                  (node) => node.id === configState[configType],
                )?.data.label
              }
            </p>
          </div>
        )}
      </CardContent>
    </Card>
  );
}
