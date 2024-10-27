import React, { useState } from 'react';
import { View, Text, Image, StyleSheet, TouchableOpacity, SafeAreaView, ScrollView, Modal, TouchableNativeFeedback } from 'react-native';
import { useRouter } from 'expo-router';
import { Colors } from '@/constants/Colors';
import { MaterialIcons, Ionicons, FontAwesome, FontAwesome5 } from '@expo/vector-icons';
import TechCarousel from '@/components/TechCarousel';
import { Video, ResizeMode } from 'expo-av';
import ButtonIcon from '@/components/ButtonIcon';
import Button from '@/components/Button';
import { ProvidersService } from '@/api/auth/providers/service';

export default function LandingPage() {
    const router = useRouter();
    const [errorMessage, setErrorMessage] = useState("");
    const [modalVisible, setModalVisible] = useState(false);

    const handleSignIn = () => {
        router.push('/SignIn');
    };

    const handleSignUp = () => {
        router.push('/SignUp');
    };

    const handleDismissError = () => {
        setModalVisible(false);
        setErrorMessage("");
    };

    const handleSocialSignIn = async (provider: string) => {
        try {
            await ProvidersService.handleOAuth(provider);
            router.push('/(tabs)/HomeScreen');
        } catch (error) {
            setErrorMessage((error as Error).message + "\nPlease try again.");
            setModalVisible(true);
        }
    }

    const data = {
        logo: require('../assets/images/logo.png'),
        slogan: "Connect and Automate Effortlessly",
        description: "Trigger empowers you to connect services seamlessly. Automate tasks and enhance productivity by turning your ideas into efficient workflows.",
        buttons: {
            email: "Start with Email",
            google: "Start with Google",
            signIn: "Sign In",
            signUp: "Sign Up",
        },
    };

    const technologies = [
        { name: 'google', icon: <Ionicons name="logo-google" size={30} color={Colors.light.tintLight} /> },
        { name: 'discord', icon: <FontAwesome5 name="discord" size={30} color={Colors.light.tintLight} /> },
        { name: 'github', icon: <Ionicons name="logo-github" size={30} color={Colors.light.tintLight} /> },
        { name: 'slack', icon: <FontAwesome name="slack" size={30} color={Colors.light.tintLight} /> },
        { name: 'outlook', icon: <Ionicons name="logo-microsoft" size={30} color={Colors.light.tintLight} /> },
    ];

    return (
        <SafeAreaView style={styles.safeArea}>
            <View style={styles.navbar}>
                {/* <Logo width={80} height={80} /> */}
                <Image source={data.logo} style={styles.logo} />
            </View>

            <ScrollView contentContainerStyle={styles.scrollContainer}>
                <Text style={styles.slogan}>{data.slogan}</Text>
                <Text style={styles.description}>{data.description}</Text>
                <View style={styles.authButtonsContainer}>
                    <ButtonIcon
                        onPress={handleSignIn}
                        title={data.buttons.email}
                        icon={<MaterialIcons name="email" size={24} color="#FFFFFF" />}
                        backgroundColor={Colors.light.tabIconSelected}
                        textColor='#FFFFFF'
                    />
                    <ButtonIcon
                        onPress={() => handleSocialSignIn('google')}
                        title={data.buttons.google}
                        icon={<Image source={require('../assets/images/google-logo.png')} style={styles.googleLogo} />}
                        backgroundColor='#FFFFFF'
                        textColor='#000000'
                    />
                </View>
                <TechCarousel technologies={technologies} />
                {/* <View style={styles.videoContainer}>
                    <Video
                        source={require('@/assets/video_placeholder.mov')}
                        style={styles.video}
                        resizeMode={ResizeMode.COVER}
                        shouldPlay
                        isLooping
                        isMuted={false}
                    />
                </View> */}
                {/* <View style={styles.footer}>
                    <Text>Footer with team info</Text>
                </View> */}
                <View style={styles.extraSpace}></View>
            </ScrollView>

            <View style={styles.bottomButtons}>
                <Button
                    onPress={handleSignIn}
                    title={data.buttons.signIn}
                    backgroundColor={Colors.light.tabIconSelected}
                    textColor='#FFFFFF'
                    buttonWidth='45%'
                />
                <Button
                    onPress={handleSignUp}
                    title={data.buttons.signUp}
                    backgroundColor={Colors.light.tabIconDefault}
                    textColor='#FFFFFF'
                    buttonWidth='45%'
                />
            </View>
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
        height: 100,
        alignItems: 'center',
        justifyContent: 'center',
        backgroundColor: '#fff',
        borderBottomWidth: 1,
        borderBottomColor: '#ccc',
        paddingTop: 20,
    },
    logo: {
        resizeMode: 'contain',
        height: 30,
    },
    scrollContainer: {
        flexGrow: 1,
        paddingHorizontal: 16,
        justifyContent: 'flex-start',
    },
    slogan: {
        textAlign: 'center',
        fontSize: 28,
        fontWeight: 'bold',
        marginVertical: 10,
    },
    description: {
        textAlign: 'center',
        fontSize: 16,
        marginBottom: 10,
    },
    authButtonsContainer: {
        marginVertical: 10,
    },
    googleLogo: {
        width: 20,
        height: 20,
    },
    videoContainer: {
        height: 200,
        marginVertical: 20,
    },
    video: {
        width: '100%',
        height: '100%',
        borderRadius: 10,
    },
    footer: {
        height: 100,
        justifyContent: 'center',
        alignItems: 'center',
        marginVertical: 20,
        backgroundColor: '#d0d0d0',
        borderRadius: 10,
    },
    extraSpace: {
        height: 60,
    },
    bottomButtons: {
        position: 'absolute',
        bottom: 20,
        flexDirection: 'row',
        width: '100%',
        justifyContent: 'space-around',
        paddingHorizontal: 16,
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
