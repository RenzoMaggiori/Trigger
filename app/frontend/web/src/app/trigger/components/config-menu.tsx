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
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";

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

const configMap = [
  { name: "Trigger", type: "trigger" },
  { name: "Reaction", type: "reaction" },
];

export function ConfigMenu({ menu, parentNodes, node }: ConfigMenuType) {
  const { triggerWorkspace, setFields } = useMenu();
  const nodeItem = triggerWorkspace?.nodes[node?.id || ""];
  const [configState, setConfigState] = React.useState<Record<string, unknown>>(
    {
      trigger: nodeItem?.fields.triggerStatus || "None",
      reaction: nodeItem?.fields.reactionStatus || "None",
    },
  );

  if (!node) return <div>{"custom node doesn't exist"}</div>;
  if (!nodeItem) return <div>could not find node</div>;

  const handleStatusChange = (
    status: Status | null,
    configType: "trigger" | "reaction",
  ) => {
    const newStatus = status?.value || "None";
    setConfigState((prev) => ({
      ...prev,
      [configType]: newStatus,
    }));
    setFields(node.id, {
      ...nodeItem.fields,
      [`${configType}Status`]: newStatus,
    });
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
        <Tabs defaultValue="trigger">
          <TabsList className="mb-4 grid grid-cols-2">
            <TabsTrigger value="trigger">Trigger Settings</TabsTrigger>
            <TabsTrigger value="reaction">Reaction Settings</TabsTrigger>
          </TabsList>
          {configMap.map((item, key) => (
            <TabsContent key={key} value={item.type}>
              <div className="mb-4">
                <Label
                  htmlFor={`${item.type}-dropdown`}
                  className="block text-sm font-medium text-gray-700"
                >
                  {item.name} Configuration
                </Label>
                <Combox
                  statuses={combinedStatuses}
                  setSelectedStatus={(status) =>
                    handleStatusChange(
                      status,
                      item.type === "trigger" ? "trigger" : "reaction",
                    )
                  }
                  selectedStatus={
                    combinedStatuses.find(
                      (status) => status.value === configState[item.type],
                    ) || null
                  }
                  label="info"
                  icon={<SiGooglegemini className="mr-2" />}
                />
              </div>

              {configState[item.type] === "Personalized" && (
                <div className="p-4 border border-gray-300 rounded-md">
                  <h4 className="text-lg font-bold mb-2">
                    Personalized {item.name} Settings
                  </h4>
                  <SettingsComponent node={nodeItem} type={item.type} />
                </div>
              )}

              {configState[item.type] !== "Personalized" &&
                configState[item.type] !== "None" && (
                  <div className="mt-4">
                    <h4 className="font-bold">Selected Parent Node ID:</h4>
                    <p>{configState[item.type] as string}</p>
                    <h4 className="font-bold">Parent Node Label:</h4>
                    <p>
                      {
                        parentNodes.find(
                          (node) => node.id === configState[item.type],
                        )?.data.label
                      }
                    </p>
                  </div>
                )}
            </TabsContent>
          ))}
        </Tabs>
      </CardContent>
    </Card>
  );
}
