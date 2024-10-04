import React, { createContext, useContext, useState, ReactNode } from 'react';

interface MenuContextType {
    fields: any[];
    setFields: React.Dispatch<React.SetStateAction<any[]>>;
    updateFields: (newFields: any[]) => void;
}

const MenuContext = createContext<MenuContextType | undefined>(undefined);

export const useMenu = () => {
    const context = useContext(MenuContext);
    if (!context) {
        throw new Error('useMenu must be used within a MenuProvider');
    }
    return context;
};

export const MenuProvider = ({ children, initialFields }: { children: ReactNode; initialFields: any[] }) => {
    const [fields, setFields] = useState<any[]>(initialFields);


    const updateFields = (newFields: any[]) => {
        setFields(newFields);
    };

    return (
        <MenuContext.Provider value={{ fields, setFields, updateFields }}>
            {children}
        </MenuContext.Provider>
    );
};
