import React, { useState } from 'react';
import { View, Text, StyleSheet, TouchableOpacity } from 'react-native';
import Animated from 'react-native-reanimated';
// @ts-ignore
import { Ionicons } from '@expo/vector-icons';

interface ActionBox {
    name: string;
    icon: React.ReactElement;
}

interface TechBoxProps extends ActionBox {
    onDelete: () => void;
}

export default function TechBox({ name, icon, onDelete }: TechBoxProps) {

    return (
    <Animated.View style={styles.contentBox}>
        <Animated.View style={styles.box}>
            <TouchableOpacity style={{ position: 'absolute', top: 5, left: 5 }} onPress={onDelete}>
                <Ionicons name="close" size={20} color="red" />
            </TouchableOpacity>
            <View style={styles.iconContainer}>{icon}</View>
            <Text style={styles.techName}>{name}</Text>
        </Animated.View>
    </Animated.View>
    );
}

const styles = StyleSheet.create({
    contentBox: {
        justifyContent: 'center',
        alignItems: 'center',
        flex: 1,
    },
    iconContainer: {
        marginBottom: 5,
    },
    techName: {
        fontSize: 16,
        fontWeight: 'bold',
    },
    box: {
        width: 100,
        height: 100,
        backgroundColor: '#ebebeb',
        borderRadius: 15,
        justifyContent: 'center',
        alignItems: 'center',
    },
});
