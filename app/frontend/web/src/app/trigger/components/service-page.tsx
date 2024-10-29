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
  const {triggerWorkspace, setTriggerWorkspace} = useMenu();

  const mutation = useMutation({
    mutationFn: send_workspace,
    onSuccess: (data) => {
      const nodes: Record<string, NodeItem> = {};
      for (const n of data.nodes) {
        nodes[n.node_id] = {
          id: n.node_id,
          type: n.action_id,
          fields: n.input || {},
          parent_ids: n.parents,
          child_ids: n.children,
          x_pos: n.x_pos,
          y_pos: n.y_pos,
        };
      }
      setTriggerWorkspace({ id: data.id, nodes });
      setLoading(false);
    },
  });

  const handleOnClick = () => {
    if (!triggerWorkspace) return;
    mutation.mutate(triggerWorkspace);
  };
  return (
    <div className="w-auto p-5">
      <Card className="h-full">
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
            className="w-full text-md rounded-full bg-gradient-to-r from-blue-500 via-violet-500 to-fuchsia-500 hover:bg-gradient-to-r hover:from-blue-600 hover:via-violet-600 hover:to-fuchsia-600 animate-gradient text-white"
            onClick={() => {
              handleOnClick();
              setLoading(true);
            }}
            disabled={loading}
          >
            {loading && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
            Deploy Trigger
          </Button>
        </CardContent>
      </Card>
    </div>
  );
};
