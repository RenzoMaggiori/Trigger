import React, { useEffect, useRef, useState } from 'react';
import { View, Text, StyleSheet, TouchableOpacity } from 'react-native';
import Animated, { FadeIn, FadeInUp, FadeOut, FadeOutUp, Layout } from 'react-native-reanimated';
import { AntDesign } from '@expo/vector-icons';

interface ActionBox {
    index: number;
    id: number;
    name: string;
    initialState: React.MutableRefObject<boolean>;
    icon: React.ReactElement;
}

interface ProviderBoxProps extends ActionBox {
    onDelete: (id: number) => void;
}

export default function ProviderBox({ index, id, name, icon, initialState, onDelete }: ProviderBoxProps) {

    return (
    <Animated.View style={styles.boxContent} entering={initialState.current ? FadeIn.delay(100 * index) : FadeIn} exiting={FadeOut} layout={Layout.delay(50)}>
        <TouchableOpacity style={styles.deleteButton} onPress={() => onDelete(id)}>
            <AntDesign name="closecircle" size={25} color="red" />
        </TouchableOpacity>
        <View style={styles.iconContainer}>{icon}</View>
        <Text style={styles.techName}>{name}</Text>
        <Text style={styles.techName}>{id}</Text>
    </Animated.View>
    );
}

const styles = StyleSheet.create({
    boxContent: {
        height: 100,
        width: '90%',
        backgroundColor: 'white',
        justifyContent: 'center',
        alignItems: 'center',
        marginVertical: 10,
        alignSelf: 'center',
        borderRadius: 10,
        elevation: 5,
    },
    deleteButton: {
        position: 'absolute',
        top: 10,
        left: 10,
        zIndex: 1,
    },
    iconContainer: {
        marginBottom: 5,
    },
    techName: {
        fontSize: 16,
        fontWeight: 'bold',
    },
});
