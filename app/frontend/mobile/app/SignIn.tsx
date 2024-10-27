import React, { useState } from 'react';
import { View, Text, TextInput, StyleSheet, SafeAreaView, ScrollView, Image, Modal, Pressable, TouchableNativeFeedback } from 'react-native';
import { useRouter } from 'expo-router';
import { Colors } from '@/constants/Colors';
import Button from '@/components/Button';
import { MaterialIcons, Ionicons } from '@expo/vector-icons';
import ButtonIcon from '@/components/ButtonIcon';
import { CredentialsService } from '@/api/auth/credentials/service';

export default function SignIn() {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [modalVisible, setModalVisible] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");

    const router = useRouter();

    const handleSignIn = async () => {
        await CredentialsService.login(email, password)
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

    const handleSocialSignIn = (service: string) => {
        router.push('/(tabs)/HomeScreen');
        console.log(`Signing up with ${service}`);
    };

    const technologies = [
        { name: 'Google', icon: <Ionicons name="logo-google" size={30} color={Colors.light.google} />},
        { name: 'Github', icon: <Ionicons name="logo-github" size={30} color={Colors.light.github} /> },
        { name: 'Outlook', icon: <Ionicons name="logo-microsoft" size={30} color={Colors.light.outlook} /> },
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
                <Button
                    onPress={handleSignIn}
                    title="SIGN IN"
                />
                <View style={styles.separatorContainer}>
                    <View style={styles.line} />
                    <Text style={styles.orText}>or</Text>
                    <View style={styles.line} />
                </View>
                {technologies.map((tech, index) => (
                    <ButtonIcon
                        key={index}
                        onPress={() => handleSocialSignIn(tech.name)}
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
        marginBottom: 20,
    },
    input: {
        borderWidth: 1,
        borderColor: '#ccc',
        padding: 10,
        marginBottom: 10,
        borderRadius: 8,
    },
    signInButton: {
        backgroundColor: Colors.light.tabIconSelected,
        padding: 15,
        borderRadius: 8,
        alignItems: 'center',
        marginBottom: 10,
    },
    signInButtonText: {
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
