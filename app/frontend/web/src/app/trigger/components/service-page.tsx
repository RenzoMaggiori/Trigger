"use client";
import React from "react";
import { Card, CardContent } from "@/components/ui/card";
import { TriggerDraggable } from "@/app/trigger/components/trigger-draggable";
import { NodeItem, Service } from "@/app/trigger/lib/types";
import { Button } from "@/components/ui/button";
import { Loader2 } from "lucide-react";
import { useMenu } from "@/app/trigger/components/MenuProvider";
import { useMutation } from "@tanstack/react-query";
import { send_workspace } from "@/app/trigger/lib/send-workspace";
import { toast } from "sonner";
import { deployWorkspace } from "../lib/deploy-workspace";
import { Input } from "@/components/ui/input";

interface ServicesProps {
  services: Service[];
  handleDragStart: (
    e: React.DragEvent<HTMLDivElement>,
    service: Service,
  ) => void;
}

export const ServicesComponent: React.FC<ServicesProps> = ({
  services,
  handleDragStart,
}) => {
  const [loading, setLoading] = React.useState<boolean>(false);
  const [loadingDeploy, setLoadingDeploy] = React.useState<boolean>(false);
  const { triggerWorkspace, setTriggerWorkspace } = useMenu();

  const action = {
    description: new Date().toLocaleString("en-US", {
      weekday: "long",
      year: "numeric",
      month: "long",
      day: "numeric",
      hour: "numeric",
      minute: "numeric",
      hour12: true,
    }),
    action: {
      label: "Undo",
      onClick: () => console.log("Undo"),
    },
  };

  const mutation = useMutation({
    mutationFn: send_workspace,
    onSuccess: (data) => {
      const nodes: Record<string, NodeItem> = {};
      for (const n of data.nodes) {
        nodes[n.node_id] = {
          id: n.node_id,
          action_id: n.action_id,
          fields: n.input || {},
          parent_ids: n.parents,
          child_ids: n.children,
          status: n.status,
          x_pos: n.x_pos,
          y_pos: n.y_pos,
        };
      }
      setTriggerWorkspace({ id: data.id, name: data.name, nodes });
      setLoading(false);
      toast("Workspace saved successfully", action);
    },
    onError: () => {
      setLoading(false);
      toast("Error while saving the workspace", action);
    },
  });

  const deployWorkspaceMutation = useMutation({
    mutationFn: deployWorkspace,
    onSuccess: () => {
      setLoadingDeploy(false);
      window.location.href = "/home";
    },
    onError: () => {
      setLoading(false);
      toast("Error while deploying the workspace", action);
    },
  });

  const handleOnClick = () => {
    if (!triggerWorkspace) return;
    mutation.mutate(triggerWorkspace);
  };

  const handleDeploy = () => {
    if (!triggerWorkspace) return;
    deployWorkspaceMutation.mutate({ id: triggerWorkspace.id });
  };

  const handleNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (!triggerWorkspace) return
    setTriggerWorkspace({ ...triggerWorkspace, name: e.target.value });
  };

  return (
    <div className="flex flex-col w-auto p-5 gap-y-2">
      <Card className="flex items-center justify-center">
          <Input placeholder="workspace name" value={triggerWorkspace?.name} onChange={handleNameChange}/>
      </Card>
      <Card className="h-full overflow-hidden">
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
            className="w-full text-md rounded-full border-black py-2"
            variant="outline"
            onClick={() => {
              setLoading(true);
              handleOnClick();
            }}
            disabled={loading}
          >
            {loading && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
            Save Trigger
          </Button>
          <Button
            className="w-full text-md rounded-full bg-gradient-to-r from-blue-500 via-violet-500 to-fuchsia-500 hover:bg-gradient-to-r hover:from-blue-600 hover:via-violet-600 hover:to-fuchsia-600 animate-gradient text-white py-2"
            onClick={() => {
              setLoadingDeploy(true);
              handleOnClick();
              handleDeploy();
            }}
            disabled={loadingDeploy}
          >
            {loadingDeploy && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
            Deploy Workspace
          </Button>
        </CardContent>
      </Card>
    </div>
  );
};
