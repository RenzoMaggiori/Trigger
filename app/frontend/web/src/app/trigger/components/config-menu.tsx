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
import { ActionType, CustomNode } from "@/app/trigger/lib/types";
import {
  BitBucketSettings,
  DiscordSettings,
  EmailSettings,
  GithubSettings,
  SpotifySettings,
  TimerSettings,
  TwitchSettings,
} from "@/app/trigger/components/service-settings";
import { Badge } from "@/components/ui/badge";
import { cn } from "@/lib/utils";

export type ConfigMenuType = {
  menu: keyof typeof settingsComponentMap;
  parentNodes: CustomNode[];
  node: CustomNode | null;
  actions: ActionType;
};

const settingsComponentMap = {
  email: EmailSettings,
  discord: DiscordSettings,
  github: GithubSettings,
  spotify: SpotifySettings,
  twitch: TwitchSettings,
  timer: TimerSettings,
  bitbucket: BitBucketSettings,
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

export function ConfigMenu({
  menu,
  parentNodes,
  node,
  actions,
}: ConfigMenuType) {
  const { triggerWorkspace } = useMenu();
  const nodeItem = triggerWorkspace?.nodes[node?.id || ""];


  const nodeConfig = React.useRef(
    new Map<string, { configType: "trigger" | "reaction"; configState: Record<string, unknown> }>()
  );

  const [currentConfig, setCurrentConfig] = React.useState<{
    configType: "trigger" | "reaction";
    configState: Record<string, unknown>;
  }>(() => ({
    configType: "trigger",
    configState: { trigger: "Personalized" },
  }));

  React.useEffect(() => {
    if (node && nodeItem) {
      const initialConfigType: "trigger" | "reaction" =
        nodeItem.fields["type"] === "reaction" ? "reaction" : "trigger";

      if (!nodeConfig.current.has(node.id)) {
        nodeConfig.current.set(node.id, {
          configType: initialConfigType,
          configState: { [initialConfigType]: "Personalized" },
        });
      }

      setCurrentConfig(nodeConfig.current.get(node.id)!);
    }
  }, [node, nodeItem]);

  if (!node) return <div>{"custom node doesn't exist"}</div>;
  if (!nodeItem) return <div>could not find node</div>;

  const handleStatusChange = (
    status: Status | null,
    selectedConfigType: "trigger" | "reaction",
  ) => {
    const newStatus = status?.value || "Personalized";
    const updatedConfig = {
      ...currentConfig,
      configState: {
        ...currentConfig.configState,
        [selectedConfigType]: newStatus,
      },
    };
    nodeConfig.current.set(node.id, updatedConfig);
    setCurrentConfig(updatedConfig);
  };

  const handleConfigTypeChange = (selectedConfigType: "trigger" | "reaction") => {
    const updatedConfig = {
      ...currentConfig,
      configType: selectedConfigType,
      configState: {
        ...currentConfig.configState,
        [selectedConfigType]: currentConfig.configState[selectedConfigType] || "Personalized",
      },
    };
    nodeConfig.current.set(node.id, updatedConfig);
    setCurrentConfig(updatedConfig);
  };

  const { configType, configState } = currentConfig;

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

  const nodeStatus: Record<string, string> = {
    "completed": "bg-green-500 hover:bg-green-600",
    "active": "bg-violet-500 hover:bg-violet-600",
    "inactive": "bg-white border-black text-black hover:bg-zinc-200",
  }

  const SettingsComponent = settingsComponentMap[menu];
  return (
    <Card className="h-full w-[500px]">
      <CardHeader>
        <CardTitle className="flex items-center justify-between text-xl font-bold">
          <div className="flex flex-row items-center text-center">
            {" "}
            {node?.data?.label} Settings
          </div>
          <Badge
            className={cn(
              "rounded-full text-sm",
              nodeStatus[nodeItem.status]
            )}
          >
            {nodeItem.status === "active" ? "in progress" : nodeItem.status}
          </Badge>
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
              setSelectedStatus={(selected) =>
                handleConfigTypeChange(
                  selected?.value === "trigger" ? "trigger" : "reaction",
                )
              }
              selectedStatus={
                configOptions.find((option) => option.value === configType) ||
                null
              }
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
              Personalized {configType === "trigger" ? "Trigger" : "Reaction"}{" "}
              Settings
            </h4>
            <SettingsComponent
              key={configType}
              node={nodeItem}
              type={configType}
              actions={actions}
            />
          </div>
        )}

        {configState[configType] !== "Personalized" && (
          <div className="mt-4">
            <h4 className="font-bold">Selected Parent Node ID:</h4>
            <p>{configState[configType] as string}</p>
            <h4 className="font-bold">Parent Node Label:</h4>
            <p>
              {
                parentNodes.find((node) => node.id === configState[configType])
                  ?.data.label
              }
            </p>
          </div>
        )}
      </CardContent>
    </Card>
  );
}
