"use client";
import React from "react";
import { NodeItem, TriggerWorkspace } from "@/app/trigger/lib/types";

type MenuContextType = {
  triggerWorkspace: TriggerWorkspace | null;
  setNodes: (nodes: Record<string, NodeItem>) => void;
  setFields: (nodeID: NodeItem["id"], fields: Record<string, unknown>) => void;
  setTriggerWorkspace: React.Dispatch<
    React.SetStateAction<TriggerWorkspace | null>
  >;
};

type MenuProviderType = {
  children: React.ReactNode;
  initialWorkspace?: TriggerWorkspace | null;
};

const MenuContext = React.createContext<MenuContextType | undefined>(undefined);

export const useMenu = () => {
  const context = React.useContext(MenuContext);
  if (!context) {
    throw new Error("useMenu must be used within a MenuProvider");
  }
  return context;
};

export function MenuProvider({
  children,
  initialWorkspace = null,
}: MenuProviderType) {
  const [triggerWorkspace, setTriggerWorkspace] =
    React.useState<TriggerWorkspace | null>(initialWorkspace);

  const setNodes = (nodes: Record<string, NodeItem>) => {
    setTriggerWorkspace((prev) => {
      if (!prev) return prev;

      const updates = { ...prev.nodes };
      Object.entries(nodes).forEach(([id, newNode]) => {
        const existingNode = updates[id];

        if (existingNode) {
          updates[id] = {
            ...existingNode,
            ...newNode,
            fields: {
              ...existingNode.fields,
              ...newNode.fields,
            },
          };
        } else {
          updates[id] = newNode;
        }
      });
      return {
        ...prev,
        nodes: updates,
      };
    });
  };

  const updateNodes = (nodes: Record<string, NodeItem>) => {
    setTriggerWorkspace((prev) => {
      if (!prev) return prev;

      const updates = { ...prev.nodes };
      Object.entries(nodes).forEach(([id, newNode]) => {
        const existingNode = updates[id];

        if (existingNode) {
          updates[id] = {
            ...existingNode,
            ...newNode,
            fields: {
              ...newNode.fields,
            },
          };
        } else {
          updates[id] = newNode;
        }
      });
      return {
        ...prev,
        nodes: updates,
      };
    });
  };

  const setFields = (
    nodeID: NodeItem["id"],
    fields: Record<string, unknown>,
  ) => {
    if (!triggerWorkspace) return;

    const node = triggerWorkspace.nodes[nodeID];
    if (!node) return;
    const updates: NodeItem = {
      ...node,
      fields: {
        ...fields,
      },
    };
    updateNodes({ [nodeID]: updates });
  };

  return (
    <MenuContext.Provider
      value={{ triggerWorkspace, setNodes, setFields, setTriggerWorkspace }}
    >
      {children}
    </MenuContext.Provider>
  );
}
