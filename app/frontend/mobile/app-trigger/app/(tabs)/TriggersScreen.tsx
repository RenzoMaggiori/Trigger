import React, { useState } from 'react';
import { View, Text, StyleSheet, ScrollView, TouchableOpacity } from 'react-native';
import { Colors } from '@/constants/Colors';
import { Ionicons, FontAwesome5 } from '@expo/vector-icons';

export default function TriggersScreen() {
    const [selectedComponent, setSelectedComponent] = useState<string | null>(null);

    const technologies = [
        { name: 'Google', icon: <Ionicons name="logo-google" size={30} color={Colors.light.google} /> },
        { name: 'Github', icon: <Ionicons name="logo-github" size={30} color={Colors.light.github} /> },
        { name: 'Outlook', icon: <Ionicons name="logo-microsoft" size={30} color={Colors.light.outlook} /> },
        { name: 'Slack', icon: <FontAwesome5 name="slack" size={30} color={Colors.light.slack} /> },
        { name: 'Discord', icon: <FontAwesome5 name="discord" size={30} color={Colors.light.discord} /> },
    ];

    return (
        <View style={styles.container}>
            <View style={styles.servicesContainer}>
                <Text style={styles.servicesTitle}>Services</Text>
                <ScrollView horizontal showsHorizontalScrollIndicator={false} contentContainerStyle={styles.servicesList}>
                    {technologies.map((tech, index) => (
                        <TouchableOpacity key={index} style={styles.techItem} onPress={() => setSelectedComponent(tech.name)}>
                            <View style={styles.iconContainer}>{tech.icon}</View>
                            <Text style={styles.techName}>{tech.name}</Text>
                        </TouchableOpacity>
                    ))}
                </ScrollView>
            </View>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#FFFFFF',
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
