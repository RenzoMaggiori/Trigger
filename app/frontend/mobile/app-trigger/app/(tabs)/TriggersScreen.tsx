import React from 'react';
import { View, Text, StyleSheet } from 'react-native';

export default function TriggersScreen() {
    return (
        <View style={styles.container}>
            <Text>Triggers screen content</Text>
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
