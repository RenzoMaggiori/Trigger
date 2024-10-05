import React, { useState } from 'react';
import { View, Text, StyleSheet, SafeAreaView, Switch, ScrollView } from 'react-native';
import { FontAwesome5, Ionicons, FontAwesome } from '@expo/vector-icons';
import { Colors } from '@/constants/Colors';

const technologies = [
    { name: 'Google', icon: <Ionicons name="logo-google" size={30} color={Colors.light.google} /> },
    { name: 'Discord', icon: <FontAwesome5 name="discord" size={30} color={Colors.light.discord} /> },
    { name: 'Github', icon: <Ionicons name="logo-github" size={30} color={Colors.light.github} /> },
    { name: 'Slack', icon: <FontAwesome name="slack" size={30} color={Colors.light.slack} /> },
    { name: 'Outlook', icon: <Ionicons name="logo-microsoft" size={30} color={Colors.light.outlook} /> },
];

export default function SettingsScreen() {
    return (
        <SafeAreaView style={styles.safeArea}>
            <ScrollView contentContainerStyle={styles.scrollContainer}>
                {technologies.map((tech, index) => (
                    <TechnologyItem key={index} technology={tech} />
                ))}
            </ScrollView>
        </SafeAreaView>
    );
}

type Technology = {
    name: string;
    icon: JSX.Element;
    connected?: boolean;
};

function TechnologyItem({ technology }: { technology: Technology }) {
    const [isProfileVisible, setIsProfileVisible] = useState(false);
    const [isConnected, setIsConnected] = useState(technology.connected);

    return (
        <View style={styles.card}>
            <View style={styles.row}>
                <View style={styles.nameContainer}>
                    <View style={styles.iconContainer}>{technology.icon}</View>
                    <Text style={styles.name}>{technology.name}</Text>
                </View>
                <View style={styles.statusContainer}>
                    <Text style={isConnected ? styles.connected : styles.disconnected}>
                        {isConnected ? 'Connected' : 'Disconnected'}
                    </Text>
                    <View
                        style={[
                            styles.statusCircle,
                            { backgroundColor: isConnected ? 'green' : 'red' },
                        ]}
                    />
                </View>
            </View>
            <View style={styles.switchRow}>
                <Text>Show on Profile</Text>
                <Switch
                    value={isProfileVisible}
                    onValueChange={() => setIsProfileVisible(!isProfileVisible)}
                />
            </View>
            <View style={styles.switchRow}>
                <Text>Connection</Text>
                <Switch
                    value={isConnected}
                    onValueChange={() => setIsConnected(!isConnected)}
                />
            </View>
        </View>
    );
}

const styles = StyleSheet.create({
    safeArea: {
        flex: 1,
        backgroundColor: Colors.light.tintLight,
        paddingHorizontal: 16,
    },
    scrollContainer: {
        flexGrow: 1,
        paddingHorizontal: 16,
        justifyContent: 'flex-start',
    },
    card: {
        backgroundColor: Colors.light.background,
        borderRadius: 10,
        padding: 20,
        marginVertical: 10,
        shadowColor: '#000',
        shadowOpacity: 0.1,
        shadowRadius: 5,
        shadowOffset: { width: 0, height: 3 },
        elevation: 3,
    },
    row: {
        flexDirection: 'row',
        alignItems: 'center',
        justifyContent: 'space-between',
        marginBottom: 20,
    },
    nameContainer: {
        flexDirection: 'row',
        alignItems: 'center',
    },
    iconContainer: {
        marginRight: 10,
    },
    name: {
        fontSize: 18,
        fontWeight: 'bold',
    },
    statusContainer: {
        flexDirection: 'row',
        alignItems: 'center',
    },
    connected: {
        color: 'green',
        fontWeight: 'bold',
        marginRight: 8,
    },
    disconnected: {
        color: 'red',
        fontWeight: 'bold',
        marginRight: 8,
    },
    statusCircle: {
        width: 12,
        height: 12,
        borderRadius: 6,
    },
    switchRow: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: 'center',
    },
});
