import React from 'react';
import { View, Text, Button, Image, StyleSheet, TouchableOpacity } from 'react-native';
import { useRouter } from 'expo-router';
import { Colors } from '@/constants/Colors';

export default function LandingPage() {
    const router = useRouter();

    const handleSignIn = () => {
        router.push('/SignIn');
    };

    const handleSignUp = () => {
        router.push('/SignUp');
    };

    return (
        <View style={styles.container}>
            <Image source={require('../assets/images/react-logo.png')} style={styles.logo} />

            <View style={styles.authButtons}>
                <TouchableOpacity style={styles.signInButton} onPress={handleSignIn}>
                    <Text style={styles.authButtonText}>Sign In</Text>
                </TouchableOpacity>
                <TouchableOpacity style={styles.signUpButton} onPress={handleSignUp}>
                    <Text style={styles.authButtonText}>Sign Up</Text>
                </TouchableOpacity>
            </View>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        padding: 16,
        justifyContent: 'space-between',
        backgroundColor: Colors.light.background,
    },
    logo: {
        width: 100,
        height: 100,
        alignSelf: 'center',
        marginVertical: 20,
    },
    authButtons: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        marginBottom: 20,
        width: '100%',
        paddingHorizontal: 16,
    },
    signInButton: {
        backgroundColor: Colors.light.tabIconSelected,
        padding: 15,
        borderRadius: 8,
        width: '45%',
        alignItems: 'center',
    },
    signUpButton: {
        backgroundColor: Colors.light.tabIconSelected,
        padding: 15,
        borderRadius: 8,
        width: '45%',
        alignItems: 'center',
    },
    authButtonText: {
        color: '#FFFFFF',
        fontSize: 16,
    },
});
