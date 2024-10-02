import React from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { ServiceSetting } from "./trigger-draggable";
import { Textarea } from "@/components/ui/textarea";
import { Label } from "@/components/ui/label";

interface ConfigMenuProps {
    menu: ServiceSetting[];
}

const InputComponent = ({ label }: { label: string }) => (
    <div>
        <label className="block font-medium mb-1">{label}</label>
        <input type="text" placeholder={label} className="border p-2 w-full rounded-md" />
    </div>
);

const TextareaComponent = ({ label, options }: { label: string, options: string[] }) => (
    <div>
        <Label className="font-lg">{options[0]}</Label>
        <Textarea placeholder={label} className="rezise-none"/>
    </div>
);

const DropdownComponent = ({ label, options }: { label: string; options: string[] }) => (
    <div>
        <label className="block font-medium mb-1">{label}</label>
        <select className="border p-2 w-full rounded-md">
            {options.map((option, idx) => (
                <option key={idx} value={option}>
                    {option}
                </option>
            ))}
        </select>
    </div>
);

const CheckboxComponent = ({ label }: { label: string }) => (
    <div>
        <label className="flex items-center space-x-2">
            <input type="checkbox" className="form-checkbox" />
            <span>{label}</span>
        </label>
    </div>
);

const ButtonComponent = ({ label }: { label: string }) => (
    <button className="bg-blue-500 text-white p-2 rounded-md w-full hover:bg-blue-600">{label}</button>
);

const componentMap: Record<string, React.FC<any>> = {
    input: InputComponent,
    textarea: TextareaComponent,
    dropdown: DropdownComponent,
    button: ButtonComponent,
    checkbox: CheckboxComponent,
};

export function ConfigMenu({ menu }: ConfigMenuProps) {
    return (
        <div className="flex flex-col gap-4">
            {menu.map((setting, index) => {

                const Component = componentMap[setting.type];
                return Component ? (
                    <Card key={index} className="w-full">
                        <CardHeader>
                            <CardTitle className="text-xl">Settings</CardTitle>
                        </CardHeader>
                        <CardContent>
                            <Component label={setting.label} options={setting.options} />
                        </CardContent>
                    </Card>
                ) : (

                    <Card key={index} className="w-full">
                        <CardContent>
                            <div className="text-gray-500">Unknown setting type: {setting.type}</div>
                        </CardContent>
                    </Card>
                );
            })}
        </div>
    );
}
