"use client"

import React from 'react'
import Image from "next/image";

import { SideMenu } from '@/app/home/components/sidemenu'
import { useMutation, useQuery } from '@tanstack/react-query';
import { Workspaces } from '@/app/home/lib/types';
import { getWorkspaces } from '@/app/home/lib/get-workspaces';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { ScrollArea } from '@/components/ui/scroll-area';
import { LogoIcon } from '@/components/ui/logoIcon';
import { Button } from '@/components/ui/button';
import { getTemplates } from '@/app/templates/lib/get-templates';
import { TemplatesArrayType, TemplatesType } from '@/app/templates/lib/types';
import { newTrigger } from '@/app/home/lib/new-trigger';

export default function Page() {
    const [workspaces, setWorkspaces] = React.useState<Workspaces>([])
    const [templates, setTemplates] = React.useState<TemplatesArrayType>([])

    const { data, isPending, error } = useQuery({
        queryKey: ["templates"],
        queryFn: async () => {
            const [workspace, templates] = await Promise.all([
                getWorkspaces(),
                getTemplates(),
            ]);
            return { workspace, templates };
        },
    });

    React.useEffect(() => {
        if (data?.workspace) {
            setWorkspaces(data.workspace);
        }
        if (data?.templates)
            setTemplates(data.templates)
    }, [data]);

    const mutation = useMutation({
        mutationFn: newTrigger,
        onSuccess: (trigger) => {
          window.location.href = `/trigger/${trigger.id}`;
        },
      });

    if (isPending) return <div>Loading...</div>
    if (error) return <div>Could not get templates.</div>

    const handleCardClick = (template: TemplatesType) => {
        mutation.mutate(template);
    };

    return (
        <div className="flex h-screen w-full overflow-hidden">
            <SideMenu workspaceLen={workspaces.length} />
            <div className="flex flex-col w-full p-5 overflow-x-auto">
                <Card className="py-6 h-full">
                    <CardContent>
                        <ScrollArea>
                            <div className="px-8 py-4 w-full">
                                <Card className="w-full h-[200px] bg-gradient-to-tr from-blue-500 via-violet-500 to-fuchsia-500 p-5 rounded-lg shadow-lg animate-gradient">
                                    <CardContent className="flex flex-col h-full items-center justify-center text-white text-center font-bold text-lg lg:text-3xl gap-y-3">
                                        <LogoIcon className="w-[150px] fill-yellow-500" />
                                        <p>Try Trigger for 30 days free</p>
                                        <Button className="bg-zinc-200 text-black p-5 hover:bg-zinc-100">
                                            Start free trial
                                        </Button>
                                    </CardContent>
                                </Card>
                            </div>
                            <p className="text-3xl font-bold p-5">Templates</p>
                            {templates.map((t, index) => (
                                <div key={index}>
                                    <Card
                                        className="flex flex-col bg-zinc-100 shadow-md rounded-lg w-[200px]"
                                        key={index}
                                        onClick={() => handleCardClick(t)}
                                    >
                                        <CardHeader className="p-4 border-b">
                                            <CardTitle className="text-xl font-bold">
                                                <Image
                                                    src="https://fakeimg.pl/300x200"
                                                    alt={t.name}
                                                    width={500}
                                                    height={500}
                                                    layout="responsive"
                                                />
                                            </CardTitle>
                                        </CardHeader>
                                    </Card>
                                    <p className="font-bold text-md text-start p-1">
                                        {t.name}
                                    </p>
                                </div>
                            ))}

                        </ScrollArea>
                    </CardContent>
                </Card>
            </div>
        </div>
    )
}
