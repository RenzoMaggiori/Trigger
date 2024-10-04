import React, { useState } from 'react';
import { View, Text, StyleSheet, ScrollView, TouchableOpacity } from 'react-native';
import { Colors } from '@/constants/Colors';
// @ts-ignore
import { Ionicons, FontAwesome5 } from '@expo/vector-icons';
import { GestureHandlerRootView } from 'react-native-gesture-handler';
import TechBox from '@/components/actions/TechBox';
import Draggable from '@/components/actions/Draggable';

interface ActionBox {
    name: string;
    icon: React.ReactElement;
    id: number;
}

const technologies: ActionBox[] = [
    { id: 0, name: 'Google', icon: <Ionicons name="logo-google" size={30} color={Colors.light.google} /> },
    { id: 0, name: 'Github', icon: <Ionicons name="logo-github" size={30} color={Colors.light.github} /> },
    { id: 0, name: 'Outlook', icon: <Ionicons name="logo-microsoft" size={30} color={Colors.light.outlook} /> },
    { id: 0, name: 'Slack', icon: <FontAwesome5 name="slack" size={30} color={Colors.light.slack} /> },
    { id: 0, name: 'Discord', icon: <FontAwesome5 name="discord" size={30} color={Colors.light.discord} /> },
];

export default function TriggersScreen() {
    const [selectedComponents, setSelectedComponents] = useState<ActionBox[]>([]);

    const handlePress = (techName: string) => {
        setSelectedComponents(prev => {
            const foundTech = technologies.find(tech => tech.name === techName);
            if (foundTech) {
                foundTech.id = Math.random();
                return [...prev, foundTech];
            }
            return prev;
        });
    };

    const handleDelete = (techId: number) => {
        setSelectedComponents(prev => prev.filter(tech => tech.id !== techId));
    };

    return (
        <View style={styles.container}>
            <View style={styles.servicesContainer}>
                <Text style={styles.servicesTitle}>Services</Text>
                <ScrollView horizontal showsHorizontalScrollIndicator={false} contentContainerStyle={styles.servicesList}>
                    {technologies.map((tech, index) => (
                        <TouchableOpacity key={index} style={styles.techItem} onPress={() => handlePress(tech.name)}>
                            <View style={styles.iconContainer}>{tech.icon}</View>
                            <Text style={styles.techName}>{tech.name}</Text>
                        </TouchableOpacity>
                    ))}
                </ScrollView>
            </View>
            <GestureHandlerRootView style={styles.manageActions}>
                {selectedComponents.map((component, index) => (
                    <Draggable key={`${component.name}-${index}-`}>
                        <TechBox
                            name={component.name}
                            icon={component.icon}
                            onDelete={() => handleDelete(component.id)}
                        />
                    </Draggable>
                ))}
            </GestureHandlerRootView>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
    },
    manageActions: {
        flex: 1,
        marginTop: 10,
        padding: 20,
        borderRadius: 10,
        marginHorizontal: 10,
        backgroundColor: '#ffdab9',
    },
    servicesContainer: {
        marginTop: 10,
        padding: 20,
        backgroundColor: Colors.light.tintLight,
        borderRadius: 10,
        marginHorizontal: 10,
    },
    servicesTitle: {
        fontSize: 18,
        fontWeight: 'bold',
        marginBottom: 10,
    },
    servicesList: {
        flexDirection: 'row',
    },
    techItem: {
        marginRight: 20,
        alignItems: 'center',
        justifyContent: 'center',
        borderWidth: 1,
        borderColor: '#ccc',
        borderRadius: 10,
        padding: 10,
        backgroundColor: '#FFFFFF',
    },
    iconContainer: {
        marginBottom: 5,
    },
    techName: {
        fontSize: 16,
        fontWeight: 'bold',
    },
});
