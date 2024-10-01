import React from 'react';
import { View, Text, StyleSheet } from 'react-native';

export default function CommunityScreen() {
    return (
        <View style={styles.container}>
            <Text>Community screen content</Text>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
    },
});
