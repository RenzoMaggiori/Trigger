import React, { useCallback, useEffect, useRef, useState } from 'react';
import { View, Text, StyleSheet, TouchableOpacity, ScrollView } from 'react-native';
import Animated, { FadeIn, FadeOut, Layout, useSharedValue } from 'react-native-reanimated';
import { AntDesign, Ionicons, MaterialIcons } from '@expo/vector-icons';
import { Colors } from '@/constants/Colors';
import TechBox from './TechBox';
import { GestureHandlerRootView } from 'react-native-gesture-handler';
import Draggable from './Draggable';

interface TechProps {
    name: string;
    icon: React.ReactElement;
}

interface TechBox extends TechProps {
    id: number;
}

interface ActionReactionListProps {
    tech?: TechProps;
}

export default function ActionReactionList({ tech }: ActionReactionListProps) {
    const initialState = useRef<boolean>(true);
    const [techs, setTechs] = useState<TechBox[]>([]);

    useEffect(() => {
        initialState.current = false;
        if (tech) {
            setTechs(prevTechs => [...prevTechs, { ...tech, id: prevTechs.length }]);
        }
    }, [tech]);

    const onDelete = useCallback((id: number) => {
        setTechs((currTechs: TechBox[]) => {
            let newItems = currTechs.filter((item: any) => item.id !== id);
            newItems.map((tech) => {
                tech.id = newItems.indexOf(tech);
            });
            return newItems;
        });
    }, []);

    return (
        <GestureHandlerRootView style={{height: "100%", minHeight: 450}}>
            {techs.length !== 0 ? (
            <View>
                {techs.map((tech, index) => (
                <Draggable key={index}>
                    <TechBox
                    index={index}
                    id={tech.id}
                    name={tech.name}
                    icon={tech.icon}
                    initialState={initialState}
                    onDelete={onDelete}
                    />
                </Draggable>
                ))}
                <TouchableOpacity style={styles.save}>
                    <MaterialIcons name="save" size={24} color="black" />
                </TouchableOpacity>
            </View>
            ) : (
                <View>
                    <Text>No techs available</Text>
                </View>
            )}
        </GestureHandlerRootView>
    );
}

const styles = StyleSheet.create({
    listItem: {
        height: 100,
        width: '90%',
        backgroundColor: 'white',
        justifyContent: 'center',
        alignItems: 'center',
        marginVertical: 10,
        alignSelf: 'center',
        borderRadius: 10,
        elevation: 5,
    },
    save: {
        height: 50,
        width: '50%',
        backgroundColor: '#73c6b6',
        justifyContent: 'center',
        alignItems: 'center',
        alignSelf: 'center',
        borderRadius: 10,
        elevation: 5,
        marginTop: 10,
    },
    deleteButton: {
        position: 'absolute',
        top: 10,
        left: 10,
        zIndex: 1,
    },
    iconContainer: {
        marginBottom: 5,
    },
    techName: {
        fontSize: 16,
        fontWeight: 'bold',
    },
});
