import React from 'react';
import { View, Text, Image, StyleSheet, TouchableOpacity, SafeAreaView, ScrollView } from 'react-native';
import { useRouter } from 'expo-router';
import { Colors } from '@/constants/Colors';
import { MaterialIcons, Ionicons, FontAwesome, FontAwesome5 } from '@expo/vector-icons';
import TechCarousel from '@/components/TechCarousel';
import { Video, ResizeMode } from 'expo-av';
import ButtonIcon from '@/components/ButtonIcon';
import Button from '@/components/Button';

export default function LandingPage() {
    const router = useRouter();

    const handleSignIn = () => {
        router.push('/SignIn');
    };

    const handleSignUp = () => {
        router.push('/SignUp');
    };

    const data = {
        logo: require('../assets/images/react-logo.png'),
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
        { name: 'Gmail', icon: <MaterialIcons name="email" size={30} color={Colors.light.tintLight} /> },
        { name: 'Discord', icon: <FontAwesome5 name="discord" size={30} color={Colors.light.tintLight} /> },
        { name: 'Github', icon: <Ionicons name="logo-github" size={30} color={Colors.light.tintLight} /> },
        { name: 'Slack', icon: <FontAwesome name="slack" size={30} color={Colors.light.tintLight} /> },
        { name: 'Outlook', icon: <Ionicons name="logo-microsoft" size={30} color={Colors.light.tintLight} /> },
    ];

    return (
        <SafeAreaView style={styles.safeArea}>
            <View style={styles.navbar}>
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
                        onPress={handleSignUp}
                        title={data.buttons.google}
                        icon={<Image source={require('../assets/images/google-logo.png')} style={styles.googleLogo} />}
                        backgroundColor='#FFFFFF'
                        textColor='#000000'
                    />
                </View>
                <TechCarousel technologies={technologies} />
                <View style={styles.videoContainer}>
                    <Video
                        source={require('@/assets/video_placeholder.mov')}
                        style={styles.video}
                        resizeMode={ResizeMode.COVER}
                        shouldPlay
                        isLooping
                        isMuted={false}
                    />
                </View>
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
        width: 80,
        height: 80,
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
});
