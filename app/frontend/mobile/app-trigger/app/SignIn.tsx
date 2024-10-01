import React, { useState } from 'react';
import { View, Text, TextInput, TouchableOpacity, StyleSheet } from 'react-native';
import { useRouter } from 'expo-router';
import { Colors } from '@/constants/Colors';
import Button from '@/components/Button';

export default function SignIn() {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const router = useRouter();

    const handleSignIn = () => {
        router.push('/(tabs)/HomeScreen');
    };

    return (
        <View style={styles.container}>
            <Text style={styles.title}>SIGN IN</Text>
            <TextInput
                style={styles.input}
                placeholder="Email"
                value={email}
                // onChangeText={setEmail}
            />
            <TextInput
                style={styles.input}
                placeholder="Password"
                secureTextEntry
                value={password}
                // onChangeText={setPassword}
            />
            <Button
                onPress={handleSignIn}
                title="SIGN IN"
            />
            <Text style={styles.orText}>or</Text>
            <Button
                onPress={handleSignIn}
                title="Services"
                backgroundColor='#FFFFFF'
                textColor='#000000'
            />
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        padding: 16,
        justifyContent: 'center',
        backgroundColor: Colors.light.background,
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
        marginVertical: 10,
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
});
