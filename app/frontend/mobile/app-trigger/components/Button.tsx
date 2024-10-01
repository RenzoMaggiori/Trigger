import React from 'react';
import { TouchableOpacity, Text, StyleSheet, ViewStyle, DimensionValue } from 'react-native';

interface ButtonProps {
    onPress: () => void;
    title: string;
    backgroundColor?: string;
    textColor?: string;
    buttonWidth?: DimensionValue;
    style?: ViewStyle;
}

const Button: React.FC<ButtonProps> = ({ onPress, title, backgroundColor = '#0a7ea4', textColor = '#FFFFFF', buttonWidth = '100%', style }) => {
    return (
        <TouchableOpacity style={[styles.button, { backgroundColor, width: buttonWidth }, style]} onPress={onPress}>
            <Text style={[styles.buttonText, { color: textColor }]}>{title}</Text>
        </TouchableOpacity>
    );
};

const styles = StyleSheet.create({
    button: {
        borderWidth: 1,
        borderColor: '#ddd',
        padding: 15,
        borderRadius: 30,
        alignItems: 'center',
        justifyContent: 'center',
    },
    buttonText: {
        fontSize: 16,
    },
});

export default Button;
