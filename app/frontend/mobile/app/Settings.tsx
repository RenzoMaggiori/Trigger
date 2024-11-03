import React, { useEffect, useState } from 'react';
import { View, Text, StyleSheet, SafeAreaView, Switch, ScrollView, Modal, TouchableOpacity, TouchableNativeFeedback } from 'react-native';
import { FontAwesome5, Ionicons } from '@expo/vector-icons';
import { Colors } from '@/constants/Colors';
import Button from '@/components/Button';
import AsyncStorage from '@react-native-async-storage/async-storage';
import { SettingsService } from '@/api/settings/service';
import { SyncService } from '@/api/sync/service';
import { ProvidersService } from '@/api/auth/providers/service';

const providers = [
    { name: 'google', icon: <Ionicons name="logo-google" size={30} color={Colors.light.google} /> },
    { name: 'twitch', icon: <FontAwesome5 name="twitch" size={30} color={Colors.light.twitch} /> },
    { name: 'discord', icon: <FontAwesome5 name="discord" size={30} color={Colors.light.discord} /> },
    { name: 'spotify', icon: <FontAwesome5 name="spotify" size={30} color={Colors.light.spotify} /> },
    { name: 'github', icon: <Ionicons name="logo-github" size={30} color={Colors.light.github} /> },
    { name: 'bitbucket', icon: <FontAwesome5 name="bitbucket" size={30} color={Colors.light.bitbucket} /> },
];

export default function Settings() {
    return (
        <SafeAreaView style={styles.safeArea}>
            <ScrollView contentContainerStyle={styles.scrollContainer}>
                <UserInfoCard />
                {providers.map((tech, index) => (
                    <ProviderItem key={index} provider={tech} />
                ))}
            </ScrollView>
        </SafeAreaView>
    );
}

function UserInfoCard() {
    const [user, setUser] = useState<{ email: string; role: string } | null>(null);

    useEffect(() => {
        const fetchUserData = async () => {
            try {
                const userData = await AsyncStorage.getItem('user');
                if (userData) {
                    setUser(JSON.parse(userData));
                }
            } catch (error) {
                console.error("Error fetching user data:", error);
            }
        };
        fetchUserData();
    }, []);

    return (
        <View style={styles.userInfoCard}>
            <Text style={styles.userInfoText}>{user ? `User Email: ${user.email}` : 'Loading...'}</Text>
        </View>
    );
}

type Provider = {
    name: string;
    icon: JSX.Element;
    connected?: boolean;
};

