import React, { useState } from 'react';
import { View, Text, TextInput, TouchableOpacity, StyleSheet, SafeAreaView, ScrollView, Image, Modal, Pressable, TouchableNativeFeedback, Linking, Alert } from 'react-native';
import { useRouter } from 'expo-router';
import { Colors } from '@/constants/Colors';
import Button from '@/components/Button';
import { MaterialIcons, Ionicons, FontAwesome, FontAwesome5 } from '@expo/vector-icons';
import ButtonIcon from '@/components/ButtonIcon';
import { CredentialsService } from '@/api/auth/credentials/service';
import * as WebBrowser from 'expo-web-browser';
import { WebView } from 'react-native-webview';

interface ProvidersProps {
    providers: {
        icon: React.JSX.Element;
        name: string;
        text: string;
        className?: string;
    }[];
}

export default function SignUp() {
    const IP = process.env['IPv4'];
    const BASE_URL = `http://${IP}:8000/`;

    const [name, setName] = useState('');
    const [surname, setSurname] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [confirmPassword, setConfirmPassword] = useState('');

    const [modalVisible, setModalVisible] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");

    const router = useRouter();

    const handleSignUp = async () => {
        await CredentialsService.register(email, password)
            .then(() => router.push('/(tabs)/HomeScreen'))
            .catch((error) => {
                setErrorMessage(error.message + "\nPlease try again.");
                setModalVisible(true);
            });
    };

    const handleDismissError = () => {
        setModalVisible(false);
        setErrorMessage("");
    };

    const handleOpenAuth = async (providerName: string) => {
        const url = `http://172.16.6.32:8000/api/oauth2/login?provider=${providerName}`;

        try {
            const result = await WebBrowser.openAuthSessionAsync(url,
                `http://172.16.6.32:8000/api/oauth2/login?provider=${providerName}`,
            );
            if (result.type === 'cancel') {
                Alert.alert('Browser Canceled', 'The browser was closed by the user.');
            } else if (result.type === 'dismiss') {
                Alert.alert('Browser Dismissed', 'The browser was dismissed.');
            } else {
                console.log('Browser opened successfully:', result);
            }
            if (result.type === 'success') {
                console.log('result: ', result?.url);
            }

        } catch (error) {
            Alert.alert('Error', 'Unable to open the URL');
            console.error('Failed to open URL:', error);
        }
    };

    const technologies = [
        { name: 'google', icon: <Ionicons name="logo-google" size={30} color={Colors.light.google} /> },
        { name: 'github', icon: <Ionicons name="logo-github" size={30} color={Colors.light.github} /> },
        { name: 'outlook', icon: <Ionicons name="logo-microsoft" size={30} color={Colors.light.outlook} /> },
    ];

    const data = {
        logo: require('../assets/images/react-logo.png'),
    };

    return (
        <SafeAreaView style={styles.safeArea}>
            <ScrollView contentContainerStyle={styles.scrollContainer}>
                <View style={styles.navbar}>
                    <Image source={data.logo} style={styles.logo} />
                </View>
                <TextInput
                    style={styles.input}
                    placeholder="Name"
                    value={name}
                    onChangeText={setName}
                />
                <TextInput
                    style={styles.input}
                    placeholder="Surname"
                    value={surname}
                    onChangeText={setSurname}
                />
                <TextInput
                    style={styles.input}
                    placeholder="Email"
                    value={email}
                    onChangeText={setEmail}
                />
                <TextInput
                    style={styles.input}
                    placeholder="Password"
                    secureTextEntry
                    value={password}
                    onChangeText={setPassword}
                />
                <TextInput
                    style={styles.input}
                    placeholder="Confirm Password"
                    secureTextEntry
                    value={confirmPassword}
                    onChangeText={setConfirmPassword}
                />
                <Button
                    onPress={handleSignUp}
                    title="SIGN UP"
                />
                <View style={styles.separatorContainer}>
                    <View style={styles.line} />
                    <Text style={styles.orText}>or</Text>
                    <View style={styles.line} />
                </View>
                {technologies.map((tech, index) => (
                    <ButtonIcon
                        key={index}
                        onPress={() => handleOpenAuth(tech.name)}
                        title={"Continue with " + tech.name}
                        icon={tech.icon}
                        backgroundColor="#FFFFFF"
                        borderCol={tech.icon.props.color}
                    />
                ))}
            </ScrollView>
            <Modal
                animationType="fade"
                transparent={true}
                visible={modalVisible}
                onRequestClose={handleDismissError}
            >
                <View style={styles.modalContainer}>
                    <View style={styles.modalContent}>
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
        </SafeAreaView>
    );
}

const styles = StyleSheet.create({
    safeArea: {
        flex: 1,
        backgroundColor: Colors.light.background,
    },
    navbar: {
        alignItems: 'center',
    },
    logo: {
        width: 80,
        height: 80,
    },
    scrollContainer: {
        flexGrow: 1,
        paddingHorizontal: 16,
        justifyContent: 'flex-start',
    },
    title: {
        fontSize: 24,
        fontWeight: 'bold',
        textAlign: 'center',
        marginVertical: 20,
    },
    input: {
        borderWidth: 1,
        borderColor: '#ccc',
        padding: 10,
        marginBottom: 10,
        borderRadius: 8,
    },
    signUpButton: {
        backgroundColor: Colors.light.tabIconSelected,
        padding: 15,
        borderRadius: 8,
        alignItems: 'center',
        marginBottom: 10,
    },
    signUpButtonText: {
        color: '#fff',
        fontSize: 16,
    },
    orText: {
        textAlign: 'center',
        marginHorizontal: 20,
    },
    separatorContainer: {
        flexDirection: 'row',
        alignItems: 'center',
        marginVertical: 20,
    },
    line: {
        flex: 1,
        height: 1,
        backgroundColor: '#ccc',
    },
    servicesButton: {
        padding: 15,
        borderColor: Colors.light.tabIconSelected,
        borderWidth: 1,
        borderRadius: 8,
        alignItems: 'center',
    },
    servicesButtonText: {
        color: Colors.light.tabIconSelected,
    },
    modalContainer: {
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
        backgroundColor: "rgba(0,0,0,0.5)",
    },
    modalContent: {
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
