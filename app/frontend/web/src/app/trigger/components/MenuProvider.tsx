"use client"
import React from 'react';
import { NodesArrayItem, TriggerWorkspace } from '@/app/trigger/lib/types';

type MenuContextType = {
    triggerWorkspace: TriggerWorkspace | null
    setNodes: (nodes: NodesArrayItem[]) => void
    setFields: (nodeID: NodesArrayItem["id"], fields: Record<string, any>) => void
    setTriggerWorkspace: React.Dispatch<React.SetStateAction<TriggerWorkspace | null>>
}

type MenuProviderType = {
    children: React.ReactNode
    initialNodes?: Record<NodesArrayItem["id"], NodesArrayItem>
    initialWorkspace?: TriggerWorkspace | null;
}

const MenuContext = React.createContext<MenuContextType | undefined>(undefined);

export const useMenu = () => {
    const context = React.useContext(MenuContext);
    if (!context) {
        throw new Error('useMenu must be used within a MenuProvider');
    }
    return context;
};

async function getWorkspaceId() {
    return Math.random() * 100
}

export function MenuProvider({ children, initialNodes = {}, initialWorkspace = null }: MenuProviderType) {
    const [triggerWorkspace, setTriggerWorkspace] = React.useState<TriggerWorkspace | null>(initialWorkspace);

    const setNodes = (nodes: NodesArrayItem[]) => {
        setTriggerWorkspace((prev) => {
            if (!prev)
                return prev
            return { ...prev, nodes }
        })
    }

    const setFields = (nodeID: NodesArrayItem["id"], fields: Record<string, any>) => {
        if (!triggerWorkspace)
            return
        const newNode = triggerWorkspace.nodes.filter((n) => n.id === nodeID)
        if (newNode.length != 1)
            return
        setNodes([...triggerWorkspace.nodes.filter((n) => n.id !== nodeID), { ...newNode[0], fields }])
    }


    return (
        <MenuContext.Provider value={{ triggerWorkspace, setNodes, setFields, setTriggerWorkspace }}>
            {children}
        </MenuContext.Provider>
    );
};