function ProviderItem({ provider }: { provider: Provider }) {
    const [isProfileVisible, setIsProfileVisible] = useState(false);
    const [isConnected, setIsConnected] = useState(false);
    const [modalVisible, setModalVisible] = useState(false);
    const [confirmActionType, setConfirmActionType] = useState<'connect' | 'disconnect' | null>(null);
    const [modalErrVisible, setErrModalVisible] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");

    useEffect(() => {
        const fetchConnectionStatus = async () => {
            try {
                const settings = await SettingsService.getSettings();
                const providerStatus = settings.find((item: any) => item.providerName === provider.name);

                setIsConnected(providerStatus?.active === true);

                const storedStatus = await AsyncStorage.getItem(provider.name);
                if (storedStatus === 'true') {
                    setIsConnected(true);
                }
            } catch (error) {
                console.error(`Error fetching connection status for ${provider.name}:`, error);
            }
        };

        fetchConnectionStatus();
    }, [provider.name]);

    const handleSignIn = async () => {
        try {
            await SyncService.handleSync(provider.name);
            setIsConnected(true);
            await AsyncStorage.setItem(provider.name, 'true');
        } catch (error) {
            setErrorMessage((error as Error).message + "\nPlease try again.");
            setErrModalVisible(true);
        }
    };

    const handleSignOut = async () => {
        try {
            await ProvidersService.disconnectOAuth(provider.name);
            setIsConnected(false);
        } catch (error) {
            console.error(`Error disconnecting from ${provider.name}:`, error);
        }
    };

    const handleConnectDisconnect = () => {
        setModalVisible(true);
        setConfirmActionType(isConnected ? 'disconnect' : 'connect');
    };

    const confirmAction = async () => {
        if (confirmActionType === 'connect') {
            await handleSignIn();
        } else if (confirmActionType === 'disconnect') {
            await handleSignOut();
        }
        setModalVisible(false);
    };

    const handleDismissError = () => {
        setErrModalVisible(false);
        setErrorMessage("");
    };

    return (
        <View style={styles.card}>
            <View style={styles.row}>
                <View style={styles.nameContainer}>
                    <View style={styles.iconContainer}>{provider.icon}</View>
                    <Text style={styles.name}>{provider.name}</Text>
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
            <View style={styles.buttonConnect}>
                <Button
                    title={isConnected ? 'Disconnect' : 'Connect'}
                    onPress={handleConnectDisconnect}
                    backgroundColor={Colors.light.tint}
                    textColor="#FFFFFF"
                    buttonWidth="45%"
                    paddingV={7.5}
                />
            </View>
            <Modal
                animationType="slide"
                transparent={true}
                visible={modalVisible}
                onRequestClose={() => setModalVisible(false)}
            >
                <View style={styles.modalContainer}>
                    <View style={styles.modalContent}>
                        <Text style={styles.modalTitle}>
                            {isConnected ? `Disconnect ${provider.name}` : `Connect ${provider.name}`}
                        </Text>
                        <Text style={styles.modalMessage}>
                            Are you sure you want to make this change?
                        </Text>
                        <View style={styles.modalButtons}>
                            <TouchableOpacity onPress={() => setModalVisible(false)} style={styles.cancelButton}>
                                <Text style={styles.buttonText}>CANCEL</Text>
                            </TouchableOpacity>
                            <TouchableOpacity onPress={confirmAction} style={styles.acceptButton}>
                                <Text style={styles.buttonText}>ACCEPT</Text>
                            </TouchableOpacity>
                        </View>
                    </View>
                </View>
            </Modal>
            <Modal
                animationType="fade"
                transparent={true}
                visible={modalErrVisible}
                onRequestClose={handleDismissError}
            >
                <View style={styles.modalErrContainer}>
                    <View style={styles.modalErrContent}>
                        <Text style={styles.errorMessage} numberOfLines={2}>
                            {errorMessage}
                        </Text>
                        <View style={styles.separator} />
                        <TouchableNativeFeedback
                            onPress={handleDismissError}
                            background={TouchableNativeFeedback.Ripple('#f2f0eb', false)}
                        >
                            <View style={styles.dismissButton}>
                                <Text style={styles.dismissButtonText}>DONE</Text>
                            </View>
                        </TouchableNativeFeedback>
                    </View>
                </View>
            </Modal>
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
    userInfoCard: {
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
    userInfoText: {
        fontSize: 16,
        fontWeight: 'bold',
        color: Colors.light.tintDark,
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
    buttonConnect: {
        marginTop: 20,
        flexDirection: 'row',
        justifyContent: 'flex-end',
    },
    modalContainer: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: 'rgba(0, 0, 0, 0.7)',
    },
    modalContent: {
        width: 300,
        backgroundColor: '#fff',
        padding: 20,
        borderRadius: 10,
        alignItems: 'center',
    },
    modalTitle: {
        fontSize: 18,
        fontWeight: 'bold',
        marginBottom: 15,
    },
    modalMessage: {
        fontSize: 16,
        marginBottom: 20,
    },
    modalButtons: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        width: '100%',
    },
    cancelButton: {
        backgroundColor: '#d3d3d3',
        padding: 10,
        borderRadius: 5,
        width: '45%',
        alignItems: 'center',
    },
    acceptButton: {
        backgroundColor: Colors.light.tint,
        padding: 10,
        borderRadius: 5,
        width: '45%',
        alignItems: 'center',
    },
    buttonText: {
        color: '#fff',
        fontWeight: 'bold',
    },
    modalErrContainer: {
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
        backgroundColor: "rgba(0,0,0,0.5)",
    },
    modalErrContent: {
        backgroundColor: "#fff",
        padding: 20,
        borderRadius: 4,
        width: "80%",
        elevation: 5,
    },
    errorMessage: {
        color: "#f25749",
        marginBottom: 10,
        marginTop: 10,
        textAlign: "center",
        fontSize: 16,
    },
    dismissButton: {
        marginTop: 10,
        padding: 10,
        alignItems: "center",
        justifyContent: "center",
    },
    dismissButtonText: {
        color: "#f25749",
        fontWeight: "bold",
    },
    separator: {
        height: 1,
        backgroundColor: "#f2f0eb",
        marginVertical: 12,
    },
});
