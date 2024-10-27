import React from 'react';
import { TouchableOpacity, Text, StyleSheet, ViewStyle, DimensionValue } from 'react-native';

interface ButtonProps {
    onPress: () => void;
    title: string;
    backgroundColor?: string;
    textColor?: string;
    borderCol?: string;
    buttonWidth?: DimensionValue;
    paddingV?: DimensionValue;
    style?: ViewStyle;
}

const Button: React.FC<ButtonProps> = ({ onPress, title, backgroundColor = '#0a7ea4', textColor = '#FFFFFF', borderCol = '#ddd', buttonWidth = '100%', paddingV = 15, style }) => {
    return (
        <TouchableOpacity style={[styles.button, { backgroundColor, borderColor: borderCol, width: buttonWidth, paddingVertical: paddingV }, style]} onPress={onPress}>
            <Text style={[styles.buttonText, { color: textColor }]}>{title}</Text>
        </TouchableOpacity>
    );
};

const styles = StyleSheet.create({
    button: {
        borderWidth: 1,
        paddingHorizontal: 15,
        borderRadius: 30,
        alignItems: 'center',
        justifyContent: 'center',
    },
    buttonText: {
        fontSize: 16,
    },
});

export default Button;
