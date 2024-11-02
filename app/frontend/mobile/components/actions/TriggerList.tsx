import React, { useCallback, useEffect, useRef, useState } from 'react';
import { View, Text, StyleSheet, TouchableOpacity, ScrollView } from 'react-native';
import Animated, { FadeIn, FadeOut, Layout, useSharedValue } from 'react-native-reanimated';
import { AntDesign, Ionicons, MaterialIcons } from '@expo/vector-icons';
import { Colors } from '@/constants/Colors';
import ProviderBox from '@/components/actions/ProviderBox';
import { GestureHandlerRootView } from 'react-native-gesture-handler';
import Draggable from './Draggable';

interface ProviderProps {
    name: string;
    icon: React.ReactElement;
    provider: string;
    action: string;
}

interface ProviderBox extends ProviderProps {
    id: number;
}

interface TriggerListProps {
    provider?: ProviderProps;
}

export default function TriggerList({ provider }: TriggerListProps) {
    const initialState = useRef<boolean>(true);
    const [providers, setProviders] = useState<ProviderBox[]>([]);

    useEffect(() => {
        initialState.current = false;
        if (provider) {
            setProviders(prevProviders => [...prevProviders, { ...provider, id: prevProviders.length }]);
        }
    }, [provider]);

    const onDelete = useCallback((id: number) => {
        setProviders((currProviders: ProviderBox[]) => {
            let newItems = currProviders.filter((item: any) => item.id !== id);
            newItems.map((provider) => {
                provider.id = newItems.indexOf(provider);
            });
            return newItems;
        });
    }, []);

    return (
        <GestureHandlerRootView style={{height: "100%", minHeight: 450}}>
            {providers.length !== 0 ? (
            <View>
                {providers.map((provider, index) => (
                <Draggable key={index}>
                    <ProviderBox
                    index={index}
                    id={provider.id}
                    name={provider.name}
                    icon={provider.icon}
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
                    <Text>No providers available</Text>
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
    providerName: {
        fontSize: 16,
        fontWeight: 'bold',
    },
});
